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
          JarvisSurface(
            padding: const EdgeInsets.all(14),
            child: Row(
              children: [
                Icon(
                  Icons.person_outline,
                  color: ownerScopeColor(owner.owner, owner.records),
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('${_title(owner.owner)} spend'),
                      const SizedBox(height: 6),
                      Wrap(
                        spacing: 8,
                        runSpacing: 8,
                        children: [
                          JarvisBadge('${owner.records} purchases'),
                          for (final state in ownerScopeStates(
                            owner.owner,
                            owner.records,
                          ))
                            JarvisBadge(state.label, tone: state.tone),
                        ],
                      ),
                    ],
                  ),
                ),
                Text(
                  _moneyText(owner.purchaseSpendMinorUnits, owner.currency),
                  style: Theme.of(context).textTheme.titleSmall,
                ),
              ],
            ),
          ),
          const SizedBox(height: 8),
        ],
      ],
    );
  }
}
