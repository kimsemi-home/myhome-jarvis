part of '../daemon_client.dart';

List<PurchaseOwner> _purchaseOwners(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  final owners = <PurchaseOwner>[];
  for (final item in value.whereType<Map<String, Object?>>()) {
    final owner = _purchaseOwner(item);
    if (owner != null) {
      owners.add(owner);
    }
  }
  return owners;
}

PurchaseOwner? _purchaseOwner(Map<String, Object?> item) {
  final owner = _string(item['owner']);
  if (owner == null || owner.isEmpty) {
    return null;
  }
  return PurchaseOwner(
    owner: owner,
    records: _int(item['records']) ?? 0,
    currency: _string(item['currency']) ?? '',
    purchaseSpendMinorUnits: _int(item['purchase_spend_minor_units']) ?? 0,
  );
}
