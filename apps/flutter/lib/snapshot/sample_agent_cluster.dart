part of '../snapshot.dart';

const _sampleAgentClusterSignals = [
  AgentClusterSignal(
    key: 'evidence_first',
    label: 'Evidence first',
    status: 'active',
    evidence: 'observation and evidence precede code',
  ),
  AgentClusterSignal(
    key: 'authority_gated',
    label: 'Authority gated',
    status: 'gated',
    evidence: 'producer, reviewer, verifier, and steward roles are separated',
  ),
  AgentClusterSignal(
    key: 'feedback_loop',
    label: 'Feedback loop',
    status: 'tracked',
    evidence: 'incidents must end in verification and knowledge update',
  ),
];
