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
   (ddd-concept "CommercePurchase" "CommerceIntelligence" "Entity"
                "A fixture-backed purchase IR item used for commerce intelligence and recurring candidates."
                #("purchase" "commerce purchase" "purchase ir")
                "lisp/ssot/commerce.lisp"
                #("generated/commerce.generated.json" "fixtures/commerce_purchases.jsonl")
                #("StorageLake" "SecurityPolicy"))
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
