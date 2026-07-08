part of '../snapshot.dart';

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
    this.fixtureOnly = false,
  });

  final int records;
  final String currency;
  final int totalSpendMinorUnits;
  final int recurringCandidateCount;
  final List<RecurringPurchase> recurringCandidates;
  final List<String> categories;
  final List<PurchaseOwner> owners;
  final bool fixtureOnly;
}
