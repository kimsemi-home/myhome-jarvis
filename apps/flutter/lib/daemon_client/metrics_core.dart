part of '../daemon_client.dart';

List<SystemMetric> _coreMetrics(_StatusMaps s) {
  final bindHost =
      _string(s.metrics['bind_host']) ??
      _string(s.health['host']) ??
      '127.0.0.1';
  final dryRun =
      _bool(s.health['dry_run']) ?? _bool(s.metrics['dry_run_default']) ?? true;
  final executeEnabled = _bool(s.metrics['execute_enabled']) ?? false;
  final publicSafetyOK = _bool(s.security['ok']);
  final securityFindings =
      (_int(s.security['current_finding_count']) ?? 0) +
      (_int(s.security['history_finding_count']) ?? 0);
  return [
    SystemMetric(
      label: 'Daemon',
      value: bindHost,
      icon: Icons.settings_ethernet,
    ),
    SystemMetric(
      label: 'Network',
      value: _networkMode(
        bindHost,
        _string(s.health['mode']),
        _bool(s.metrics['lan_bind_allowed']) ?? false,
      ),
      icon: _bool(s.metrics['lan_bind_allowed']) == true
          ? Icons.lan_outlined
          : Icons.wifi_off_outlined,
    ),
    SystemMetric(
      label: 'LAN Auth',
      value: _bool(s.auth['configured']) == true ? 'Configured' : 'Missing',
      icon: Icons.vpn_key_outlined,
    ),
    SystemMetric(
      label: 'Mode',
      value: dryRun
          ? 'Dry-run'
          : (executeEnabled ? 'Execute-gated' : 'Execute-ready'),
      icon: Icons.shield_outlined,
    ),
    _qualityMetric(s.quality),
    SystemMetric(
      label: 'Public Safety',
      value: publicSafetyOK == false ? 'Findings ($securityFindings)' : 'Clear',
      icon: publicSafetyOK == false
          ? Icons.report_problem_outlined
          : Icons.verified_user_outlined,
    ),
    SystemMetric(
      label: 'Code Shape',
      value: _codeShapeValue(
        _int(s.codeShape['budget_regression_count']) ?? 0,
        _int(s.codeShape['legacy_debt_count']) ?? 0,
        _int(s.codeShape['max_file_lines']) ?? 75,
      ),
      icon: Icons.format_line_spacing,
    ),
  ];
}
