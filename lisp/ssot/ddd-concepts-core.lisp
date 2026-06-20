(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry-core*
  (vector
   (ddd-concept "HomeCommand" "HomeControl" "ValueObject"
                "A local-first home-control intent with a validated dry-run execution plan."
                #("command" "home command" "intent" "home action")
                "internal/commands"
                #("generated/commands.generated.json" "internal/commands/registry.go" "harness/golden/home_control.golden.json")
                #("SecurityPolicy" "ClosedLoopPlanner"))
   (ddd-concept "HouseholdTransaction" "HouseholdFinance" "Entity"
                "A fixture-backed household finance transaction IR item without raw bank data."
                #("transaction" "finance transaction" "transaction ir")
                "lisp/ssot/finance.lisp"
                #("generated/finance.generated.json" "fixtures/finance_transactions.jsonl")
                #("StorageLake" "SecurityPolicy"))
   (ddd-concept "HouseholdFinanceConsent" "HouseholdFinance" "Policy"
                "A private consent ledger and public-safe readiness gate for real finance connectors and shared household scopes."
                #("finance consent" "spouse consent" "household consent" "consent ledger")
                "internal/financeconsent"
                #("generated/finance_consent.generated.json" "internal/financeconsent/status.go" "docs/finance-consent.md")
                #("HouseholdTransaction" "AuthorityGate" "SecurityPolicy"))
   (ddd-concept "CommercePurchase" "CommerceIntelligence" "Entity"
                "A fixture-backed purchase IR item used for commerce intelligence and recurring candidates."
                #("purchase" "commerce purchase" "purchase ir")
                "lisp/ssot/commerce.lisp"
                #("generated/commerce.generated.json" "fixtures/commerce_purchases.jsonl")
                #("StorageLake" "SecurityPolicy"))
   (ddd-concept "MonetizationExperiment" "CommerceIntelligence" "Repository"
                "A private experiment ledger for revenue hypotheses, review gates, evidence, expected value bands, and cost estimates."
                #("monetization" "revenue experiment" "experiment ledger" "revenue hypothesis")
                "internal/monetization"
                #("generated/monetization.generated.json" "internal/monetization/status.go" "docs/monetization-experiments.md")
                #("CommercePurchase" "CodexCostGovernor" "AuthorityGate" "EvidenceGraph"))
   (ddd-concept "CodexSustainabilityLoop" "AgentOps" "Policy"
                "A private evidence loop that measures Codex sustainability, trend posture, cycle speed, cache value, and proposal proof."
                #("codex sustainability" "trend evidence" "usage sustainability" "optimization evidence")
                "internal/codexsustainability"
                #("generated/codex_sustainability.generated.json" "internal/codexsustainability/status.go" "docs/codex-sustainability.md")
                #("CodexCostGovernor" "ClosedLoopPlanner" "MonetizationExperiment" "EvidenceGraph"))
   (ddd-concept "CrossRepoContextPack" "AgentOps" "Policy"
                "A public-safe context, ontology, SSOT, authority, security, and version handoff contract for downstream repositories."
                #("context pack" "cross repo context" "ontology handoff" "version handoff")
                "internal/contextpack"
                #("generated/context_pack.generated.json" "internal/contextpack/status.go" "docs/context-pack.md")
                #("PublicSafeRepoFactory" "ConceptRegistry" "AuthorityGate" "SecurityPolicy"))
   (ddd-concept "ConnectorCatalog" "ConnectorReadiness" "Policy"
                "A public-safe catalog of planned read-only connectors that stays fixture-only until explicit connector work begins."
                #("connector" "connector catalog" "connector readiness" "connector status")
                "lisp/ssot/connectors.lisp"
                #("generated/connectors.generated.json" "internal/connectors/status.go" "docs/connectors.md")
                #("HouseholdTransaction" "CommercePurchase" "SecurityPolicy"))
   (ddd-concept "AgentClusterPolicy" "AgentCluster" "Policy"
                "A public-safe evidence-first learning loop policy for agent roles, sidecars, confidence gates, leases, quarantine, and knowledge updates."
                #("agent cluster" "learning loop" "evidence-first workflow" "governed compiler farm" "agent governance")
                "lisp/ssot/agent-cluster.lisp"
                #("generated/agent_cluster.generated.json" "internal/agentcluster/status.go" "docs/agent-cluster.md")
                #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy" "ConceptRegistry"))))
