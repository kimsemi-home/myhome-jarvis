part of '../main.dart';

class HouseholdScopeTile extends StatelessWidget {
  const HouseholdScopeTile({
    super.key,
    required this.icon,
    required this.iconColor,
    required this.title,
    required this.subtitle,
  });

  final IconData icon;
  final Color iconColor;
  final String title;
  final String subtitle;

  @override
  Widget build(BuildContext context) {
    return JarvisSurface(
      padding: const EdgeInsets.all(14),
      child: Row(
        children: [
          Icon(icon, color: iconColor),
          const SizedBox(width: 12),
          Expanded(child: Text(title, overflow: TextOverflow.ellipsis)),
          JarvisBadge(subtitle),
        ],
      ),
    );
  }
}
