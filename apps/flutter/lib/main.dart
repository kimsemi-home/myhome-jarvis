import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void main() {
  runApp(JarvisApp(client: DaemonSnapshotClient.local()));
}

class JarvisApp extends StatelessWidget {
  const JarvisApp({
    super.key,
    this.client = const StaticSnapshotClient(JarvisSnapshot.sample),
  });

  final JarvisClient client;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'myhome-jarvis',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(
          seedColor: const Color(0xFF2F6F5E),
          brightness: Brightness.light,
        ),
        useMaterial3: true,
      ),
      home: JarvisHome(client: client),
    );
  }
}

class JarvisHome extends StatefulWidget {
  const JarvisHome({super.key, required this.client});

  final JarvisClient client;

  @override
  State<JarvisHome> createState() => _JarvisHomeState();
}

class _JarvisHomeState extends State<JarvisHome> {
  late Future<JarvisSnapshot> _snapshot;

  @override
  void initState() {
    super.initState();
    _snapshot = widget.client.load();
  }

  @override
  void didUpdateWidget(JarvisHome oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.client != widget.client) {
      _snapshot = widget.client.load();
    }
  }

  void _reload() {
    setState(() {
      _snapshot = widget.client.load();
    });
  }

  @override
  Widget build(BuildContext context) {
    return FutureBuilder<JarvisSnapshot>(
      future: _snapshot,
      initialData: JarvisSnapshot.sample,
      builder: (context, state) {
        final snapshot = state.hasError
            ? JarvisSnapshot.offlineFallback()
            : state.data ?? JarvisSnapshot.sample;
        return JarvisScaffold(
          snapshot: snapshot,
          client: widget.client,
          loading: state.connectionState == ConnectionState.waiting,
          onRefresh: _reload,
        );
      },
    );
  }
}

class JarvisScaffold extends StatelessWidget {
  const JarvisScaffold({
    super.key,
    required this.snapshot,
    required this.client,
    required this.loading,
    required this.onRefresh,
  });

  final JarvisSnapshot snapshot;
  final JarvisCommandClient client;
  final bool loading;
  final VoidCallback onRefresh;

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: 6,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('myhome-jarvis'),
          actions: [
            IconButton(
              tooltip: 'Refresh',
              onPressed: loading ? null : onRefresh,
              icon: const Icon(Icons.refresh),
            ),
          ],
          bottom: const TabBar(
            isScrollable: true,
            tabAlignment: TabAlignment.start,
            tabs: [
              Tab(icon: Icon(Icons.monitor_heart_outlined), text: 'Status'),
              Tab(icon: Icon(Icons.tune_outlined), text: 'Commands'),
              Tab(icon: Icon(Icons.hub_outlined), text: 'Linear'),
              Tab(icon: Icon(Icons.storage_outlined), text: 'Storage'),
              Tab(icon: Icon(Icons.groups_outlined), text: 'Household'),
              Tab(icon: Icon(Icons.auto_graph_outlined), text: 'Optimize'),
            ],
          ),
        ),
        body: Stack(
          children: [
            TabBarView(
              children: [
                StatusView(metrics: snapshot.metrics),
                CommandsView(commands: snapshot.commands, client: client),
                PlainListView(title: 'Linear', items: snapshot.linearItems),
                PlainListView(title: 'Storage', items: snapshot.storageItems),
                HouseholdView(scopes: snapshot.householdScopes),
                PlainListView(
                  title: 'Optimize',
                  items: snapshot.recommendationItems,
                ),
              ],
            ),
            if (loading) const LinearProgressIndicator(),
          ],
        ),
      ),
    );
  }
}

class StatusView extends StatelessWidget {
  const StatusView({super.key, required this.metrics});

  final List<SystemMetric> metrics;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: GridView.builder(
        padding: const EdgeInsets.all(16),
        gridDelegate: const SliverGridDelegateWithMaxCrossAxisExtent(
          maxCrossAxisExtent: 320,
          mainAxisExtent: 112,
          crossAxisSpacing: 12,
          mainAxisSpacing: 12,
        ),
        itemCount: metrics.length,
        itemBuilder: (context, index) => MetricTile(metric: metrics[index]),
      ),
    );
  }
}

class MetricTile extends StatelessWidget {
  const MetricTile({super.key, required this.metric});

  final SystemMetric metric;

