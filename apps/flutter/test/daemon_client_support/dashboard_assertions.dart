import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void expectDomainDashboards(JarvisSnapshot snapshot) {
  expectPurchaseDashboard(snapshot);
  expectFinanceDashboard(snapshot);
  expectHouseholdScopes(snapshot);
}

void expectPurchaseDashboard(JarvisSnapshot snapshot) {
  expect(snapshot.purchaseDashboard.totalSpendMinorUnits, 26800);
  expect(snapshot.purchaseDashboard.recurringCandidateCount, 1);
  expect(snapshot.purchaseDashboard.categories, contains('grocery'));
  expect(snapshot.purchaseDashboard.owners.map((owner) => owner.owner), [
    'household',
    'user',
  ]);
  final candidate = snapshot.purchaseDashboard.recurringCandidates.single;
  expect(candidate.itemName, 'Bottled water 2L x 6');
  expect(candidate.latestTotalMinorUnits, 11800);
}

void expectFinanceDashboard(JarvisSnapshot snapshot) {
  expect(snapshot.financeDashboard.netMinorUnits, 4346800);
  expect(snapshot.financeDashboard.subscriptionMinorUnits, 65900);
  expect(snapshot.financeDashboard.cardDebitCount, 2);
  expect(snapshot.financeDashboard.categories, contains('subscription'));
  expect(snapshot.financeDashboard.owners.map((owner) => owner.owner), [
    'household',
    'user',
  ]);
  expect(snapshot.financeDashboard.owners.first.netMinorUnits, 4434100);
}

void expectHouseholdScopes(JarvisSnapshot snapshot) {
  expect(snapshot.householdScopes.map((scope) => scope.scope), [
    'user',
    'spouse',
    'household',
  ]);
  expect(snapshot.householdScopes.first.financeNetMinorUnits, -87300);
  expect(snapshot.householdScopes.last.purchaseSpendMinorUnits, 26800);
}
