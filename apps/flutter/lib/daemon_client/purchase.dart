part of '../daemon_client.dart';

PurchaseDashboard _purchaseDashboard(Map<String, Object?> domain) {
  final commerce = _object(domain['commerce']);
  if (commerce == null) {
    return JarvisSnapshot.sample.purchaseDashboard;
  }
  return PurchaseDashboard(
    records: _int(commerce['records']) ?? 0,
    currency: _string(commerce['currency']) ?? '',
    totalSpendMinorUnits: _int(commerce['total_spend_minor_units']) ?? 0,
    recurringCandidateCount: _int(commerce['recurring_candidate_count']) ?? 0,
    recurringCandidates: _recurringPurchases(commerce['recurring_candidates']),
    categories: _stringList(commerce['categories']),
    owners: _purchaseOwners(commerce['owner_breakdown']),
  );
}

List<RecurringPurchase> _recurringPurchases(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  final candidates = <RecurringPurchase>[];
  for (final item in value.whereType<Map<String, Object?>>()) {
    final candidate = _recurringPurchase(item);
    if (candidate != null) {
      candidates.add(candidate);
    }
  }
  return candidates;
}

RecurringPurchase? _recurringPurchase(Map<String, Object?> item) {
  final merchantName = _string(item['merchant_name']);
  final itemName = _string(item['item_name']);
  if (merchantName == null ||
      merchantName.isEmpty ||
      itemName == null ||
      itemName.isEmpty) {
    return null;
  }
  return RecurringPurchase(
    merchantName: merchantName,
    itemName: itemName,
    currency: _string(item['currency']) ?? '',
    purchaseCount: _int(item['purchase_count']) ?? 0,
    latestTotalMinorUnits: _int(item['latest_total_minor_units']) ?? 0,
    latestPurchasedAt: _string(item['latest_purchased_at']) ?? '',
  );
}
