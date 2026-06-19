import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows dry-run home commands', (tester) async {
    await openCommandsTab(tester);

    for (final command in const [
      'open-youtube',
      'open-netflix',
      'open-disney-plus',
      'open-tving',
      'open-wavve',
      'open-coupang-play',
    ]) {
      expect(find.text(command), findsOneWidget);
    }

    for (final field in const [
      ('open-youtube-search', 'query'),
      ('open-url', 'url'),
    ]) {
      await scrollCommandIntoView(tester, field.$1);
      expect(find.text(field.$1), findsOneWidget);
      expect(find.widgetWithText(TextField, field.$2), findsOneWidget);
    }

    await scrollCommandIntoView(tester, 'open-ott');
    expect(find.text('open-ott'), findsOneWidget);
    expect(find.text('service'), findsOneWidget);

    await scrollCommandIntoView(tester, 'volume-set');
    expect(find.text('volume-set'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'level'), findsOneWidget);
    expect(find.text('30'), findsOneWidget);
    await scrollCommandIntoView(tester, 'volume-up');
    expect(find.text('volume-up'), findsOneWidget);
    await scrollCommandIntoView(tester, 'volume-down');
    expect(find.text('volume-down'), findsOneWidget);
    expect(find.widgetWithText(TextField, 'step'), findsWidgets);

    for (final command in const ['volume-mute', 'display-sleep', 'mac-sleep']) {
      await scrollCommandIntoView(tester, command);
      expect(find.text(command), findsOneWidget);
    }
    expect(find.byTooltip('Dry-run'), findsWidgets);
  });
}
