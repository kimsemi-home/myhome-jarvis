part of '../snapshot.dart';

@immutable
class FinanceOwner {
  const FinanceOwner({
    required this.owner,
    required this.records,
    required this.currency,
    required this.creditMinorUnits,
    required this.debitMinorUnits,
    required this.netMinorUnits,
  });

  final String owner;
  final int records;
  final String currency;
  final int creditMinorUnits;
  final int debitMinorUnits;
  final int netMinorUnits;
}

@immutable
class FinanceDashboard {
  const FinanceDashboard({
    required this.records,
    required this.currency,
    required this.creditMinorUnits,
    required this.debitMinorUnits,
    required this.netMinorUnits,
    required this.subscriptionMinorUnits,
    required this.subscriptionCount,
    required this.cardDebitMinorUnits,
    required this.cardDebitCount,
    required this.categories,
    required this.owners,
    this.fixtureOnly = false,
  });

  final int records;
  final String currency;
  final int creditMinorUnits;
  final int debitMinorUnits;
  final int netMinorUnits;
  final int subscriptionMinorUnits;
  final int subscriptionCount;
  final int cardDebitMinorUnits;
  final int cardDebitCount;
  final List<String> categories;
  final List<FinanceOwner> owners;
  final bool fixtureOnly;
}
