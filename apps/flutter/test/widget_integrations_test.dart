import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows Linear, storage, and connector tabs', (tester) async {
    await pumpJarvis(tester);

    await openTab(tester, 'Linear');
    expect(find.text('Active queue'), findsOneWidget);
    expect(find.text('Next issue'), findsOneWidget);
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('queued'), findsWidgets);
    expect(find.text('verified'), findsWidgets);

    await openTab(tester, 'Storage');
    expect(find.text('finance_transactions'), findsOneWidget);
    expect(find.text('commerce_purchases'), findsOneWidget);
    expect(find.text('local fixture'), findsWidgets);
    expect(find.text('format'), findsWidgets);

    await openTab(tester, 'Connectors');
    expect(find.text('MyData aggregator'), findsOneWidget);
    await scrollConnectorIntoView(tester, 'External evidence lake');
    expect(find.text('External evidence lake'), findsOneWidget);
    expect(find.text('public_evidence_boundary'), findsOneWidget);
    expect(find.textContaining('raw_payload_import'), findsOneWidget);
    expect(find.text('fixture-only'), findsWidgets);
    expect(find.textContaining('Allowed:'), findsWidgets);
    expect(find.textContaining('Blocked:'), findsWidgets);
  });
}
