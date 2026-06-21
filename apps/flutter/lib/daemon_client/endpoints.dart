part of '../daemon_client.dart';

const _commandsPath = '/commands';

const _objectEndpointPaths = <String, String>{
  'health': '/health',
  'assistant': '/assistant/status',
  'auth': '/auth/status',
  'linear': '/linear/status',
  'repo': '/repo/status',
  'security': '/security/status',
  'codeShape': '/code-shape/status',
  'domain': '/domain/summary',
  'connectors': '/connectors/status',
  'agentCluster': '/agent-cluster/status',
  'learning': '/learning/status',
  'evidence': '/evidence/status',
  'confidence': '/confidence/status',
  'translation': '/translation/status',
  'controlPlane': '/control-plane/status',
  'incidents': '/incidents/status',
  'evidenceQuality': '/evidence-quality/status',
  'review': '/review/status',
  'authority': '/authority/status',
  'authorityReviewDecision': '/authority-review/decision-packet',
  'metrics': '/metrics',
  'events': '/events',
  'supervisor': '/supervisor/status',
  'audit': '/audit/status',
  'quality': '/quality/status',
  'planner': '/planner/status',
};
