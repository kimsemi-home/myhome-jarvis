(in-package #:myhome-jarvis.ssot)

(defparameter *monetization-experiment-states*
  #("backlog" "designed" "review_required" "running" "paused" "measuring"
    "closed" "rejected"))

(defparameter *monetization-decision-kinds*
  #("hypothesis_created" "experiment_started" "scale_requested"
    "pause_requested" "close_experiment" "reject_experiment"))

(defparameter *monetization-review-statuses*
  #("requested" "approved" "rejected" "not_required"))

(defparameter *monetization-expected-value-bands*
  #("unknown" "low" "medium" "high"))

(defparameter *monetization-policy*
  (list :context "MonetizationExperimentLedger"
        :version "v1"
        :generated_artifact "generated/monetization.generated.json"
        :private_experiment_ledger "data/private/monetization/experiments.jsonl"
        :append_only t
        :public_status_redacted t
        :raw_revenue_public_allowed nil
        :decision_evidence_required t
        :cost_estimate_required t
        :experiment_states *monetization-experiment-states*
        :decision_kinds *monetization-decision-kinds*
        :review_statuses *monetization-review-statuses*
        :expected_value_bands *monetization-expected-value-bands*
        :cost_unit_kinds #("codex_tokens" "codex_coin" "github_actions_minutes"
                           "external_tool_cost")
        :required_fields #("at" "experiment_id" "hypothesis_key" "state"
                           "decision_kind" "review_status"
                           "expected_value_band" "cost_estimate_units"
                           "cost_unit_kind" "evidence_refs")
        :allowed_evidence_prefixes #("generated/" "docs/" ".github/"
                                     "data/private/" "internal/" "cmd/")
        :public_summary_fields #("experiment_count" "decision_count"
                                 "invalid_record_count" "review_required_count"
                                 "missing_evidence_count"
                                 "missing_cost_estimate_count"
                                 "expected_value_unknown_count" "by_state"
                                 "by_decision_kind" "by_review_status"
                                 "by_expected_value_band" "checked_at")
        :forbidden_public_fields #("private_revenue_notes" "private_counterparty"
                                   "raw_revenue_amount" "evidence_refs"
                                   "raw_prompt" "token" "secret" "credential")
        :commands #("mhj monetization status")))
