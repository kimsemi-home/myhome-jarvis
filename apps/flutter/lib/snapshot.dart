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
class JarvisSnapshot {
  const JarvisSnapshot({
    required this.metrics,
    required this.commands,
    required this.linearItems,
    required this.storageItems,
    required this.recommendationItems,
    required this.householdScopes,
  });

  final List<SystemMetric> metrics;
  final List<HomeCommand> commands;
  final List<String> linearItems;
  final List<String> storageItems;
  final List<String> recommendationItems;
  final List<HouseholdScope> householdScopes;

  static const sample = JarvisSnapshot(
    metrics: [
      SystemMetric(
        label: 'Daemon',
        value: '127.0.0.1:3888',
        icon: Icons.settings_ethernet,
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
        label: 'Linear',
        value: 'Online-ready',
        icon: Icons.hub_outlined,
      ),
    ],
    commands: [
      HomeCommand(
        name: 'open-youtube',
        payload: '{}',
        icon: Icons.play_circle_outline,
      ),
      HomeCommand(
        name: 'volume-set',
        payload: '{"level":30}',
        icon: Icons.volume_up_outlined,
        payloadFields: ['level'],
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
      ],
      commands: sample.commands,
      linearItems: const ['Offline queue', 'Local fallback'],
      storageItems: sample.storageItems,
      recommendationItems: sample.recommendationItems,
      householdScopes: sample.householdScopes,
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
