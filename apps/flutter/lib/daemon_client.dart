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
      final auth = await _getObject(client, '/auth/status');
      final commands = await _getArray(client, '/commands');
      final linear = await _getObject(client, '/linear/status');
      final repo = await _getObject(client, '/repo/status');
      final security = await _getObject(client, '/security/status');
      final domain = await _getObject(client, '/domain/summary');
      final connectors = await _getObject(client, '/connectors/status');
      final agentCluster = await _getObject(client, '/agent-cluster/status');
      final learning = await _getObject(client, '/learning/status');
      final metrics = await _getObject(client, '/metrics');
      final events = await _getObject(client, '/events');
      final supervisor = await _getObject(client, '/supervisor/status');
      final audit = await _getObject(client, '/audit/status');
      final quality = await _getObject(client, '/quality/status');
      final planner = await _getObject(client, '/planner/status');
      return buildSnapshot(
        health: health,
        auth: auth,
        commands: commands,
        linear: linear,
        repo: repo,
        security: security,
        domain: domain,
        connectors: connectors,
        agentCluster: agentCluster,
        learning: learning,
        metrics: metrics,
        events: events,
        supervisor: supervisor,
        audit: audit,
        quality: quality,
        planner: planner,
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
  required Map<String, Object?> auth,
  required List<Object?> commands,
  required Map<String, Object?> linear,
  required Map<String, Object?> repo,
  required Map<String, Object?> security,
  required Map<String, Object?> domain,
  required Map<String, Object?> connectors,
  required Map<String, Object?> agentCluster,
  required Map<String, Object?> learning,
  required Map<String, Object?> metrics,
  required Map<String, Object?> events,
  required Map<String, Object?> supervisor,
  required Map<String, Object?> audit,
  required Map<String, Object?> quality,
  required Map<String, Object?> planner,
}) {
  final bindHost =
      _string(metrics['bind_host']) ?? _string(health['host']) ?? '127.0.0.1';
  final healthMode = _string(health['mode']);
  final lanBindAllowed = _bool(metrics['lan_bind_allowed']) ?? false;
  final dryRun =
      _bool(health['dry_run']) ?? _bool(metrics['dry_run_default']) ?? true;
  final executeEnabled = _bool(metrics['execute_enabled']) ?? false;
  final requestCount = _int(metrics['requests']);
  final eventCount = _int(events['count']) ?? _int(metrics['event_count']);
  final goroutineCount = _int(metrics['goroutine_count']);
  final heapAllocBytes = _int(metrics['heap_alloc_bytes']);
  final auditCount = _int(audit['count']);
  final qualityCount = _int(quality['count']) ?? 0;
  final qualityLast = quality['last'];
  final qualityOK = qualityLast is Map<String, Object?>
      ? _bool(qualityLast['ok'])
      : null;
  final supervisorStale = _bool(supervisor['stale']);
  final supervisorRecorded = _bool(supervisor['recorded']) ?? false;
  final plannerReady = _int(planner['ready_count']) ?? 0;
  final plannerCompleted = _int(planner['completed_count']) ?? 0;
  final plannerBlockedExternal =
      _int(planner['blocked_external_write_count']) ?? 0;
  final plannerGate = _plannerGate(planner['blocked_external_write_tasks']);
  final plannerTasks = _int(planner['task_count']) ?? 0;
  final linearMode = _string(linear['mode']) ?? 'offline';
  final repoClean = _bool(repo['worktree_clean']);
  final authConfigured = _bool(auth['configured']);
  final publicSafetyOK = _bool(security['ok']);
  final securityFindings =
      (_int(security['current_finding_count']) ?? 0) +
      (_int(security['history_finding_count']) ?? 0);
  final connectorCount = _int(connectors['connector_count']) ?? 0;
  final fixtureConnectorCount = _int(connectors['fixture_mode_count']) ?? 0;
  final agentRoleCount = _int(agentCluster['role_count']) ?? 0;
  final authorityGateRequired =
      _bool(agentCluster['authority_gate_required']) ?? false;
  final learningOpenCount = _int(learning['open_count']) ?? 0;
  final learningCount = _int(learning['count']) ?? 0;

  return JarvisSnapshot(
    metrics: [
      SystemMetric(
        label: 'Daemon',
        value: bindHost,
        icon: Icons.settings_ethernet,
      ),
      SystemMetric(
        label: 'Network',
        value: _networkMode(bindHost, healthMode, lanBindAllowed),
        icon: lanBindAllowed ? Icons.lan_outlined : Icons.wifi_off_outlined,
      ),
      SystemMetric(
        label: 'LAN Auth',
        value: authConfigured == true ? 'Configured' : 'Missing',
        icon: Icons.vpn_key_outlined,
      ),
      SystemMetric(
        label: 'Mode',
        value: dryRun
            ? 'Dry-run'
            : (executeEnabled ? 'Execute-gated' : 'Execute-ready'),
        icon: Icons.shield_outlined,
      ),
      SystemMetric(
        label: 'Quality',
        value: qualityCount == 0
            ? 'Unrecorded'
            : '${qualityOK == false ? 'Failing' : 'Passing'} ($qualityCount)',
        icon: Icons.verified_outlined,
      ),
      SystemMetric(
        label: 'Public Safety',
        value: publicSafetyOK == false
            ? 'Findings ($securityFindings)'
            : 'Clear',
        icon: publicSafetyOK == false
            ? Icons.report_problem_outlined
            : Icons.verified_user_outlined,
      ),
      SystemMetric(
        label: 'Requests',
        value: requestCount == null ? '0' : '$requestCount',
        icon: Icons.query_stats_outlined,
      ),
      SystemMetric(
        label: 'Events',
        value: eventCount == null ? '0' : '$eventCount',
        icon: Icons.receipt_long_outlined,
      ),
      if (goroutineCount != null)
        SystemMetric(
          label: 'Runtime',
          value: '$goroutineCount goroutines',
          icon: Icons.memory_outlined,
        ),
      if (heapAllocBytes != null)
        SystemMetric(
          label: 'Heap',
          value: _formatBytes(heapAllocBytes),
          icon: Icons.storage_outlined,
        ),
      SystemMetric(
        label: 'Supervisor',
        value: supervisorRecorded
            ? (supervisorStale == true ? 'Stale' : 'Reachable')
            : 'Unrecorded',
        icon: Icons.memory_outlined,
      ),
      SystemMetric(
        label: 'Command Audit',
        value: auditCount == null ? '0' : '$auditCount',
        icon: Icons.fact_check_outlined,
      ),
      SystemMetric(
        label: 'Planner',
        value: _plannerProgress(
          plannerReady,
          plannerCompleted,
          plannerBlockedExternal,
          plannerTasks,
        ),
        icon: Icons.schema_outlined,
      ),
      if (plannerGate != null)
        SystemMetric(
          label: 'Planner Gate',
          value: plannerGate,
          icon: Icons.lock_outline,
        ),
      SystemMetric(
        label: 'Linear',
        value: _title(linearMode),
        icon: Icons.hub_outlined,
      ),
      SystemMetric(
        label: 'Connectors',
        value: connectorCount == 0
            ? 'Fixture-only'
            : '$fixtureConnectorCount/$connectorCount fixture',
        icon: Icons.cable_outlined,
      ),
      SystemMetric(
        label: 'Agent Cluster',
        value: authorityGateRequired
            ? (agentRoleCount == 0 ? 'Governed' : '$agentRoleCount roles gated')
            : 'Ungated',
        icon: Icons.account_tree_outlined,
      ),
      SystemMetric(
        label: 'Learning',
        value: learningOpenCount == 0
            ? '$learningCount observed'
            : '$learningOpenCount open',
        icon: Icons.psychology_alt_outlined,
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
    recommendations: _recommendations(domain),
    householdScopes: _householdScopes(domain),
    financeDashboard: _financeDashboard(domain),
    purchaseDashboard: _purchaseDashboard(domain),
    connectors: _connectors(connectors),
    agentClusterSignals: _agentClusterSignals(agentCluster),
  );
}

String _plannerProgress(
  int ready,
  int completed,
  int blockedExternal,
  int tasks,
) {
  if (completed > 0 && ready == 0 && blockedExternal > 0) {
    return '$completed/$tasks done, $blockedExternal gated';
  }
  if (completed > 0 && ready == 0) {
    return '$completed/$tasks done';
  }
  return '$ready/$tasks ready';
}

String? _plannerGate(Object? tasks) {
  if (tasks is! List<Object?> || tasks.isEmpty) {
    return null;
  }
  final first = tasks.first;
  if (first is! Map<String, Object?>) {
    return null;
  }
  final id = _string(first['id']);
  if (id != null && id.isNotEmpty) {
    return _titleWords(id);
  }
  final title = _string(first['title']);
  if (title != null && title.isNotEmpty) {
    return title;
  }
  return null;
}

String _titleWords(String value) {
  return value
      .split(RegExp(r'[_\-\s]+'))
      .where((part) => part.isNotEmpty)
      .map(_title)
      .join(' ');
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
  final teamCount = _int(linear['team_count']);
  if (teamCount != null) {
    items.add('Teams: $teamCount');
  }
  final viewerConfigured = _bool(linear['viewer_configured']);
  if (viewerConfigured != null) {
    items.add('Viewer configured: $viewerConfigured');
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

List<RecommendationInsight> _recommendations(Map<String, Object?> domain) {
  final recommendations = _object(domain['recommendations']);
  if (recommendations == null) {
    return JarvisSnapshot.sample.recommendations;
  }
  final rawItems = recommendations['items'];
  if (rawItems is! List<Object?>) {
    return JarvisSnapshot.sample.recommendations;
  }
  final items = <RecommendationInsight>[];
  for (final item in rawItems.whereType<Map<String, Object?>>()) {
    final title = _string(item['title']);
    if (title == null || title.isEmpty) {
      continue;
    }
    items.add(
      RecommendationInsight(
        kind: _string(item['kind']) ?? 'review',
        title: title,
        rationale: _string(item['rationale']) ?? '',
        score: _int(item['score']) ?? 0,
        currency: _string(item['currency']) ?? '',
        estimatedMonthlyMinorUnits:
            _int(item['estimated_monthly_minor_units']) ?? 0,
        evidenceCount: _int(item['evidence_count']) ?? 0,
      ),
    );
  }
  return items.isEmpty ? JarvisSnapshot.sample.recommendations : items;
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

FinanceDashboard _financeDashboard(Map<String, Object?> domain) {
  final finance = _object(domain['finance']);
  if (finance == null) {
    return JarvisSnapshot.sample.financeDashboard;
  }
  final owners = <FinanceOwner>[];
  final rawOwners = finance['owner_breakdown'];
  if (rawOwners is List<Object?>) {
    for (final item in rawOwners.whereType<Map<String, Object?>>()) {
      final owner = _string(item['owner']);
      if (owner == null || owner.isEmpty) {
        continue;
      }
      owners.add(
        FinanceOwner(
          owner: owner,
          records: _int(item['records']) ?? 0,
          currency: _string(item['currency']) ?? '',
          creditMinorUnits: _int(item['credit_minor_units']) ?? 0,
          debitMinorUnits: _int(item['debit_minor_units']) ?? 0,
          netMinorUnits: _int(item['net_minor_units']) ?? 0,
        ),
      );
    }
  }
  return FinanceDashboard(
    records: _int(finance['records']) ?? 0,
    currency: _string(finance['currency']) ?? '',
    creditMinorUnits: _int(finance['credit_minor_units']) ?? 0,
    debitMinorUnits: _int(finance['debit_minor_units']) ?? 0,
    netMinorUnits: _int(finance['net_minor_units']) ?? 0,
    subscriptionMinorUnits: _int(finance['subscription_minor_units']) ?? 0,
    subscriptionCount: _int(finance['subscription_count']) ?? 0,
    cardDebitMinorUnits: _int(finance['card_debit_minor_units']) ?? 0,
    cardDebitCount: _int(finance['card_debit_count']) ?? 0,
    categories: _stringList(finance['categories']),
    owners: owners,
  );
}

PurchaseDashboard _purchaseDashboard(Map<String, Object?> domain) {
  final commerce = _object(domain['commerce']);
  if (commerce == null) {
    return JarvisSnapshot.sample.purchaseDashboard;
  }
  final candidates = <RecurringPurchase>[];
  final rawCandidates = commerce['recurring_candidates'];
  if (rawCandidates is List<Object?>) {
    for (final item in rawCandidates.whereType<Map<String, Object?>>()) {
      final merchantName = _string(item['merchant_name']);
      final itemName = _string(item['item_name']);
      if (merchantName == null ||
          merchantName.isEmpty ||
          itemName == null ||
          itemName.isEmpty) {
        continue;
      }
      candidates.add(
        RecurringPurchase(
          merchantName: merchantName,
          itemName: itemName,
          currency: _string(item['currency']) ?? '',
          purchaseCount: _int(item['purchase_count']) ?? 0,
          latestTotalMinorUnits: _int(item['latest_total_minor_units']) ?? 0,
          latestPurchasedAt: _string(item['latest_purchased_at']) ?? '',
        ),
      );
    }
  }

  final owners = <PurchaseOwner>[];
  final rawOwners = commerce['owner_breakdown'];
  if (rawOwners is List<Object?>) {
    for (final item in rawOwners.whereType<Map<String, Object?>>()) {
      final owner = _string(item['owner']);
      if (owner == null || owner.isEmpty) {
        continue;
      }
      owners.add(
        PurchaseOwner(
          owner: owner,
          records: _int(item['records']) ?? 0,
          currency: _string(item['currency']) ?? '',
          purchaseSpendMinorUnits:
              _int(item['purchase_spend_minor_units']) ?? 0,
        ),
      );
    }
  }

  return PurchaseDashboard(
    records: _int(commerce['records']) ?? 0,
    currency: _string(commerce['currency']) ?? '',
    totalSpendMinorUnits: _int(commerce['total_spend_minor_units']) ?? 0,
    recurringCandidateCount: _int(commerce['recurring_candidate_count']) ?? 0,
    recurringCandidates: candidates,
    categories: _stringList(commerce['categories']),
    owners: owners,
  );
}

List<ConnectorReadiness> _connectors(Map<String, Object?> status) {
  final rawConnectors = status['connectors'];
  if (rawConnectors is! List<Object?>) {
    return JarvisSnapshot.sample.connectors;
  }
  final connectors = <ConnectorReadiness>[];
  for (final item in rawConnectors.whereType<Map<String, Object?>>()) {
    final key = _string(item['key']);
    final label = _string(item['label']);
    if (key == null || key.isEmpty || label == null || label.isEmpty) {
      continue;
    }
    connectors.add(
      ConnectorReadiness(
        key: key,
        label: label,
        category: _string(item['category']) ?? '',
        status: _string(item['status']) ?? 'planned',
        fixtureMode: _bool(item['fixture_mode']) ?? true,
        dataClasses: _stringList(item['data_classes']),
        allowedOperations: _stringList(item['allowed_operations']),
        forbiddenOperations: _stringList(item['forbidden_operations']),
        nextStep: _string(item['next_step']) ?? '',
      ),
    );
  }
  return connectors.isEmpty ? JarvisSnapshot.sample.connectors : connectors;
}

List<AgentClusterSignal> _agentClusterSignals(Map<String, Object?> status) {
  final rawSignals = status['signals'];
  if (rawSignals is! List<Object?>) {
    return JarvisSnapshot.sample.agentClusterSignals;
  }
  final signals = <AgentClusterSignal>[];
  for (final item in rawSignals.whereType<Map<String, Object?>>()) {
    final key = _string(item['key']);
    final label = _string(item['label']);
    if (key == null || key.isEmpty || label == null || label.isEmpty) {
      continue;
    }
    signals.add(
      AgentClusterSignal(
        key: key,
        label: label,
        status: _string(item['status']) ?? 'tracked',
        evidence: _string(item['evidence']) ?? '',
      ),
    );
  }
  return signals.isEmpty ? JarvisSnapshot.sample.agentClusterSignals : signals;
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
    case 'open_coupang_play':
    case 'open_netflix':
      return Icons.movie_filter_outlined;
    case 'open_disney_plus':
      return Icons.auto_awesome_outlined;
    case 'open_tving':
      return Icons.live_tv_outlined;
    case 'open_wavve':
      return Icons.waves_outlined;
    case 'open_youtube':
    case 'open_youtube_search':
      return Icons.play_circle_outline;
    case 'open_ott':
    case 'movie_mode':
      return Icons.theaters_outlined;
    case 'volume_set':
    case 'volume_up':
    case 'volume_down':
      return Icons.volume_up_outlined;
    case 'volume_mute':
      return Icons.volume_off_outlined;
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
