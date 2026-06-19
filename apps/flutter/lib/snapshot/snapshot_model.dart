part of '../snapshot.dart';

@immutable
class JarvisSnapshot {
  const JarvisSnapshot({
    required this.metrics,
    required this.commands,
    required this.linearItems,
    required this.storageItems,
    required this.recommendationItems,
    required this.recommendations,
    required this.householdScopes,
    required this.financeDashboard,
    required this.purchaseDashboard,
    required this.connectors,
    required this.agentClusterSignals,
  });

  final List<SystemMetric> metrics;
  final List<HomeCommand> commands;
  final List<String> linearItems;
  final List<String> storageItems;
  final List<String> recommendationItems;
  final List<RecommendationInsight> recommendations;
  final List<HouseholdScope> householdScopes;
  final FinanceDashboard financeDashboard;
  final PurchaseDashboard purchaseDashboard;
  final List<ConnectorReadiness> connectors;
  final List<AgentClusterSignal> agentClusterSignals;

  static const sample = _sampleJarvisSnapshot;

  factory JarvisSnapshot.offlineFallback() => _offlineFallbackSnapshot();
}
