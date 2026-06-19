import 'dart:io';

import 'fixtures.dart';
import 'intent.dart';
import 'json_response.dart';

Future<HttpServer> startDaemonFixtureServer() async {
  final server = await HttpServer.bind(InternetAddress.loopbackIPv4, 0);
  server.listen(_handleDaemonFixtureRequest);
  return server;
}

Uri daemonUri(HttpServer server) {
  return Uri.parse('http://${server.address.address}:${server.port}');
}

Future<void> _handleDaemonFixtureRequest(HttpRequest request) async {
  if (request.uri.path == '/intent') {
    await handleIntentRequest(request);
    return;
  }
  final payload = await daemonFixturePayload(request.uri.path);
  if (payload == null) {
    await notFound(request);
    return;
  }
  writeJson(request, payload);
}
