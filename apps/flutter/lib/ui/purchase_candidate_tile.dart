part of '../main.dart';

class RecurringCandidateTile extends StatelessWidget {
  const RecurringCandidateTile(this.candidate, {super.key});

  final RecurringPurchase candidate;

  @override
  Widget build(BuildContext context) {
    final shad = ShadTheme.of(context);
    return Padding(
      padding: const EdgeInsets.only(bottom: 8),
      child: JarvisSurface(
        padding: const EdgeInsets.all(14),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Icon(Icons.repeat_outlined, color: shad.colorScheme.primary),
            const SizedBox(width: 12),
            Expanded(child: _RecurringCandidateBody(candidate)),
            const SizedBox(width: 12),
            Text(
              _moneyText(candidate.latestTotalMinorUnits, candidate.currency),
              style: Theme.of(context).textTheme.titleSmall,
            ),
          ],
        ),
      ),
    );
  }
}

class _RecurringCandidateBody extends StatelessWidget {
  const _RecurringCandidateBody(this.candidate);

  final RecurringPurchase candidate;

  @override
  Widget build(BuildContext context) {
    final details =
        '${candidate.merchantName} / ${candidate.purchaseCount} purchases / '
        '${candidate.latestPurchasedAt}';
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(candidate.itemName, style: Theme.of(context).textTheme.titleSmall),
        const SizedBox(height: 6),
        Text(details),
        const SizedBox(height: 8),
        JarvisBadge(
          '${candidate.purchaseCount} purchases',
          tone: JarvisBadgeTone.secondary,
        ),
      ],
    );
  }
}
