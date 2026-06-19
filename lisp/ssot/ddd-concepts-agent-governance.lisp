(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry-agent-governance*
  (vector
   (ddd-concept "IncidentLifecycle" "AgentCluster" "Repository"
                "A private incident lifecycle ledger that classifies observed failures, assigns owner roles, tracks quarantine state, and exposes only redacted incident debt."
                #("incident lifecycle" "incident ledger" "quarantine report" "incident debt" "feedback loop incident")
                "internal/incidents"
                #("generated/incidents.generated.json" "internal/incidents/status.go" "docs/incident-lifecycle.md")
                #("LearningLedger" "EvidenceGraph" "ConfidenceAssessor" "AgentClusterPolicy" "SecurityPolicy"))
   (ddd-concept "EvidenceQualityAssessor" "AgentCluster" "Policy"
                "A private evidence quality snapshot assessor that tracks staleness, mapping confidence, and reassessment debt without exposing raw evidence."
                #("evidence quality" "evidence quality assessor" "quality snapshot" "mapping confidence" "reassessment debt")
                "internal/evidencequality"
                #("generated/evidence_quality.generated.json" "internal/evidencequality/status.go" "docs/evidence-quality.md")
                #("EvidenceGraph" "ConfidenceAssessor" "IncidentLifecycle" "TranslationManifest" "AgentClusterPolicy" "SecurityPolicy"))
   (ddd-concept "AuthorityGate" "AgentCluster" "Policy"
                "A public-safe Reasoning RBAC and Domain ABAC gate that limits authority based on evidence, confidence, quality, incidents, control-plane state, translation debt, and public safety."
                #("authority gate" "authority status gate" "reasoning rbac" "domain abac" "permission gate" "automation authority")
                "internal/authority"
                #("generated/authority.generated.json" "internal/authority/status.go" "docs/authority-gate.md")
                #("AgentClusterPolicy" "ConfidenceAssessor" "EvidenceQualityAssessor" "IncidentLifecycle" "ControlPlaneManifest" "TranslationManifest" "HumanReviewCapacity" "SecurityPolicy"))
   (ddd-concept "HumanReviewCapacity" "AgentCluster" "Policy"
                "A private human review queue and capacity status that treats review as finite, risk-prioritized, and public-safe."
                #("human review capacity" "review capacity" "review queue" "backup steward" "review debt")
                "internal/review"
                #("generated/review.generated.json" "internal/review/status.go" "docs/human-review-capacity.md")
                #("AuthorityGate" "IncidentLifecycle" "EvidenceQualityAssessor" "ControlPlaneManifest" "AgentClusterPolicy" "SecurityPolicy"))
   (ddd-concept "CodeShapeBudget" "AgentCluster" "Policy"
                "A generated public-safe line budget that blocks new oversized files while tracking existing legacy refactor debt."
                #("code shape budget" "line budget" "75 line budget" "legacy code debt" "file size guard")
                "internal/codeshape"
                #("generated/code_shape.generated.json" "internal/codeshape/status.go" "docs/code-shape-budget.md")
                #("AgentClusterPolicy" "HumanReviewCapacity" "KnowledgeIndex" "SecurityPolicy"))))
