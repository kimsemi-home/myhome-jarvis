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
    (:name "ConnectorReadiness"
     :owner "connectors"
     :description "Public-safe planned connector catalog and read-only readiness status.")
    (:name "AgentCluster"
     :owner "agent-cluster"
     :description "Evidence-first multi-agent learning loop policy, authority gates, and verification sidecars.")
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
     :ddd_kind "ValueObject"
     :description "A local-first home-control intent with a validated dry-run execution plan."
     :allowed_aliases #("command" "home command" "intent" "home action")
     :owner "internal/commands"
     :generated_targets #("generated/commands.generated.json"
                          "internal/commands/registry.go"
                          "harness/golden/home_control.golden.json")
     :related_concepts #("SecurityPolicy" "ClosedLoopPlanner"))
    (:canonical_name "HouseholdTransaction"
     :bounded_context "HouseholdFinance"
     :ddd_kind "Entity"
     :description "A fixture-backed household finance transaction IR item without raw bank data."
     :allowed_aliases #("transaction" "finance transaction" "transaction ir")
     :owner "lisp/ssot/finance.lisp"
     :generated_targets #("generated/finance.generated.json"
                          "fixtures/finance_transactions.jsonl")
     :related_concepts #("StorageLake" "SecurityPolicy"))
    (:canonical_name "CommercePurchase"
     :bounded_context "CommerceIntelligence"
     :ddd_kind "Entity"
     :description "A fixture-backed purchase IR item used for commerce intelligence and recurring candidates."
     :allowed_aliases #("purchase" "commerce purchase" "purchase ir")
     :owner "lisp/ssot/commerce.lisp"
     :generated_targets #("generated/commerce.generated.json"
                          "fixtures/commerce_purchases.jsonl")
     :related_concepts #("StorageLake" "SecurityPolicy"))
    (:canonical_name "ConnectorCatalog"
     :bounded_context "ConnectorReadiness"
     :ddd_kind "Policy"
     :description "A public-safe catalog of planned read-only connectors that stays fixture-only until explicit connector work begins."
     :allowed_aliases #("connector" "connector catalog" "connector readiness" "connector status")
     :owner "lisp/ssot/connectors.lisp"
     :generated_targets #("generated/connectors.generated.json"
                          "internal/connectors/status.go"
                          "docs/connectors.md")
     :related_concepts #("HouseholdTransaction" "CommercePurchase" "SecurityPolicy"))
    (:canonical_name "AgentClusterPolicy"
     :bounded_context "AgentCluster"
     :ddd_kind "Policy"
     :description "A public-safe evidence-first learning loop policy for agent roles, sidecars, confidence gates, leases, quarantine, and knowledge updates."
     :allowed_aliases #("agent cluster" "learning loop" "evidence-first workflow" "governed compiler farm" "authority gate")
     :owner "lisp/ssot/agent-cluster.lisp"
     :generated_targets #("generated/agent_cluster.generated.json"
                          "internal/agentcluster/status.go"
                          "docs/agent-cluster.md")
     :related_concepts #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy" "ConceptRegistry"))
    (:canonical_name "StorageLake"
     :bounded_context "StorageLake"
     :ddd_kind "Aggregate"
     :description "The local-only storage policy, lake layers, and generated storage contract."
     :allowed_aliases #("lake" "local lake" "storage policy")
     :owner "lisp/ssot/storage.lisp"
     :generated_targets #("generated/storage.generated.json"
                          "docs/storage.md")
     :related_concepts #("HouseholdTransaction" "CommercePurchase" "SecurityPolicy"))
    (:canonical_name "SecurityPolicy"
     :bounded_context "SecurityPolicy"
     :ddd_kind "Policy"
     :description "The public-repository boundary for secrets, private markers, auth, and allowed languages."
     :allowed_aliases #("security" "public safety" "secret scan" "allowed languages")
     :owner "lisp/ssot/security.lisp"
     :generated_targets #("generated/security.generated.json"
                          "internal/security/security.go"
                          "docs/security.md")
     :related_concepts #("HomeCommand" "StorageLake" "LinearWorkQueue"))
    (:canonical_name "LinearWorkQueue"
     :bounded_context "AgentOps"
     :ddd_kind "Port"
     :description "The Linear or offline queue work source used by the closed loop."
     :allowed_aliases #("linear" "work queue" "offline queue")
     :owner "internal/linear"
     :generated_targets #("generated/linear.generated.json"
                          "docs/linear-workflow.md")
     :related_concepts #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))
    (:canonical_name "ClosedLoopPlanner"
     :bounded_context "AgentOps"
     :ddd_kind "Aggregate"
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
     :ddd_kind "Repository"
     :description "The SSOT-owned registry of canonical concepts, aliases, owners, and generated targets."
     :allowed_aliases #("concept registry" "canonical concepts" "aliases")
     :owner "lisp/ssot/ddd.lisp"
     :generated_targets #("generated/concepts.generated.json"
                          "lisp/ssot/ddd.lisp")
     :related_concepts #("KnowledgeIndex" "ClosedLoopPlanner"))
    (:canonical_name "KnowledgeIndex"
     :bounded_context "KnowledgeIndex"
     :ddd_kind "Repository"
     :description "A local lexical index over SSOT, generated artifacts, source, docs, fixtures, harness, backlog, working log, and private offline Linear records."
     :allowed_aliases #("knowledge index" "local index" "lexical index" "knowledge search")
     :owner "internal/knowledge"
     :generated_targets #("generated/concepts.generated.json"
                          "internal/knowledge/index.go"
                          "docs/knowledge-index.md")
     :related_concepts #("ConceptRegistry" "ClosedLoopPlanner" "LinearWorkQueue"))
    (:canonical_name "LinearGraphQLAdapter"
     :bounded_context "AgentOps"
     :ddd_kind "Adapter"
     :description "The direct Go GraphQL adapter that talks to Linear without a Node or TypeScript SDK."
     :allowed_aliases #("linear graphql adapter" "graphql client")
     :owner "internal/linear"
     :generated_targets #("internal/linear/status.go"
                          "docs/linear-workflow.md")
     :related_concepts #("LinearWorkQueue" "SecurityPolicy"))
    (:canonical_name "LinearOfflineFallback"
     :bounded_context "AgentOps"
     :ddd_kind "AntiCorruptionLayer"
     :description "The offline action boundary that prevents failed external Linear sync from being reported as success."
     :allowed_aliases #("offline fallback" "linear offline fallback" "offline action")
     :owner "internal/linear"
     :generated_targets #("internal/linear/status.go"
                          "docs/linear-workflow.md")
     :related_concepts #("LinearWorkQueue" "ClosedLoopPlanner"))
    (:canonical_name "CheckpointRecorded"
     :bounded_context "AgentOps"
     :ddd_kind "DomainEvent"
     :description "A private closed-loop checkpoint event containing redacted Linear, planner, KnowledgeIndex, and public-safety evidence."
     :allowed_aliases #("checkpoint recorded" "checkpoint event" "loop checkpoint event")
     :owner "internal/orchestrator"
     :generated_targets #("internal/orchestrator/checkpoint.go"
                          "docs/closed-loop.md")
     :related_concepts #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))))

