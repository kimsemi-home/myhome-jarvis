part of '../main.dart';

class PayloadFieldsEditor extends StatelessWidget {
  const PayloadFieldsEditor({
    super.key,
    required this.fields,
    required this.controllers,
    required this.service,
    required this.onServiceChanged,
  });

  final List<String> fields;
  final Map<String, TextEditingController> controllers;
  final String? service;
  final ValueChanged<String?> onServiceChanged;

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.centerLeft,
      child: Wrap(
        spacing: 8,
        runSpacing: 8,
        children: [
          for (final field in fields)
            SizedBox(
              width: 220,
              child: field == 'service'
                  ? PayloadServiceField(
                      service: service,
                      onChanged: onServiceChanged,
                    )
                  : PayloadTextField(
                      field: field,
                      controller: controllers[field],
                    ),
            ),
        ],
      ),
    );
  }
}
