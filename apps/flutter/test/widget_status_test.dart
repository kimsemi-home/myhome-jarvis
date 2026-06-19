import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows local operations status', (tester) async {
    await pumpJarvis(tester);

    expect(find.text('myhome-jarvis'), findsOneWidget);
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
    expect(find.text('Learning'), findsOneWidget);
    expect(find.text('0 observed'), findsOneWidget);

    await tester.drag(find.byType(GridView).first, const Offset(0, -260));
    await tester.pumpAndSettle();

    expect(find.text('Evidence Graph'), findsOneWidget);
    expect(find.text('0 nodes'), findsOneWidget);
    expect(find.text('Confidence'), findsOneWidget);
    expect(find.text('Local'), findsWidgets);
    expect(find.text('Authority Gate'), findsOneWidget);
    expect(find.text('6 blocked'), findsOneWidget);
    expect(find.text('Review Capacity'), findsOneWidget);
    expect(find.text('Available'), findsOneWidget);
    expect(find.byIcon(Icons.refresh), findsOneWidget);
  });
}
