part of '../daemon_client.dart';

Future<CommandPlan> _dryRunCommandWithDaemon(
  DaemonSnapshotClient config,
  HomeCommand command,
) async {
  final client = HttpClient()..connectionTimeout = config.timeout;
  try {
    final decoded = await _postObject(config, client, '/intent', {
      'command': command.name,
      'payload': jsonDecode(command.payload) as Object?,
      'execute': false,
    });
    return commandPlanFromJson(decoded);
  } finally {
    client.close(force: true);
  }
}

CommandPlan commandPlanFromJson(Map<String, Object?> json) {
  return CommandPlan(
    name: _string(json['name']) ?? 'unknown',
    dryRun: _bool(json['dry_run']) ?? true,
    executeAllowed: _bool(json['execute_allowed']) ?? false,
    invocations: _invocations(json['invocations']),
    warnings: _stringList(json['warnings']),
  );
}

HomeCommand? _commandFromSpec(Map<String, Object?> spec) {
  final name = _string(spec['name']);
  if (name == null || name.isEmpty) {
    return null;
  }
  return HomeCommand(
    name: name.replaceAll('_', '-'),
    payload: _payloadExample(name),
    icon: _commandIcon(name),
    payloadFields: _stringList(spec['payload_fields']),
  );
}

List<HomeCommand> _commandsFromSpecs(List<Object?> commands) {
  return commands
      .whereType<Map<String, Object?>>()
      .map(_commandFromSpec)
      .whereType<HomeCommand>()
      .toList(growable: false);
}

List<CommandInvocation> _invocations(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  return value
      .whereType<Map<String, Object?>>()
      .map(
        (item) => CommandInvocation(
          label: _string(item['label']) ?? 'command',
          argv: _stringList(item['argv']),
          url: _string(item['url']),
        ),
      )
      .toList(growable: false);
}
