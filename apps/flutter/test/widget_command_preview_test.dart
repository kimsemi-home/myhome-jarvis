import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows dry-run preview dialog', (tester) async {
    await openCommandsTab(tester);

    await tester.tap(find.byTooltip('Dry-run').first);
    await tester.pumpAndSettle();

    expect(find.text('Dry-run plan'), findsOneWidget);
    expect(find.textContaining('mhj command open-youtube'), findsOneWidget);
    expect(find.text('Close'), findsOneWidget);
  });

  testWidgets('edits command payload before dry-run preview', (tester) async {
    await openCommandsTab(tester);
    await scrollCommandIntoView(tester, 'volume-set');

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
}
