import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows cluster and household tabs', (tester) async {
    await pumpJarvis(tester);

    await openTab(tester, 'Cluster');
    expect(find.text('Evidence first'), findsOneWidget);
    expect(find.text('Authority gated'), findsOneWidget);
    expect(find.text('Feedback loop'), findsOneWidget);

    await openTab(tester, 'Household');
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
  });
}
