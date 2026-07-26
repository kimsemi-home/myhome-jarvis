part of '../main.dart';

class FinanceOwnerSection extends StatelessWidget {
  const FinanceOwnerSection({
    super.key,
    required this.owners,
    required this.totalDebitMinorUnits,
  });

  final List<FinanceOwner> owners;
  final int totalDebitMinorUnits;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('누가 어떻게 쓰고 있나요?', style: Theme.of(context).textTheme.titleMedium),
        const SizedBox(height: 4),
        Text(
          '공동 지출과 각자 지출의 비중을 비교해요.',
          style: Theme.of(context).textTheme.bodySmall,
        ),
        const SizedBox(height: 10),
        const Wrap(
          spacing: 8,
          runSpacing: 8,
          children: [
            JarvisBadge('household scoped', tone: JarvisBadgeTone.success),
            JarvisBadge('owner scoped', tone: JarvisBadgeTone.success),
          ],
        ),
        const SizedBox(height: 10),
        for (final owner in owners) ...[
          JarvisSurface(
            padding: const EdgeInsets.all(14),
            child: FinanceOwnerCard(
              owner: owner,
              totalDebitMinorUnits: totalDebitMinorUnits,
            ),
          ),
          const SizedBox(height: 8),
        ],
      ],
    );
  }
}
