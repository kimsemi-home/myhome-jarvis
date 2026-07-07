part of '../main.dart';

class ConnectorTile extends StatelessWidget {
  const ConnectorTile({super.key, required this.connector});

  final ConnectorReadiness connector;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return JarvisSurface(
      padding: const EdgeInsets.all(14),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              Icon(Icons.cable_outlined, color: shad.colorScheme.primary),
              const SizedBox(width: 10),
              Expanded(
                child: Text(
                  connector.label,
                  style: Theme.of(context).textTheme.titleSmall,
                ),
              ),
              JarvisBadge(connector.status, tone: JarvisBadgeTone.secondary),
            ],
          ),
          const SizedBox(height: 8),
          JarvisBadgeWrap(labels: _connectorLabels),
          if (connector.allowedOperations.isNotEmpty) ...[
            const SizedBox(height: 10),
            _ConnectorLine('Allowed', connector.allowedOperations),
          ],
          if (connector.forbiddenOperations.isNotEmpty) ...[
            const SizedBox(height: 6),
            _ConnectorLine('Blocked', connector.forbiddenOperations),
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
    );
  }

  List<String> get _connectorLabels {
    return [
      connector.category,
      connector.fixtureMode ? 'fixture-only' : 'off',
      ...connector.dataClasses,
    ];
  }
}

class _ConnectorLine extends StatelessWidget {
  const _ConnectorLine(this.label, this.values);

  final String label;
  final List<String> values;

  @override
  Widget build(BuildContext context) {
    return Text(
      '$label: ${values.join(', ')}',
      style: Theme.of(context).textTheme.bodyMedium,
    );
  }
}
