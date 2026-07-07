part of '../main.dart';

class FinanceOwnerSection extends StatelessWidget {
  const FinanceOwnerSection({super.key, required this.owners});

  final List<FinanceOwner> owners;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
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
                Icon(Icons.person_outline, color: shad.colorScheme.primary),
                const SizedBox(width: 12),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('${_title(owner.owner)} net'),
                      const SizedBox(height: 6),
                      JarvisBadge('${owner.records} records'),
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
