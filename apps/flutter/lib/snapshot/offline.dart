part of '../snapshot.dart';

JarvisSnapshot _offlineFallbackSnapshot() {
  return JarvisSnapshot(
    metrics: const [..._offlineCoreMetrics, ..._offlineGovernanceMetrics],
    commands: JarvisSnapshot.sample.commands,
    linearItems: const ['Offline queue', 'Local fallback'],
    storageItems: JarvisSnapshot.sample.storageItems,
    recommendationItems: JarvisSnapshot.sample.recommendationItems,
    recommendations: JarvisSnapshot.sample.recommendations,
    householdScopes: JarvisSnapshot.sample.householdScopes,
    financeDashboard: JarvisSnapshot.sample.financeDashboard,
    purchaseDashboard: JarvisSnapshot.sample.purchaseDashboard,
    connectors: JarvisSnapshot.sample.connectors,
    agentClusterSignals: JarvisSnapshot.sample.agentClusterSignals,
  );
}
