import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('maps cluster signals to agent-readable states', () {
    const active = AgentClusterSignal(
      key: 'a',
      label: 'Active',
      status: 'active',
      evidence: '',
    );
    const gated = AgentClusterSignal(
      key: 'g',
      label: 'Gated',
      status: 'gated',
      evidence: '',
    );
    const stale = AgentClusterSignal(
      key: 's',
      label: 'Stale',
      status: 'stale',
      evidence: '',
    );

    expect(active.clusterState, AgentClusterState.active);
    expect(gated.clusterState, AgentClusterState.gated);
    expect(stale.clusterState, AgentClusterState.stale);
    expect(AgentClusterState.active.tone, JarvisBadgeTone.success);
    expect(AgentClusterState.gated.tone, JarvisBadgeTone.warning);
    expect(AgentClusterState.tracked.tone, JarvisBadgeTone.muted);
    expect(AgentClusterState.stale.tone, JarvisBadgeTone.destructive);
  });
}