  @override
  Widget build(BuildContext context) {
    final colors = Theme.of(context).colorScheme;
    return DecoratedBox(
      decoration: BoxDecoration(
        border: Border.all(color: colors.outlineVariant),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Row(
          children: [
            Icon(metric.icon, color: colors.primary),
            const SizedBox(width: 12),
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    metric.label,
                    style: Theme.of(context).textTheme.labelLarge,
                  ),
                  const SizedBox(height: 6),
                  Text(
                    metric.value,
                    overflow: TextOverflow.ellipsis,
                    style: Theme.of(context).textTheme.titleMedium,
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}

class CommandsView extends StatelessWidget {
  const CommandsView({super.key, required this.commands, required this.client});

  final List<HomeCommand> commands;
  final JarvisCommandClient client;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
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

class CommandRow extends StatefulWidget {
  const CommandRow({super.key, required this.command, required this.onDryRun});

  final HomeCommand command;
  final ValueChanged<HomeCommand> onDryRun;

  @override
  State<CommandRow> createState() => _CommandRowState();
}

class _CommandRowState extends State<CommandRow> {
  final Map<String, TextEditingController> _controllers = {};
  String? _service;

  @override
  void initState() {
    super.initState();
    _syncControllers();
  }

  @override
  void didUpdateWidget(CommandRow oldWidget) {
    super.didUpdateWidget(oldWidget);
    if (oldWidget.command.name != widget.command.name ||
        oldWidget.command.payload != widget.command.payload ||
        oldWidget.command.payloadFields != widget.command.payloadFields) {
      _disposeControllers();
      _syncControllers();
    }
  }

  @override
  void dispose() {
    _disposeControllers();
    super.dispose();
  }

  void _disposeControllers() {
    for (final controller in _controllers.values) {
      controller.dispose();
    }
    _controllers.clear();
  }

  void _syncControllers() {
    final values = _decodePayload(widget.command.payload);
    _service = values['service'] is String ? values['service'] as String : null;
    for (final field in widget.command.payloadFields) {
      if (field == 'service') {
        continue;
      }
      _controllers[field] = TextEditingController(
        text: _payloadText(values[field]),
      );
    }
  }

  HomeCommand _editedCommand() {
    if (widget.command.payloadFields.isEmpty) {
      return widget.command;
    }
    final payload = <String, Object?>{};
    for (final field in widget.command.payloadFields) {
      if (field == 'service') {
        payload[field] = _service ?? 'netflix';
        continue;
      }
      final text = _controllers[field]?.text.trim() ?? '';
      payload[field] = _payloadValue(field, text);
    }
    return widget.command.copyWith(payload: jsonEncode(payload));
  }

  @override
  Widget build(BuildContext context) {
    final colors = Theme.of(context).colorScheme;
    return DecoratedBox(
      decoration: BoxDecoration(
        border: Border.all(color: colors.outlineVariant),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
        child: Column(
          children: [
            Row(
              children: [
                Icon(widget.command.icon, color: colors.primary),
                const SizedBox(width: 16),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(
                        widget.command.name,
                        style: Theme.of(context).textTheme.titleMedium,
                      ),
                      if (widget.command.payloadFields.isEmpty) ...[
                        const SizedBox(height: 4),
                        Text(
                          widget.command.payload,
                          maxLines: 1,
                          overflow: TextOverflow.ellipsis,
                          style: Theme.of(context).textTheme.bodyMedium,
                        ),
                      ],
                    ],
                  ),
                ),
                IconButton(
                  tooltip: 'Dry-run',
                  onPressed: () => widget.onDryRun(_editedCommand()),
                  icon: const Icon(Icons.play_arrow),
                ),
              ],
            ),
            if (widget.command.payloadFields.isNotEmpty) ...[
              const SizedBox(height: 12),
              PayloadFieldsEditor(
                fields: widget.command.payloadFields,
                controllers: _controllers,
                service: _service,
                onServiceChanged: (value) {
                  setState(() {
                    _service = value;
                  });
                },
              ),
            ],
          ],
        ),
      ),
    );
  }
}

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

  static const _serviceOptions = [
    'netflix',
    'youtube',
    'tving',
    'wavve',
    'disney',
    'coupangplay',
  ];

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
                  ? DropdownButtonFormField<String>(
                      initialValue: service ?? _serviceOptions.first,
                      decoration: const InputDecoration(
                        border: OutlineInputBorder(),
                        isDense: true,
                        labelText: 'service',
                      ),
                      items: [
                        for (final option in _serviceOptions)
                          DropdownMenuItem(value: option, child: Text(option)),
                      ],
                      onChanged: onServiceChanged,
                    )
                  : TextField(
                      controller: controllers[field],
                      decoration: InputDecoration(
                        border: const OutlineInputBorder(),
                        isDense: true,
                        labelText: field,
                      ),
                      keyboardType: _numericPayloadField(field)
                          ? TextInputType.number
                          : TextInputType.text,
                    ),
            ),
        ],
      ),
    );
  }
}

