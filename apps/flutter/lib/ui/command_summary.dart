part of '../main.dart';

class CommandSummary extends StatelessWidget {
  const CommandSummary({
    super.key,
    required this.command,
    required this.onDryRun,
  });

  final HomeCommand command;
  final VoidCallback onDryRun;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return Row(
      children: [
        Icon(command.icon, color: shad.colorScheme.primary),
        const SizedBox(width: 16),
        Expanded(child: CommandSummaryText(command: command)),
        JarvisIconAction(
          tooltip: 'Dry-run',
          icon: Icons.play_arrow,
          onPressed: onDryRun,
        ),
      ],
    );
  }
}

class CommandSummaryText extends StatelessWidget {
  const CommandSummaryText({super.key, required this.command});

  final HomeCommand command;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(command.name, style: Theme.of(context).textTheme.titleMedium),
        if (command.payloadFields.isEmpty) ...[
          const SizedBox(height: 4),
          Text(
            command.payload,
            maxLines: 1,
            overflow: TextOverflow.ellipsis,
            style: Theme.of(context).textTheme.bodyMedium,
          ),
        ],
      ],
    );
  }
}
