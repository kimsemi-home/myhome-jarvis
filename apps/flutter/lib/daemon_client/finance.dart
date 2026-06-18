part of '../daemon_client.dart';

FinanceDashboard _financeDashboard(Map<String, Object?> domain) {
  final finance = _object(domain['finance']);
  if (finance == null) {
    return JarvisSnapshot.sample.financeDashboard;
  }
  return FinanceDashboard(
    records: _int(finance['records']) ?? 0,
    currency: _string(finance['currency']) ?? '',
    creditMinorUnits: _int(finance['credit_minor_units']) ?? 0,
    debitMinorUnits: _int(finance['debit_minor_units']) ?? 0,
    netMinorUnits: _int(finance['net_minor_units']) ?? 0,
    subscriptionMinorUnits: _int(finance['subscription_minor_units']) ?? 0,
    subscriptionCount: _int(finance['subscription_count']) ?? 0,
    cardDebitMinorUnits: _int(finance['card_debit_minor_units']) ?? 0,
    cardDebitCount: _int(finance['card_debit_count']) ?? 0,
    categories: _stringList(finance['categories']),
    owners: _financeOwners(finance['owner_breakdown']),
  );
}

List<FinanceOwner> _financeOwners(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  final owners = <FinanceOwner>[];
  for (final item in value.whereType<Map<String, Object?>>()) {
    final owner = _financeOwner(item);
    if (owner != null) {
      owners.add(owner);
    }
  }
  return owners;
}

FinanceOwner? _financeOwner(Map<String, Object?> item) {
  final owner = _string(item['owner']);
  if (owner == null || owner.isEmpty) {
    return null;
  }
  return FinanceOwner(
    owner: owner,
    records: _int(item['records']) ?? 0,
    currency: _string(item['currency']) ?? '',
    creditMinorUnits: _int(item['credit_minor_units']) ?? 0,
    debitMinorUnits: _int(item['debit_minor_units']) ?? 0,
    netMinorUnits: _int(item['net_minor_units']) ?? 0,
  );
}
