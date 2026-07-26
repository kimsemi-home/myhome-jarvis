part of '../snapshot.dart';

const _sampleFinanceDashboard = FinanceDashboard(
  records: 4,
  fixtureOnly: true,
  currency: 'KRW',
  creditMinorUnits: 4500000,
  debitMinorUnits: 195200,
  netMinorUnits: 4304800,
  subscriptionMinorUnits: 65900,
  subscriptionCount: 1,
  cardDebitMinorUnits: 195200,
  cardDebitCount: 3,
  connectorProvider: 'toss_mydata',
  connectorParityPercent: 100,
  connectorParityPass: true,
  categories: ['dining', 'grocery', 'income', 'subscription'],
  owners: [
    FinanceOwner(
      owner: 'household',
      records: 2,
      currency: 'KRW',
      creditMinorUnits: 4500000,
      debitMinorUnits: 65900,
      netMinorUnits: 4434100,
    ),
    FinanceOwner(
      owner: 'user',
      records: 1,
      currency: 'KRW',
      creditMinorUnits: 0,
      debitMinorUnits: 87300,
      netMinorUnits: -87300,
    ),
    FinanceOwner(
      owner: 'spouse',
      records: 1,
      currency: 'KRW',
      creditMinorUnits: 0,
      debitMinorUnits: 42000,
      netMinorUnits: -42000,
    ),
  ],
);
