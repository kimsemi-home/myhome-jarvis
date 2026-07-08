part of '../snapshot.dart';

const _sampleFinanceDashboard = FinanceDashboard(
  records: 3,
  fixtureOnly: true,
  currency: 'KRW',
  creditMinorUnits: 4500000,
  debitMinorUnits: 153200,
  netMinorUnits: 4346800,
  subscriptionMinorUnits: 65900,
  subscriptionCount: 1,
  cardDebitMinorUnits: 153200,
  cardDebitCount: 2,
  categories: ['income', 'subscription', 'utilities'],
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
  ],
);
