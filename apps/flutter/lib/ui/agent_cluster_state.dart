part of '../main.dart';

enum AgentClusterState { active, gated, tracked, stale }

extension AgentClusterStateUi on AgentClusterState {
  String get label => switch (this) {
    AgentClusterState.active => 'active',
    AgentClusterState.gated => 'gated',
    AgentClusterState.tracked => 'tracked',
    AgentClusterState.stale => 'stale',
  };

  JarvisBadgeTone get tone => switch (this) {
    AgentClusterState.active => JarvisBadgeTone.primary,
    AgentClusterState.gated => JarvisBadgeTone.secondary,
    AgentClusterState.tracked => JarvisBadgeTone.outline,
    AgentClusterState.stale => JarvisBadgeTone.destructive,
  };

  Color get iconColor => switch (this) {
    AgentClusterState.active => JarvisAstryxTokens.success,
    AgentClusterState.gated => JarvisAstryxTokens.warning,
    AgentClusterState.tracked => JarvisAstryxTokens.accent,
    AgentClusterState.stale => JarvisAstryxTokens.error,
  };
}

extension AgentClusterStateForSignal on AgentClusterSignal {
  AgentClusterState get clusterState {
    final text = status.toLowerCase();
    if (text.contains('stale') || text.contains('blocked')) {
      return AgentClusterState.stale;
    }
    if (text.contains('gated')) {
      return AgentClusterState.gated;
    }
    if (text.contains('active')) {
      return AgentClusterState.active;
    }
    return AgentClusterState.tracked;
  }
}
