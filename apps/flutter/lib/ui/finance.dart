part of '../main.dart';

class FinanceView extends StatelessWidget {
  const FinanceView({
    super.key,
    required this.dashboard,
  });

  final FinanceDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    final sections = <Widget>[
      Text('가계 자금 한눈에', style: Theme.of(context).textTheme.headlineSmall),
      const SizedBox(height: 4),
      Text(
        '김주윤·김세미·공동 지출을 한 화면에서 비교해요.',
        style: Theme.of(context).textTheme.bodyMedium,
      ),
      const SizedBox(height: 16),
      FinanceConnectionCard(
        fixtureOnly: dashboard.fixtureOnly,
        dashboard: dashboard,
      ),
      const SizedBox(height: 12),
      FinanceDashboardStateBadges(dashboard: dashboard),
      const SizedBox(height: 12),
      FinanceAtAGlance(dashboard: dashboard),
      const SizedBox(height: 20),
      FinanceOwnerSection(
        owners: dashboard.owners,
        totalDebitMinorUnits: dashboard.debitMinorUnits,
      ),
      const SizedBox(height: 20),
      FinanceMetricsGrid(dashboard: dashboard),
    ];
    if (dashboard.categories.isNotEmpty) {
      sections.addAll([
        const SizedBox(height: 20),
        Text('Categories', style: Theme.of(context).textTheme.titleMedium),
        const SizedBox(height: 8),
        CategoryChips(categories: dashboard.categories),
      ]);
    }
    return SafeArea(
      child: ListView(padding: const EdgeInsets.all(16), children: sections),
    );
  }
}
