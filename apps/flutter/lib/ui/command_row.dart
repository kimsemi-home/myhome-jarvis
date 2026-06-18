part of '../main.dart';

class CommandRow extends StatefulWidget {
  const CommandRow({super.key, required this.command, required this.onDryRun});

  final HomeCommand command;
  final ValueChanged<HomeCommand> onDryRun;

  @override
  State<CommandRow> createState() => _CommandRowState();
}
