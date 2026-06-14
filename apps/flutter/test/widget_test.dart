import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';

void main() {
  testWidgets('shows local operations status', (tester) async {
    await tester.pumpWidget(const JarvisApp());

    expect(find.text('myhome-jarvis'), findsOneWidget);
    expect(find.text('Daemon'), findsOneWidget);
    expect(find.text('127.0.0.1:3888'), findsOneWidget);
    expect(find.text('Network'), findsOneWidget);
    expect(find.text('Local-only'), findsOneWidget);
    expect(find.text('LAN Auth'), findsOneWidget);
    expect(find.text('Configured'), findsOneWidget);
    expect(find.text('Dry-run'), findsOneWidget);
    expect(find.byIcon(Icons.refresh), findsOneWidget);
  });

  testWidgets('shows dry-run home commands', (tester) async {
    await tester.pumpWidget(const JarvisApp());
    await tester.tap(find.text('Commands'));
    await tester.pumpAndSettle();

    expect(find.text('open-youtube'), findsOneWidget);
    expect(find.text('open-netflix'), findsOneWidget);
    expect(find.text('open-disney-plus'), findsOneWidget);
    expect(find.text('open-tving'), findsOneWidget);
    expect(find.text('open-wavve'), findsOneWidget);
    expect(find.text('open-coupang-play'), findsOneWidget);
    await _scrollCommandIntoView(tester, 'open-youtube-search');
    expect(find.text('open-youtube-search'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'query'), findsOneWidget);
    await _scrollCommandIntoView(tester, 'open-url');
    expect(find.text('open-url'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'url'), findsOneWidget);
    await _scrollCommandIntoView(tester, 'open-ott');
    expect(find.text('open-ott'), findsOneWidget);
    expect(find.text('service'), findsOneWidget);

    await _scrollCommandIntoView(tester, 'volume-set');
    expect(find.text('volume-set'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'level'), findsOneWidget);
    expect(find.text('30'), findsOneWidget);
    await _scrollCommandIntoView(tester, 'volume-up');
    expect(find.text('volume-up'), findsOneWidget);
    await _scrollCommandIntoView(tester, 'volume-down');
    expect(find.text('volume-down'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'step'), findsWidgets);
    await _scrollCommandIntoView(tester, 'display-sleep');
    expect(find.text('display-sleep'), findsOneWidget);
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

    await _scrollCommandIntoView(tester, 'volume-set');

    await tester.enterText(find.widgetWithText(TextField, 'level'), '42');
    final volumeSetRow = find.ancestor(
      of: find.text('volume-set'),
      matching: find.byType(CommandRow),
    );
    await tester.tap(
      find.descendant(of: volumeSetRow, matching: find.byTooltip('Dry-run')),
    );
    await tester.pumpAndSettle();

    expect(
      find.textContaining('mhj command volume-set {"level":42}'),
      findsOneWidget,
    );
  });

  testWidgets(
    'shows finance, purchases, Linear, storage, household, and optimization tabs',
    (tester) async {
      await tester.pumpWidget(const JarvisApp());

      await tester.tap(
        find.descendant(
          of: find.byType(TabBar),
          matching: find.text('Finance'),
        ),
      );
      await tester.pumpAndSettle();
      expect(find.text('Net'), findsOneWidget);
      expect(find.text('4346800 KRW'), findsOneWidget);
      expect(find.text('Subscriptions'), findsOneWidget);
      expect(find.text('1 / 65900 KRW'), findsOneWidget);
      expect(find.text('Card-linked'), findsOneWidget);
      expect(find.text('2 / 153200 KRW'), findsOneWidget);
      expect(find.text('Owner Breakdown'), findsOneWidget);
      expect(find.text('Household net'), findsOneWidget);
      expect(find.text('Categories'), findsOneWidget);

      await tester.tap(
        find.descendant(
          of: find.byType(TabBar),
          matching: find.text('Purchases'),
        ),
      );
      await tester.pumpAndSettle();
      expect(find.text('Spend'), findsOneWidget);
      expect(find.text('26800 KRW'), findsOneWidget);
      expect(find.text('Recurring Candidates'), findsOneWidget);
      expect(find.text('Bottled water 2L x 6'), findsOneWidget);
      expect(find.text('Coupang / 2 purchases / 2026-06-10'), findsOneWidget);
      expect(find.text('Owner Spend'), findsOneWidget);
      expect(find.text('Household spend'), findsOneWidget);

      await tester.tap(
        find.descendant(of: find.byType(TabBar), matching: find.text('Linear')),
      );
      await tester.pumpAndSettle();
      expect(find.text('Active queue'), findsOneWidget);
      expect(find.text('Next issue'), findsOneWidget);

      await tester.tap(
        find.descendant(
          of: find.byType(TabBar),
          matching: find.text('Storage'),
        ),
      );
      await tester.pumpAndSettle();
      expect(find.text('finance_transactions'), findsOneWidget);
      expect(find.text('commerce_purchases'), findsOneWidget);

      await tester.tap(
        find.descendant(
          of: find.byType(TabBar),
          matching: find.text('Household'),
        ),
      );
      await tester.pumpAndSettle();
      expect(find.text('Finance net: -87300 KRW'), findsOneWidget);
      expect(find.text('Purchase spend: 3200 KRW'), findsOneWidget);

      await tester.tap(
        find.descendant(
          of: find.byType(SegmentedButton<String>),
          matching: find.text('Household'),
        ),
      );
      await tester.pumpAndSettle();
      expect(find.text('Finance net: 4346800 KRW'), findsOneWidget);
      expect(find.text('Purchase spend: 26800 KRW'), findsOneWidget);

      await tester.tap(
        find.descendant(
          of: find.byType(TabBar),
          matching: find.text('Optimize'),
        ),
      );
      await tester.pumpAndSettle();
      expect(
        find.text('Compare recurring purchase: Bottled water 2L x 6'),
        findsOneWidget,
      );
      expect(find.text('81'), findsOneWidget);
      expect(find.text('11800 KRW'), findsOneWidget);
      expect(find.text('2 evidence'), findsWidgets);
      expect(find.text('Review card-linked household spend'), findsOneWidget);
      expect(find.text('153200 KRW'), findsOneWidget);
      expect(find.text('Review household subscriptions'), findsOneWidget);
      expect(find.text('65900 KRW'), findsOneWidget);
    },
  );
}

Future<void> _scrollCommandIntoView(WidgetTester tester, String name) async {
  await tester.scrollUntilVisible(
    find.text(name),
    240,
    scrollable: find.descendant(
      of: find.byType(CommandsView),
      matching: find.byWidgetPredicate(
        (widget) =>
            widget is Scrollable && widget.axisDirection == AxisDirection.down,
      ),
    ),
  );
  await tester.pumpAndSettle();
}
