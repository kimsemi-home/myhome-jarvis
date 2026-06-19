import 'dart:convert';
import 'dart:io';

import 'package:flutter_test/flutter_test.dart';

import 'json_response.dart';

Future<void> handleIntentRequest(HttpRequest request) async {
  final body = jsonDecode(await utf8.decoder.bind(request).join());
  expect(body, isA<Map<String, Object?>>());
  final payload = body as Map<String, Object?>;
  expect(payload['command'], 'volume-set');
  expect(payload['payload'], {'level': 30});
  expect(payload['execute'], false);
  writeJson(request, {
    'name': 'volume_set',
    'dry_run': true,
    'execute_allowed': false,
    'invocations': [
      {
        'label': 'volume_set',
        'argv': ['osascript', '-e', 'set volume output volume 30'],
      },
    ],
  });
}
