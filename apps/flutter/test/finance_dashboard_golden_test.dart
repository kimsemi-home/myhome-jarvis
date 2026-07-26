import 'dart:ui';
import 'dart:io';

import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';

import 'widget_helpers.dart';

void main() {
  test('root README references the generated finance golden', () {
    final readme = File('../../README.md').readAsStringSync();
    expect(readme, contains('apps/flutter/test/golden/finance_dashboard.png'));
  });

  testWidgets('finance dashboard matches the README golden', (tester) async {
    await tester.binding.setSurfaceSize(const Size(430, 932));
    addTearDown(() => tester.binding.setSurfaceSize(null));

    await pumpJarvis(tester);
    await openTab(tester, 'Finance');

    await expectLater(
      find.byType(JarvisScaffold),
      matchesGoldenFile('golden/finance_dashboard.png'),
    );
  });
}
