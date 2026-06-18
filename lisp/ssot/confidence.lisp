(in-package #:myhome-jarvis.ssot)

(defparameter *confidence-policy*
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/confidence.generated.json"
        :assessor_key "confidence_assessor"
        :confidence_is_cap t
        :self_report_allowed nil
        :public_status_redacted t
        :raw_evidence_public_allowed nil
        :levels #("blocked" "low" "medium" "high")
        :inputs #("evidence_graph"
                  "learning_ledger"
                  "quality_gate"
                  "public_safety")
        :cap_rules #((:key "public_safety_findings"
                      :when "public_safety_not_ok"
                      :cap "blocked")
                     (:key "quality_failing"
                      :when "latest_quality_failed"
                      :cap "blocked")
                     (:key "missing_evidence_links"
                      :when "evidence_edge_count_zero"
                      :cap "low")
                     (:key "dangling_refs"
                      :when "dangling_evidence_ref_count_positive"
                      :cap "low")
                     (:key "open_learning_debt"
                      :when "open_learning_count_positive"
                      :cap "medium")
                     (:key "quality_unrecorded"
                      :when "latest_quality_missing"
                      :cap "medium")
                     (:key "evidence_backed"
                      :when "evidence_links_and_verification_clear"
                      :cap "high"))
        :public_summary_fields #("policy_path"
                                 "assessor_key"
                                 "level_cap"
                                 "blocked"
                                 "self_report_allowed"
                                 "evidence_link_count"
                                 "dangling_evidence_ref_count"
                                 "open_learning_count"
                                 "quality_recorded"
                                 "quality_ok"
                                 "public_safety_ok"
                                 "active_rule"
                                 "checked_at")
        :forbidden_public_fields #("summary"
                                   "next_action"
                                   "evidence_refs"
                                   "raw_evidence"
                                   "raw_prompt"
                                   "raw_transcript"
                                   "token"
                                   "secret"
                                   "credential"
                                   "cookie"
                                   "account_id"
                                   "card_number"
                                   "local_absolute_path")
        :commands #("mhj confidence status")))
