import 'package:flutter/widgets.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  test('external evidence connector mirrors the public upstream fixture', () {
    final connector = JarvisSnapshot.sample.connectors.singleWhere(
      (item) => item.key == 'external-evidence-lake',
    );

    expect(connector.label, 'External evidence lake');
    expect(connector.category, 'public_evidence_boundary');
    expect(connector.status, 'bootstrap');
    expect(connector.fixtureMode, isTrue);
    expect(connector.dataClasses, [
      'context_pack',
      'ui_status_metadata',
      'validation_summary',
    ]);
    expect(connector.allowedOperations, [
      'read_public_fixture',
      'show_status',
      'link_upstream',
    ]);
    expect(connector.forbiddenOperations, [
      'raw_payload_import',
      'credential_request',
      'private_archive',
      'collector_write',
    ]);
    expect(
      connector.nextStep,
      'Render public status only from the evidence-lake UI metadata fixture.',
    );
  });

  testWidgets('connectors tab is the read-only shadcn pilot', (tester) async {
    await pumpJarvis(tester);
    await openTab(tester, 'Connectors');

    expect(find.byKey(const Key('connectors-list')), findsOneWidget);
    expect(find.byType(ConnectorTile), findsWidgets);
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('fixture-only'), findsWidgets);

    await scrollConnectorIntoView(tester, 'External evidence lake');

    expect(find.text('External evidence lake'), findsOneWidget);
    expect(find.text('public_evidence_boundary'), findsOneWidget);
    expect(find.text('context_pack'), findsOneWidget);
    expect(find.text('ui_status_metadata'), findsOneWidget);
    expect(find.text('validation_summary'), findsOneWidget);
    expect(find.textContaining('read_public_fixture'), findsOneWidget);
    expect(find.textContaining('show_status'), findsOneWidget);
    expect(find.textContaining('link_upstream'), findsOneWidget);
    expect(find.textContaining('raw_payload_import'), findsOneWidget);
    expect(find.textContaining('credential_request'), findsWidgets);
    expect(find.textContaining('private_archive'), findsOneWidget);
    expect(find.textContaining('collector_write'), findsOneWidget);
  });
}
