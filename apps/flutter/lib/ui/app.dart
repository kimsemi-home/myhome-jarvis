part of '../main.dart';

class JarvisApp extends StatelessWidget {
  const JarvisApp({
    super.key,
    this.client = const StaticSnapshotClient(JarvisSnapshot.sample),
  });

  final JarvisClient client;

  @override
  Widget build(BuildContext context) {
    return ShadApp.custom(
      theme: JarvisShadTheme.light,
      appBuilder: (context) {
        return MaterialApp(
          debugShowCheckedModeBanner: false,
          title: 'myhome-jarvis',
          theme: JarvisShadTheme.material(context),
          home: JarvisHome(client: client),
          builder: (_, child) => ShadAppBuilder(child: child),
          scrollBehavior: const ShadScrollBehavior(),
        );
      },
    );
  }
}
