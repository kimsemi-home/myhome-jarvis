part of '../main.dart';

class PlainListView extends StatelessWidget {
  const PlainListView({super.key, required this.title, required this.items});

  final String title;
  final List<String> items;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
        padding: const EdgeInsets.all(16),
        itemCount: items.length,
        separatorBuilder: (_, _) => const Divider(height: 1),
        itemBuilder: (context, index) {
          return ListTile(
            leading: const Icon(Icons.circle_outlined),
            title: Text(items[index]),
          );
        },
      ),
    );
  }
}
