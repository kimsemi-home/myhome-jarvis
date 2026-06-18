part of '../main.dart';

class CategoryChips extends StatelessWidget {
  const CategoryChips({super.key, required this.categories});

  final List<String> categories;

  @override
  Widget build(BuildContext context) {
    return Wrap(
      spacing: 8,
      runSpacing: 8,
      children: [
        for (final category in categories) Chip(label: Text(category)),
      ],
    );
  }
}
