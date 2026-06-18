part of '../daemon_client.dart';

List<SystemMetric> _learningMetrics(_StatusMaps s) {
  final learningOpen = _int(s.learning['open_count']) ?? 0;
  final learningCount = _int(s.learning['count']) ?? 0;
  final evidenceEdges = _int(s.evidence['edge_count']) ?? 0;
  final evidenceNodes = _int(s.evidence['node_count']) ?? 0;
  final danglingRefs = _int(s.evidence['dangling_evidence_ref_count']) ?? 0;
  final confidenceCap = _string(s.confidence['level_cap']) ?? 'unknown';
  final confidenceBlocked = _bool(s.confidence['blocked']) ?? false;
  return [
    SystemMetric(
      label: 'Learning',
      value: learningOpen == 0
          ? '$learningCount observed'
          : '$learningOpen open',
      icon: Icons.psychology_alt_outlined,
    ),
    SystemMetric(
      label: 'Evidence Graph',
      value: danglingRefs > 0
          ? '$danglingRefs dangling'
          : (evidenceEdges == 0
                ? '$evidenceNodes nodes'
                : '$evidenceEdges links'),
      icon: Icons.device_hub_outlined,
    ),
    SystemMetric(
      label: 'Confidence',
      value: confidenceBlocked ? 'Blocked' : _title(confidenceCap),
      icon: Icons.rule_folder_outlined,
    ),
    ..._opsMetrics(s),
  ];
}
