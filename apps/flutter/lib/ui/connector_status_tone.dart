part of '../main.dart';

extension ConnectorStatusTone on ConnectorReadiness {
  JarvisBadgeTone get statusTone {
    final value = status.toLowerCase();
    if (value.contains('blocked') ||
        value.contains('fail') ||
        value.contains('error')) {
      return JarvisBadgeTone.destructive;
    }
    if (value.contains('ready') ||
        value.contains('verified') ||
        value.contains('ok')) {
      return JarvisBadgeTone.success;
    }
    if (value.contains('fixture') ||
        value.contains('local') ||
        value.contains('off')) {
      return JarvisBadgeTone.muted;
    }
    return JarvisBadgeTone.warning;
  }
}
