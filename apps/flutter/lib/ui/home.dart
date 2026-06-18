part of '../main.dart';

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
