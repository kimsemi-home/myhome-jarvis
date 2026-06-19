part of '../snapshot.dart';

const _sampleJarvisSnapshot = JarvisSnapshot(
  metrics: _sampleMetrics,
  commands: _sampleCommands,
  linearItems: ['Active queue', 'Next issue', 'Offline queue'],
  storageItems: ['finance_transactions', 'commerce_purchases', 'Parquet+Zstd'],
  recommendationItems: _sampleRecommendationItems,
  recommendations: _sampleRecommendations,
  householdScopes: _sampleHouseholdScopes,
  financeDashboard: _sampleFinanceDashboard,
  purchaseDashboard: _samplePurchaseDashboard,
  connectors: _sampleConnectors,
  agentClusterSignals: _sampleAgentClusterSignals,
);

const _sampleMetrics = [..._sampleCoreMetrics, ..._sampleGovernanceMetrics];

const _sampleCommands = [
  ..._sampleMediaCommands,
  ..._sampleOpenCommands,
  ..._sampleSystemCommands,
];
