import 'dart:io';

import 'package:flutter_test/flutter_test.dart';

import 'json_response.dart';

Future<HttpServer> startAuthHeaderCaptureServer(String expectedToken) async {
  final server = await HttpServer.bind(InternetAddress.loopbackIPv4, 0);
  server.listen((request) async {
    expect(
      request.headers.value(HttpHeaders.authorizationHeader),
      'Bearer $expectedToken',
    );
    expect(request.uri.path, '/intent');
    writeJson(request, {
      'name': 'volume_set',
      'dry_run': true,
      'execute_allowed': false,
      'invocations': <Object?>[],
    });
  });
  return server;
}
