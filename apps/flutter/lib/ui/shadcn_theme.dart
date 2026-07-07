part of '../main.dart';

class JarvisShadTheme {
  static const primary = Color(0xFF176B5B);
  static const accent = Color(0xFF2F6F5E);
  static const background = Color(0xFFF8FAF7);
  static const foreground = Color(0xFF18211F);
  static const muted = Color(0xFFEFF4F1);
  static const border = Color(0xFFD8E2DE);
  static const destructive = Color(0xFFBA3A3A);

  static ShadThemeData get light {
    return ShadThemeData(
      colorScheme: const ShadNeutralColorScheme.light(
        background: background,
        foreground: foreground,
        card: Color(0xFFFFFFFF),
        cardForeground: foreground,
        primary: primary,
        primaryForeground: Color(0xFFFFFFFF),
        secondary: muted,
        secondaryForeground: foreground,
        muted: muted,
        mutedForeground: Color(0xFF60706B),
        accent: Color(0xFFE4EFEA),
        accentForeground: foreground,
        destructive: destructive,
        border: border,
        input: border,
        ring: accent,
        selection: Color(0xFFBFE6D7),
      ),
      radius: const BorderRadius.all(Radius.circular(8)),
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
