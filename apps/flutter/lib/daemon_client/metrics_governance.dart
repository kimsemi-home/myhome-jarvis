part of '../daemon_client.dart';

List<SystemMetric> _governanceMetrics(_StatusMaps s) {
  final connectorCount = _int(s.connectors['connector_count']) ?? 0;
  final fixtureCount = _int(s.connectors['fixture_mode_count']) ?? 0;
  final roleCount = _int(s.agentCluster['role_count']) ?? 0;
  final gateRequired =
      _bool(s.agentCluster['authority_gate_required']) ?? false;
  final outcome = _string(s.authority['outcome']) ?? 'unknown';
  final authorityDebt = _int(s.authority['authority_debt_count']) ?? 0;
  final blocked = _int(s.authority['blocked_decision_count']) ?? 0;
  final reviewState = _string(s.review['capacity_state']) ?? 'unknown';
  final reviewDebt = _int(s.review['review_debt_count']) ?? 0;
  final reviewOpen = _int(s.review['open_count']) ?? 0;
  final decisionContract =
      _object(s.authorityReviewDecision['decision_contract']) ?? {};
  return [
    SystemMetric(
      label: 'Linear',
      value: _title(_string(s.linear['mode']) ?? 'offline'),
      icon: Icons.hub_outlined,
    ),
    SystemMetric(
      label: 'Connectors',
      value: connectorCount == 0
          ? 'Fixture-only'
          : '$fixtureCount/$connectorCount fixture',
      icon: Icons.cable_outlined,
    ),
    SystemMetric(
      label: 'Agent Cluster',
      value: gateRequired
          ? (roleCount == 0 ? 'Governed' : '$roleCount roles gated')
          : 'Ungated',
      icon: Icons.account_tree_outlined,
    ),
    SystemMetric(
      label: 'Authority Gate',
      value: _authorityGateValue(outcome, authorityDebt, blocked),
      icon: Icons.admin_panel_settings_outlined,
    ),
    SystemMetric(
      label: 'Authority Decision',
      value: _authorityDecisionValue(decisionContract),
      icon: Icons.fact_check_outlined,
    ),
    SystemMetric(
      label: 'Review Capacity',
      value: _reviewCapacityValue(reviewState, reviewDebt, reviewOpen),
      icon: Icons.rate_review_outlined,
    ),
  ];
}
