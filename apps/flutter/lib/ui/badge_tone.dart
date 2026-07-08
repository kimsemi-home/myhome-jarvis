part of '../main.dart';

enum JarvisBadgeTone {
  primary,
  secondary,
  outline,
  destructive,
  success,
  warning,
  muted,
}

extension JarvisBadgeToneStyle on JarvisBadgeTone {
  ShadBadgeVariant get variant => switch (this) {
    JarvisBadgeTone.primary => ShadBadgeVariant.primary,
    JarvisBadgeTone.secondary => ShadBadgeVariant.secondary,
    JarvisBadgeTone.outline => ShadBadgeVariant.outline,
    JarvisBadgeTone.destructive => ShadBadgeVariant.destructive,
    JarvisBadgeTone.success => ShadBadgeVariant.secondary,
    JarvisBadgeTone.warning => ShadBadgeVariant.secondary,
    JarvisBadgeTone.muted => ShadBadgeVariant.secondary,
  };

  String get semanticLabel => switch (this) {
    JarvisBadgeTone.primary => 'primary',
    JarvisBadgeTone.secondary => 'secondary',
    JarvisBadgeTone.outline => 'outline',
    JarvisBadgeTone.destructive => 'danger',
    JarvisBadgeTone.success => 'success',
    JarvisBadgeTone.warning => 'warning',
    JarvisBadgeTone.muted => 'muted',
  };

  Color? get backgroundColor => switch (this) {
    JarvisBadgeTone.success => JarvisAstryxTokens.successMuted,
    JarvisBadgeTone.warning => JarvisAstryxTokens.warningMuted,
    JarvisBadgeTone.destructive => JarvisAstryxTokens.errorMuted,
    JarvisBadgeTone.muted => JarvisAstryxTokens.backgroundMuted,
    _ => null,
  };

  Color? get foregroundColor => switch (this) {
    JarvisBadgeTone.success => JarvisAstryxTokens.success,
    JarvisBadgeTone.warning => JarvisAstryxTokens.warning,
    JarvisBadgeTone.destructive => JarvisAstryxTokens.error,
    JarvisBadgeTone.muted => JarvisAstryxTokens.textSecondary,
    _ => null,
  };
}
