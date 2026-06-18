part of '../main.dart';

class FinanceOwnerSection extends StatelessWidget {
  const FinanceOwnerSection({super.key, required this.owners});

  final List<FinanceOwner> owners;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('Owner Breakdown', style: Theme.of(context).textTheme.titleMedium),
        const SizedBox(height: 8),
        for (final owner in owners) ...[
          ListTile(
            leading: const Icon(Icons.person_outline),
            title: Text('${_title(owner.owner)} net'),
            subtitle: Text('${owner.records} records'),
            trailing: Text(
              _moneyText(owner.netMinorUnits, owner.currency),
              style: Theme.of(context).textTheme.titleSmall,
            ),
          ),
          const Divider(height: 1),
        ],
      ],
    );
  }
}
