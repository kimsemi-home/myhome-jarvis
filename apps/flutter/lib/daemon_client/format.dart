part of '../daemon_client.dart';

String _title(String value) {
  if (value.isEmpty) {
    return value;
  }
  return value[0].toUpperCase() + value.substring(1).toLowerCase();
}

String _titleWords(String value) {
  return value
      .split(RegExp(r'[_\-\s]+'))
      .where((part) => part.isNotEmpty)
      .map(_title)
      .join(' ');
}

String _networkMode(String bindHost, String? healthMode, bool lanBindAllowed) {
  if (lanBindAllowed) {
    return 'LAN token-gated';
  }
  if (healthMode == 'local' ||
      bindHost == '127.0.0.1' ||
      bindHost == 'localhost' ||
      bindHost == '::1') {
    return 'Local-only';
  }
  return 'Remote';
}

String _formatBytes(int bytes) {
  if (bytes < 1024) {
    return '$bytes B';
  }
  final kib = bytes / 1024;
  if (kib < 1024) {
    return '${kib.toStringAsFixed(1)} KiB';
  }
  final mib = kib / 1024;
  if (mib < 1024) {
    return '${mib.toStringAsFixed(1)} MiB';
  }
  return '${(mib / 1024).toStringAsFixed(1)} GiB';
}
