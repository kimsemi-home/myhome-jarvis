part of '../daemon_client.dart';

class DaemonSnapshotClient implements JarvisClient {
  const DaemonSnapshotClient({
    required this.baseUri,
    this.timeout = const Duration(seconds: 2),
    this.authToken,
  });

  factory DaemonSnapshotClient.local() {
    return DaemonSnapshotClient(baseUri: Uri.parse('http://127.0.0.1:3888'));
  }

  final Uri baseUri;
  final Duration timeout;
  final String? authToken;

  @override
  Future<JarvisSnapshot> load() => _loadSnapshotFromDaemon(this);

  @override
  Future<CommandPlan> dryRun(HomeCommand command) {
    return _dryRunCommandWithDaemon(this, command);
  }
}
