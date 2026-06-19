import 'package:flutter_test/flutter_test.dart';
import 'package:myhome_jarvis_app/snapshot.dart';

void expectLinearAndStorage(JarvisSnapshot snapshot) {
  expect(snapshot.linearItems, contains('Teams: 1'));
  expect(snapshot.linearItems, contains('Viewer configured: true'));
  expect(snapshot.linearItems, contains('Synced: true'));
  expect(snapshot.storageItems, contains('Finance: 3 transactions'));
  expect(snapshot.storageItems, contains('Finance net: 4346800 KRW'));
  expect(snapshot.storageItems, contains('Commerce: 3 purchases'));
  expect(snapshot.storageItems, contains('Storage: parquet+zstd'));
}
