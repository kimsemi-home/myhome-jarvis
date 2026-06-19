(in-package #:myhome-jarvis.ssot)

(defun authority-policy-core ()
  (list :context "AgentCluster"
        :version "v1"
        :generated_artifact "generated/authority.generated.json"
        :public_status_redacted t
        :self_authority_allowed nil
        :reasoning_tier_grants_approval nil
        :public_repo_high_risk_blocked t
        :required_inputs #("confidence_assessor"
                           "evidence_quality"
                           "incident_lifecycle"
                           "control_plane"
                           "translation"
                           "human_review"
                           "public_safety")))
