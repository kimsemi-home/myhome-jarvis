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
        case '/commands':
          _writeJson(request, [
            {
              'name': 'open_youtube',
              'summary': 'Open YouTube',
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
            'teams': [
              {'name': 'Home Ops'},
            ],
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
        case '/domain/summary':
          _writeJson(request, {
            'finance': {
              'records': 3,
              'currency': 'KRW',
              'net_minor_units': 4346800,
            },
            'commerce': {'records': 3, 'recurring_candidate_count': 1},
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
              'count': 2,
              'items': [
                {
                  'kind': 'recurring_purchase_review',
                  'title': 'Compare recurring purchase: Bottled water',
                  'score': 81,
                },
                {
                  'kind': 'subscription_review',
                  'title': 'Review household subscriptions',
                  'score': 61,
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
        case '/metrics':
          _writeJson(request, {
            'bind_host': '127.0.0.1',
            'requests': 7,
            'event_count': 2,
            'dry_run_default': true,
          });
          return;
        case '/events':
          _writeJson(request, {'count': 2, 'events': <Object?>[]});
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
    expect(snapshot.metrics.map((metric) => metric.value), contains('7'));
    expect(
      snapshot.metrics.singleWhere((metric) => metric.label == 'Events').value,
      '2',
    );
    expect(snapshot.metrics.map((metric) => metric.value), contains('Dirty'));
    expect(
      snapshot.commands.map((command) => command.name),
      contains('open-youtube'),
    );
    expect(
      snapshot.commands.map((command) => command.payload),
      contains('{"level":30}'),
    );
    expect(snapshot.linearItems, contains('Team: Home Ops'));
    expect(snapshot.linearItems, contains('Synced: true'));
    expect(snapshot.storageItems, contains('Finance: 3 transactions'));
    expect(snapshot.storageItems, contains('Finance net: 4346800 KRW'));
    expect(snapshot.storageItems, contains('Commerce: 3 purchases'));
    expect(snapshot.storageItems, contains('Storage: parquet+zstd'));
    expect(
      snapshot.recommendationItems,
      contains('81 - Compare recurring purchase: Bottled water'),
    );
    expect(
      snapshot.recommendationItems,
      contains('61 - Review household subscriptions'),
    );
    expect(snapshot.householdScopes.map((scope) => scope.scope), [
      'user',
      'spouse',
      'household',
    ]);
    expect(snapshot.householdScopes.first.financeNetMinorUnits, -87300);
    expect(snapshot.householdScopes.last.purchaseSpendMinorUnits, 26800);

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
}

void _writeJson(HttpRequest request, Object body) {
  request.response.headers.contentType = ContentType.json;
  request.response.write(jsonEncode(body));
  request.response.close();
}
