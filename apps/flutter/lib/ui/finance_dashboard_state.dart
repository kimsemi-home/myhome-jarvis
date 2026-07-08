part of '../main.dart';

enum FinanceDashboardState {
  summaryOnly,
  fixtureOnly,
  verifiedMetadata,
  cardLinkedReview,
  empty,
}

extension FinanceDashboardStateLabel on FinanceDashboardState {
  String get label => switch (this) {
    FinanceDashboardState.summaryOnly => 'summary-only',
    FinanceDashboardState.fixtureOnly => 'fixture-only',
    FinanceDashboardState.verifiedMetadata => 'verified metadata',
    FinanceDashboardState.cardLinkedReview => 'card-linked review',
    FinanceDashboardState.empty => 'empty',
  };

  JarvisBadgeTone get tone => switch (this) {
    FinanceDashboardState.summaryOnly => JarvisBadgeTone.muted,
    FinanceDashboardState.fixtureOnly => JarvisBadgeTone.muted,
    FinanceDashboardState.verifiedMetadata => JarvisBadgeTone.success,
    FinanceDashboardState.cardLinkedReview => JarvisBadgeTone.warning,
    FinanceDashboardState.empty => JarvisBadgeTone.muted,
  };
}

extension FinanceDashboardStateRules on FinanceDashboard {
  List<FinanceDashboardState> get dashboardStates => [
    FinanceDashboardState.summaryOnly,
    if (fixtureOnly) FinanceDashboardState.fixtureOnly,
    if (records == 0)
      FinanceDashboardState.empty
    else
      FinanceDashboardState.verifiedMetadata,
    if (cardDebitCount > 0) FinanceDashboardState.cardLinkedReview,
  ];
}

class FinanceDashboardStateBadges extends StatelessWidget {
  const FinanceDashboardStateBadges({super.key, required this.dashboard});

  final FinanceDashboard dashboard;

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
