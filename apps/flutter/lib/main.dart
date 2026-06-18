import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:myhome_jarvis_app/daemon_client.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

part 'ui/home.dart';
part 'ui/scaffold.dart';
part 'ui/status.dart';
part 'ui/finance.dart';
part 'ui/finance_metrics_grid.dart';
part 'ui/finance_metric_tile.dart';
part 'ui/finance_owner_section.dart';
part 'ui/purchases.dart';
part 'ui/purchase_metrics_grid.dart';
part 'ui/purchase_sections.dart';
part 'ui/purchase_owner_section.dart';
part 'ui/category_chips.dart';
part 'ui/optimize_view.dart';
part 'ui/connectors_view.dart';
part 'ui/agent_cluster_view.dart';
part 'ui/simple_lists.dart';
part 'ui/agent_cluster_tile.dart';
part 'ui/connector_tile.dart';
part 'ui/recommendation_tile.dart';
part 'ui/recommendation_parts.dart';
part 'ui/recommendation_chips.dart';
part 'ui/commands_view.dart';
part 'ui/command_row.dart';
part 'ui/command_row_state.dart';
part 'ui/command_row_content.dart';
part 'ui/command_payload_state.dart';
part 'ui/command_summary.dart';
part 'ui/payload_fields_editor.dart';
part 'ui/payload_field_widgets.dart';
part 'ui/command_payload.dart';
part 'ui/formatters.dart';
part 'ui/command_plan.dart';
part 'ui/household.dart';
part 'ui/household_scope_body.dart';
part 'ui/household_scope_tile.dart';

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
