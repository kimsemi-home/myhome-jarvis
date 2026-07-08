import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

import 'widget_helpers.dart';

void main() {
  testWidgets('mounts shadcn theme and migrated surfaces', (tester) async {
    await pumpJarvis(tester);

    expect(find.byType(ShadApp), findsOneWidget);
    expect(find.byType(ShadAppBuilder), findsOneWidget);
    expect(find.byType(MaterialApp), findsOneWidget);
    expect(find.byType(Scaffold), findsOneWidget);
    expect(find.byType(AppBar), findsOneWidget);
    expect(find.byType(TabBar), findsOneWidget);
    expect(find.byType(TabBarView), findsOneWidget);
    expect(find.byType(JarvisIconAction), findsOneWidget);
    expect(find.byType(JarvisSurface), findsWidgets);

    final context = tester.element(find.byType(MaterialApp));
    final shad = ShadTheme.of(context);
    expect(shad.colorScheme.background, JarvisAstryxTokens.backgroundBody);
    expect(shad.colorScheme.foreground, JarvisAstryxTokens.textPrimary);
    expect(shad.colorScheme.primary, JarvisAstryxTokens.accent);
    expect(shad.radius, const BorderRadius.all(JarvisAstryxTokens.radius));
    for (final label in [
      'Status',
      'Commands',
      'Finance',
      'Purchases',
      'Linear',
      'Storage',
      'Connectors',
      'Cluster',
      'Household',
      'Optimize',
    ]) {
      expect(find.text(label), findsWidgets);
    }

    await openTab(tester, 'Commands');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadIconButton), findsWidgets);
    await scrollCommandIntoView(tester, 'open-ott');
    expect(find.byType(ShadSelect<String>), findsWidgets);

    await openTab(tester, 'Finance');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);

    await openTab(tester, 'Purchases');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);

    await openTab(tester, 'Linear');
    expect(find.byType(JarvisSurface), findsWidgets);

    await openTab(tester, 'Cluster');
    expect(find.byType(ShadBadge), findsWidgets);

    await openTab(tester, 'Household');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);

    await openTab(tester, 'Connectors');
    expect(find.byType(ShadBadge), findsWidgets);
    expect(find.text('fixture-only'), findsWidgets);

    await openTab(tester, 'Optimize');
    expect(find.byType(JarvisSurface), findsWidgets);
    expect(find.byType(ShadBadge), findsWidgets);
  });
}
