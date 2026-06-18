part of '../main.dart';

class AgentClusterTile extends StatelessWidget {
  const AgentClusterTile({super.key, required this.signal});

  final AgentClusterSignal signal;

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
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Icon(Icons.account_tree_outlined, color: colors.primary),
            const SizedBox(width: 12),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Row(
                    children: [
                      Expanded(
                        child: Text(
                          signal.label,
                          style: Theme.of(context).textTheme.titleSmall,
                        ),
                      ),
                      Chip(label: Text(signal.status)),
                    ],
                  ),
                  if (signal.evidence.isNotEmpty) ...[
                    const SizedBox(height: 8),
                    Text(
                      signal.evidence,
                      style: Theme.of(context).textTheme.bodyMedium,
                    ),
                  ],
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
