(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry-ops-adapters*
  (vector
   (ddd-concept "KnowledgeIndex" "KnowledgeIndex" "Repository"
                "A local lexical index over SSOT, generated artifacts, source, docs, fixtures, harness, backlog, working log, and private offline Linear records."
                #("knowledge index" "local index" "lexical index" "knowledge search")
                "internal/knowledge"
                #("generated/concepts.generated.json" "internal/knowledge/index.go" "docs/knowledge-index.md")
                #("ConceptRegistry" "ClosedLoopPlanner" "LinearWorkQueue"))
   (ddd-concept "LinearGraphQLAdapter" "AgentOps" "Adapter"
                "The direct Go GraphQL adapter that talks to Linear without a Node or TypeScript SDK."
                #("linear graphql adapter" "graphql client")
                "internal/linear"
                #("internal/linear/status.go" "docs/linear-workflow.md")
                #("LinearWorkQueue" "SecurityPolicy"))
   (ddd-concept "LinearOfflineFallback" "AgentOps" "AntiCorruptionLayer"
                "The offline action boundary that prevents failed external Linear sync from being reported as success."
                #("offline fallback" "linear offline fallback" "offline action")
                "internal/linear"
                #("internal/linear/status.go" "docs/linear-workflow.md")
                #("LinearWorkQueue" "ClosedLoopPlanner"))
   (ddd-concept "CheckpointRecorded" "AgentOps" "DomainEvent"
                "A private closed-loop checkpoint event containing redacted Linear, planner, KnowledgeIndex, and public-safety evidence."
                #("checkpoint recorded" "checkpoint event" "loop checkpoint event")
                "internal/orchestrator"
                #("internal/orchestrator/checkpoint.go" "docs/closed-loop.md")
                #("ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))))
