part of '../snapshot.dart';

const _sampleHouseholdScopes = [
  HouseholdScope(
    scope: 'user',
    label: 'User',
    currency: 'KRW',
    financeRecords: 1,
    financeNetMinorUnits: -87300,
    purchaseRecords: 1,
    purchaseSpendMinorUnits: 3200,
  ),
  HouseholdScope(
    scope: 'spouse',
    label: 'Spouse',
    currency: 'KRW',
    financeRecords: 0,
    financeNetMinorUnits: 0,
    purchaseRecords: 0,
    purchaseSpendMinorUnits: 0,
  ),
  HouseholdScope(
    scope: 'household',
    label: 'Household',
    currency: 'KRW',
    financeRecords: 3,
    financeNetMinorUnits: 4346800,
    purchaseRecords: 3,
    purchaseSpendMinorUnits: 26800,
  ),
];
