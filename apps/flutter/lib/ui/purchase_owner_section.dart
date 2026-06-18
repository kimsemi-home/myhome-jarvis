part of '../main.dart';

class PurchaseOwnerSection extends StatelessWidget {
  const PurchaseOwnerSection({super.key, required this.owners});

  final List<PurchaseOwner> owners;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('Owner Spend', style: Theme.of(context).textTheme.titleMedium),
        const SizedBox(height: 8),
        for (final owner in owners) ...[
          ListTile(
            leading: const Icon(Icons.person_outline),
            title: Text('${_title(owner.owner)} spend'),
            subtitle: Text('${owner.records} purchases'),
            trailing: Text(
              _moneyText(owner.purchaseSpendMinorUnits, owner.currency),
              style: Theme.of(context).textTheme.titleSmall,
            ),
          ),
          const Divider(height: 1),
        ],
      ],
    );
  }
}
