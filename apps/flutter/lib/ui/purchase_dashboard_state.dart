part of '../main.dart';

enum PurchaseDashboardState {
  summaryOnly,
  verifiedMetadata,
  recurringCandidate,
  empty,
}

extension PurchaseDashboardStateLabel on PurchaseDashboardState {
  String get label => switch (this) {
    PurchaseDashboardState.summaryOnly => 'summary-only',
    PurchaseDashboardState.verifiedMetadata => 'verified metadata',
    PurchaseDashboardState.recurringCandidate => 'recurring candidate',
    PurchaseDashboardState.empty => 'empty',
  };

  JarvisBadgeTone get tone => switch (this) {
    PurchaseDashboardState.summaryOnly => JarvisBadgeTone.outline,
    PurchaseDashboardState.verifiedMetadata => JarvisBadgeTone.primary,
    PurchaseDashboardState.recurringCandidate => JarvisBadgeTone.secondary,
    PurchaseDashboardState.empty => JarvisBadgeTone.outline,
  };
}

extension PurchaseDashboardStateRules on PurchaseDashboard {
  List<PurchaseDashboardState> get dashboardStates => [
    PurchaseDashboardState.summaryOnly,
    if (records == 0)
      PurchaseDashboardState.empty
    else
      PurchaseDashboardState.verifiedMetadata,
    if (recurringCandidateCount > 0) PurchaseDashboardState.recurringCandidate,
  ];
}

class PurchaseDashboardStateBadges extends StatelessWidget {
  const PurchaseDashboardStateBadges({super.key, required this.dashboard});

  final PurchaseDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [
        for (final state in dashboard.dashboardStates)
          JarvisBadge(state.label, tone: state.tone),
      ],
    );
  }
}
