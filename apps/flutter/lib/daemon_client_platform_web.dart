import 'dart:async';

class ContentType {
  const ContentType._(this.mimeType);

  static const json = ContentType._('application/json');
  final String mimeType;
}

class HttpHeaders {
  static const authorizationHeader = 'authorization';

  ContentType? contentType;

  void set(String name, Object value) {}
}

class HttpClient {
  Duration? connectionTimeout;

  Future<HttpClientRequest> getUrl(Uri uri) => _unsupported(uri);

  Future<HttpClientRequest> postUrl(Uri uri) => _unsupported(uri);

  void close({bool force = false}) {}
}

class HttpClientRequest {
  HttpClientRequest(this.uri);

  final Uri uri;
  final HttpHeaders headers = HttpHeaders();

  void write(Object? value) {}

  Future<HttpClientResponse> close() => _unsupported(uri);
}

class HttpClientResponse {
  int get statusCode => 500;

  Stream<String> transform(StreamTransformer<List<int>, String> transformer) =>
      Stream<String>.empty();
}

class HttpException implements Exception {
  HttpException(this.message, {this.uri});

  final String message;
  final Uri? uri;

  @override
  String toString() => message;
}

Future<T> _unsupported<T>(Uri uri) {
  return Future<T>.error(
    UnsupportedError('daemon access is disabled in the public web demo: $uri'),
  );
}
