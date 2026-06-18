import 'package:flutter/material.dart';

@immutable
class SystemMetric {
  const SystemMetric({
    required this.label,
    required this.value,
    required this.icon,
  });

  final String label;
  final String value;
  final IconData icon;
}

@immutable
class HomeCommand {
  const HomeCommand({
    required this.name,
    required this.payload,
    required this.icon,
    this.payloadFields = const [],
  });

  final String name;
  final String payload;
  final IconData icon;
  final List<String> payloadFields;

  HomeCommand copyWith({String? payload}) {
    return HomeCommand(
      name: name,
      payload: payload ?? this.payload,
      icon: icon,
      payloadFields: payloadFields,
    );
  }
}

@immutable
class CommandInvocation {
  const CommandInvocation({required this.label, required this.argv, this.url});

  final String label;
  final List<String> argv;
  final String? url;
}

@immutable
class CommandPlan {
  const CommandPlan({
    required this.name,
    required this.dryRun,
    required this.executeAllowed,
    required this.invocations,
    required this.warnings,
  });

  final String name;
  final bool dryRun;
  final bool executeAllowed;
  final List<CommandInvocation> invocations;
  final List<String> warnings;

  factory CommandPlan.preview(HomeCommand command) {
    return CommandPlan(
      name: command.name,
      dryRun: true,
      executeAllowed: false,
      invocations: [
        CommandInvocation(
          label: command.name,
          argv: ['mhj', 'command', command.name, command.payload],
        ),
      ],
      warnings: const [],
    );
  }
}

@immutable
class HouseholdScope {
  const HouseholdScope({
    required this.scope,
    required this.label,
    required this.currency,
    required this.financeRecords,
    required this.financeNetMinorUnits,
    required this.purchaseRecords,
    required this.purchaseSpendMinorUnits,
  });

  final String scope;
  final String label;
  final String currency;
  final int financeRecords;
  final int financeNetMinorUnits;
  final int purchaseRecords;
  final int purchaseSpendMinorUnits;
}

@immutable
class FinanceOwner {
  const FinanceOwner({
    required this.owner,
    required this.records,
    required this.currency,
    required this.creditMinorUnits,
    required this.debitMinorUnits,
    required this.netMinorUnits,
  });

  final String owner;
  final int records;
  final String currency;
  final int creditMinorUnits;
  final int debitMinorUnits;
  final int netMinorUnits;
}

@immutable
class FinanceDashboard {
  const FinanceDashboard({
    required this.records,
    required this.currency,
    required this.creditMinorUnits,
    required this.debitMinorUnits,
    required this.netMinorUnits,
    required this.subscriptionMinorUnits,
    required this.subscriptionCount,
    required this.cardDebitMinorUnits,
    required this.cardDebitCount,
    required this.categories,
    required this.owners,
  });

  final int records;
  final String currency;
  final int creditMinorUnits;
  final int debitMinorUnits;
  final int netMinorUnits;
  final int subscriptionMinorUnits;
  final int subscriptionCount;
  final int cardDebitMinorUnits;
  final int cardDebitCount;
  final List<String> categories;
  final List<FinanceOwner> owners;
}

@immutable
class PurchaseOwner {
  const PurchaseOwner({
    required this.owner,
    required this.records,
    required this.currency,
    required this.purchaseSpendMinorUnits,
  });

  final String owner;
  final int records;
  final String currency;
  final int purchaseSpendMinorUnits;
}

@immutable
class RecurringPurchase {
  const RecurringPurchase({
    required this.merchantName,
    required this.itemName,
    required this.currency,
    required this.purchaseCount,
    required this.latestTotalMinorUnits,
    required this.latestPurchasedAt,
  });

  final String merchantName;
  final String itemName;
  final String currency;
  final int purchaseCount;
  final int latestTotalMinorUnits;
  final String latestPurchasedAt;
}

@immutable
class PurchaseDashboard {
  const PurchaseDashboard({
    required this.records,
    required this.currency,
    required this.totalSpendMinorUnits,
    required this.recurringCandidateCount,
    required this.recurringCandidates,
    required this.categories,
    required this.owners,
  });

  final int records;
  final String currency;
  final int totalSpendMinorUnits;
  final int recurringCandidateCount;
  final List<RecurringPurchase> recurringCandidates;
  final List<String> categories;
  final List<PurchaseOwner> owners;
}

