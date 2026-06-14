import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';

void main() {
  testWidgets('shows local operations status', (tester) async {
    await tester.pumpWidget(const JarvisApp());

    expect(find.text('myhome-jarvis'), findsOneWidget);
    expect(find.text('Daemon'), findsOneWidget);
    expect(find.text('127.0.0.1:3888'), findsOneWidget);
    expect(find.text('Dry-run'), findsOneWidget);
    expect(find.byIcon(Icons.refresh), findsOneWidget);
  });

  testWidgets('shows dry-run home commands', (tester) async {
    await tester.pumpWidget(const JarvisApp());
    await tester.tap(find.text('Commands'));
    await tester.pumpAndSettle();

    expect(find.text('open-youtube'), findsOneWidget);
    expect(find.text('volume-set'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'level'), findsOneWidget);
    expect(find.text('30'), findsOneWidget);
    expect(find.byTooltip('Dry-run'), findsWidgets);
  });

  testWidgets('shows dry-run preview dialog', (tester) async {
    await tester.pumpWidget(const JarvisApp());
    await tester.tap(find.text('Commands'));
    await tester.pumpAndSettle();

    await tester.tap(find.byTooltip('Dry-run').first);
    await tester.pumpAndSettle();

    expect(find.text('Dry-run plan'), findsOneWidget);
    expect(find.textContaining('mhj command open-youtube'), findsOneWidget);
    expect(find.text('Close'), findsOneWidget);
  });

  testWidgets('edits command payload before dry-run preview', (tester) async {
    await tester.pumpWidget(const JarvisApp());
    await tester.tap(find.text('Commands'));
    await tester.pumpAndSettle();

    await tester.enterText(find.widgetWithText(TextField, 'level'), '42');
    await tester.tap(find.byTooltip('Dry-run').at(1));
    await tester.pumpAndSettle();

    expect(
      find.textContaining('mhj command volume-set {"level":42}'),
      findsOneWidget,
    );
  });

  testWidgets('shows Linear and storage tabs', (tester) async {
    await tester.pumpWidget(const JarvisApp());

    await tester.tap(
      find.descendant(of: find.byType(TabBar), matching: find.text('Linear')),
    );
    await tester.pumpAndSettle();
    expect(find.text('Active queue'), findsOneWidget);
    expect(find.text('Next issue'), findsOneWidget);

    await tester.tap(
      find.descendant(of: find.byType(TabBar), matching: find.text('Storage')),
    );
    await tester.pumpAndSettle();
    expect(find.text('finance_transactions'), findsOneWidget);
    expect(find.text('commerce_purchases'), findsOneWidget);
  });
}
