part of '../main.dart';

enum HouseholdScopeState {
  selected,
  ownerScoped,
  householdScoped,
  summaryOnly,
  empty,
}

extension HouseholdScopeStateLabel on HouseholdScopeState {
  String get label => switch (this) {
    HouseholdScopeState.selected => 'selected scope',
    HouseholdScopeState.ownerScoped => 'owner scoped',
    HouseholdScopeState.householdScoped => 'household scoped',
    HouseholdScopeState.summaryOnly => 'summary-only',
    HouseholdScopeState.empty => 'empty',
  };

  JarvisBadgeTone get tone => switch (this) {
    HouseholdScopeState.selected => JarvisBadgeTone.success,
    HouseholdScopeState.ownerScoped => JarvisBadgeTone.success,
    HouseholdScopeState.householdScoped => JarvisBadgeTone.success,
    HouseholdScopeState.summaryOnly => JarvisBadgeTone.muted,
    HouseholdScopeState.empty => JarvisBadgeTone.warning,
  };
}

extension HouseholdScopeStateRules on HouseholdScope {
  bool get isEmptyScope => financeRecords == 0 && purchaseRecords == 0;

  List<HouseholdScopeState> get stateBadges => [
    HouseholdScopeState.selected,
    if (scope == 'household')
      HouseholdScopeState.householdScoped
    else
      HouseholdScopeState.ownerScoped,
    HouseholdScopeState.summaryOnly,
    if (isEmptyScope) HouseholdScopeState.empty,
  ];

  Color get stateColor {
    if (isEmptyScope) {
      return JarvisAstryxTokens.warning;
    }
    return scope == 'household'
        ? JarvisAstryxTokens.success
        : JarvisAstryxTokens.accent;
  }
}

class HouseholdScopeStateBadges extends StatelessWidget {
  const HouseholdScopeStateBadges({super.key, required this.scope});

  final HouseholdScope scope;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [
        for (final state in scope.stateBadges)
          JarvisBadge(state.label, tone: state.tone),
      ],
    );
  }
}
