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
    (unless (> (length (getf *planner-policy* :task_graph)) 0)
      (error "Planner policy must include a task graph"))
    (unless (find "linear_sync"
                  (coerce (map 'vector (lambda (task) (getf task :id))
                               (getf *planner-policy* :task_graph))
                          'list)
                  :test #'string=)
      (error "Planner task graph must include Linear sync boundary"))
    t))

(defun write-generated-artifacts (root)
  (validate-ssot)
  (write-json-file (merge-pathnames "generated/commands.generated.json" root)
                   (list :project *project*
                         :commands (coerce *commands* 'vector)))
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
