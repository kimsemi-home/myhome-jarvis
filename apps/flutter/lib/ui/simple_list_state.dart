part of '../main.dart';

enum SimpleListItemState {
  offline,
  synced,
  queued,
  configured,
  localFixture,
  verified,
  format,
}

extension SimpleListItemStateUi on SimpleListItemState {
  String get label => switch (this) {
    SimpleListItemState.offline => 'offline',
    SimpleListItemState.synced => 'synced',
    SimpleListItemState.queued => 'queued',
    SimpleListItemState.configured => 'configured',
    SimpleListItemState.localFixture => 'local fixture',
    SimpleListItemState.verified => 'verified',
    SimpleListItemState.format => 'format',
  };

  JarvisBadgeTone get tone => switch (this) {
    SimpleListItemState.offline => JarvisBadgeTone.warning,
    SimpleListItemState.synced => JarvisBadgeTone.success,
    SimpleListItemState.queued => JarvisBadgeTone.warning,
    SimpleListItemState.configured => JarvisBadgeTone.success,
    SimpleListItemState.localFixture => JarvisBadgeTone.muted,
    SimpleListItemState.verified => JarvisBadgeTone.success,
    SimpleListItemState.format => JarvisBadgeTone.muted,
  };
}

SimpleListItemState simpleListItemState(String title, String item) {
  final text = '${title.toLowerCase()} ${item.toLowerCase()}';
  if (text.contains('offline') || text.contains('local fallback')) {
    return SimpleListItemState.offline;
  }
  if (text.contains('synced: true')) {
    return SimpleListItemState.synced;
  }
  if (text.contains('queue')) {
    return SimpleListItemState.queued;
  }
  if (text.contains('configured: true')) {
    return SimpleListItemState.configured;
  }
  if (text.contains('parquet') || text.contains('zstd')) {
    return SimpleListItemState.format;
  }
  if (text.contains('finance') || text.contains('commerce')) {
    return SimpleListItemState.localFixture;
  }
  return SimpleListItemState.verified;
}
