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
    final selected = service ?? options.first;
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text('service', style: Theme.of(context).textTheme.labelMedium),
        const SizedBox(height: 6),
        ShadSelect<String>(
          key: const Key('payload-field-service'),
          initialValue: selected,
          minWidth: double.infinity,
          maxWidth: double.infinity,
          placeholder: const Text('service'),
          selectedOptionBuilder: (context, value) => Text(value),
          options: [
            for (final option in options)
              ShadOption(value: option, child: Text(option)),
          ],
          onChanged: onChanged,
        ),
      ],
    );
  }
}

class PayloadTextField extends StatelessWidget {
  const PayloadTextField({super.key, required this.field, this.controller});

  final String field;
  final TextEditingController? controller;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Text(field, style: Theme.of(context).textTheme.labelMedium),
        const SizedBox(height: 6),
        ShadInput(
          key: Key('payload-field-$field'),
          controller: controller,
          placeholder: Text(field),
          keyboardType: _numericPayloadField(field)
              ? TextInputType.number
              : TextInputType.text,
        ),
      ],
    );
  }
}
