import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/main.dart';
import 'package:shadcn_ui/shadcn_ui.dart';

void main() {
  testWidgets('maps JarvisBadge semantic tones to Astryx colors', (
    tester,
  ) async {
    await tester.pumpWidget(
      ShadApp.custom(
        theme: JarvisShadTheme.light,
        appBuilder: (context) {
          return MaterialApp(
            theme: JarvisShadTheme.material(context),
            home: const Scaffold(
              body: Wrap(
                children: [
                  JarvisBadge('ok', tone: JarvisBadgeTone.success),
                  JarvisBadge('review', tone: JarvisBadgeTone.warning),
                  JarvisBadge('risk', tone: JarvisBadgeTone.destructive),
                  JarvisBadge('quiet', tone: JarvisBadgeTone.muted),
                ],
              ),
            ),
            builder: (_, child) => ShadAppBuilder(child: child),
          );
        },
      ),
    );

    ShadBadge badgeFor(String label) {
      return tester.widget<ShadBadge>(
        find.ancestor(of: find.text(label), matching: find.byType(ShadBadge)),
      );
    }

    expect(badgeFor('ok').variant, ShadBadgeVariant.secondary);
    expect(badgeFor('ok').backgroundColor, JarvisAstryxTokens.successMuted);
    expect(badgeFor('ok').foregroundColor, JarvisAstryxTokens.success);
    expect(find.bySemanticsLabel('success badge: ok'), findsOneWidget);

    expect(badgeFor('review').backgroundColor, JarvisAstryxTokens.warningMuted);
    expect(badgeFor('review').foregroundColor, JarvisAstryxTokens.warning);
    expect(find.bySemanticsLabel('warning badge: review'), findsOneWidget);

    expect(badgeFor('risk').variant, ShadBadgeVariant.destructive);
    expect(badgeFor('risk').backgroundColor, JarvisAstryxTokens.errorMuted);
    expect(badgeFor('risk').foregroundColor, JarvisAstryxTokens.error);

    expect(
      badgeFor('quiet').backgroundColor,
      JarvisAstryxTokens.backgroundMuted,
    );
    expect(badgeFor('quiet').foregroundColor, JarvisAstryxTokens.textSecondary);
  });
}