@immutable
class RecommendationInsight {
  const RecommendationInsight({
    required this.kind,
    required this.title,
    required this.rationale,
    required this.score,
    required this.currency,
    required this.estimatedMonthlyMinorUnits,
    required this.evidenceCount,
  });

  final String kind;
  final String title;
  final String rationale;
  final int score;
  final String currency;
  final int estimatedMonthlyMinorUnits;
  final int evidenceCount;
}

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

  static const sample = JarvisSnapshot(
    metrics: [
      SystemMetric(
        label: 'Daemon',
        value: '127.0.0.1:3888',
        icon: Icons.settings_ethernet,
      ),
      SystemMetric(
        label: 'Network',
        value: 'Local-only',
        icon: Icons.wifi_off_outlined,
      ),
      SystemMetric(
        label: 'LAN Auth',
        value: 'Configured',
        icon: Icons.vpn_key_outlined,
      ),
      SystemMetric(
        label: 'Mode',
        value: 'Dry-run',
        icon: Icons.shield_outlined,
      ),
      SystemMetric(
        label: 'Quality',
        value: 'Passing',
        icon: Icons.verified_outlined,
      ),
      SystemMetric(
        label: 'Public Safety',
        value: 'Clear',
        icon: Icons.verified_user_outlined,
      ),
      SystemMetric(
        label: 'Linear',
        value: 'Online-ready',
        icon: Icons.hub_outlined,
      ),
      SystemMetric(
        label: 'Agent Cluster',
        value: '5 roles gated',
        icon: Icons.account_tree_outlined,
      ),
      SystemMetric(
        label: 'Learning',
        value: '0 observed',
        icon: Icons.psychology_alt_outlined,
      ),
      SystemMetric(
        label: 'Evidence Graph',
        value: '0 nodes',
        icon: Icons.device_hub_outlined,
      ),
    ],
    commands: [
      HomeCommand(
        name: 'open-youtube',
        payload: '{}',
        icon: Icons.play_circle_outline,
      ),
      HomeCommand(
        name: 'open-netflix',
        payload: '{}',
        icon: Icons.movie_filter_outlined,
      ),
      HomeCommand(
        name: 'open-disney-plus',
        payload: '{}',
        icon: Icons.auto_awesome_outlined,
      ),
      HomeCommand(
        name: 'open-tving',
        payload: '{}',
        icon: Icons.live_tv_outlined,
      ),
      HomeCommand(
        name: 'open-wavve',
        payload: '{}',
        icon: Icons.waves_outlined,
      ),
      HomeCommand(
        name: 'open-coupang-play',
        payload: '{}',
        icon: Icons.play_circle_outline,
      ),
      HomeCommand(
        name: 'open-youtube-search',
        payload: '{"query":"lofi music"}',
        icon: Icons.play_circle_outline,
        payloadFields: ['query'],
      ),
      HomeCommand(
        name: 'open-url',
        payload: '{"url":"https://www.youtube.com"}',
        icon: Icons.public_outlined,
        payloadFields: ['url'],
      ),
      HomeCommand(
        name: 'open-ott',
        payload: '{"service":"netflix"}',
        icon: Icons.theaters_outlined,
        payloadFields: ['service'],
      ),
      HomeCommand(
        name: 'volume-set',
        payload: '{"level":30}',
        icon: Icons.volume_up_outlined,
        payloadFields: ['level'],
      ),
      HomeCommand(
        name: 'volume-up',
        payload: '{"step":10}',
        icon: Icons.volume_up_outlined,
        payloadFields: ['step'],
      ),
      HomeCommand(
        name: 'volume-down',
        payload: '{"step":10}',
        icon: Icons.volume_down_outlined,
        payloadFields: ['step'],
      ),
      HomeCommand(
        name: 'volume-mute',
        payload: '{}',
        icon: Icons.volume_off_outlined,
      ),
      HomeCommand(
        name: 'display-sleep',
        payload: '{}',
        icon: Icons.monitor_outlined,
      ),
      HomeCommand(
        name: 'mac-sleep',
        payload: '{}',
        icon: Icons.bedtime_outlined,
      ),
      HomeCommand(
        name: 'movie-mode',
        payload: '{}',
        icon: Icons.theaters_outlined,
      ),
      HomeCommand(
        name: 'sleep-mode',
        payload: '{}',
        icon: Icons.bedtime_outlined,
      ),
    ],
    linearItems: ['Active queue', 'Next issue', 'Offline queue'],
    storageItems: [
      'finance_transactions',
      'commerce_purchases',
      'Parquet+Zstd',
    ],
    recommendationItems: [
      '81 - Compare recurring purchase: Bottled water 2L x 6',
      '67 - Review card-linked household spend',
      '61 - Review household subscriptions',
      '49 - Keep household cash buffer',
    ],
    recommendations: [
      RecommendationInsight(
        kind: 'recurring_purchase_review',
        title: 'Compare recurring purchase: Bottled water 2L x 6',
        rationale: 'Coupang appears repeatedly in local purchase fixtures.',
        score: 81,
        currency: 'KRW',
        estimatedMonthlyMinorUnits: 11800,
        evidenceCount: 2,
      ),
      RecommendationInsight(
        kind: 'card_usage_review',
        title: 'Review card-linked household spend',
        rationale:
            'Card-linked debit fixtures exist; keep this as a review-only recommendation, not a card action.',
        score: 67,
        currency: 'KRW',
        estimatedMonthlyMinorUnits: 153200,
        evidenceCount: 2,
      ),
      RecommendationInsight(
        kind: 'subscription_review',
        title: 'Review household subscriptions',
        rationale:
            'Subscription-like debit fixtures exist; keep this as a review-only recommendation.',
        score: 61,
        currency: 'KRW',
        estimatedMonthlyMinorUnits: 65900,
        evidenceCount: 1,
      ),
      RecommendationInsight(
        kind: 'cash_buffer',
        title: 'Keep household cash buffer',
        rationale:
            'Fixture cashflow is positive; reserve surplus before recommendations become executable.',
        score: 49,
        currency: 'KRW',
        estimatedMonthlyMinorUnits: 4346800,
        evidenceCount: 3,
      ),
    ],
    householdScopes: [
      HouseholdScope(
        scope: 'user',
        label: 'User',
        currency: 'KRW',
        financeRecords: 1,
        financeNetMinorUnits: -87300,
        purchaseRecords: 1,
        purchaseSpendMinorUnits: 3200,
      ),
      HouseholdScope(
        scope: 'spouse',
        label: 'Spouse',
        currency: 'KRW',
        financeRecords: 0,
        financeNetMinorUnits: 0,
        purchaseRecords: 0,
        purchaseSpendMinorUnits: 0,
      ),
      HouseholdScope(
        scope: 'household',
        label: 'Household',
        currency: 'KRW',
        financeRecords: 3,
        financeNetMinorUnits: 4346800,
        purchaseRecords: 3,
        purchaseSpendMinorUnits: 26800,
      ),
    ],
    financeDashboard: FinanceDashboard(
      records: 3,
      currency: 'KRW',
      creditMinorUnits: 4500000,
      debitMinorUnits: 153200,
      netMinorUnits: 4346800,
      subscriptionMinorUnits: 65900,
      subscriptionCount: 1,
      cardDebitMinorUnits: 153200,
      cardDebitCount: 2,
      categories: ['income', 'subscription', 'utilities'],
      owners: [
        FinanceOwner(
          owner: 'household',
          records: 2,
          currency: 'KRW',
          creditMinorUnits: 4500000,
          debitMinorUnits: 65900,
          netMinorUnits: 4434100,
        ),
        FinanceOwner(
          owner: 'user',
          records: 1,
          currency: 'KRW',
          creditMinorUnits: 0,
          debitMinorUnits: 87300,
          netMinorUnits: -87300,
        ),
      ],
    ),
    purchaseDashboard: PurchaseDashboard(
      records: 3,
      currency: 'KRW',
      totalSpendMinorUnits: 26800,
      recurringCandidateCount: 1,
      recurringCandidates: [
        RecurringPurchase(
          merchantName: 'Coupang',
          itemName: 'Bottled water 2L x 6',
          currency: 'KRW',
          purchaseCount: 2,
          latestTotalMinorUnits: 11800,
          latestPurchasedAt: '2026-06-10',
        ),
      ],
      categories: ['grocery', 'household'],
      owners: [
        PurchaseOwner(
          owner: 'household',
          records: 2,
          currency: 'KRW',
          purchaseSpendMinorUnits: 23600,
        ),
        PurchaseOwner(
          owner: 'user',
          records: 1,
          currency: 'KRW',
          purchaseSpendMinorUnits: 3200,
        ),
      ],
    ),
    connectors: [
      ConnectorReadiness(
        key: 'mydata',
        label: 'MyData aggregator',
        category: 'finance_aggregation',
        status: 'planned',
        fixtureMode: true,
        dataClasses: ['accounts', 'cards', 'transactions'],
        allowedOperations: ['read_fixture', 'summarize'],
        forbiddenOperations: [
          'credential_request',
          'external_api_call',
          'transfer',
          'trade',
          'card_action',
        ],
        nextStep:
            'Define consent and local vault boundaries before any real connector.',
      ),
      ConnectorReadiness(
        key: 'commerce',
        label: 'Commerce purchases',
        category: 'commerce',
        status: 'planned',
        fixtureMode: true,
        dataClasses: ['orders', 'items', 'recurring_candidates'],
        allowedOperations: ['read_fixture', 'recommend_review', 'summarize'],
        forbiddenOperations: [
          'credential_request',
          'cookie_import',
          'scraping',
          'purchase',
        ],
        nextStep:
            'Extend local purchase fixtures and avoid scraping/cookie capture.',
      ),
    ],
    agentClusterSignals: [
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
        evidence:
            'producer, reviewer, verifier, and steward roles are separated',
      ),
      AgentClusterSignal(
        key: 'feedback_loop',
        label: 'Feedback loop',
        status: 'tracked',
        evidence: 'incidents must end in verification and knowledge update',
      ),
    ],
  );

  factory JarvisSnapshot.offlineFallback() {
    return JarvisSnapshot(
      metrics: const [
        SystemMetric(
          label: 'Daemon',
          value: 'Offline fallback',
          icon: Icons.settings_ethernet,
        ),
        SystemMetric(
          label: 'Network',
          value: 'Local-only',
          icon: Icons.wifi_off_outlined,
        ),
        SystemMetric(
          label: 'LAN Auth',
          value: 'Unavailable',
          icon: Icons.vpn_key_outlined,
        ),
        SystemMetric(
          label: 'Mode',
          value: 'Dry-run',
          icon: Icons.shield_outlined,
        ),
        SystemMetric(
          label: 'Quality',
          value: 'Local',
          icon: Icons.verified_outlined,
        ),
        SystemMetric(
          label: 'Linear',
          value: 'Offline',
          icon: Icons.hub_outlined,
        ),
        SystemMetric(
          label: 'Agent Cluster',
          value: 'Local',
          icon: Icons.account_tree_outlined,
        ),
        SystemMetric(
          label: 'Learning',
          value: 'Local',
          icon: Icons.psychology_alt_outlined,
        ),
        SystemMetric(
          label: 'Evidence Graph',
          value: 'Local',
          icon: Icons.device_hub_outlined,
        ),
      ],
      commands: sample.commands,
      linearItems: const ['Offline queue', 'Local fallback'],
      storageItems: sample.storageItems,
      recommendationItems: sample.recommendationItems,
      recommendations: sample.recommendations,
      householdScopes: sample.householdScopes,
      financeDashboard: sample.financeDashboard,
      purchaseDashboard: sample.purchaseDashboard,
      connectors: sample.connectors,
      agentClusterSignals: sample.agentClusterSignals,
    );
  }
}

abstract interface class JarvisSnapshotClient {
  Future<JarvisSnapshot> load();
}

abstract interface class JarvisCommandClient {
  Future<CommandPlan> dryRun(HomeCommand command);
}

abstract interface class JarvisClient
    implements JarvisSnapshotClient, JarvisCommandClient {}

class StaticSnapshotClient implements JarvisClient {
  const StaticSnapshotClient(this.snapshot);

  final JarvisSnapshot snapshot;

  @override
  Future<JarvisSnapshot> load() async => snapshot;

  @override
  Future<CommandPlan> dryRun(HomeCommand command) async {
    return CommandPlan.preview(command);
  }
}
