import 'package:myhome_jarvis_app/snapshot.dart';

String metricValue(JarvisSnapshot snapshot, String label) {
  return snapshot.metrics.singleWhere((metric) => metric.label == label).value;
}
