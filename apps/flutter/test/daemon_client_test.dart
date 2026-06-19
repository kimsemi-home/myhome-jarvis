import 'daemon_client_cases/auth_header_case.dart';
import 'daemon_client_cases/lan_mode_case.dart';
import 'daemon_client_cases/load_snapshot_case.dart';

void main() {
  runLoadSnapshotTest();
  runAuthHeaderTest();
  runLanModeTest();
}
