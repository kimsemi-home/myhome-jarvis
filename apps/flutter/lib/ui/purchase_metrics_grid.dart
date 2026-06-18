part of '../main.dart';

class PurchaseMetricsGrid extends StatelessWidget {
  const PurchaseMetricsGrid({super.key, required this.dashboard});

  final PurchaseDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    return GridView.count(
      crossAxisCount: MediaQuery.sizeOf(context).width >= 760 ? 3 : 2,
      childAspectRatio: 2.5,
      crossAxisSpacing: 12,
      mainAxisSpacing: 12,
      shrinkWrap: true,
      physics: const NeverScrollableScrollPhysics(),
      children: [
        FinanceMetricTile(
          label: 'Spend',
          value: _moneyText(dashboard.totalSpendMinorUnits, dashboard.currency),
          icon: Icons.shopping_bag_outlined,
        ),
        FinanceMetricTile(
          label: 'Recurring',
          value: '${dashboard.recurringCandidateCount}',
          icon: Icons.repeat_outlined,
        ),
        FinanceMetricTile(
          label: 'Records',
          value: '${dashboard.records}',
          icon: Icons.receipt_long_outlined,
        ),
      ],
    );
  }
}
