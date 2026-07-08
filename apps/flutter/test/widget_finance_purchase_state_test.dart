import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('maps finance dashboard states for agents', () {
    final states = JarvisSnapshot.sample.financeDashboard.dashboardStates;

    expect(states, contains(FinanceDashboardState.summaryOnly));
    expect(states, contains(FinanceDashboardState.verifiedMetadata));
    expect(states, contains(FinanceDashboardState.cardLinkedReview));
  });

  test('maps purchase dashboard states for agents', () {
    final states = JarvisSnapshot.sample.purchaseDashboard.dashboardStates;

    expect(states, contains(PurchaseDashboardState.summaryOnly));
    expect(states, contains(PurchaseDashboardState.verifiedMetadata));
    expect(states, contains(PurchaseDashboardState.recurringCandidate));
  });

  test('maps owner scopes for agents', () {
    expect(
      ownerScopeStates('household', 2),
      contains(OwnerScopeState.householdScoped),
    );
    expect(ownerScopeStates('user', 1), contains(OwnerScopeState.ownerScoped));
    expect(ownerScopeStates('spouse', 0), contains(OwnerScopeState.empty));
  });
}
