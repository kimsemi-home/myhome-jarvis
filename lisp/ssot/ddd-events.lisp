(in-package #:myhome-jarvis.ssot)

(defparameter *domain-events*
  #((:name "CheckpointRecorded"
     :bounded_context "AgentOps"
     :description "Emitted when a closed-loop cycle writes private checkpoint evidence."
     :emitted_by "ClosedLoopPlanner"
     :payload_fields #("linear_status" "planner_status"
                       "knowledge_evidence" "security_status"))
    (:name "KnowledgeLookupRecorded"
     :bounded_context "KnowledgeIndex"
     :description "Recorded in planner status and checkpoints when a pre-plan KnowledgeIndex lookup runs."
     :emitted_by "KnowledgeIndex"
     :payload_fields #("query" "concept_count" "hit_count"
                       "linear_issues" "must_read"))))

(defparameter *harness-case-contracts*
  #((:name "home_control_golden"
     :bounded_context "HomeControl"
     :command "mhj harness home"
     :evidence_target "harness/golden/home_control.golden.json"
     :description "Home command harness must stay aligned with generated command catalog.")
    (:name "finance_fixture"
     :bounded_context "HouseholdFinance"
     :command "mhj harness finance"
     :evidence_target "fixtures/finance_transactions.jsonl"
     :description "Finance harness must use fixture-first transaction IR without raw finance data.")
    (:name "commerce_fixture"
     :bounded_context "CommerceIntelligence"
     :command "mhj harness commerce"
     :evidence_target "fixtures/commerce_purchases.jsonl"
     :description "Commerce harness must use fixture-first purchase IR without raw commerce exports.")))
