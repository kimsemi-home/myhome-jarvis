(in-package #:myhome-jarvis.ssot)

(defparameter *authority-outcomes*
  #("limited" "review_required" "blocked"))

(defparameter *authority-debt-classes*
  #("public_safety" "confidence_cap" "evidence_quality" "incident"
    "control_plane" "translation" "human_review"))

(defparameter *authority-public-summary-fields*
  #("policy_path" "outcome" "active_rule" "input_count" "decision_count"
    "allowed_decision_count" "blocked_decision_count" "authority_debt_count"
    "public_repo_mode" "reasoning_tier_grants_approval"
    "self_authority_allowed" "public_safety_ok" "confidence_cap"
    "evidence_quality_debt_count" "incident_debt_count"
    "control_plane_debt_count" "translation_debt_count"
    "human_review_debt_count" "human_review_capacity_state"
    "allowed_decisions" "blocked_decisions" "by_risk" "profile_count"
    "review_required_profile_count" "public_safety_gated_profile_count"
    "public_repo_change_gated_profile_count" "workflow_change_gated_profile_count"
    "self_approval_blocked_profile_count" "profile_keys"
    "review_required_profiles" "public_safety_gated_profiles" "checked_at"))

(defparameter *authority-forbidden-public-fields*
  #("raw_rationale" "raw_evidence" "evidence_ref" "evidence_refs"
    "raw_prompt" "raw_transcript" "token" "secret" "credential" "cookie"
    "account_id" "card_number" "local_absolute_path" "linear_url"
    "private_evidence"))
