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
    final colors = Theme.of(context).colorScheme;
    return Row(
      children: [
        Icon(command.icon, color: colors.primary),
        const SizedBox(width: 16),
        Expanded(child: CommandSummaryText(command: command)),
        IconButton(
          tooltip: 'Dry-run',
          onPressed: onDryRun,
          icon: const Icon(Icons.play_arrow),
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
