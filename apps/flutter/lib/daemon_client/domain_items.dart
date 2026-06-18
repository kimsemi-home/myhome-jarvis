part of '../daemon_client.dart';

List<String> _domainItems(Map<String, Object?> domain) {
  final items = <String>[];
  items.addAll(_financeItems(_object(domain['finance'])));
  items.addAll(_commerceItems(_object(domain['commerce'])));
  items.addAll(_storageItems(_object(domain['storage'])));
  return items.isEmpty ? JarvisSnapshot.sample.storageItems : items;
}

List<String> _financeItems(Map<String, Object?>? finance) {
  if (finance == null) {
    return const [];
  }
  final items = <String>[];
  final records = _int(finance['records']);
  final net = _int(finance['net_minor_units']);
  final currency = _string(finance['currency']) ?? '';
  if (records != null) {
    items.add('Finance: $records transactions');
  }
  if (net != null) {
    items.add('Finance net: $net $currency'.trim());
  }
  return items;
}

List<String> _commerceItems(Map<String, Object?>? commerce) {
  if (commerce == null) {
    return const [];
  }
  final items = <String>[];
  final records = _int(commerce['records']);
  final recurring = _int(commerce['recurring_candidate_count']);
  if (records != null) {
    items.add('Commerce: $records purchases');
  }
  if (recurring != null) {
    items.add('Recurring candidates: $recurring');
  }
  return items;
}

List<String> _storageItems(Map<String, Object?>? storage) {
  if (storage == null) {
    return const [];
  }
  final items = <String>[];
  final datasets = _stringList(storage['datasets']);
  final layers = _stringList(storage['lake_layers']);
  final format = _string(storage['long_term_format']);
  final compression = _string(storage['compression']);
  if (datasets.isNotEmpty) {
    items.add('Datasets: ${datasets.join(', ')}');
  }
  if (layers.isNotEmpty) {
    items.add('Lake layers: ${layers.join(', ')}');
  }
  if (format != null && compression != null) {
    items.add('Storage: $format+$compression');
  }
  return items;
}
