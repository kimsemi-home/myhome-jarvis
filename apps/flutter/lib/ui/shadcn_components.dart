part of '../main.dart';

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
