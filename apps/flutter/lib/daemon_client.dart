import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

class DaemonSnapshotClient implements JarvisClient {
  const DaemonSnapshotClient({
    required this.baseUri,
    this.timeout = const Duration(seconds: 2),
    this.authToken,
  });

  factory DaemonSnapshotClient.local() {
    return DaemonSnapshotClient(baseUri: Uri.parse('http://127.0.0.1:3888'));
  }

  final Uri baseUri;
  final Duration timeout;
  final String? authToken;

  @override
  Future<JarvisSnapshot> load() async {
    final client = HttpClient()..connectionTimeout = timeout;
    try {
      final health = await _getObject(client, '/health');
      final commands = await _getArray(client, '/commands');
      final linear = await _getObject(client, '/linear/status');
      final repo = await _getObject(client, '/repo/status');
      final domain = await _getObject(client, '/domain/summary');
      final metrics = await _getObject(client, '/metrics');
      return buildSnapshot(
        health: health,
        commands: commands,
        linear: linear,
        repo: repo,
        domain: domain,
        metrics: metrics,
      );
    } finally {
      client.close(force: true);
    }
  }

  @override
  Future<CommandPlan> dryRun(HomeCommand command) async {
    final client = HttpClient()..connectionTimeout = timeout;
    try {
      final decoded = await _postObject(client, '/intent', {
        'command': command.name,
        'payload': jsonDecode(command.payload) as Object?,
        'execute': false,
      });
      return commandPlanFromJson(decoded);
    } finally {
      client.close(force: true);
    }
  }

  Future<Map<String, Object?>> _getObject(
    HttpClient client,
    String path,
  ) async {
    final decoded = await _getJson(client, path);
    if (decoded is Map<String, Object?>) {
      return decoded;
    }
    throw FormatException('expected object response from $path');
  }

  Future<List<Object?>> _getArray(HttpClient client, String path) async {
    final decoded = await _getJson(client, path);
    if (decoded is List<Object?>) {
      return decoded;
    }
    throw FormatException('expected array response from $path');
  }

  Future<Object?> _getJson(HttpClient client, String path) async {
    final request = await client.getUrl(baseUri.resolve(path)).timeout(timeout);
    _applyAuthHeader(request);
    final response = await request.close().timeout(timeout);
    return _decodeResponse(response, request.uri);
  }

  Future<Map<String, Object?>> _postObject(
    HttpClient client,
    String path,
    Object body,
  ) async {
    final request = await client
        .postUrl(baseUri.resolve(path))
        .timeout(timeout);
    request.headers.contentType = ContentType.json;
    _applyAuthHeader(request);
    request.write(jsonEncode(body));
    final response = await request.close().timeout(timeout);
    final decoded = await _decodeResponse(response, request.uri);
    if (decoded is Map<String, Object?>) {
      return decoded;
    }
    throw FormatException('expected object response from $path');
  }

  Future<Object?> _decodeResponse(HttpClientResponse response, Uri uri) async {
    final body = await response.transform(utf8.decoder).join().timeout(timeout);
    if (response.statusCode < 200 || response.statusCode >= 300) {
      throw HttpException('daemon returned ${response.statusCode}', uri: uri);
    }
    return jsonDecode(body) as Object?;
  }

  void _applyAuthHeader(HttpClientRequest request) {
    final token = authToken?.trim();
    if (token != null && token.isNotEmpty) {
      request.headers.set(HttpHeaders.authorizationHeader, 'Bearer $token');
    }
  }
}

CommandPlan commandPlanFromJson(Map<String, Object?> json) {
  return CommandPlan(
    name: _string(json['name']) ?? 'unknown',
    dryRun: _bool(json['dry_run']) ?? true,
    executeAllowed: _bool(json['execute_allowed']) ?? false,
    invocations: _invocations(json['invocations']),
    warnings: _stringList(json['warnings']),
  );
}

