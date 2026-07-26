part of '../main.dart';

class FinanceAtAGlance extends StatelessWidget {
  const FinanceAtAGlance({super.key, required this.dashboard});

  final FinanceDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    return JarvisSurface(
      padding: const EdgeInsets.all(18),
      child: Row(
        crossAxisAlignment: CrossAxisAlignment.end,
        children: [
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text('이번 달 요약', style: Theme.of(context).textTheme.titleMedium),
                const SizedBox(height: 12),
                Text('순현금흐름', style: Theme.of(context).textTheme.labelMedium),
                const SizedBox(height: 4),
                Text(
                  _wonText(dashboard.netMinorUnits),
                  style: Theme.of(context).textTheme.headlineSmall,
                ),
              ],
            ),
          ),
          Column(
            crossAxisAlignment: CrossAxisAlignment.end,
            children: [
              Text('전체 지출', style: Theme.of(context).textTheme.labelMedium),
              const SizedBox(height: 4),
              Text(_wonText(dashboard.debitMinorUnits)),
              const SizedBox(height: 8),
              Text('구독 ${_wonText(dashboard.subscriptionMinorUnits)}'),
            ],
          ),
        ],
      ),
    );
  }
}
