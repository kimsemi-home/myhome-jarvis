part of '../main.dart';

class AgentClusterView extends StatelessWidget {
  const AgentClusterView({super.key, required this.signals});

  final List<AgentClusterSignal> signals;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
        padding: const EdgeInsets.all(16),
        itemCount: signals.length,
        separatorBuilder: (_, _) => const SizedBox(height: 8),
        itemBuilder: (context, index) {
          return AgentClusterTile(signal: signals[index]);
        },
      ),
    );
  }
}
