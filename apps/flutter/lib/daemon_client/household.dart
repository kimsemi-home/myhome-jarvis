part of '../daemon_client.dart';

List<HouseholdScope> _householdScopes(Map<String, Object?> domain) {
  final household = _object(domain['household']);
  if (household == null) {
    return JarvisSnapshot.sample.householdScopes;
  }
  final rawScopes = household['scopes'];
  if (rawScopes is! List<Object?>) {
    return JarvisSnapshot.sample.householdScopes;
  }
  final scopes = <HouseholdScope>[];
  for (final item in rawScopes.whereType<Map<String, Object?>>()) {
    final scope = _householdScope(item);
    if (scope != null) {
      scopes.add(scope);
    }
  }
  return scopes.isEmpty ? JarvisSnapshot.sample.householdScopes : scopes;
}

HouseholdScope? _householdScope(Map<String, Object?> item) {
  final scope = _string(item['scope']);
  final label = _string(item['label']);
  if (scope == null || label == null) {
    return null;
  }
  return HouseholdScope(
    scope: scope,
    label: label,
    currency: _string(item['currency']) ?? '',
    financeRecords: _int(item['finance_records']) ?? 0,
    financeNetMinorUnits: _int(item['finance_net_minor_units']) ?? 0,
    purchaseRecords: _int(item['purchase_records']) ?? 0,
    purchaseSpendMinorUnits: _int(item['purchase_spend_minor_units']) ?? 0,
  );
}
