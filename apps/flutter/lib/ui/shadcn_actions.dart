part of '../main.dart';

class JarvisIconAction extends StatelessWidget {
  const JarvisIconAction({
    super.key,
    required this.tooltip,
    required this.icon,
    required this.onPressed,
  });

  final String tooltip;
  final IconData icon;
  final VoidCallback? onPressed;

  @override
  Widget build(BuildContext context) {
    return Tooltip(
      message: tooltip,
      child: ShadIconButton.outline(
        icon: Icon(icon),
        iconSize: 18,
        width: 36,
        height: 36,
        onPressed: onPressed,
      ),
    );
  }
}
