import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows cluster and household tabs', (tester) async {
    await pumpJarvis(tester);

    await openTab(tester, 'Cluster');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('Evidence first'), findsOneWidget);
    expect(find.text('active'), findsOneWidget);
    expect(find.text('Authority gated'), findsOneWidget);
    expect(find.text('gated'), findsOneWidget);
    expect(find.text('Feedback loop'), findsOneWidget);
    expect(find.text('tracked'), findsOneWidget);

    await openTab(tester, 'Household');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('selected scope'), findsOneWidget);
    expect(find.text('owner scoped'), findsOneWidget);
    expect(find.text('summary-only'), findsOneWidget);
    expect(find.text('Finance net: -87300 KRW'), findsOneWidget);
    expect(find.text('Purchase spend: 3200 KRW'), findsOneWidget);

    await tester.tap(
      find.descendant(
        of: find.byType(SegmentedButton<String>),
        matching: find.text('Spouse'),
      ),
    );
    await tester.pumpAndSettle();

    expect(find.text('empty'), findsOneWidget);
    expect(find.text('Finance net: 0 KRW'), findsOneWidget);
    expect(find.text('Purchase spend: 0 KRW'), findsOneWidget);

    await tester.tap(
      find.descendant(
        of: find.byType(SegmentedButton<String>),
        matching: find.text('Household'),
      ),
    );
    await tester.pumpAndSettle();

    expect(find.text('household scoped'), findsOneWidget);
    expect(find.text('Finance net: 4346800 KRW'), findsOneWidget);
    expect(find.text('Purchase spend: 26800 KRW'), findsOneWidget);
  });
}
