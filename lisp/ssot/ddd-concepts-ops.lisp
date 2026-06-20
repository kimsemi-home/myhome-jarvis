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
   (ddd-concept "CodexCostGovernor" "AgentOps" "Policy"
                "A private usage ledger and public-safe budget status for Codex tokens, coins, actions minutes, and external tool costs."
                #("codex cost" "cost governor" "codex coin" "usage budget")
                "internal/codexcost"
                #("generated/codex_cost.generated.json" "internal/codexcost/status.go" "docs/codex-cost-governor.md")
                #("ClosedLoopPlanner" "SecurityPolicy" "EvidenceGraph"))
   (ddd-concept "PublicSafeRepoFactory" "AgentOps" "Policy"
                "The SSOT template and authority-gated bootstrap contract for creating public Shorts Factory repos and Codex projects."
                #("repo factory" "shorts repo" "codex project bootstrap" "public repo template")
                "internal/repofactory"
                #("generated/repo_factory.generated.json" "internal/repofactory/status.go" "docs/repo-factory.md")
                #("SecurityPolicy" "AuthorityGate" "EvidenceGraph" "CodexCostGovernor"))
   (ddd-concept "MergeEvidencePolicy" "AgentOps" "Policy"
                "The public-safe rule that eligible PRs are merged and completion evidence records PR, commit, checks, and Linear status."
                #("merge evidence" "merge when eligible" "main quality evidence")
                "internal/mergeevidence"
                #("generated/merge_evidence.generated.json" "internal/mergeevidence/status.go" "docs/merge-evidence.md")
                #("LinearWorkQueue" "SecurityPolicy" "EvidenceGraph"))
   (ddd-concept "ConceptRegistry" "KnowledgeIndex" "Repository"
                "The SSOT-owned registry of canonical concepts, aliases, owners, and generated targets."
                #("concept registry" "canonical concepts" "aliases")
                "lisp/ssot/ddd.lisp"
                #("generated/concepts.generated.json" "lisp/ssot/ddd.lisp")
                #("KnowledgeIndex" "ClosedLoopPlanner"))))
