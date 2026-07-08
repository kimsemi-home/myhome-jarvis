import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('maps finance dashboard states for agents', () {
    final states = JarvisSnapshot.sample.financeDashboard.dashboardStates;

    expect(states, contains(FinanceDashboardState.summaryOnly));
    expect(states, contains(FinanceDashboardState.fixtureOnly));
    expect(states, contains(FinanceDashboardState.verifiedMetadata));
    expect(states, contains(FinanceDashboardState.cardLinkedReview));
    expect(
      FinanceDashboardState.verifiedMetadata.tone,
      JarvisBadgeTone.success,
    );
    expect(
      FinanceDashboardState.cardLinkedReview.tone,
      JarvisBadgeTone.warning,
    );
    expect(FinanceDashboardState.fixtureOnly.tone, JarvisBadgeTone.muted);
  });

  test('maps purchase dashboard states for agents', () {
    final states = JarvisSnapshot.sample.purchaseDashboard.dashboardStates;

    expect(states, contains(PurchaseDashboardState.summaryOnly));
    expect(states, contains(PurchaseDashboardState.fixtureOnly));
    expect(states, contains(PurchaseDashboardState.verifiedMetadata));
    expect(states, contains(PurchaseDashboardState.recurringCandidate));
    expect(
      PurchaseDashboardState.verifiedMetadata.tone,
      JarvisBadgeTone.success,
    );
    expect(
      PurchaseDashboardState.recurringCandidate.tone,
      JarvisBadgeTone.warning,
    );
    expect(PurchaseDashboardState.fixtureOnly.tone, JarvisBadgeTone.muted);
  });

  test('maps owner scopes for agents', () {
    expect(
      ownerScopeStates('household', 2),
      contains(OwnerScopeState.householdScoped),
    );
    expect(ownerScopeStates('user', 1), contains(OwnerScopeState.ownerScoped));
    expect(ownerScopeStates('spouse', 0), contains(OwnerScopeState.empty));
    expect(OwnerScopeState.ownerScoped.tone, JarvisBadgeTone.success);
    expect(OwnerScopeState.householdScoped.tone, JarvisBadgeTone.success);
    expect(OwnerScopeState.empty.tone, JarvisBadgeTone.warning);
  });
}
