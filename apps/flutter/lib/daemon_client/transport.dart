part of '../daemon_client.dart';

Future<Map<String, Object?>> _getObject(
  DaemonSnapshotClient config,
  HttpClient client,
  String path,
) async {
  final decoded = await _getJson(config, client, path);
  if (decoded is Map<String, Object?>) {
    return decoded;
  }
  throw FormatException('expected object response from $path');
}

Future<List<Object?>> _getArray(
  DaemonSnapshotClient config,
  HttpClient client,
  String path,
) async {
  final decoded = await _getJson(config, client, path);
  if (decoded is List<Object?>) {
    return decoded;
  }
  throw FormatException('expected array response from $path');
}

Future<Object?> _getJson(
  DaemonSnapshotClient config,
  HttpClient client,
  String path,
) async {
  final request = await client
      .getUrl(config.baseUri.resolve(path))
      .timeout(config.timeout);
  _applyAuthHeader(config, request);
  final response = await request.close().timeout(config.timeout);
  return _decodeResponse(config, response, request.uri);
}

Future<Map<String, Object?>> _postObject(
  DaemonSnapshotClient config,
  HttpClient client,
  String path,
  Object body,
) async {
  final request = await client
      .postUrl(config.baseUri.resolve(path))
      .timeout(config.timeout);
  request.headers.contentType = ContentType.json;
  _applyAuthHeader(config, request);
  request.write(jsonEncode(body));
  final response = await request.close().timeout(config.timeout);
  final decoded = await _decodeResponse(config, response, request.uri);
  if (decoded is Map<String, Object?>) {
    return decoded;
  }
  throw FormatException('expected object response from $path');
}
