part of '../main.dart';

class StatusView extends StatelessWidget {
  const StatusView({super.key, required this.metrics});

  final List<SystemMetric> metrics;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: GridView.builder(
        key: const Key('status-grid'),
        padding: const EdgeInsets.all(16),
        gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
          maxCrossAxisExtent: 320,
          mainAxisExtent: 128,
          crossAxisSpacing: 12,
          mainAxisSpacing: 12,
        ),
        itemCount: metrics.length,
        itemBuilder: (context, index) => MetricTile(metric: metrics[index]),
      ),
    );
  }
}
