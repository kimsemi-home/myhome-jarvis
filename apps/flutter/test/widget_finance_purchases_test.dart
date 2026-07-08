import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows finance and purchase tabs', (tester) async {
    await pumpJarvis(tester);

    await openTab(tester, 'Finance');
    expect(find.byType(FinanceMetricTile), findsNWidgets(6));
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('summary-only'), findsOneWidget);
    expect(find.text('fixture-only'), findsOneWidget);
    expect(find.text('verified metadata'), findsOneWidget);
    expect(find.text('card-linked review'), findsOneWidget);
    expect(find.text('household scoped'), findsOneWidget);
    expect(find.text('owner scoped'), findsOneWidget);
    expect(find.text('Net'), findsOneWidget);
    expect(find.text('4346800 KRW'), findsOneWidget);
    expect(find.text('Subscriptions'), findsOneWidget);
    expect(find.text('1 / 65900 KRW'), findsOneWidget);
    expect(find.text('Card-linked'), findsOneWidget);
    expect(find.text('2 / 153200 KRW'), findsOneWidget);
    expect(find.text('Owner Breakdown'), findsOneWidget);
    expect(find.text('Household net'), findsOneWidget);
    await tester.drag(find.byType(ListView), const Offset(0, -300));
    await tester.pumpAndSettle();
    expect(find.text('Categories'), findsOneWidget);

    await openTab(tester, 'Purchases');
    expect(find.byType(FinanceMetricTile), findsNWidgets(3));
    expect(find.byType(RecurringCandidateTile), findsOneWidget);
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('summary-only'), findsOneWidget);
    expect(find.text('fixture-only'), findsOneWidget);
    expect(find.text('verified metadata'), findsOneWidget);
    expect(find.text('recurring candidate'), findsOneWidget);
    expect(find.text('household scoped'), findsOneWidget);
    expect(find.text('owner scoped'), findsOneWidget);
    expect(find.text('Spend'), findsOneWidget);
    expect(find.text('26800 KRW'), findsOneWidget);
    expect(find.text('Recurring Candidates'), findsOneWidget);
    expect(find.text('Bottled water 2L x 6'), findsOneWidget);
    expect(find.text('Coupang / 2 purchases / 2026-06-10'), findsOneWidget);
    expect(find.text('Owner Spend'), findsOneWidget);
    expect(find.text('Household spend'), findsOneWidget);
  });
}
