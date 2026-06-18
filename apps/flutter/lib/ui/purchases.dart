part of '../main.dart';

class PurchasesView extends StatelessWidget {
  const PurchasesView({super.key, required this.dashboard});

  final PurchaseDashboard dashboard;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          PurchaseMetricsGrid(dashboard: dashboard),
          RecurringCandidatesSection(candidates: dashboard.recurringCandidates),
          const SizedBox(height: 20),
          PurchaseOwnerSection(owners: dashboard.owners),
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
