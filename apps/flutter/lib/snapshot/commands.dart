part of '../snapshot.dart';

@immutable
class HomeCommand {
  const HomeCommand({
    required this.name,
    required this.payload,
    required this.icon,
    this.payloadFields = const [],
  });

  final String name;
  final String payload;
  final IconData icon;
  final List<String> payloadFields;

  HomeCommand copyWith({String? payload}) {
    return HomeCommand(
      name: name,
      payload: payload ?? this.payload,
      icon: icon,
      payloadFields: payloadFields,
    );
  }
}

@immutable
class CommandInvocation {
  const CommandInvocation({required this.label, required this.argv, this.url});

  final String label;
  final List<String> argv;
  final String? url;
}

@immutable
class CommandPlan {
  const CommandPlan({
    required this.name,
    required this.dryRun,
    required this.executeAllowed,
    required this.invocations,
    required this.warnings,
  });

  final String name;
  final bool dryRun;
  final bool executeAllowed;
  final List<CommandInvocation> invocations;
  final List<String> warnings;

  factory CommandPlan.preview(HomeCommand command) {
    return CommandPlan(
      name: command.name,
      dryRun: true,
      executeAllowed: false,
      invocations: [
        CommandInvocation(
          label: command.name,
          argv: ['mhj', 'command', command.name, command.payload],
        ),
      ],
      warnings: const [],
    );
  }
}
