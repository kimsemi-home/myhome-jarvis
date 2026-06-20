(in-package #:myhome-jarvis.ssot)

(defparameter *assistant-universal-terms*
  #((:key "intent" :meaning "user goal captured without private payload")
    (:key "capability" :meaning "bounded assistant skill with authority rules")
    (:key "evidence" :meaning "repo-relative proof for a claim")
    (:key "decision" :meaning "authority-gated choice with lease and reviewer")
    (:key "work_item" :meaning "public-safe closed-loop unit of intent and evidence")
    (:key "cost_unit" :meaning "Codex token or paid service usage unit")
    (:key "monetization_loop" :meaning "experiment tied to revenue evidence")
    (:key "repo_factory" :meaning "repeatable public-safe repo creation flow")
    (:key "merge_evidence" :meaning "proof that eligible PR work reached main")
    (:key "evidence_retention" :meaning "local log compression archive lifecycle")
    (:key "evidence_noise_budget" :meaning "bounded low-signal evidence budget")
    (:key "household_scope" :meaning "user spouse household or shared view")))

(defparameter *assistant-capability-pillars*
  #((:key "local_media_concierge"
     :features #("youtube_search" "ott_launch" "playback_readiness")
     :authority "local_interactive" :evidence "harness home")
    (:key "household_finance_copilot"
     :features #("spouse_scope" "subscription_review" "cashflow_summary")
     :authority "read_only_review" :evidence "harness finance")
    (:key "shorts_factory_control_plane"
     :features #("repo_bootstrap" "codex_project_setup" "publishing_checks")
     :authority "review_required" :evidence "verification graph")
    (:key "monetization_console"
     :features #("experiment_backlog" "revenue_hypothesis" "unit_economics")
     :authority "review_required" :evidence "Linear project")
    (:key "codex_cost_governor"
     :features #("usage_ledger" "budget_alert" "roi_review")
     :authority "local_readonly" :evidence "generated codex cost policy")
    (:key "self_improvement_loop"
     :features #("pdca" "learning_ledger" "incident_feedback")
     :authority "authority_gated" :evidence "generated pdca")))

(defparameter *assistant-guardrails*
  #("local_first" "public_safe_by_default" "private_data_stays_private"
    "finance_actions_are_review_only" "no_cookie_or_credential_import"
    "no_self_approval" "cost_tracking_before_external_scale"
    "repo_creation_requires_public_safety_scan"
    "evidence_noise_budget_blocks_archive_promotion"))

(defparameter *assistant-linear-epics*
  #("Universal Language and Governance"
    "Local Media Concierge"
    "Household Finance Copilot"
    "Shorts Factory Repo Control Plane"
    "Monetization Console"
    "Codex Cost Governor"
    "Self-Improvement Command Center"
    "Evidence Retention and Noise Budget"
    "Security Privacy and Authority Hardening"))

(defparameter *assistant-vision-policy*
  (list :context "AssistantVision"
        :version "v1"
        :mission "local-first household executive assistant"
        :generated_artifact "generated/assistant_vision.generated.json"
        :closed_loop_until "trusted daily oversight assistant"
        :operating_mode "observe-plan-act-verify-learn"
        :universal_terms *assistant-universal-terms*
        :capability_pillars *assistant-capability-pillars*
        :guardrails *assistant-guardrails*
        :linear_epics *assistant-linear-epics*
        :public_summary_fields #("context" "version" "mission"
                                 "operating_mode" "capability_pillars"
                                 "guardrails" "linear_epics")
        :forbidden_public_fields #("token" "secret" "credential" "cookie"
                                   "account_id" "card_number" "phone" "email"
                                   "local_absolute_path" "raw_prompt"
                                   "raw_transcript" "reviewer_identity")))
