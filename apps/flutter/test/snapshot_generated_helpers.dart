import 'dart:convert';
import 'dart:io';

Future<List<Map<String, Object?>>> generatedCommands() => generatedEntries(
  'generated/commands.generated.json',
  'commands',
  'generated command catalog',
);

Future<List<Map<String, Object?>>> generatedConnectors() => generatedEntries(
  'generated/connectors.generated.json',
  'connectors',
  'generated connector catalog',
);

Future<List<Map<String, Object?>>> generatedAgentClusterSignals() =>
    generatedEntries(
      'generated/agent_cluster.generated.json',
      'signals',
      'generated agent cluster policy',
    );

Future<List<Map<String, Object?>>> generatedEntries(
  String relativePath,
  String key,
  String label,
) async {
  final decoded = jsonDecode(
    await findGeneratedFile(relativePath).readAsString(),
  );
  if (decoded is! Map<String, Object?>) {
    throw FormatException('$label must be an object');
  }
  final entries = decoded[key];
  if (entries is! List<Object?>) {
    throw FormatException('$label $key must be a list');
  }
  return entries.whereType<Map<String, Object?>>().toList(growable: false);
}

File findGeneratedFile(String relativePath) {
  var directory = Directory.current;
  for (;;) {
    final candidate = File('${directory.path}/$relativePath');
    if (candidate.existsSync()) {
      return candidate;
    }
    final parent = directory.parent;
    if (parent.path == directory.path) {
      throw FileSystemException(
        'could not find generated file',
        Directory.current.path,
      );
    }
    directory = parent;
  }
}

String displayName(String name) {
  return name.replaceAll('_', '-');
}

String requiredString(Object? value) {
  if (value is String && value.isNotEmpty) {
    return value;
  }
  throw FormatException('expected non-empty string');
}

List<String> requiredStringList(Object? value) {
  if (value is! List<Object?>) {
    throw FormatException('expected string list');
  }
  return value.whereType<String>().toList(growable: false);
}
