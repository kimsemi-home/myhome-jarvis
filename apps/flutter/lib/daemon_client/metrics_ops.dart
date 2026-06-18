part of '../daemon_client.dart';

List<SystemMetric> _opsMetrics(_StatusMaps s) {
  final translationDebt = _int(s.translation['open_debt_count']) ?? 0;
  final forbidden = _int(s.translation['forbidden_loss_count']) ?? 0;
  final manifests = _int(s.translation['manifest_count']) ?? 0;
  return [
    _translationMetric(forbidden, translationDebt, manifests),
    _countDebtMetric(
      'Control Plane',
      _int(s.controlPlane['manifest_debt_count']) ?? 0,
      _int(s.controlPlane['count']) ?? 0,
      'manifest debt',
      'manifests',
      Icons.route_outlined,
    ),
    _countDebtMetric(
      'Incidents',
      _int(s.incidents['incident_debt_count']) ?? 0,
      _int(s.incidents['count']) ?? 0,
      'incident debt',
      'incidents',
      Icons.crisis_alert_outlined,
    ),
    _countDebtMetric(
      'Evidence Quality',
      _int(s.evidenceQuality['reassessment_debt_count']) ?? 0,
      _int(s.evidenceQuality['snapshot_count']) ?? 0,
      'reassess',
      'snapshots',
      Icons.fact_check_outlined,
    ),
  ];
}

SystemMetric _translationMetric(int forbidden, int debt, int manifests) {
  return SystemMetric(
    label: 'Translation',
    value: forbidden > 0
        ? '$forbidden forbidden'
        : (debt > 0 ? '$debt open debt' : '$manifests manifests'),
    icon: Icons.translate_outlined,
  );
}
