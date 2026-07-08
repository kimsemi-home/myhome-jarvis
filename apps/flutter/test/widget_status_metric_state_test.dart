import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('maps status metric values to agent-readable badge states', () {
    const gated = SystemMetric(
      label: 'Assistant Center',
      value: '2 gated',
      icon: Icons.assistant_direction_outlined,
    );
    const blocked = SystemMetric(
      label: 'Authority Gate',
      value: '6 blocked',
      icon: Icons.admin_panel_settings_outlined,
    );
    const verified = SystemMetric(
      label: 'Public Safety',
      value: 'Clear',
      icon: Icons.verified_user_outlined,
    );
    const local = SystemMetric(
      label: 'Network',
      value: 'Local-only',
      icon: Icons.wifi_off_outlined,
    );

    expect(gated.statusState, StatusMetricState.warning);
    expect(blocked.statusState, StatusMetricState.blocked);
    expect(verified.statusState, StatusMetricState.verified);
    expect(local.statusState, StatusMetricState.local);
  });
}
