part of '../snapshot.dart';

extension JarvisSnapshotCopy on JarvisSnapshot {
  JarvisSnapshot copyWithFinance(FinanceDashboard financeDashboard) {
    return JarvisSnapshot(
      metrics: metrics,
      commands: commands,
      linearItems: linearItems,
      storageItems: storageItems,
      recommendationItems: recommendationItems,
      recommendations: recommendations,
      householdScopes: householdScopes,
      financeDashboard: financeDashboard,
      purchaseDashboard: purchaseDashboard,
      connectors: connectors,
      agentClusterSignals: agentClusterSignals,
    );
  }
}
