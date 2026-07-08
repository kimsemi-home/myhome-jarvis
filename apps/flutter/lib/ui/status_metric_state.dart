part of '../main.dart';

enum StatusMetricState { blocked, warning, verified, local, observed }

extension StatusMetricStateUi on StatusMetricState {
  String get badgeLabel => switch (this) {
    StatusMetricState.blocked => 'blocked',
    StatusMetricState.warning => 'warning',
    StatusMetricState.verified => 'verified',
    StatusMetricState.local => 'local',
    StatusMetricState.observed => 'observed',
  };

  JarvisBadgeTone get badgeTone => switch (this) {
    StatusMetricState.blocked => JarvisBadgeTone.destructive,
    StatusMetricState.warning => JarvisBadgeTone.secondary,
    StatusMetricState.verified => JarvisBadgeTone.primary,
    StatusMetricState.local => JarvisBadgeTone.outline,
    StatusMetricState.observed => JarvisBadgeTone.outline,
  };

  Color get iconColor => switch (this) {
    StatusMetricState.blocked => JarvisAstryxTokens.error,
    StatusMetricState.warning => JarvisAstryxTokens.warning,
    StatusMetricState.verified => JarvisAstryxTokens.success,
    StatusMetricState.local => JarvisAstryxTokens.accent,
    StatusMetricState.observed => JarvisAstryxTokens.textSecondary,
  };
}

extension StatusMetricStateForMetric on SystemMetric {
  StatusMetricState get statusState => statusMetricStateFor(label, value);
}
