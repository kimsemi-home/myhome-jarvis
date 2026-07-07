part of '../main.dart';

class RecurringCandidatesSection extends StatelessWidget {
  const RecurringCandidatesSection({super.key, required this.candidates});

  final List<RecurringPurchase> candidates;

  @override
  Widget build(BuildContext context) {
    if (candidates.isEmpty) {
      return const SizedBox.shrink();
    }
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        const SizedBox(height: 20),
        Text(
          'Recurring Candidates',
          style: Theme.of(context).textTheme.titleMedium,
        ),
        const SizedBox(height: 8),
        for (final candidate in candidates) RecurringCandidateTile(candidate),
      ],
    );
  }
}
