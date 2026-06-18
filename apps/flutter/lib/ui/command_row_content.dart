part of '../main.dart';

class CommandRowContent extends StatelessWidget {
  const CommandRowContent({
    super.key,
    required this.command,
    required this.payload,
    required this.onServiceChanged,
    required this.onDryRun,
  });

  final HomeCommand command;
  final _CommandPayloadState payload;
  final ValueChanged<String?> onServiceChanged;
  final VoidCallback onDryRun;

  @override
  Widget build(BuildContext context) {
    final colors = Theme.of(context).colorScheme;
    return DecoratedBox(
      decoration: BoxDecoration(
        border: Border.all(color: colors.outlineVariant),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
        child: Column(
          children: [
            CommandSummary(command: command, onDryRun: onDryRun),
            if (command.payloadFields.isNotEmpty) ...[
              const SizedBox(height: 12),
              PayloadFieldsEditor(
                fields: command.payloadFields,
                controllers: payload.controllers,
                service: payload.service,
                onServiceChanged: onServiceChanged,
              ),
            ],
          ],
        ),
      ),
    );
  }
}
