part of '../main.dart';

class ConnectorTile extends StatelessWidget {
  const ConnectorTile({super.key, required this.connector});

  final ConnectorReadiness connector;

  @override
  Widget build(BuildContext context) {
    final colors = Theme.of(context).colorScheme;
    return DecoratedBox(
      decoration: BoxDecoration(
        border: Border.all(color: colors.outlineVariant),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.all(14),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                Icon(Icons.cable_outlined, color: colors.primary),
                const SizedBox(width: 10),
                Expanded(
                  child: Text(
                    connector.label,
                    style: Theme.of(context).textTheme.titleSmall,
                  ),
                ),
                Chip(label: Text(connector.status)),
              ],
            ),
            const SizedBox(height: 8),
            Wrap(
              spacing: 8,
              runSpacing: 8,
              children: [
                Chip(label: Text(connector.category)),
                Chip(
                  label: Text(connector.fixtureMode ? 'fixture-only' : 'off'),
                ),
                for (final dataClass in connector.dataClasses)
                  Chip(label: Text(dataClass)),
              ],
            ),
            if (connector.allowedOperations.isNotEmpty) ...[
              const SizedBox(height: 10),
              Text(
                'Allowed: ${connector.allowedOperations.join(', ')}',
                style: Theme.of(context).textTheme.bodyMedium,
              ),
            ],
            if (connector.forbiddenOperations.isNotEmpty) ...[
              const SizedBox(height: 6),
              Text(
                'Blocked: ${connector.forbiddenOperations.join(', ')}',
                style: Theme.of(context).textTheme.bodyMedium,
              ),
            ],
            if (connector.nextStep.isNotEmpty) ...[
              const SizedBox(height: 8),
              Text(
                connector.nextStep,
                style: Theme.of(context).textTheme.bodyMedium,
              ),
            ],
          ],
        ),
      ),
    );
  }
}
