(in-package #:myhome-jarvis.ssot)

(defun authority-decision (key risk public review blocked)
  (list :key key :risk risk :public_repo_allowed public
        :requires_human_review review :allowed_when_blocked blocked))

(defparameter *authority-decisions*
  (vector
   (authority-decision "read_status" "low" t nil t)
   (authority-decision "evidence_collection" "low" t nil t)
   (authority-decision "deterministic_verification" "low" t nil t)
   (authority-decision "revalidation" "low" t nil t)
   (authority-decision "low_risk_fixture_change" "medium" t nil nil)
   (authority-decision "incident_response" "medium" t t t)
   (authority-decision "major_ontology_change" "high" nil t nil)
   (authority-decision "security_boundary_change" "high" nil t nil)
   (authority-decision "production_change" "high" nil t nil)
   (authority-decision "evidence_pruning" "high" nil t nil)
   (authority-decision "quarantine_release" "high" nil t nil)
   (authority-decision "high_risk_automation" "high" nil t nil)))
