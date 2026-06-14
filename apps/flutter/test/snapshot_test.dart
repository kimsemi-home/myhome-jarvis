import 'dart:convert';
import 'dart:io';

import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('offline fallback commands match generated command catalog', () async {
    final generated = await _generatedCommands();
    final fallbackByName = {
      for (final command in JarvisSnapshot.sample.commands)
        command.name: command,
    };
    final generatedNames = {
      for (final command in generated) _displayName(_string(command['name'])),
    };

    expect(fallbackByName.keys, unorderedEquals(generatedNames));

    for (final command in generated) {
      final name = _displayName(_string(command['name']));
      final fallback = fallbackByName[name];
      expect(fallback, isNotNull, reason: '$name missing from fallback');
      final fields = _stringList(command['payload_fields']);
      expect(fallback!.payloadFields, fields, reason: '$name payload fields');

      final decoded = jsonDecode(fallback.payload);
      expect(
        decoded,
        isA<Map<String, Object?>>(),
        reason: '$name payload JSON',
      );
      final payload = (decoded as Map).cast<String, Object?>();
      expect(
        payload.keys,
        unorderedEquals(fields),
        reason: '$name payload keys',
      );
    }
  });
}

Future<List<Map<String, Object?>>> _generatedCommands() async {
  final path = _findGeneratedCommandCatalog();
  final decoded = jsonDecode(await path.readAsString());
  if (decoded is! Map<String, Object?>) {
    throw FormatException('generated command catalog must be an object');
  }
  final commands = decoded['commands'];
  if (commands is! List<Object?>) {
    throw FormatException('generated command catalog commands must be a list');
  }
  return commands.whereType<Map<String, Object?>>().toList(growable: false);
}

File _findGeneratedCommandCatalog() {
  var directory = Directory.current;
  for (;;) {
    final candidate = File(
      '${directory.path}/generated/commands.generated.json',
    );
    if (candidate.existsSync()) {
      return candidate;
    }
    final parent = directory.parent;
    if (parent.path == directory.path) {
      throw FileSystemException(
        'could not find generated command catalog',
        Directory.current.path,
      );
    }
    directory = parent;
  }
}

String _displayName(String name) {
  return name.replaceAll('_', '-');
}

String _string(Object? value) {
  if (value is String && value.isNotEmpty) {
    return value;
  }
  throw FormatException('expected non-empty string');
}

List<String> _stringList(Object? value) {
  if (value is! List<Object?>) {
    throw FormatException('expected string list');
  }
  return value.whereType<String>().toList(growable: false);
}
