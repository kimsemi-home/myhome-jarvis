part of '../main.dart';

class _CommandPayloadState {
  final Map<String, TextEditingController> controllers = {};
  String? service;

  void sync(HomeCommand command) {
    dispose();
    final values = _decodePayload(command.payload);
    service = values['service'] is String ? values['service'] as String : null;
    for (final field in command.payloadFields) {
      if (field == 'service') {
        continue;
      }
      controllers[field] = TextEditingController(
        text: _payloadText(values[field]),
      );
    }
  }

  void dispose() {
    for (final controller in controllers.values) {
      controller.dispose();
    }
    controllers.clear();
  }

  HomeCommand edited(HomeCommand command) {
    if (command.payloadFields.isEmpty) {
      return command;
    }
    final payload = <String, Object?>{};
    for (final field in command.payloadFields) {
      if (field == 'service') {
        payload[field] = service ?? 'netflix';
        continue;
      }
      final text = controllers[field]?.text.trim() ?? '';
      payload[field] = _payloadValue(field, text);
    }
    return command.copyWith(payload: jsonEncode(payload));
  }
}
