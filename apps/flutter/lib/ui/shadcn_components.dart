part of '../main.dart';

enum JarvisBadgeTone { primary, secondary, outline, destructive }

class JarvisSurface extends StatelessWidget {
  const JarvisSurface({super.key, required this.child, this.padding});

  final Widget child;
  final EdgeInsetsGeometry? padding;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return DecoratedBox(
      decoration: BoxDecoration(
        color: shad.colorScheme.card,
        border: Border.all(color: shad.colorScheme.border),
        borderRadius: shad.radius,
      ),
      child: Padding(
        padding: padding ?? const EdgeInsets.all(16),
        child: child,
      ),
    );
  }
}

class JarvisBadge extends StatelessWidget {
  const JarvisBadge(
    this.label, {
    super.key,
    this.tone = JarvisBadgeTone.outline,
  });

  final String label;
  final JarvisBadgeTone tone;

  @override
  Widget build(BuildContext context) {
    final child = Text(label, overflow: TextOverflow.ellipsis);
    return switch (tone) {
      JarvisBadgeTone.primary => ShadBadge(child: child),
      JarvisBadgeTone.secondary => ShadBadge.secondary(child: child),
      JarvisBadgeTone.destructive => ShadBadge.destructive(child: child),
      JarvisBadgeTone.outline => ShadBadge.outline(child: child),
    };
  }
}

class JarvisBadgeWrap extends StatelessWidget {
  const JarvisBadgeWrap({super.key, required this.labels});

  final Iterable<String> labels;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [for (final label in labels) JarvisBadge(label)],
    );
  }
}
