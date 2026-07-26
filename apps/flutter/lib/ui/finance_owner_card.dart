part of '../main.dart';

class FinanceOwnerCard extends StatelessWidget {
  const FinanceOwnerCard({
    super.key,
    required this.owner,
    required this.totalDebitMinorUnits,
  });

  final FinanceOwner owner;
  final int totalDebitMinorUnits;

  @override
  Widget build(BuildContext context) {
    final ratio = totalDebitMinorUnits <= 0
        ? 0.0
        : (owner.debitMinorUnits / totalDebitMinorUnits).clamp(0.0, 1.0);
    final percent = (ratio * 100).round();
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Row(
          children: [
            Icon(
              owner.owner == 'household'
                  ? Icons.groups_outlined
                  : Icons.person_outline,
              color: ownerScopeColor(owner.owner, owner.records),
            ),
            const SizedBox(width: 10),
            Expanded(
              child: Text(
                _ownerDisplayName(owner.owner),
                style: Theme.of(context).textTheme.titleSmall,
              ),
            ),
            Text('순증감 ${_wonText(owner.netMinorUnits)}'),
          ],
        ),
        const SizedBox(height: 12),
        Row(
          children: [
            Expanded(child: Text('지출 ${_wonText(owner.debitMinorUnits)}')),
            Text('$percent%'),
          ],
        ),
        const SizedBox(height: 6),
        LinearProgressIndicator(
          value: ratio,
          minHeight: 8,
          backgroundColor: JarvisAstryxTokens.backgroundMuted,
          color: ownerScopeColor(owner.owner, owner.records),
        ),
        const SizedBox(height: 8),
        Wrap(
          spacing: 8,
          children: [
            JarvisBadge('${owner.records}건'),
            JarvisBadge('수입 ${_wonText(owner.creditMinorUnits)}'),
          ],
        ),
      ],
    );
  }
}

String _ownerDisplayName(String owner) => switch (owner) {
  'user' => '김주윤',
  'spouse' => '김세미',
  'household' => '공동 생활비',
  _ => _title(owner),
};