JarvisSnapshot buildSnapshot({
  required Map<String, Object?> health,
  required List<Object?> commands,
  required Map<String, Object?> linear,
  required Map<String, Object?> repo,
  required Map<String, Object?> domain,
  required Map<String, Object?> metrics,
}) {
  final bindHost =
      _string(metrics['bind_host']) ?? _string(health['host']) ?? '127.0.0.1';
  final dryRun =
      _bool(health['dry_run']) ?? _bool(metrics['dry_run_default']) ?? true;
  final requestCount = _int(metrics['requests']);
  final linearMode = _string(linear['mode']) ?? 'offline';
  final repoClean = _bool(repo['worktree_clean']);

  return JarvisSnapshot(
    metrics: [
      SystemMetric(
        label: 'Daemon',
        value: bindHost,
        icon: Icons.settings_ethernet,
      ),
      SystemMetric(
        label: 'Mode',
        value: dryRun ? 'Dry-run' : 'Execute-ready',
        icon: Icons.shield_outlined,
      ),
      SystemMetric(
        label: 'Requests',
        value: requestCount == null ? '0' : '$requestCount',
        icon: Icons.query_stats_outlined,
      ),
      SystemMetric(
        label: 'Linear',
        value: _title(linearMode),
        icon: Icons.hub_outlined,
      ),
      SystemMetric(
        label: 'Repo',
        value: repoClean == false ? 'Dirty' : 'Clean',
        icon: Icons.account_tree_outlined,
      ),
    ],
    commands: commands
        .whereType<Map<String, Object?>>()
        .map(_commandFromSpec)
        .whereType<HomeCommand>()
        .toList(growable: false),
    linearItems: _linearItems(linear),
    storageItems: _domainItems(domain),
    recommendationItems: _recommendationItems(domain),
    householdScopes: _householdScopes(domain),
  );
}

HomeCommand? _commandFromSpec(Map<String, Object?> spec) {
  final name = _string(spec['name']);
  if (name == null || name.isEmpty) {
    return null;
  }
  final displayName = name.replaceAll('_', '-');
  return HomeCommand(
    name: displayName,
    payload: _payloadExample(name),
    icon: _commandIcon(name),
    payloadFields: _stringList(spec['payload_fields']),
  );
}

List<String> _linearItems(Map<String, Object?> linear) {
  final items = <String>[];
  final mode = _string(linear['mode']);
  if (mode != null && mode.isNotEmpty) {
    items.add('Mode: ${_title(mode)}');
  }
  final synced = _bool(linear['synced']);
  if (synced != null) {
    items.add('Synced: $synced');
  }
  final teams = linear['teams'];
  if (teams is List<Object?>) {
    for (final team in teams.whereType<Map<String, Object?>>()) {
      final name = _string(team['name']);
      if (name != null && name.isNotEmpty) {
        items.add('Team: $name');
      }
    }
  }
  final queuePath = _string(linear['queue_path']);
  if (queuePath != null && queuePath.isNotEmpty) {
    items.add('Queue: linear-offline-queue.jsonl');
  }
  return items.isEmpty ? const ['Offline queue'] : items;
}

List<String> _domainItems(Map<String, Object?> domain) {
  final items = <String>[];
  final finance = _object(domain['finance']);
  if (finance != null) {
    final records = _int(finance['records']);
    final net = _int(finance['net_minor_units']);
    final currency = _string(finance['currency']) ?? '';
    if (records != null) {
      items.add('Finance: $records transactions');
    }
    if (net != null) {
      items.add('Finance net: $net $currency'.trim());
    }
  }

  final commerce = _object(domain['commerce']);
  if (commerce != null) {
    final records = _int(commerce['records']);
    final recurring = _int(commerce['recurring_candidate_count']);
    if (records != null) {
      items.add('Commerce: $records purchases');
    }
    if (recurring != null) {
      items.add('Recurring candidates: $recurring');
    }
  }

  final storage = _object(domain['storage']);
  if (storage != null) {
    final datasets = _stringList(storage['datasets']);
    final layers = _stringList(storage['lake_layers']);
    final format = _string(storage['long_term_format']);
    final compression = _string(storage['compression']);
    if (datasets.isNotEmpty) {
      items.add('Datasets: ${datasets.join(', ')}');
    }
    if (layers.isNotEmpty) {
      items.add('Lake layers: ${layers.join(', ')}');
    }
    if (format != null && compression != null) {
      items.add('Storage: $format+$compression');
    }
  }

  return items.isEmpty ? JarvisSnapshot.sample.storageItems : items;
}

