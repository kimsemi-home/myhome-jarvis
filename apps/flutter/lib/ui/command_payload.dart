part of '../main.dart';

Map<String, Object?> _decodePayload(String payload) {
  try {
    final decoded = jsonDecode(payload);
    if (decoded is Map<String, Object?>) {
      return decoded;
    }
  } on FormatException {
    return const {};
  }
  return const {};
}

String _payloadText(Object? value) {
  return value == null ? '' : '$value';
}

Object _payloadValue(String field, String text) {
  if (_numericPayloadField(field)) {
    return int.tryParse(text) ?? 0;
  }
  return text;
}

bool _numericPayloadField(String field) {
  return field == 'level' || field == 'step';
}
