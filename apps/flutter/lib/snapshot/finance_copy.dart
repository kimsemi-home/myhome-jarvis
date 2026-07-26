part of '../snapshot.dart';

extension FinanceDashboardCopy on FinanceDashboard {
  FinanceDashboard copyWith({bool? fixtureOnly}) {
    return FinanceDashboard(
      records: records,
      currency: currency,
      creditMinorUnits: creditMinorUnits,
      debitMinorUnits: debitMinorUnits,
      netMinorUnits: netMinorUnits,
      subscriptionMinorUnits: subscriptionMinorUnits,
      subscriptionCount: subscriptionCount,
      cardDebitMinorUnits: cardDebitMinorUnits,
      cardDebitCount: cardDebitCount,
      categories: categories,
      owners: owners,
      fixtureOnly: fixtureOnly ?? this.fixtureOnly,
      connectorProvider: connectorProvider,
      connectorParityPercent: connectorParityPercent,
      connectorParityPass: connectorParityPass,
    );
  }
}
