part of '../main.dart';

class AgentClusterTile extends StatelessWidget {
  const AgentClusterTile({super.key, required this.signal});

  final AgentClusterSignal signal;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return JarvisSurface(
      padding: const EdgeInsets.all(14),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Icon(Icons.account_tree_outlined, color: shad.colorScheme.primary),
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
                    JarvisBadge(signal.status, tone: JarvisBadgeTone.secondary),
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
    );
  }
}
