(in-package #:myhome-jarvis.ssot)

(defun write-all-generated-artifacts (root)
  (write-json-file (artifact-path root "commands")
                   (list :project *project*
                         :commands (coerce *commands* 'vector)))
  (write-json-file (artifact-path root "concepts")
                   (list :bounded_contexts *bounded-contexts*
                         :ddd_patterns *ddd-patterns*
                         :concepts *concept-registry*
                         :domain_events *domain-events*
                         :harness_case_contracts *harness-case-contracts*
                         :generated_artifact_contracts *generated-artifact-contracts*
                         :planning_rules *planning-rules*
                         :knowledge_index_schema *knowledge-index-schema*))
  (write-json-file (artifact-path root "finance")
                   (list :entities *finance-entities*
                         :transaction_ir *transaction-ir*))
  (write-json-file (artifact-path root "commerce")
                   (list :entities *commerce-entities*
                         :purchase_ir *purchase-ir*))
  (dolist (artifact (policy-artifacts))
    (write-json-file (artifact-path root (first artifact)) (second artifact)))
  (write-control-plane-verification root)
  (write-quality-workflows root)
  (write-verification-backends root)
  (write-verification-schema root)
  (write-verification-evidence root)
  (write-verification-conformance root)
  (write-verification-tests root)
  (write-release-pipeline root)
  (write-verification-doc root))

(defun artifact-path (root name)
  (merge-pathnames (format nil "generated/~A.generated.json" name) root))

(defun policy-artifacts ()
  `(("storage" ,*storage-policy*)
    ("household" ,*household-policy*)
    ("finance_consent" ,*finance-consent-policy*)
    ("recommendations" ,*recommendation-policy*)
    ("scheduler" ,*scheduler-policy*)
    ("security" ,*security-policy*)
    ("connectors" ,*connector-policy*)
    ("agent_cluster" ,*agent-cluster-policy*)
    ("learning" ,*learning-policy*)
    ("evidence" ,*evidence-graph-policy*)
    ("confidence" ,*confidence-policy*)
    ("translation" ,*translation-policy*)
    ("control_plane" ,*control-plane-policy*)
    ("incidents" ,*incident-policy*)
    ("evidence_quality" ,*evidence-quality-policy*)
    ("review" ,*review-policy*)
    ("assistant_vision" ,*assistant-vision-policy*)
    ("codex_cost" ,*codex-cost-policy*)
    ("codex_sustainability" ,*codex-sustainability-policy*)
    ("context_pack" ,*context-pack-policy*)
    ("media_readiness" ,*media-readiness-policy*)
    ("merge_evidence" ,*merge-evidence-policy*)
    ("monetization" ,*monetization-policy*)
    ("repo_factory" ,*repo-factory-policy*)
    ("authority" ,*authority-policy*)
    ("pdca" ,*pdca-policy*)
    ("code_shape" ,*code-shape-policy*)
    ("verification_graph" ,*verification-graph*)
    ("linear" ,*linear-policy*)
    ("planner" ,*planner-policy*)))
