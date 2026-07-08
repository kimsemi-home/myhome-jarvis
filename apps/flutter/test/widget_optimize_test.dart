import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows optimization tab', (tester) async {
    await pumpJarvis(tester);
    await openTab(tester, 'Optimize');

    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
    expect(
      find.text('Compare recurring purchase: Bottled water 2L x 6'),
      findsOneWidget,
    );
    expect(find.text('review-only'), findsWidgets);
    expect(find.text('warning'), findsOneWidget);
    expect(find.text('scored'), findsWidgets);
    expect(find.text('evidence-backed'), findsWidgets);
    expect(find.text('81'), findsOneWidget);
    expect(find.text('11800 KRW'), findsOneWidget);
    expect(find.text('2 evidence'), findsWidgets);
    expect(find.text('Review card-linked household spend'), findsOneWidget);
    expect(find.text('153200 KRW'), findsOneWidget);
    expect(find.text('Review household subscriptions'), findsOneWidget);
    expect(find.text('65900 KRW'), findsOneWidget);
  });
}
