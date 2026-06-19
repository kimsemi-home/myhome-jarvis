import 'dart:convert';
import 'dart:io';

void writeJson(HttpRequest request, Object? body) {
  request.response.headers.contentType = ContentType.json;
  request.response.write(jsonEncode(body));
  request.response.close();
}

Future<void> notFound(HttpRequest request) async {
  request.response.statusCode = HttpStatus.notFound;
  await request.response.close();
}
