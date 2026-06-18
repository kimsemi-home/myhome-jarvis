part of '../main.dart';

class PayloadServiceField extends StatelessWidget {
  const PayloadServiceField({
    super.key,
    required this.service,
    required this.onChanged,
  });

  static const options = [
    'netflix',
    'youtube',
    'tving',
    'wavve',
    'disney',
    'coupangplay',
  ];

  final String? service;
  final ValueChanged<String?> onChanged;

  @override
  Widget build(BuildContext context) {
    return DropdownButtonFormField<String>(
      initialValue: service ?? options.first,
      isExpanded: true,
      decoration: const InputDecoration(
        border: OutlineInputBorder(),
        isDense: true,
        labelText: 'service',
      ),
      items: [
        for (final option in options)
          DropdownMenuItem(value: option, child: Text(option)),
      ],
      onChanged: onChanged,
    );
  }
}

class PayloadTextField extends StatelessWidget {
  const PayloadTextField({super.key, required this.field, this.controller});

  final String field;
  final TextEditingController? controller;

  @override
  Widget build(BuildContext context) {
    return TextField(
      controller: controller,
      decoration: InputDecoration(
        border: const OutlineInputBorder(),
        isDense: true,
        labelText: field,
      ),
      keyboardType: _numericPayloadField(field)
          ? TextInputType.number
          : TextInputType.text,
    );
  }
}
