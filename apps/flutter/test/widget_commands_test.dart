import 'package:flutter/widgets.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

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
      await scrollCommandIntoView(tester, command);
      expect(find.text(command), findsOneWidget);
    }
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('dry-run'), findsWidgets);
    expect(find.text('execute blocked'), findsWidgets);

    for (final field in const [
      ('open-youtube-search', 'query'),
      ('open-url', 'url'),
    ]) {
      await scrollCommandIntoView(tester, field.$1);
      expect(find.text(field.$1), findsOneWidget);
      expect(find.byKey(Key('payload-field-${field.$2}')), findsOneWidget);
      expect(find.text('payload editable'), findsWidgets);
    }

    await scrollCommandIntoView(tester, 'open-ott');
    expect(find.text('open-ott'), findsOneWidget);
    expect(find.text('service'), findsOneWidget);
    expect(find.byKey(const Key('payload-field-service')), findsOneWidget);
    expect(find.byType(ShadSelect<String>), findsWidgets);

    await scrollCommandIntoView(tester, 'volume-set');
    expect(find.text('volume-set'), findsOneWidget);
    expect(find.byKey(const Key('payload-field-level')), findsOneWidget);
    expect(_payloadText(tester, 'level'), '30');
    await scrollCommandIntoView(tester, 'volume-up');
    expect(find.text('volume-up'), findsOneWidget);
    await scrollCommandIntoView(tester, 'volume-down');
    expect(find.text('volume-down'), findsOneWidget);
    expect(find.byType(ShadInput), findsWidgets);

    for (final command in const ['volume-mute', 'display-sleep', 'mac-sleep']) {
      await scrollCommandIntoView(tester, command);
      expect(find.text(command), findsOneWidget);
    }
    expect(find.byTooltip('Dry-run'), findsWidgets);
  });
}

String _payloadText(WidgetTester tester, String field) {
  final input = find.byKey(Key('payload-field-$field'));
  final editable = find.descendant(
    of: input,
    matching: find.byType(EditableText),
  );
  return tester.widget<EditableText>(editable).controller.text;
}
