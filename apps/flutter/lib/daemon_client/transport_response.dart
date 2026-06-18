part of '../daemon_client.dart';

Future<Object?> _decodeResponse(
  DaemonSnapshotClient config,
  HttpClientResponse response,
  Uri uri,
) async {
  final body = await response
      .transform(utf8.decoder)
      .join()
      .timeout(config.timeout);
  if (response.statusCode < 200 || response.statusCode >= 300) {
    throw HttpException('daemon returned ${response.statusCode}', uri: uri);
  }
  return jsonDecode(body) as Object?;
}

void _applyAuthHeader(DaemonSnapshotClient config, HttpClientRequest request) {
  final token = config.authToken?.trim();
  if (token != null && token.isNotEmpty) {
    request.headers.set(HttpHeaders.authorizationHeader, 'Bearer $token');
  }
}
