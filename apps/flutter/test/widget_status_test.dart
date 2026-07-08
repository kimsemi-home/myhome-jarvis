import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows local operations status', (tester) async {
    await pumpJarvis(tester);

    expect(find.text('myhome-jarvis'), findsOneWidget);
    expect(find.text('Assistant Center'), findsOneWidget);
    expect(find.text('2 gated'), findsOneWidget);
    expect(find.text('Next Safe Action'), findsOneWidget);
    expect(find.text('Resolve authority'), findsOneWidget);
    expect(find.text('Daemon'), findsOneWidget);
    expect(find.text('127.0.0.1:3888'), findsOneWidget);
    expect(find.text('Network'), findsOneWidget);
    expect(find.text('Local-only'), findsOneWidget);
    expect(find.text('LAN Auth'), findsOneWidget);
    expect(find.text('Configured'), findsOneWidget);
    expect(find.text('Dry-run'), findsOneWidget);
    expect(find.text('Public Safety'), findsOneWidget);
    expect(find.text('Clear'), findsOneWidget);
    expect(find.text('Code Shape'), findsOneWidget);
    expect(find.text('Tracked'), findsOneWidget);
    expect(find.text('Agent Cluster'), findsOneWidget);
    expect(find.text('5 roles gated'), findsOneWidget);
    expect(find.text('Codex Cost'), findsOneWidget);
    expect(find.text('0 units'), findsOneWidget);
    expect(find.byKey(const Key('status-grid')), findsOneWidget);
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('warning'), findsWidgets);
    expect(find.text('local'), findsWidgets);
    expect(find.text('verified'), findsWidgets);

    await tester.drag(find.byType(GridView).first, const Offset(0, -520));
    await tester.pumpAndSettle();

    expect(find.text('Learning'), findsOneWidget);
    expect(find.text('0 observed'), findsOneWidget);
    expect(find.text('Evidence Graph'), findsOneWidget);
    expect(find.text('0 nodes'), findsOneWidget);
    expect(find.text('Confidence'), findsOneWidget);
    expect(find.text('Local'), findsWidgets);
    expect(find.text('Authority Gate'), findsOneWidget);
    expect(find.text('6 blocked'), findsOneWidget);
    expect(find.text('blocked'), findsWidgets);
    expect(find.text('Review Capacity'), findsOneWidget);
    expect(find.text('Available'), findsOneWidget);
    expect(find.byIcon(Icons.refresh), findsOneWidget);
  });
}
