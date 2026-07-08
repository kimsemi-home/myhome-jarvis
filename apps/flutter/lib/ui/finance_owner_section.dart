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
                      Text('${_title(owner.owner)} net'),
                      const SizedBox(height: 6),
                      Wrap(
                        spacing: 8,
                        runSpacing: 8,
                        children: [
                          JarvisBadge('${owner.records} records'),
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
                  _moneyText(owner.netMinorUnits, owner.currency),
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
