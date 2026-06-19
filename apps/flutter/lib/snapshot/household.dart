part of '../snapshot.dart';

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
