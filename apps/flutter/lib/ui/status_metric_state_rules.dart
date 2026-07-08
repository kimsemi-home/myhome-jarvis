part of '../main.dart';

StatusMetricState statusMetricStateFor(String label, String value) {
  final text = '${label.toLowerCase()} ${value.toLowerCase()}';
  if (_hasAny(text, const [
    'blocked',
    'failing',
    'findings',
    'regression',
    'dirty',
    'stale',
    'missing',
    'forbidden',
  ])) {
    return StatusMetricState.blocked;
  }
  if (text.contains('ungated')) {
    return StatusMetricState.verified;
  }
  if (_hasAny(text, const [
    'gated',
    'debt',
    'open',
    'overloaded',
    'constrained',
    'unrecorded',
    'dangling',
  ])) {
    return StatusMetricState.warning;
  }
  if (_hasAny(text, const [
    'clear',
    'passing',
    'configured',
    'available',
    'tracked',
    'clean',
    'reachable',
    'ready',
    '<=',
  ])) {
    return StatusMetricState.verified;
  }
  if (_hasAny(text, const ['local-only', 'dry-run', 'local', 'fixture-only'])) {
    return StatusMetricState.local;
  }
  return StatusMetricState.observed;
}

bool _hasAny(String text, List<String> terms) {
  for (final term in terms) {
    if (text.contains(term)) {
      return true;
    }
  }
  return false;
}
