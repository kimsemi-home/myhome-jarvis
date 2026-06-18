part of '../main.dart';

class ConnectorsView extends StatelessWidget {
  const ConnectorsView({super.key, required this.connectors});

  final List<ConnectorReadiness> connectors;

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: ListView.separated(
        padding: const EdgeInsets.all(16),
        itemCount: connectors.length,
        separatorBuilder: (_, _) => const SizedBox(height: 8),
        itemBuilder: (context, index) {
          return ConnectorTile(connector: connectors[index]);
        },
      ),
    );
  }
}