Map<String, Object?> _decodePayload(String payload) {
  try {
    final decoded = jsonDecode(payload);
    if (decoded is Map<String, Object?>) {
      return decoded;
    }
  } on FormatException {
    return const {};
  }
  return const {};
}

String _payloadText(Object? value) {
  return value == null ? '' : '$value';
}

Object _payloadValue(String field, String text) {
  if (_numericPayloadField(field)) {
    return int.tryParse(text) ?? 0;
  }
  return text;
}

bool _numericPayloadField(String field) {
  return field == 'level' || field == 'step';
}

Future<void> showDryRunPreview(
  BuildContext context,
  JarvisCommandClient client,
  HomeCommand command,
) async {
  final messenger = ScaffoldMessenger.of(context);
  try {
    final plan = await client.dryRun(command);
    if (!context.mounted) {
      return;
    }
    await showDialog<void>(
      context: context,
      builder: (dialogContext) => CommandPlanDialog(plan: plan),
    );
  } catch (error) {
    messenger.showSnackBar(SnackBar(content: Text('Dry-run failed: $error')));
  }
}

class CommandPlanDialog extends StatelessWidget {
  const CommandPlanDialog({super.key, required this.plan});

  final CommandPlan plan;

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Text(plan.name),
      content: SizedBox(
        width: 420,
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            mainAxisSize: MainAxisSize.min,
            children: [
              Text(plan.dryRun ? 'Dry-run plan' : 'Execution plan'),
              const SizedBox(height: 12),
              for (final invocation in plan.invocations) ...[
                Text(
                  invocation.label,
                  style: Theme.of(context).textTheme.labelLarge,
                ),
                const SizedBox(height: 4),
                SelectableText(invocation.argv.join(' ')),
                if (invocation.url != null) ...[
                  const SizedBox(height: 4),
                  SelectableText(invocation.url!),
                ],
                const SizedBox(height: 12),
              ],
              for (final warning in plan.warnings)
                Text(
                  warning,
                  style: TextStyle(color: Theme.of(context).colorScheme.error),
                ),
            ],
          ),
        ),
      ),
      actions: [
        TextButton(
          onPressed: () => Navigator.of(context).pop(),
          child: const Text('Close'),
        ),
      ],
    );
  }
}

class HouseholdView extends StatefulWidget {
  const HouseholdView({super.key, required this.scopes});

  final List<HouseholdScope> scopes;

  @override
  State<HouseholdView> createState() => _HouseholdViewState();
}

class _HouseholdViewState extends State<HouseholdView> {
  String? _selectedScope;

  @override
  void initState() {
    super.initState();
    _selectedScope = _firstScope(widget.scopes);
  }

  @override
  void didUpdateWidget(HouseholdView oldWidget) {
    super.didUpdateWidget(oldWidget);
    final scopes = widget.scopes.map((scope) => scope.scope).toSet();
    if (_selectedScope == null || !scopes.contains(_selectedScope)) {
      _selectedScope = _firstScope(widget.scopes);
    }
  }

  @override
  Widget build(BuildContext context) {
    if (widget.scopes.isEmpty) {
      return const PlainListView(title: 'Household', items: ['No scope data']);
    }
    final selected = widget.scopes.firstWhere(
      (scope) => scope.scope == _selectedScope,
      orElse: () => widget.scopes.first,
    );
    return SafeArea(
      child: ListView(
        padding: const EdgeInsets.all(16),
        children: [
          Align(
            alignment: Alignment.centerLeft,
            child: SegmentedButton<String>(
              segments: [
                for (final scope in widget.scopes)
                  ButtonSegment(value: scope.scope, label: Text(scope.label)),
              ],
              selected: {selected.scope},
              onSelectionChanged: (selection) {
                setState(() {
                  _selectedScope = selection.first;
                });
              },
            ),
          ),
          const SizedBox(height: 16),
          ListTile(
            leading: const Icon(Icons.account_balance_wallet_outlined),
            title: Text(
              'Finance net: ${selected.financeNetMinorUnits} ${selected.currency}',
            ),
            subtitle: Text('${selected.financeRecords} records'),
          ),
          const Divider(height: 1),
          ListTile(
            leading: const Icon(Icons.shopping_bag_outlined),
            title: Text(
              'Purchase spend: ${selected.purchaseSpendMinorUnits} ${selected.currency}',
            ),
            subtitle: Text('${selected.purchaseRecords} records'),
          ),
        ],
      ),
    );
  }

  String? _firstScope(List<HouseholdScope> scopes) {
    return scopes.isEmpty ? null : scopes.first.scope;
  }
}

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
