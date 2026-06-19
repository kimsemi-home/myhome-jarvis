part of '../snapshot.dart';

@immutable
class ConnectorReadiness {
  const ConnectorReadiness({
    required this.key,
    required this.label,
    required this.category,
    required this.status,
    required this.fixtureMode,
    required this.dataClasses,
    required this.allowedOperations,
    required this.forbiddenOperations,
    required this.nextStep,
  });

  final String key;
  final String label;
  final String category;
  final String status;
  final bool fixtureMode;
  final List<String> dataClasses;
  final List<String> allowedOperations;
  final List<String> forbiddenOperations;
  final String nextStep;
}

@immutable
class AgentClusterSignal {
  const AgentClusterSignal({
    required this.key,
    required this.label,
    required this.status,
    required this.evidence,
  });

  final String key;
  final String label;
  final String status;
  final String evidence;
}
