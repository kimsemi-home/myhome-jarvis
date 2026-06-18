part of '../main.dart';

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
      length: 10,
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
              Tab(
                icon: Icon(Icons.account_balance_wallet_outlined),
                text: 'Finance',
              ),
              Tab(icon: Icon(Icons.shopping_bag_outlined), text: 'Purchases'),
              Tab(icon: Icon(Icons.hub_outlined), text: 'Linear'),
              Tab(icon: Icon(Icons.storage_outlined), text: 'Storage'),
              Tab(icon: Icon(Icons.cable_outlined), text: 'Connectors'),
              Tab(icon: Icon(Icons.account_tree_outlined), text: 'Cluster'),
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
                FinanceView(dashboard: snapshot.financeDashboard),
                PurchasesView(dashboard: snapshot.purchaseDashboard),
                PlainListView(title: 'Linear', items: snapshot.linearItems),
                PlainListView(title: 'Storage', items: snapshot.storageItems),
                ConnectorsView(connectors: snapshot.connectors),
                AgentClusterView(signals: snapshot.agentClusterSignals),
                HouseholdView(scopes: snapshot.householdScopes),
                OptimizeView(recommendations: snapshot.recommendations),
              ],
            ),
            if (loading) const LinearProgressIndicator(),
          ],
        ),
      ),
    );
  }
}
