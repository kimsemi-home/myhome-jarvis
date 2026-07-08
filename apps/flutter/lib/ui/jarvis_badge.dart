part of '../main.dart';

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
    return Semantics(
      label: '${tone.semanticLabel} badge: $label',
      child: ShadBadge.raw(
        variant: tone.variant,
        backgroundColor: tone.backgroundColor,
        foregroundColor: tone.foregroundColor,
        hoverBackgroundColor: tone.backgroundColor,
        child: child,
      ),
    );
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
