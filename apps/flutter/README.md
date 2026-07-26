# myhome-jarvis Flutter App

This is the local client skeleton for `myhome-jarvis`.

It is intentionally Dart-only at this stage. Platform runner files will be added
when the app needs device-specific packaging.

## Finance screen golden

The household finance dashboard is documented by the root README and protected
by `test/finance_dashboard_golden_test.dart`.

```sh
flutter test --update-goldens test/finance_dashboard_golden_test.dart
flutter test test/finance_dashboard_golden_test.dart
```

Golden images: [`test/golden/finance_dashboard.png`](test/golden/finance_dashboard.png)
for local macOS runs and
[`test/golden/finance_dashboard_linux.png`](test/golden/finance_dashboard_linux.png)
for the Linux CI runner.

The local closed loop also checks that this README and the repository README
continue to point at the generated image:

```sh
cd ../..
make -f generated/local_quality.generated.mk verify-flutter
```
