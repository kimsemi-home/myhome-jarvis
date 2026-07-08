part of '../main.dart';

class MetricTile extends StatelessWidget {
  const MetricTile({super.key, required this.metric});

  final SystemMetric metric;

  @override
  Widget build(BuildContext context) {
    final state = metric.statusState;
    return Semantics(
      label: '${metric.label}: ${metric.value}, ${state.badgeLabel}',
      child: JarvisSurface(
        key: ValueKey('status-metric-${metric.label}'),
        child: Row(
          children: [
            Icon(metric.icon, color: state.iconColor),
            const SizedBox(width: 12),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Row(
                    children: [
                      Expanded(
                        child: Text(
                          metric.label,
                          maxLines: 1,
                          overflow: TextOverflow.ellipsis,
                          style: Theme.of(context).textTheme.labelLarge,
                        ),
                      ),
                      const SizedBox(width: 8),
                      JarvisBadge(state.badgeLabel, tone: state.badgeTone),
                    ],
                  ),
                  const SizedBox(height: 8),
                  Text(
                    metric.value,
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                    style: Theme.of(context).textTheme.titleMedium,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
