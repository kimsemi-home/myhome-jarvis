part of '../snapshot.dart';

const _offlineCoreMetrics = [
  SystemMetric(
    label: 'Daemon',
    value: 'Offline fallback',
    icon: Icons.settings_ethernet,
  ),
  SystemMetric(
    label: 'Network',
    value: 'Local-only',
    icon: Icons.wifi_off_outlined,
  ),
  SystemMetric(
    label: 'LAN Auth',
    value: 'Unavailable',
    icon: Icons.vpn_key_outlined,
  ),
  SystemMetric(label: 'Mode', value: 'Dry-run', icon: Icons.shield_outlined),
  SystemMetric(label: 'Quality', value: 'Local', icon: Icons.verified_outlined),
  SystemMetric(
    label: 'Code Shape',
    value: 'Local',
    icon: Icons.format_line_spacing,
  ),
  SystemMetric(label: 'Linear', value: 'Offline', icon: Icons.hub_outlined),
];
