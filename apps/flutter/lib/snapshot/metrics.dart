part of '../snapshot.dart';

@immutable
class SystemMetric {
  const SystemMetric({
    required this.label,
    required this.value,
    required this.icon,
  });

  final String label;
  final String value;
  final IconData icon;
}
