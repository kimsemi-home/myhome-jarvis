(in-package #:myhome-jarvis.ssot)

(defparameter *bounded-contexts*
  #((:name "HomeControl"
     :owner "commands"
     :description "Local home-control commands, dry-run execution plans, and Flutter command surface.")
    (:name "HouseholdFinance"
     :owner "finance"
     :description "Fixture-first household finance entities, transaction IR, and recommendation inputs.")
    (:name "CommerceIntelligence"
     :owner "commerce"
     :description "Fixture-first commerce purchase IR, recurring purchase candidates, and price intelligence.")
    (:name "StorageLake"
     :owner "storage"
     :description "Local lake layout, retention policy, and generated storage contracts.")
    (:name "SecurityPolicy"
     :owner "security"
     :description "Public-repo safety, secret scanning, auth boundaries, and allowed-language policy.")
    (:name "AgentOps"
     :owner "planner"
     :description "Closed-loop planner, Linear/offline work queue, scheduler, quality, and checkpoint evidence.")
    (:name "KnowledgeIndex"
     :owner "knowledge"
     :description "Local lexical concept index that reduces semantic duplication before code changes.")))

(defparameter *ddd-patterns*
  #("Entity"
    "ValueObject"
    "Aggregate"
    "DomainEvent"
    "Repository"
    "Policy"
    "Port"
    "Adapter"
    "AntiCorruptionLayer"))

(defparameter *concept-registry*
  #((:canonical_name "HomeCommand"
     :bounded_context "HomeControl"
     :description "A local-first home-control intent with a validated dry-run execution plan."
     :allowed_aliases #("command" "home command" "intent" "home action")
     :owner "internal/commands"
     :generated_targets #("generated/commands.generated.json"
                          "internal/commands/registry.go"
                          "harness/golden/home_control.golden.json")
     :related_concepts #("SecurityPolicy" "ClosedLoopPlanner"))
    (:canonical_name "HouseholdTransaction"
     :bounded_context "HouseholdFinance"
     :description "A fixture-backed household finance transaction IR item without raw bank data."
     :allowed_aliases #("transaction" "finance transaction" "transaction ir")
     :owner "lisp/ssot/finance.lisp"
     :generated_targets #("generated/finance.generated.json"
                          "fixtures/finance_transactions.jsonl")
     :related_concepts #("StorageLake" "SecurityPolicy"))
    (:canonical_name "CommercePurchase"
     :bounded_context "CommerceIntelligence"
     :description "A fixture-backed purchase IR item used for commerce intelligence and recurring candidates."
     :allowed_aliases #("purchase" "commerce purchase" "purchase ir")
     :owner "lisp/ssot/commerce.lisp"
     :generated_targets #("generated/commerce.generated.json"
                          "fixtures/commerce_purchases.jsonl")
     :related_concepts #("StorageLake" "SecurityPolicy"))
    (:canonical_name "StorageLake"
     :bounded_context "StorageLake"
     :description "The local-only storage policy, lake layers, and generated storage contract."
     :allowed_aliases #("lake" "local lake" "storage policy")
     :owner "lisp/ssot/storage.lisp"
     :generated_targets #("generated/storage.generated.json"
                          "docs/storage.md")
     :related_concepts #("HouseholdTransaction" "CommercePurchase" "SecurityPolicy"))
    (:canonical_name "SecurityPolicy"
     :bounded_context "SecurityPolicy"
     :description "The public-repository boundary for secrets, private markers, auth, and allowed languages."
     :allowed_aliases #("security" "public safety" "secret scan" "allowed languages")
     :owner "lisp/ssot/security.lisp"
     :generated_targets #("generated/security.generated.json"
                          "internal/security/security.go"
                          "docs/security.md")
     :related_concepts #("HomeCommand" "StorageLake" "LinearWorkQueue"))
    (:canonical_name "LinearWorkQueue"
     :bounded_context "AgentOps"
     :description "The Linear or offline queue work source used by the closed loop."
     :allowed_aliases #("linear" "work queue" "offline queue")
     :owner "internal/linear"
     :generated_targets #("generated/linear.generated.json"
                          "docs/linear-workflow.md")
     :related_concepts #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))
    (:canonical_name "ClosedLoopPlanner"
     :bounded_context "AgentOps"
     :description "The observe-plan-change-validate-index-sync-decide loop and checkpoint evidence."
     :allowed_aliases #("planner" "closed loop" "loop" "checkpoint")
     :owner "internal/planner"
     :generated_targets #("generated/planner.generated.json"
                          "internal/planner/status.go"
                          "internal/orchestrator/checkpoint.go"
                          "docs/closed-loop.md")
     :related_concepts #("KnowledgeIndex" "LinearWorkQueue" "SecurityPolicy"))
    (:canonical_name "ConceptRegistry"
     :bounded_context "KnowledgeIndex"
     :description "The SSOT-owned registry of canonical concepts, aliases, owners, and generated targets."
     :allowed_aliases #("concept registry" "canonical concepts" "aliases")
     :owner "lisp/ssot/ddd.lisp"
     :generated_targets #("generated/concepts.generated.json"
                          "lisp/ssot/ddd.lisp")
     :related_concepts #("KnowledgeIndex" "ClosedLoopPlanner"))
    (:canonical_name "KnowledgeIndex"
     :bounded_context "KnowledgeIndex"
     :description "A local lexical index over SSOT, generated artifacts, source, docs, fixtures, harness, backlog, working log, and private offline Linear records."
     :allowed_aliases #("knowledge index" "local index" "lexical index" "knowledge search")
     :owner "internal/knowledge"
     :generated_targets #("generated/concepts.generated.json"
                          "internal/knowledge/index.go"
                          "docs/knowledge-index.md")
     :related_concepts #("ConceptRegistry" "ClosedLoopPlanner" "LinearWorkQueue"))))

