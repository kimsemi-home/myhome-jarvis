part of '../main.dart';

class FinanceView extends StatelessWidget {
  const FinanceView({super.key, required this.dashboard});

  final FinanceDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          FinanceMetricsGrid(dashboard: dashboard),
          const SizedBox(height: 20),
          FinanceOwnerSection(owners: dashboard.owners),
          if (dashboard.categories.isNotEmpty) ...[
            const SizedBox(height: 20),
            Text('Categories', style: Theme.of(context).textTheme.titleMedium),
            const SizedBox(height: 8),
            CategoryChips(categories: dashboard.categories),
          ],
        ],
      ),
    );
  }
}
