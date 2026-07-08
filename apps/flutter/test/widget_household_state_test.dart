import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('maps household scopes to agent-readable states', () {
    const user = HouseholdScope(
      scope: 'user',
      label: 'User',
      currency: 'KRW',
      financeRecords: 1,
      financeNetMinorUnits: -100,
      purchaseRecords: 1,
      purchaseSpendMinorUnits: 25,
    );
    const spouse = HouseholdScope(
      scope: 'spouse',
      label: 'Spouse',
      currency: 'KRW',
      financeRecords: 0,
      financeNetMinorUnits: 0,
      purchaseRecords: 0,
      purchaseSpendMinorUnits: 0,
    );
    const household = HouseholdScope(
      scope: 'household',
      label: 'Household',
      currency: 'KRW',
      financeRecords: 2,
      financeNetMinorUnits: 100,
      purchaseRecords: 2,
      purchaseSpendMinorUnits: 50,
    );

    expect(user.stateBadges, contains(HouseholdScopeState.ownerScoped));
    expect(spouse.stateBadges, contains(HouseholdScopeState.empty));
    expect(
      household.stateBadges,
      contains(HouseholdScopeState.householdScoped),
    );
  });
}
