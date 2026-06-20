(in-package #:myhome-jarvis.ssot)

(defun validate-evidence-graph-policy (policy)
  (require-string-equal (getf policy :context) "AgentCluster"
                        "Evidence graph policy must belong to AgentCluster")
  (require-string-equal (getf policy :private_root) "data/private"
                        "Evidence graph private root must stay data/private")
  (require-true (getf policy :private_graph_required)
                "Evidence graph must require private graph evidence")
  (require-true (getf policy :public_status_redacted)
                "Evidence graph public status must stay redacted")
  (require-false (getf policy :raw_evidence_public_allowed)
                 "Evidence graph must not expose raw evidence publicly")
  (require-members '("learning_observation" "evidence_artifact")
                   (policy-list policy :node_kinds)
                   "Evidence graph missing node kind: ~A")
  (require-member "supports" (policy-list policy :edge_kinds)
                  "Evidence graph must include supports edges")
  (validate-evidence-sources policy)
  (require-members '("node_count" "edge_count" "by_node_kind"
                     "by_edge_kind" "checked_at")
                   (policy-list policy :public_summary_fields)
                   "Evidence graph summary missing field: ~A")
  (require-member "data/private/" (policy-list policy :allowed_evidence_prefixes)
                  "Evidence graph evidence refs must allow private paths")
  (require-command policy "mhj evidence status")
  (require-command policy "mhj evidence-integrity status"))

(defun validate-evidence-sources (policy)
  (let ((sources (policy-list policy :private_sources))
        (node-kinds (policy-list policy :node_kinds)))
    (require-true (> (length sources) 0)
                  "Evidence graph sources are required")
    (dolist (source sources)
      (let ((key (getf source :key))
            (path (getf source :path))
            (kind (getf source :node_kind))
            (format (getf source :format)))
        (require-string-value key "Evidence graph source key is required")
        (require-private-path path
                              "Evidence graph source must stay private")
        (require-member kind node-kinds
                        "Evidence graph source ~A has unknown node kind ~A"
                        key kind)
        (require-member format '("jsonl" "directory")
                        "Evidence graph source ~A has unsupported format ~A"
                        key format)))))
