part of '../main.dart';

class FinanceMetricsGrid extends StatelessWidget {
  const FinanceMetricsGrid({super.key, required this.dashboard});

  final FinanceDashboard dashboard;

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
          label: 'Net',
          value: _moneyText(dashboard.netMinorUnits, dashboard.currency),
          icon: Icons.account_balance_wallet_outlined,
        ),
        FinanceMetricTile(
          label: 'Credits',
          value: _moneyText(dashboard.creditMinorUnits, dashboard.currency),
          icon: Icons.trending_up_outlined,
        ),
        FinanceMetricTile(
          label: 'Debits',
          value: _moneyText(dashboard.debitMinorUnits, dashboard.currency),
          icon: Icons.trending_down_outlined,
        ),
        FinanceMetricTile(
          label: 'Subscriptions',
          value:
              '${dashboard.subscriptionCount} / ${_moneyText(dashboard.subscriptionMinorUnits, dashboard.currency)}',
          icon: Icons.subscriptions_outlined,
        ),
        FinanceMetricTile(
          label: 'Card-linked',
          value:
              '${dashboard.cardDebitCount} / ${_moneyText(dashboard.cardDebitMinorUnits, dashboard.currency)}',
          icon: Icons.credit_card_outlined,
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
