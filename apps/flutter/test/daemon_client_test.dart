import 'dart:convert';
import 'dart:io';

import 'package:flutter/widgets.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  test('loads a snapshot and dry-run plan from daemon endpoints', () async {
    final server = await HttpServer.bind(InternetAddress.loopbackIPv4, 0);
    addTearDown(() => server.close(force: true));

    server.listen((request) async {
      switch (request.uri.path) {
        case '/health':
          _writeJson(request, {
            'ok': true,
            'dry_run': true,
            'host': '127.0.0.1',
          });
          return;
        case '/auth/status':
          _writeJson(request, {
            'configured': true,
            'path': 'data/private/local-token.txt',
            'mode': '-rw-------',
            'message': 'local LAN token is configured',
          });
          return;
        case '/commands':
          _writeJson(request, [
            {
              'name': 'open_youtube',
              'summary': 'Open YouTube',
              'payload_fields': <String>[],
            },
            {
              'name': 'open_netflix',
              'summary': 'Open Netflix',
              'payload_fields': <String>[],
            },
            {
              'name': 'open_disney_plus',
              'summary': 'Open Disney+',
              'payload_fields': <String>[],
            },
            {
              'name': 'open_tving',
              'summary': 'Open TVING',
              'payload_fields': <String>[],
            },
            {
              'name': 'open_wavve',
              'summary': 'Open Wavve',
              'payload_fields': <String>[],
            },
            {
              'name': 'open_coupang_play',
              'summary': 'Open Coupang Play',
              'payload_fields': <String>[],
            },
            {
              'name': 'volume_set',
              'summary': 'Set output volume',
              'payload_fields': ['level'],
            },
          ]);
          return;
        case '/linear/status':
          _writeJson(request, {
            'mode': 'online',
            'synced': true,
            'queue_path': 'data/private/linear-offline-queue.jsonl',
            'viewer_configured': true,
            'team_count': 1,
          });
          return;
        case '/repo/status':
          _writeJson(request, {
            'branch': 'main',
            'short_sha': 'abc123',
            'worktree_clean': false,
            'tracked_changes': [
              {'code': ' M', 'path': 'README.md'},
            ],
            'untracked_files': ['docs/new.md'],
          });
          return;
        case '/security/status':
          _writeJson(request, {
            'ok': true,
            'current_ok': true,
            'current_finding_count': 0,
            'history_ok': true,
            'history_finding_count': 0,
            'checked_at': '2026-06-15T00:00:00Z',
          });
          return;
        case '/domain/summary':
          _writeJson(request, {
            'finance': {
              'records': 3,
              'currency': 'KRW',
              'credit_minor_units': 4500000,
              'debit_minor_units': 153200,
              'net_minor_units': 4346800,
              'subscription_minor_units': 65900,
              'subscription_count': 1,
              'card_debit_minor_units': 153200,
              'card_debit_count': 2,
              'categories': ['income', 'subscription', 'utilities'],
              'owner_breakdown': [
                {
                  'owner': 'household',
                  'records': 2,
                  'currency': 'KRW',
                  'credit_minor_units': 4500000,
                  'debit_minor_units': 65900,
                  'net_minor_units': 4434100,
                },
                {
                  'owner': 'user',
                  'records': 1,
                  'currency': 'KRW',
                  'credit_minor_units': 0,
                  'debit_minor_units': 87300,
                  'net_minor_units': -87300,
                },
              ],
            },
            'commerce': {
              'records': 3,
              'currency': 'KRW',
              'total_spend_minor_units': 26800,
              'recurring_candidate_count': 1,
              'recurring_candidates': [
                {
                  'merchant_name': 'Coupang',
                  'item_name': 'Bottled water 2L x 6',
                  'currency': 'KRW',
                  'purchase_count': 2,
                  'latest_total_minor_units': 11800,
                  'latest_purchased_at': '2026-06-10',
                },
              ],
              'categories': ['grocery', 'household'],
              'owner_breakdown': [
                {
                  'owner': 'household',
                  'records': 2,
                  'currency': 'KRW',
                  'purchase_spend_minor_units': 23600,
                },
                {
                  'owner': 'user',
                  'records': 1,
                  'currency': 'KRW',
                  'purchase_spend_minor_units': 3200,
                },
              ],
            },
            'household': {
              'scopes': [
                {
                  'scope': 'user',
                  'label': 'User',
                  'currency': 'KRW',
                  'finance_records': 1,
                  'finance_net_minor_units': -87300,
                  'purchase_records': 1,
                  'purchase_spend_minor_units': 3200,
                },
                {
                  'scope': 'spouse',
                  'label': 'Spouse',
                  'currency': 'KRW',
                  'finance_records': 0,
                  'finance_net_minor_units': 0,
                  'purchase_records': 0,
                  'purchase_spend_minor_units': 0,
                },
                {
                  'scope': 'household',
                  'label': 'Household',
                  'currency': 'KRW',
                  'finance_records': 3,
                  'finance_net_minor_units': 4346800,
                  'purchase_records': 3,
                  'purchase_spend_minor_units': 26800,
                },
              ],
            },
            'recommendations': {
              'count': 3,
              'items': [
                {
                  'kind': 'recurring_purchase_review',
                  'title': 'Compare recurring purchase: Bottled water',
                  'rationale': 'Coupang appears repeatedly.',
                  'score': 81,
                  'currency': 'KRW',
                  'estimated_monthly_minor_units': 11800,
                  'evidence_count': 2,
                },
                {
                  'kind': 'card_usage_review',
                  'title': 'Review card-linked household spend',
                  'rationale': 'Card-linked debit fixtures exist.',
                  'score': 67,
                  'currency': 'KRW',
                  'estimated_monthly_minor_units': 153200,
                  'evidence_count': 2,
                },
                {
                  'kind': 'subscription_review',
                  'title': 'Review household subscriptions',
                  'rationale': 'Subscription-like debit fixtures exist.',
                  'score': 61,
                  'currency': 'KRW',
                  'estimated_monthly_minor_units': 65900,
                  'evidence_count': 1,
                },
              ],
            },
            'storage': {
              'datasets': ['finance_transactions', 'commerce_purchases'],
              'lake_layers': ['raw', 'bronze', 'silver', 'gold'],
              'long_term_format': 'parquet',
              'compression': 'zstd',
            },
          });
          return;
        case '/connectors/status':
          _writeJson(request, {
            'fixture_only': true,
            'real_credentials_allowed': false,
            'external_api_calls_allowed': false,
            'connector_count': 2,
            'planned_count': 2,
            'fixture_mode_count': 2,
            'generated_path': 'generated/connectors.generated.json',
            'connectors': [
              {
                'key': 'mydata',
                'label': 'MyData aggregator',
                'category': 'finance_aggregation',
                'status': 'planned',
                'fixture_mode': true,
                'data_classes': ['accounts', 'cards', 'transactions'],
                'allowed_operations': ['read_fixture', 'summarize'],
                'forbidden_operations': [
                  'credential_request',
                  'external_api_call',
                  'transfer',
                ],
                'next_step': 'Define consent and local vault boundaries.',
              },
              {
                'key': 'commerce',
                'label': 'Commerce purchases',
                'category': 'commerce',
                'status': 'planned',
                'fixture_mode': true,
                'data_classes': ['orders', 'items'],
                'allowed_operations': [
                  'read_fixture',
                  'recommend_review',
                  'summarize',
                ],
                'forbidden_operations': ['scraping', 'purchase'],
                'next_step': 'Extend local purchase fixtures.',
              },
            ],
          });
          return;
        case '/agent-cluster/status':
          _writeJson(request, {
            'context': 'AgentCluster',
            'version': 'v1',
            'public_safe': true,
            'external_agent_execution_allowed': false,
            'self_approval_allowed': false,
            'authority_gate_required': true,
            'evidence_stage_count': 10,
            'role_count': 5,
            'sidecar_count': 6,
            'generated_path': 'generated/agent_cluster.generated.json',
            'signals': [
              {
                'key': 'evidence_first',
                'label': 'Evidence first',
                'status': 'active',
                'evidence': 'observation and evidence precede code',
              },
              {
                'key': 'authority_gated',
                'label': 'Authority gated',
                'status': 'gated',
                'evidence':
                    'producer, reviewer, verifier, and steward roles are separated',
              },
            ],
          });
          return;
        case '/learning/status':
          _writeJson(request, {
            'path': 'data/private/learning/observations.jsonl',
            'policy_path': 'generated/learning.generated.json',
            'exists': true,
            'count': 3,
            'open_count': 1,
            'closed_count': 2,
            'by_kind': {'loop_gap': 1, 'evidence_debt': 2},
            'by_stage': {'evidence_recorded': 1, 'knowledge_updated': 2},
            'last_kind': 'loop_gap',
            'last_stage': 'evidence_recorded',
            'last_status': 'open',
            'last_observed_at': '2026-06-18T00:00:00Z',
          });
          return;
        case '/evidence/status':
          _writeJson(request, {
            'policy_path': 'generated/evidence.generated.json',
            'private_root': 'data/private',
            'source_count': 5,
            'present_source_count': 2,
            'node_count': 4,
            'edge_count': 2,
            'dangling_evidence_ref_count': 0,
            'open_learning_count': 1,
            'by_node_kind': {
              'learning_observation': 1,
              'evidence_artifact': 2,
              'quality_run': 1,
            },
            'by_edge_kind': {'supports': 2},
            'sources': [
              {
                'key': 'learning',
                'node_kind': 'learning_observation',
                'format': 'jsonl',
                'present': true,
                'count': 1,
              },
            ],
          });
          return;
        case '/metrics':
          _writeJson(request, {
            'bind_host': '127.0.0.1',
            'requests': 7,
            'event_count': 2,
            'goroutine_count': 9,
            'heap_alloc_bytes': 1048576,
            'heap_sys_bytes': 2097152,
            'stack_inuse_bytes': 32768,
            'gc_count': 1,
            'dry_run_default': true,
            'execute_enabled': false,
            'lan_bind_allowed': false,
          });
          return;
        case '/events':
          _writeJson(request, {'count': 2, 'events': <Object?>[]});
          return;
        case '/supervisor/status':
          _writeJson(request, {
            'recorded': true,
            'stale': false,
            'message': 'daemon is reachable',
            'state_path': 'data/private/supervisor/daemon-state.json',
          });
          return;
        case '/audit/status':
          _writeJson(request, {
            'path': 'data/private/audit/command-intents.jsonl',
            'exists': true,
            'count': 4,
          });
          return;
        case '/quality/status':
          _writeJson(request, {
            'path': 'data/private/quality/runs.jsonl',
            'exists': true,
            'count': 3,
            'last': {
              'ok': true,
              'step_count': 12,
              'pass_count': 12,
              'fail_count': 0,
              'skip_count': 0,
            },
          });
          return;
        case '/planner/status':
          _writeJson(request, {
            'loop_mode': 'closed-loop',
            'task_count': 6,
            'ready_count': 0,
            'completed_count': 5,
            'blocked_external_write_count': 1,
            'blocked_external_write_tasks': [
              {
                'id': 'linear_sync',
                'title':
                    'Sync Linear only after explicit external-write approval',
                'owner': 'go',
                'status': 'blocked_external_write',
                'depends_on': ['quality_gate'],
              },
            ],
            'linear_template_count': 2,
            'quality_required': true,
            'linear_offline_fallback': true,
            'checkpoint_root': 'data/private/checkpoints',
          });
          return;
        case '/intent':
          final body = jsonDecode(await utf8.decoder.bind(request).join());
          expect(body, isA<Map<String, Object?>>());
          final payload = body as Map<String, Object?>;
          expect(payload['command'], 'volume-set');
          expect(payload['payload'], {'level': 30});
          expect(payload['execute'], false);
          _writeJson(request, {
            'name': 'volume_set',
            'dry_run': true,
            'execute_allowed': false,
            'invocations': [
              {
                'label': 'volume_set',
                'argv': ['osascript', '-e', 'set volume output volume 30'],
              },
            ],
          });
          return;
        default:
          request.response.statusCode = HttpStatus.notFound;
          request.response.close();
          return;
      }
    });

    final client = DaemonSnapshotClient(
      baseUri: Uri.parse('http://${server.address.address}:${server.port}'),
    );
    final snapshot = await client.load();

    expect(
      snapshot.metrics.map((metric) => metric.value),
      contains('127.0.0.1'),
    );
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Network').value,
      'Local-only',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'LAN Auth')
          .value,
      'Configured',
    );
    expect(snapshot.metrics.map((metric) => metric.value), contains('7'));
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Quality').value,
      'Passing (3)',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Public Safety')
          .value,
      'Clear',
    );
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Events').value,
      '2',
    );
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Runtime').value,
      '9 goroutines',
    );
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Heap').value,
      '1.0 MiB',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Supervisor')
          .value,
      'Reachable',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Command Audit')
          .value,
      '4',
    );
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Planner').value,
      '5/6 done, 1 gated',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Planner Gate')
          .value,
      'Linear Sync',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Connectors')
          .value,
      '2/2 fixture',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Agent Cluster')
          .value,
      '5 roles gated',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Learning')
          .value,
      '1 open',
    );
    expect(
      snapshot.metrics
          .singleWhere((metric) => metric.label == 'Evidence Graph')
          .value,
      '2 links',
    );
    expect(snapshot.metrics.map((metric) => metric.value), contains('Dirty'));
    expect(
      snapshot.commands.map((command) => command.name),
      contains('open-youtube'),
    );
    expect(
      snapshot.commands.map((command) => command.name),
      containsAll([
        'open-netflix',
        'open-disney-plus',
        'open-tving',
        'open-wavve',
        'open-coupang-play',
      ]),
    );
    expect(
      snapshot.commands.map((command) => command.payload),
      contains('{"level":30}'),
    );
    expect(snapshot.linearItems, contains('Teams: 1'));
    expect(snapshot.linearItems, contains('Viewer configured: true'));
    expect(snapshot.linearItems, contains('Synced: true'));
    expect(snapshot.storageItems, contains('Finance: 3 transactions'));
    expect(snapshot.storageItems, contains('Finance net: 4346800 KRW'));
    expect(snapshot.storageItems, contains('Commerce: 3 purchases'));
    expect(snapshot.storageItems, contains('Storage: parquet+zstd'));
    expect(snapshot.purchaseDashboard.totalSpendMinorUnits, 26800);
    expect(snapshot.purchaseDashboard.recurringCandidateCount, 1);
    expect(snapshot.purchaseDashboard.categories, contains('grocery'));
    expect(snapshot.purchaseDashboard.owners.map((owner) => owner.owner), [
      'household',
      'user',
    ]);
    expect(
      snapshot.purchaseDashboard.recurringCandidates.single.itemName,
      'Bottled water 2L x 6',
    );
    expect(
      snapshot
          .purchaseDashboard
          .recurringCandidates
          .single
          .latestTotalMinorUnits,
      11800,
    );
    expect(snapshot.financeDashboard.netMinorUnits, 4346800);
    expect(snapshot.financeDashboard.subscriptionMinorUnits, 65900);
    expect(snapshot.financeDashboard.cardDebitCount, 2);
    expect(snapshot.financeDashboard.categories, contains('subscription'));
    expect(snapshot.financeDashboard.owners.map((owner) => owner.owner), [
      'household',
      'user',
    ]);
    expect(snapshot.financeDashboard.owners.first.netMinorUnits, 4434100);
    expect(
      snapshot.recommendationItems,
      contains('81 - Compare recurring purchase: Bottled water'),
    );
    expect(
      snapshot.recommendationItems,
      contains('67 - Review card-linked household spend'),
    );
    expect(
      snapshot.recommendationItems,
      contains('61 - Review household subscriptions'),
    );
    expect(snapshot.recommendations.map((item) => item.kind), [
      'recurring_purchase_review',
      'card_usage_review',
      'subscription_review',
    ]);
    expect(snapshot.recommendations.first.score, 81);
    expect(snapshot.recommendations.first.estimatedMonthlyMinorUnits, 11800);
    expect(snapshot.recommendations.first.evidenceCount, 2);
    expect(snapshot.recommendations[1].rationale, contains('Card-linked'));
    expect(snapshot.householdScopes.map((scope) => scope.scope), [
      'user',
      'spouse',
      'household',
    ]);
    expect(snapshot.householdScopes.first.financeNetMinorUnits, -87300);
    expect(snapshot.householdScopes.last.purchaseSpendMinorUnits, 26800);
    expect(snapshot.connectors.map((connector) => connector.key), [
      'mydata',
      'commerce',
    ]);
    expect(snapshot.connectors.first.fixtureMode, isTrue);
    expect(
      snapshot.connectors.first.forbiddenOperations,
      contains('external_api_call'),
    );
    expect(snapshot.agentClusterSignals.map((signal) => signal.key), [
      'evidence_first',
      'authority_gated',
    ]);
    expect(snapshot.agentClusterSignals.first.status, 'active');

    final command = snapshot.commands.singleWhere(
      (item) => item.name == 'volume-set',
    );
    expect(command.payloadFields, ['level']);
    final plan = await client.dryRun(command);
    expect(plan.name, 'volume_set');
    expect(plan.dryRun, isTrue);
    expect(plan.executeAllowed, isFalse);
    expect(
      plan.invocations.single.argv,
      contains('set volume output volume 30'),
    );
  });

  test('sends bearer token when configured', () async {
    final server = await HttpServer.bind(InternetAddress.loopbackIPv4, 0);
    addTearDown(() => server.close(force: true));

    server.listen((request) async {
      expect(
        request.headers.value(HttpHeaders.authorizationHeader),
        'Bearer household-token',
      );
      expect(request.uri.path, '/intent');
      _writeJson(request, {
        'name': 'volume_set',
        'dry_run': true,
        'execute_allowed': false,
        'invocations': <Object?>[],
      });
    });

    final client = DaemonSnapshotClient(
      baseUri: Uri.parse('http://${server.address.address}:${server.port}'),
      authToken: 'household-token',
    );
    final plan = await client.dryRun(
      const HomeCommand(
        name: 'volume-set',
        payload: '{"level":30}',
        icon: IconData(0),
      ),
    );

    expect(plan.name, 'volume_set');
  });

  test('marks LAN daemon mode as token-gated', () {
    final snapshot = buildSnapshot(
      health: <String, Object?>{
        'mode': 'local',
        'dry_run': true,
        'host': '192.168.1.10',
      },
      auth: const <String, Object?>{'configured': true},
      commands: const <Object?>[],
      linear: const <String, Object?>{},
      repo: const <String, Object?>{},
      security: const <String, Object?>{'ok': true},
      domain: const <String, Object?>{},
      connectors: const <String, Object?>{},
      agentCluster: const <String, Object?>{},
      learning: const <String, Object?>{},
      evidence: const <String, Object?>{},
      metrics: <String, Object?>{
        'bind_host': '192.168.1.10',
        'dry_run_default': true,
        'execute_enabled': false,
        'lan_bind_allowed': true,
      },
      events: const <String, Object?>{},
      supervisor: const <String, Object?>{},
      audit: const <String, Object?>{},
      quality: const <String, Object?>{},
      planner: const <String, Object?>{},
    );

    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Network').value,
      'LAN token-gated',
    );
  });
}

void _writeJson(HttpRequest request, Object body) {
  request.response.headers.contentType = ContentType.json;
  request.response.write(jsonEncode(body));
  request.response.close();
}
