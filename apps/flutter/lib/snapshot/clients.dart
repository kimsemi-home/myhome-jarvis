part of '../snapshot.dart';

abstract interface class JarvisSnapshotClient {
  Future<JarvisSnapshot> load();
}

abstract interface class JarvisCommandClient {
  Future<CommandPlan> dryRun(HomeCommand command);
}

abstract interface class JarvisClient
    implements JarvisSnapshotClient, JarvisCommandClient {}

class StaticSnapshotClient implements JarvisClient {
  const StaticSnapshotClient(this.snapshot);

  final JarvisSnapshot snapshot;

  @override
  Future<JarvisSnapshot> load() async => snapshot;

  @override
  Future<CommandPlan> dryRun(HomeCommand command) async {
    return CommandPlan.preview(command);
  }
}
