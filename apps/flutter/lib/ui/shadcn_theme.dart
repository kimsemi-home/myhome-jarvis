part of '../main.dart';

class JarvisShadTheme {
  static const primary = JarvisAstryxTokens.accent;
  static const accent = JarvisAstryxTokens.tealMuted;
  static const background = JarvisAstryxTokens.backgroundBody;
  static const foreground = JarvisAstryxTokens.textPrimary;
  static const muted = JarvisAstryxTokens.backgroundMuted;
  static const border = JarvisAstryxTokens.border;
  static const destructive = JarvisAstryxTokens.error;

  static ShadThemeData get light {
    return ShadThemeData(
      colorScheme: const ShadNeutralColorScheme.light(
        background: background,
        foreground: foreground,
        card: JarvisAstryxTokens.backgroundCard,
        cardForeground: foreground,
        primary: primary,
        primaryForeground: JarvisAstryxTokens.onAccent,
        secondary: muted,
        secondaryForeground: foreground,
        muted: muted,
        mutedForeground: JarvisAstryxTokens.textSecondary,
        accent: accent,
        accentForeground: foreground,
        destructive: destructive,
        border: border,
        input: border,
        ring: primary,
        selection: JarvisAstryxTokens.tealMuted,
      ),
      radius: const BorderRadius.all(JarvisAstryxTokens.radius),
    );
  }

  static ThemeData material(BuildContext context) {
    final shad = ShadTheme.of(context);
    return ThemeData(
      colorScheme: ColorScheme.fromSeed(
        seedColor: primary,
        brightness: Brightness.light,
      ),
      scaffoldBackgroundColor: shad.colorScheme.background,
      dividerColor: shad.colorScheme.border,
      useMaterial3: true,
      chipTheme: ChipThemeData(
        side: BorderSide(color: shad.colorScheme.border),
        shape: RoundedRectangleBorder(borderRadius: shad.radius),
      ),
      inputDecorationTheme: InputDecorationTheme(
        border: OutlineInputBorder(borderRadius: shad.radius),
        enabledBorder: OutlineInputBorder(
          borderRadius: shad.radius,
          borderSide: BorderSide(color: shad.colorScheme.border),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: shad.radius,
          borderSide: BorderSide(color: shad.colorScheme.ring),
        ),
        isDense: true,
      ),
    );
  }
}