List<String> _recommendationItems(Map<String, Object?> domain) {
  final recommendations = _object(domain['recommendations']);
  if (recommendations == null) {
    return JarvisSnapshot.sample.recommendationItems;
  }
  final rawItems = recommendations['items'];
  if (rawItems is! List<Object?>) {
    return JarvisSnapshot.sample.recommendationItems;
  }
  final items = <String>[];
  for (final item in rawItems.whereType<Map<String, Object?>>()) {
    final title = _string(item['title']);
    if (title == null || title.isEmpty) {
      continue;
    }
    final score = _int(item['score']);
    items.add(score == null ? title : '$score - $title');
  }
  return items.isEmpty ? JarvisSnapshot.sample.recommendationItems : items;
}

List<HouseholdScope> _householdScopes(Map<String, Object?> domain) {
  final household = _object(domain['household']);
  if (household == null) {
    return JarvisSnapshot.sample.householdScopes;
  }
  final rawScopes = household['scopes'];
  if (rawScopes is! List<Object?>) {
    return JarvisSnapshot.sample.householdScopes;
  }
  final scopes = <HouseholdScope>[];
  for (final item in rawScopes.whereType<Map<String, Object?>>()) {
    final scope = _string(item['scope']);
    final label = _string(item['label']);
    final currency = _string(item['currency']) ?? '';
    if (scope == null || label == null) {
      continue;
    }
    scopes.add(
      HouseholdScope(
        scope: scope,
        label: label,
        currency: currency,
        financeRecords: _int(item['finance_records']) ?? 0,
        financeNetMinorUnits: _int(item['finance_net_minor_units']) ?? 0,
        purchaseRecords: _int(item['purchase_records']) ?? 0,
        purchaseSpendMinorUnits: _int(item['purchase_spend_minor_units']) ?? 0,
      ),
    );
  }
  return scopes.isEmpty ? JarvisSnapshot.sample.householdScopes : scopes;
}

String _payloadExample(String name) {
  switch (name) {
    case 'open_ott':
      return '{"service":"netflix"}';
    case 'open_url':
      return '{"url":"https://www.youtube.com"}';
    case 'open_youtube_search':
      return '{"query":"lofi music"}';
    case 'volume_set':
      return '{"level":30}';
    case 'volume_up':
    case 'volume_down':
      return '{"step":10}';
    default:
      return '{}';
  }
}

IconData _commandIcon(String name) {
  switch (name) {
    case 'open_youtube':
    case 'open_youtube_search':
      return Icons.play_circle_outline;
    case 'open_ott':
    case 'movie_mode':
      return Icons.theaters_outlined;
    case 'volume_set':
    case 'volume_up':
    case 'volume_down':
    case 'volume_mute':
      return Icons.volume_up_outlined;
    case 'sleep_mode':
    case 'mac_sleep':
      return Icons.bedtime_outlined;
    case 'display_sleep':
      return Icons.monitor_outlined;
    default:
      return Icons.terminal_outlined;
  }
}

String? _string(Object? value) {
  return value is String ? value : null;
}

bool? _bool(Object? value) {
  return value is bool ? value : null;
}

int? _int(Object? value) {
  if (value is int) {
    return value;
  }
  return null;
}

Map<String, Object?>? _object(Object? value) {
  if (value is Map<String, Object?>) {
    return value;
  }
  return null;
}

String _title(String value) {
  if (value.isEmpty) {
    return value;
  }
  return value[0].toUpperCase() + value.substring(1).toLowerCase();
}

List<CommandInvocation> _invocations(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  return value
      .whereType<Map<String, Object?>>()
      .map(
        (item) => CommandInvocation(
          label: _string(item['label']) ?? 'command',
          argv: _stringList(item['argv']),
          url: _string(item['url']),
        ),
      )
      .toList(growable: false);
}

List<String> _stringList(Object? value) {
  if (value is! List<Object?>) {
    return const [];
  }
  return value.whereType<String>().toList(growable: false);
}
