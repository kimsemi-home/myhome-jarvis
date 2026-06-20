(in-package #:myhome-jarvis.ssot)

(defparameter *finance-consent-statuses*
  #("requested" "granted" "revoked" "expired" "denied"))

(defparameter *finance-consent-review-statuses*
  #("requested" "approved" "rejected"))

(defparameter *finance-consent-kinds*
  #("finance_connector" "spouse_scope" "household_scope"))

(defparameter *finance-consent-policy*
  (list :context "HouseholdFinanceConsentLedger"
        :version "v1"
        :generated_artifact "generated/finance_consent.generated.json"
        :private_consent_ledger "data/private/finance/consent.jsonl"
        :append_only t
        :public_status_redacted t
        :finance_mode "read_only_review_only"
        :read_only t
        :review_only t
        :fixture_only_until_consent t
        :external_writes_allowed nil
        :transfer_actions_allowed nil
        :payment_actions_allowed nil
        :trade_actions_allowed nil
        :card_actions_allowed nil
        :real_connector_requires_active_consent t
        :spouse_scope_requires_active_consent t
        :household_scope_requires_active_consent t
        :consent_kinds *finance-consent-kinds*
        :consent_statuses *finance-consent-statuses*
        :review_statuses *finance-consent-review-statuses*
        :authority_profiles #("finance_review_only")
        :required_fields #("at" "consent_kind" "subject_scope" "status"
                           "review_status" "authority_profile"
                           "evidence_refs")
        :allowed_evidence_prefixes #("data/private/" "generated/" "docs/"
                                     "internal/" "cmd/" "fixtures/")
        :public_summary_fields #("readiness_state" "finance_mode" "exists"
                                 "record_count" "active_consent_count"
                                 "missing_required_consent_count"
                                 "review_required_count"
                                 "missing_evidence_count"
                                 "invalid_record_count"
                                 "revoked_or_expired_count"
                                 "forbidden_action_enabled_count"
                                 "consent_debt_count" "checked_at")
        :forbidden_public_fields #("subject_name" "account_id" "card_number"
                                   "raw_balance" "raw_transaction"
                                   "private_notes" "evidence_refs"
                                   "credential" "cookie" "token" "secret")
        :commands #("mhj finance-consent status")))
