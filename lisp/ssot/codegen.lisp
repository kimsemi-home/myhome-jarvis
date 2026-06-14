(in-package #:myhome-jarvis.ssot)

(defun validate-ssot ()
  (let ((command-names (mapcar (lambda (command) (getf command :name)) *commands*)))
    (dolist (required '("open_youtube"
                        "open_youtube_search"
                        "open_netflix"
                        "open_disney_plus"
                        "open_tving"
                        "open_wavve"
                        "open_coupang_play"
                        "open_ott"
                        "open_url"
                        "volume_set"
                        "volume_up"
                        "volume_down"
                        "display_sleep"
                        "movie_mode"
                        "sleep_mode"))
      (unless (member required command-names :test #'string=)
        (error "Missing required command: ~A" required)))
    (when (find "Python" *allowed-languages* :test #'string=)
      (error "Python must not be an allowed language"))
    (unless (find "raw" (getf *storage-policy* :lake_layers) :test #'string=)
      (error "Storage policy must include raw layer"))
    (unless (find "spouse" (coerce (getf *household-policy* :scopes) 'list) :test #'string=)
      (error "Household policy must include spouse scope"))
    (unless (find "subscription_review" (coerce (getf *recommendation-policy* :kinds) 'list) :test #'string=)
      (error "Recommendation policy must include subscription review"))
    (unless (getf *scheduler-policy* :crash_recovery)
      (error "Scheduler policy must require crash recovery"))
    (unless (getf *security-policy* :lan_requires_bearer_token)
      (error "LAN daemon access must require a bearer token"))
    (unless (getf *security-policy* :current_content_scan)
      (error "Current-tree content scanning must stay enabled"))
    (unless (getf *security-policy* :current_content_scan_skips_private_paths)
      (error "Current-tree content scanning must skip private paths"))
    (when (getf *security-policy* :report_matched_secret_contents)
      (error "Security reports must not expose matched secret contents"))
    (unless (getf *linear-policy* :pull_active_only)
      (error "Linear pull must stay scoped to active issues by default"))
    (unless (getf *linear-policy* :team_scope_private)
      (error "Linear team scope must stay private"))
    (unless (string= (getf *linear-policy* :project_issue_title_prefix)
                     "[myhome-jarvis]")
      (error "Linear project issue title prefix must identify myhome-jarvis work"))
    (unless (getf *linear-policy* :next_prefers_project_issues)
      (error "Linear next must prefer project issues"))
    (unless (getf *linear-policy* :next_requires_project_issue)
      (error "Linear next must not select unrelated active issues"))
    (unless (getf *linear-policy* :backlog_seed_current_project_only)
      (error "Linear backlog seeds must represent current project work"))
    (unless (getf *linear-policy* :backlog_seed_dedupes_by_title)
      (error "Linear backlog seeding must dedupe existing issue titles"))
    (unless (getf *linear-policy* :backlog_seed_queries_existing_titles)
      (error "Linear backlog seeding must query existing issue titles"))
    (validate-ddd-registry)
    (unless (> (length (getf *planner-policy* :task_graph)) 0)
      (error "Planner policy must include a task graph"))
    (unless (find "linear_sync"
                  (coerce (map 'vector (lambda (task) (getf task :id))
                               (getf *planner-policy* :task_graph))
                          'list)
                  :test #'string=)
      (error "Planner task graph must include Linear sync boundary"))
    (unless (getf *planner-policy* :knowledge_index_required_before_planning)
      (error "Planner must require KnowledgeIndex lookup before planning"))
    t))

(defun validate-ddd-registry ()
  (let ((contexts (make-hash-table :test #'equal))
        (concepts (make-hash-table :test #'equal))
        (aliases (make-hash-table :test #'equal)))
    (dolist (context (coerce *bounded-contexts* 'list))
      (let ((name (getf context :name)))
        (when (or (null name) (string= name ""))
          (error "Bounded context name is required"))
        (when (gethash name contexts)
          (error "Duplicate bounded context: ~A" name))
        (setf (gethash name contexts) t)))
    (dolist (pattern (coerce *ddd-patterns* 'list))
      (unless (find pattern '("Entity" "ValueObject" "Aggregate" "DomainEvent"
                              "Repository" "Policy" "Port" "Adapter"
                              "AntiCorruptionLayer")
                    :test #'string=)
        (error "Unknown DDD pattern: ~A" pattern)))
    (dolist (concept (coerce *concept-registry* 'list))
      (let ((name (getf concept :canonical_name))
            (context (getf concept :bounded_context))
            (targets (getf concept :generated_targets)))
        (when (or (null name) (string= name ""))
          (error "Concept canonical_name is required"))
        (when (gethash name concepts)
          (error "Duplicate concept: ~A" name))
        (setf (gethash name concepts) t)
        (unless (gethash context contexts)
          (error "Concept ~A references unknown bounded context ~A" name context))
        (unless (> (length targets) 0)
          (error "Concept ~A must declare generated targets" name))
        (dolist (alias (coerce (getf concept :allowed_aliases) 'list))
          (let ((key (string-downcase alias)))
            (when (or (string= key "") (gethash key aliases))
              (error "Duplicate or empty concept alias: ~A" alias))
            (setf (gethash key aliases) name)))))
    (dolist (concept (coerce *concept-registry* 'list))
      (dolist (related (coerce (getf concept :related_concepts) 'list))
        (unless (gethash related concepts)
          (error "Concept ~A references unknown related concept ~A"
                 (getf concept :canonical_name) related))))
    t))

(defun write-generated-artifacts (root)
  (validate-ssot)
  (write-json-file (merge-pathnames "generated/commands.generated.json" root)
                   (list :project *project*
                         :commands (coerce *commands* 'vector)))
  (write-json-file (merge-pathnames "generated/concepts.generated.json" root)
                   (list :bounded_contexts *bounded-contexts*
                         :ddd_patterns *ddd-patterns*
                         :concepts *concept-registry*
                         :generated_artifact_contracts *generated-artifact-contracts*
                         :planning_rules *planning-rules*
                         :knowledge_index_schema *knowledge-index-schema*))
  (write-json-file (merge-pathnames "generated/finance.generated.json" root)
                   (list :entities *finance-entities*
                         :transaction_ir *transaction-ir*))
  (write-json-file (merge-pathnames "generated/commerce.generated.json" root)
                   (list :entities *commerce-entities*
                         :purchase_ir *purchase-ir*))
  (write-json-file (merge-pathnames "generated/storage.generated.json" root)
                   *storage-policy*)
  (write-json-file (merge-pathnames "generated/household.generated.json" root)
                   *household-policy*)
  (write-json-file (merge-pathnames "generated/recommendations.generated.json" root)
                   *recommendation-policy*)
  (write-json-file (merge-pathnames "generated/scheduler.generated.json" root)
                   *scheduler-policy*)
  (write-json-file (merge-pathnames "generated/security.generated.json" root)
                   *security-policy*)
  (write-json-file (merge-pathnames "generated/linear.generated.json" root)
                   *linear-policy*)
  (write-json-file (merge-pathnames "generated/planner.generated.json" root)
                   *planner-policy*)
  t)

(defun write-json-file (path value)
  (ensure-directories-exist path)
  (with-open-file (stream path
                          :direction :output
                          :if-exists :supersede
                          :if-does-not-exist :create)
    (write-json-value value stream)
    (terpri stream)))

(defun write-json-value (value stream)
  (cond
    ((stringp value) (write-json-string value stream))
    ((integerp value) (princ value stream))
    ((eq value t) (princ "true" stream))
    ((null value) (princ "false" stream))
    ((vectorp value) (write-json-array value stream))
    ((plist-object-p value) (write-json-object value stream))
    ((listp value) (write-json-array (coerce value 'vector) stream))
    (t (write-json-string (princ-to-string value) stream))))

(defun write-json-object (plist stream)
  (princ "{" stream)
  (loop for (key value) on plist by #'cddr
        for first = t then nil
        do (progn
             (unless first (princ "," stream))
             (write-json-string (json-key key) stream)
             (princ ":" stream)
             (write-json-value value stream)))
  (princ "}" stream))

(defun write-json-array (vector stream)
  (princ "[" stream)
  (loop for index from 0 below (length vector)
        do (progn
             (unless (= index 0) (princ "," stream))
             (write-json-value (aref vector index) stream)))
  (princ "]" stream))

(defun write-json-string (value stream)
  (princ "\"" stream)
  (loop for char across value
        do (case char
             (#\" (princ "\\\"" stream))
             (#\\ (princ "\\\\" stream))
             (#\Newline (princ "\\n" stream))
             (#\Return (princ "\\r" stream))
             (#\Tab (princ "\\t" stream))
             (otherwise (princ char stream))))
  (princ "\"" stream))

(defun plist-object-p (value)
  (and (listp value)
       (evenp (length value))
       (loop for tail on value by #'cddr
             always (keywordp (first tail)))))

(defun json-key (key)
  (string-downcase (symbol-name key)))
