part of '../snapshot.dart';

const _samplePurchaseDashboard = PurchaseDashboard(
  records: 3,
  fixtureOnly: true,
  currency: 'KRW',
  totalSpendMinorUnits: 26800,
  recurringCandidateCount: 1,
  recurringCandidates: [
    RecurringPurchase(
      merchantName: 'Coupang',
      itemName: 'Bottled water 2L x 6',
      currency: 'KRW',
      purchaseCount: 2,
      latestTotalMinorUnits: 11800,
      latestPurchasedAt: '2026-06-10',
    ),
  ],
  categories: ['grocery', 'household'],
  owners: [
    PurchaseOwner(
      owner: 'household',
      records: 2,
      currency: 'KRW',
      purchaseSpendMinorUnits: 23600,
    ),
    PurchaseOwner(
      owner: 'user',
      records: 1,
      currency: 'KRW',
      purchaseSpendMinorUnits: 3200,
    ),
  ],
);
