import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';

Future<void> pumpJarvis(WidgetTester tester) async {
  await tester.pumpWidget(const JarvisApp());
}

Future<void> openCommandsTab(WidgetTester tester) async {
  await pumpJarvis(tester);
  await tester.tap(find.text('Commands'));
  await tester.pumpAndSettle();
}

Future<void> openTab(WidgetTester tester, String label) async {
  final tab = find.descendant(
    of: find.byType(TabBar),
    matching: find.text(label),
  );
  await tester.ensureVisible(tab);
  await tester.pumpAndSettle();
  await tester.tap(tab);
  await tester.pumpAndSettle();
}

Future<void> scrollCommandIntoView(WidgetTester tester, String name) async {
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
