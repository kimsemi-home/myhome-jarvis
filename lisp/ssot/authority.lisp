(in-package #:myhome-jarvis.ssot)

(defparameter *authority-policy*
  (append (authority-policy-core)
          (list :reasoning_tiers *authority-reasoning-tiers*
                :role_permissions *authority-role-permissions*
                :domain_attributes *authority-domain-attributes*
                :decisions *authority-decisions*
                :assistant_authority_profiles *assistant-authority-profiles*
                :outcomes *authority-outcomes*
                :authority_debt_classes *authority-debt-classes*
                :public_summary_fields *authority-public-summary-fields*
                :forbidden_public_fields *authority-forbidden-public-fields*
                :commands #("mhj authority status"
                            "mhj authority-review status"))))
