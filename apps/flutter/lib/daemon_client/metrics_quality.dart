part of '../daemon_client.dart';

SystemMetric _qualityMetric(Map<String, Object?> quality) {
  final qualityCount = _int(quality['count']) ?? 0;
  final qualityLast = quality['last'];
  final qualityOK = qualityLast is Map<String, Object?>
      ? _bool(qualityLast['ok'])
      : null;
  return SystemMetric(
    label: 'Quality',
    value: qualityCount == 0
        ? 'Unrecorded'
        : '${qualityOK == false ? 'Failing' : 'Passing'} ($qualityCount)',
    icon: Icons.verified_outlined,
  );
}
