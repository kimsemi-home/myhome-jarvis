part of '../daemon_client.dart';

String _authorityDecisionValue(Map<String, Object?> contract) {
  final canApply = _bool(contract['can_apply_decision']) ?? false;
  if (canApply) {
    return 'Apply enabled';
  }
  final items = contract['contract_items'];
  final itemCount = items is List<Object?> ? items.length : 0;
  if (itemCount > 0) {
    return '$itemCount scoped';
  }
  final reviewOnly = _bool(contract['review_only']) ?? false;
  return reviewOnly ? 'Review-only' : 'Unavailable';
}
