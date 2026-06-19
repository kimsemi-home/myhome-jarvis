(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry-agent-observation*
  (vector
   (ddd-concept "LearningLedger" "AgentCluster" "Repository"
                "A private append-only observation ledger that turns loop gaps, evidence debt, and verification misses into tracked self-improvement evidence."
                #("learning ledger" "loop gap ledger" "observation ledger" "evidence debt journal")
                "internal/learning"
                #("generated/learning.generated.json" "internal/learning/ledger.go" "docs/learning-ledger.md")
                #("AgentClusterPolicy" "ClosedLoopPlanner" "KnowledgeIndex" "SecurityPolicy"))
   (ddd-concept "EvidenceGraph" "AgentCluster" "Repository"
                "A private local graph summary that connects observations to evidence artifacts so loop gaps can be traced without exposing raw private evidence."
                #("evidence graph" "evidence board" "evidence links" "traceable evidence")
                "internal/evidence"
                #("generated/evidence.generated.json" "internal/evidence/status.go" "docs/evidence-graph.md")
                #("LearningLedger" "AgentClusterPolicy" "ClosedLoopPlanner" "SecurityPolicy"))
   (ddd-concept "ConfidenceAssessor" "AgentCluster" "Policy"
                "An external evidence-based assessor that returns a confidence cap instead of accepting agent self-reported confidence."
                #("confidence assessor" "confidence cap" "external confidence" "confidence gate")
                "internal/confidence"
                #("generated/confidence.generated.json" "internal/confidence/status.go" "docs/confidence-assessor.md")
                #("EvidenceGraph" "LearningLedger" "AgentClusterPolicy" "SecurityPolicy"))
   (ddd-concept "TranslationManifest" "AgentCluster" "ValueObject"
                "A private context-translation manifest and loss ledger summary that keeps semantic movement and meaning loss traceable without exposing raw private notes."
                #("translation manifest" "loss ledger" "semantic loss" "translation debt" "context map")
                "internal/translation"
                #("generated/translation.generated.json" "internal/translation/status.go" "docs/translation-manifest.md")
                #("ConceptRegistry" "EvidenceGraph" "LearningLedger" "AgentClusterPolicy" "SecurityPolicy"))
   (ddd-concept "ControlPlaneManifest" "AgentOps" "ValueObject"
                "A private orchestration decision receipt that records local closed-loop routing policy, authority, lease, verifier separation, evidence inputs, and output refs without exposing raw rationale."
                #("control plane manifest" "orchestration manifest" "routing receipt" "control-plane verifier" "lease manifest")
                "internal/controlplane"
                #("generated/control_plane.generated.json" "internal/controlplane/status.go" "docs/control-plane-manifest.md")
                #("ClosedLoopPlanner" "CheckpointRecorded" "EvidenceGraph" "AgentClusterPolicy" "SecurityPolicy"))))
