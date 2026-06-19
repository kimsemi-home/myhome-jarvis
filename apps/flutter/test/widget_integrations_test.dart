import 'package:flutter_test/flutter_test.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('shows Linear, storage, and connector tabs', (tester) async {
    await pumpJarvis(tester);

    await openTab(tester, 'Linear');
    expect(find.text('Active queue'), findsOneWidget);
    expect(find.text('Next issue'), findsOneWidget);

    await openTab(tester, 'Storage');
    expect(find.text('finance_transactions'), findsOneWidget);
    expect(find.text('commerce_purchases'), findsOneWidget);

    await openTab(tester, 'Connectors');
    expect(find.text('MyData aggregator'), findsOneWidget);
    expect(find.text('fixture-only'), findsWidgets);
    expect(find.textContaining('Allowed:'), findsWidgets);
    expect(find.textContaining('Blocked:'), findsWidgets);
  });
}
