part of '../main.dart';

class CommandStateBadges extends StatelessWidget {
  const CommandStateBadges({super.key, required this.command});

  final HomeCommand command;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 6,
      runSpacing: 6,
      children: [
        const JarvisBadge('dry-run', tone: JarvisBadgeTone.outline),
        const JarvisBadge('execute blocked', tone: JarvisBadgeTone.secondary),
        if (command.payloadFields.isNotEmpty)
          const JarvisBadge('payload editable', tone: JarvisBadgeTone.outline),
      ],
    );
  }
}
