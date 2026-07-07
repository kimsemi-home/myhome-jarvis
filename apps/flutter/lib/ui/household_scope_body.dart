part of '../main.dart';

class HouseholdScopeBody extends StatelessWidget {
  const HouseholdScopeBody({
    super.key,
    required this.scopes,
    required this.selected,
    required this.onSelectionChanged,
  });

  final List<HouseholdScope> scopes;
  final HouseholdScope selected;
  final ValueChanged<Set<String>> onSelectionChanged;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          Align(
            alignment: Alignment.centerLeft,
            child: SegmentedButton<String>(
              segments: [
                for (final scope in scopes)
                  ButtonSegment(value: scope.scope, label: Text(scope.label)),
              ],
              selected: {selected.scope},
              onSelectionChanged: onSelectionChanged,
            ),
          ),
          const SizedBox(height: 16),
          HouseholdScopeTile(
            icon: Icons.account_balance_wallet_outlined,
            title:
                'Finance net: ${selected.financeNetMinorUnits} ${selected.currency}',
            subtitle: '${selected.financeRecords} records',
          ),
          const SizedBox(height: 8),
          HouseholdScopeTile(
            icon: Icons.shopping_bag_outlined,
            title:
                'Purchase spend: ${selected.purchaseSpendMinorUnits} ${selected.currency}',
            subtitle: '${selected.purchaseRecords} records',
          ),
        ],
      ),
    );
  }
}
