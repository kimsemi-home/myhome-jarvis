part of '../main.dart';

class CommandsView extends StatelessWidget {
  const CommandsView({super.key, required this.commands, required this.client});

  final List<HomeCommand> commands;
  final JarvisCommandClient client;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
        key: const Key('commands-list'),
        padding: const EdgeInsets.all(16),
        itemCount: commands.length,
        separatorBuilder: (_, _) => const SizedBox(height: 8),
        itemBuilder: (context, index) => CommandRow(
          command: commands[index],
          onDryRun: (command) => showDryRunPreview(context, client, command),
        ),
      ),
    );
  }
}
