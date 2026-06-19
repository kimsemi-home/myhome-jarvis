(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry-ops*
  (vector
   (ddd-concept "StorageLake" "StorageLake" "Aggregate"
                "The local-only storage policy, lake layers, and generated storage contract."
                #("lake" "local lake" "storage policy")
                "lisp/ssot/storage.lisp"
                #("generated/storage.generated.json" "docs/storage.md")
                #("HouseholdTransaction" "CommercePurchase" "SecurityPolicy"))
   (ddd-concept "SecurityPolicy" "SecurityPolicy" "Policy"
                "The public-repository boundary for secrets, private markers, auth, and allowed languages."
                #("security" "public safety" "secret scan" "allowed languages")
                "lisp/ssot/security.lisp"
                #("generated/security.generated.json" "internal/security/status.go" "docs/security.md")
                #("HomeCommand" "StorageLake" "LinearWorkQueue"))
   (ddd-concept "LinearWorkQueue" "AgentOps" "Port"
                "The Linear or offline queue work source used by the closed loop."
                #("linear" "work queue" "offline queue")
                "internal/linear"
                #("generated/linear.generated.json" "docs/linear-workflow.md")
                #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))
   (ddd-concept "ClosedLoopPlanner" "AgentOps" "Aggregate"
                "The observe-plan-change-validate-index-sync-decide loop and checkpoint evidence."
                #("planner" "closed loop" "loop" "checkpoint")
                "internal/planner"
                #("generated/planner.generated.json" "internal/planner/status.go" "internal/orchestrator/checkpoint.go" "docs/closed-loop.md")
                #("KnowledgeIndex" "LinearWorkQueue" "SecurityPolicy"))
   (ddd-concept "ConceptRegistry" "KnowledgeIndex" "Repository"
                "The SSOT-owned registry of canonical concepts, aliases, owners, and generated targets."
                #("concept registry" "canonical concepts" "aliases")
                "lisp/ssot/ddd.lisp"
                #("generated/concepts.generated.json" "lisp/ssot/ddd.lisp")
                #("KnowledgeIndex" "ClosedLoopPlanner"))))
