part of '../main.dart';

String _wonText(int minorUnits) {
  final absolute = minorUnits.abs().toString();
  final grouped = absolute.replaceAllMapped(
    RegExp(r'(?=(\d{3})+(?!\d))'),
    (_) => ',',
  );
  return '${minorUnits < 0 ? '-' : ''}${grouped}원';
}
