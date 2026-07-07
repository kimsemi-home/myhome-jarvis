part of '../main.dart';

class CategoryChips extends StatelessWidget {
  const CategoryChips({super.key, required this.categories});

  final List<String> categories;

  @override
  Widget build(BuildContext context) {
    return JarvisBadgeWrap(labels: categories);
  }
}
