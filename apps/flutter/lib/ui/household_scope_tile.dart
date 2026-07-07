part of '../main.dart';

class HouseholdScopeTile extends StatelessWidget {
  const HouseholdScopeTile({
    super.key,
    required this.icon,
    required this.title,
    required this.subtitle,
  });

  final IconData icon;
  final String title;
  final String subtitle;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return JarvisSurface(
      padding: const EdgeInsets.all(14),
      child: Row(
        children: [
          Icon(icon, color: shad.colorScheme.primary),
          const SizedBox(width: 12),
          Expanded(child: Text(title, overflow: TextOverflow.ellipsis)),
          JarvisBadge(subtitle),
        ],
      ),
    );
  }
}
