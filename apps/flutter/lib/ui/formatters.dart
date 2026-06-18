part of '../main.dart';

String _moneyText(int minorUnits, String currency) {
  final suffix = currency.trim();
  if (suffix.isEmpty) {
    return '$minorUnits';
  }
  return '$minorUnits $suffix';
}

String _title(String value) {
  if (value.isEmpty) {
    return value;
  }
  return value[0].toUpperCase() + value.substring(1).toLowerCase();
}

IconData _recommendationIcon(String kind) {
  switch (kind) {
    case 'recurring_purchase_review':
      return Icons.repeat_outlined;
    case 'card_usage_review':
      return Icons.credit_card_outlined;
    case 'subscription_review':
      return Icons.subscriptions_outlined;
    case 'cash_buffer':
      return Icons.account_balance_wallet_outlined;
    default:
      return Icons.auto_graph_outlined;
  }
}