(defparameter *generated-artifact-contracts*
  #((:name "concepts" :path "generated/concepts.generated.json" :owner "KnowledgeIndex")
    (:name "commands" :path "generated/commands.generated.json" :owner "HomeControl")
    (:name "finance" :path "generated/finance.generated.json" :owner "HouseholdFinance")
    (:name "commerce" :path "generated/commerce.generated.json" :owner "CommerceIntelligence")
    (:name "storage" :path "generated/storage.generated.json" :owner "StorageLake")
    (:name "household" :path "generated/household.generated.json" :owner "HouseholdFinance")
    (:name "recommendations" :path "generated/recommendations.generated.json" :owner "CommerceIntelligence")
    (:name "scheduler" :path "generated/scheduler.generated.json" :owner "AgentOps")
    (:name "security" :path "generated/security.generated.json" :owner "SecurityPolicy")
    (:name "linear" :path "generated/linear.generated.json" :owner "AgentOps")
    (:name "planner" :path "generated/planner.generated.json" :owner "AgentOps")))

(defparameter *planning-rules*
  (list :knowledge_index_required_before_planning t
        :default_knowledge_query "planner KnowledgeIndex Linear closed loop"
        :semantic_changes_require_ssot_first t
        :ssot_change_requires_codegen t
        :small_cohesive_change_required t
        :validation_steps #("focused test"
                            "harness"
                            "codegen verify"
                            "ddd verify"
                            "security check")))

(defparameter *knowledge-index-schema*
  (list :kind "local-lexical"
        :external_vector_db_allowed nil
        :cloud_rag_allowed nil
        :index_roots #("lisp/ssot"
                       "generated"
                       "cmd"
                       "internal"
                       "apps/flutter"
                       "docs"
                       "fixtures"
                       "harness/golden"
                       "data/private/linear-offline-queue.jsonl")
        :query_types #("concept definition location"
                       "bounded context owner"
                       "related implementation files"
                       "related tests and generated files"
                       "related Linear issue"
                       "semantic duplication"
                       "must-read files before change")
        :evidence_fields #("canonical_name"
                           "bounded_context"
                           "owner"
                           "matched_terms"
                           "paths"
                           "linear_issues"
                           "duplicate_suspicions"
                           "must_read_files")))
