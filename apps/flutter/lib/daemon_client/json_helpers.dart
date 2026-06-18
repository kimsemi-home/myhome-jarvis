part of '../daemon_client.dart';

String? _string(Object? value) {
  return value is String ? value : null;
}

bool? _bool(Object? value) {
  return value is bool ? value : null;
}

int? _int(Object? value) {
  if (value is int) {
    return value;
  }
  return null;
}

Map<String, Object?>? _object(Object? value) {
  if (value is Map<String, Object?>) {
    return value;
  }
  return null;
}

List<String> _stringList(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  return value.whereType<String>().toList(growable: false);
}