(defparameter *domain-events*
  #((:name "CheckpointRecorded"
     :bounded_context "AgentOps"
     :description "Emitted when a closed-loop cycle writes private checkpoint evidence."
     :emitted_by "ClosedLoopPlanner"
     :payload_fields #("linear_status"
                       "planner_status"
                       "knowledge_evidence"
                       "security_status"))
    (:name "KnowledgeLookupRecorded"
     :bounded_context "KnowledgeIndex"
     :description "Recorded in planner status and checkpoints when a pre-plan KnowledgeIndex lookup runs."
     :emitted_by "KnowledgeIndex"
     :payload_fields #("query"
                       "concept_count"
                       "hit_count"
                       "linear_issues"
                       "must_read"))))

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

(defparameter *generated-artifact-contracts*
  #((:name "concepts" :path "generated/concepts.generated.json" :owner "KnowledgeIndex")
    (:name "commands" :path "generated/commands.generated.json" :owner "HomeControl")
    (:name "finance" :path "generated/finance.generated.json" :owner "HouseholdFinance")
    (:name "commerce" :path "generated/commerce.generated.json" :owner "CommerceIntelligence")
    (:name "connectors" :path "generated/connectors.generated.json" :owner "ConnectorReadiness")
    (:name "agent_cluster" :path "generated/agent_cluster.generated.json" :owner "AgentCluster")
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
