part of '../daemon_client.dart';

String _plannerProgress(
  int ready,
  int completed,
  int blockedExternal,
  int tasks,
) {
  if (completed > 0 && ready == 0 && blockedExternal > 0) {
    return '$completed/$tasks done, $blockedExternal gated';
  }
  if (completed > 0 && ready == 0) {
    return '$completed/$tasks done';
  }
  return '$ready/$tasks ready';
}

String _authorityGateValue(String outcome, int debt, int blockedDecisions) {
  switch (outcome) {
    case 'blocked':
      return 'Blocked';
    case 'review_required':
      return debt > 0 ? '$debt debt' : 'Review required';
    case 'limited':
      return '$blockedDecisions blocked';
    default:
      return _titleWords(outcome);
  }
}

String _reviewCapacityValue(String state, int debt, int openCount) {
  switch (state) {
    case 'available':
      return 'Available';
    case 'overloaded':
      return debt > 0 ? '$debt debt' : 'Overloaded';
    case 'constrained':
      return debt > 0 ? '$debt debt' : _openOrState(openCount, 'Constrained');
    default:
      return _titleWords(state);
  }
}

String _openOrState(int openCount, String state) {
  return openCount > 0 ? '$openCount open' : state;
}

String _codeShapeValue(int regressions, int legacyDebt, int maxLines) {
  if (regressions > 0) {
    return regressions == 1 ? '1 regression' : '$regressions regressions';
  }
  if (legacyDebt > 0) {
    return '$legacyDebt debt';
  }
  return '<= $maxLines lines';
}

SystemMetric _countDebtMetric(
  String label,
  int debt,
  int count,
  String debtUnit,
  String countUnit,
  IconData icon,
) {
  return SystemMetric(
    label: label,
    value: debt > 0 ? '$debt $debtUnit' : '$count $countUnit',
    icon: icon,
  );
}
