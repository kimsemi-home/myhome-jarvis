import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows the consented live finance state', (tester) async {
    final liveFinance = JarvisSnapshot.sample.financeDashboard.copyWith(
      fixtureOnly: false,
    );
    final snapshot = JarvisSnapshot.sample.copyWithFinance(liveFinance);
    await tester.pumpWidget(
      JarvisApp(client: StaticSnapshotClient(snapshot)),
    );

    await openTab(tester, 'Finance');

    expect(find.text('연결됨'), findsOneWidget);
    expect(find.text('서버 live · 읽기 전용 · 정규화 패리티 100%'), findsOneWidget);
    expect(find.text('consented live'), findsOneWidget);
    expect(find.text('fixture-only'), findsNothing);
  });
}
