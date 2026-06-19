import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

JarvisSnapshot buildLanSnapshot() {
  return buildSnapshot(
    health: <String, Object?>{
      'mode': 'local',
      'dry_run': true,
      'host': '192.168.1.10',
    },
    auth: const <String, Object?>{'configured': true},
    commands: const <Object?>[],
    linear: const <String, Object?>{},
    repo: const <String, Object?>{},
    security: const <String, Object?>{'ok': true},
    codeShape: const <String, Object?>{},
    domain: const <String, Object?>{},
    connectors: const <String, Object?>{},
    agentCluster: const <String, Object?>{},
    learning: const <String, Object?>{},
    evidence: const <String, Object?>{},
    confidence: const <String, Object?>{},
    translation: const <String, Object?>{},
    controlPlane: const <String, Object?>{},
    incidents: const <String, Object?>{},
    evidenceQuality: const <String, Object?>{},
    review: const <String, Object?>{},
    authority: const <String, Object?>{},
    metrics: <String, Object?>{
      'bind_host': '192.168.1.10',
      'dry_run_default': true,
      'execute_enabled': false,
      'lan_bind_allowed': true,
    },
    events: const <String, Object?>{},
    supervisor: const <String, Object?>{},
    audit: const <String, Object?>{},
    quality: const <String, Object?>{},
    planner: const <String, Object?>{},
  );
}
