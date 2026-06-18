import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import 'snapshot_generated_helpers.dart';

void main() {
  test('offline fallback commands match generated command catalog', () async {
    final generated = await generatedCommands();
    final fallbackByName = {
      for (final command in JarvisSnapshot.sample.commands)
        command.name: command,
    };
    final generatedNames = {
      for (final command in generated)
        displayName(requiredString(command['name'])),
    };

    expect(fallbackByName.keys, unorderedEquals(generatedNames));

    for (final command in generated) {
      final name = displayName(requiredString(command['name']));
      final fallback = fallbackByName[name];
      expect(fallback, isNotNull, reason: '$name missing from fallback');
      final fields = requiredStringList(command['payload_fields']);
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
