part of '../main.dart';

Future<void> showDryRunPreview(
  BuildContext context,
  JarvisCommandClient client,
  HomeCommand command,
) async {
  final messenger = ScaffoldMessenger.of(context);
  try {
    final plan = await client.dryRun(command);
    if (!context.mounted) {
      return;
    }
    await showShadDialog<void>(
      context: context,
      builder: (dialogContext) => CommandPlanDialog(plan: plan),
    );
  } catch (error) {
    messenger.showSnackBar(SnackBar(content: Text('Dry-run failed: $error')));
  }
}

class CommandPlanDialog extends StatelessWidget {
  const CommandPlanDialog({super.key, required this.plan});

  final CommandPlan plan;

  @override
  Widget build(BuildContext context) {
    return ShadDialog(
      title: Text(plan.name),
      description: Text(plan.dryRun ? 'Dry-run plan' : 'Execution plan'),
      child: SizedBox(
        width: 420,
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisSize: MainAxisSize.min,
            children: [
              for (final invocation in plan.invocations) ...[
                Text(
                  invocation.label,
                  style: Theme.of(context).textTheme.labelLarge,
                ),
                const SizedBox(height: 4),
                SelectableText(invocation.argv.join(' ')),
                if (invocation.url != null) ...[
                  const SizedBox(height: 4),
                  SelectableText(invocation.url!),
                ],
                const SizedBox(height: 12),
              ],
              for (final warning in plan.warnings)
                Text(
                  warning,
                  style: TextStyle(color: Theme.of(context).colorScheme.error),
                ),
            ],
          ),
        ),
      ),
      actions: [
        ShadButton.outline(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Close'),
        ),
      ],
    );
  }
}
