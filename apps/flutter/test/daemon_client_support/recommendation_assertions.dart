import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void expectRecommendations(JarvisSnapshot snapshot) {
  expect(
    snapshot.recommendationItems,
    contains('81 - Compare recurring purchase: Bottled water'),
  );
  expect(
    snapshot.recommendationItems,
    contains('67 - Review card-linked household spend'),
  );
  expect(
    snapshot.recommendationItems,
    contains('61 - Review household subscriptions'),
  );
  expect(snapshot.recommendations.map((item) => item.kind), [
    'recurring_purchase_review',
    'card_usage_review',
    'subscription_review',
  ]);
  expect(snapshot.recommendations.first.score, 81);
  expect(snapshot.recommendations.first.estimatedMonthlyMinorUnits, 11800);
  expect(snapshot.recommendations.first.evidenceCount, 2);
  expect(snapshot.recommendations[1].rationale, contains('Card-linked'));
}
