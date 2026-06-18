package main

func ciWorkflowFixture() string {
	return `name: quality
on:
  push:
  pull_request:
concurrency:
  cancel-in-progress: true
permissions:
  contents: read
env:
  LISP: "sbcl-bin"
jobs:
  ssot:
    steps:
      - uses: 40ants/setup-lisp@v4
      - run: ros -Q run -- --script lisp/scripts/validate-ssot.lisp
      - run: ros -Q run -- --script lisp/scripts/codegen.lisp
  public-safety:
    steps:
      - uses: actions/checkout@v6
        with:
          fetch-depth: 0
      - run: |
          go run ./cmd/mhj security check
          go run ./cmd/mhj security history
  go:
    steps:
      - uses: actions/cache/restore@v5
        with:
          key: go-${{ hashFiles('.github/workflows/quality.yml', '.go-version', 'rust-toolchain.toml', 'generated/*.json') }}
      - run: go run ./cmd/mhj ci verify
      - run: go run ./cmd/mhj code-shape status
      - run: go run ./cmd/mhj toolchain verify
      - uses: actions/cache/save@v5
        if: steps.unit-cache.outputs.cache-hit != 'true' && github.event_name == 'push' && github.repository == 'kimsemi-home/myhome-jarvis'
  flutter:
    steps:
      - uses: actions/cache/restore@v5
        with:
          key: flutter-${{ hashFiles('generated/commands.generated.json', 'generated/connectors.generated.json', 'generated/agent_cluster.generated.json', 'generated/learning.generated.json', 'generated/evidence.generated.json', 'generated/confidence.generated.json', 'generated/translation.generated.json', 'generated/control_plane.generated.json', 'generated/incidents.generated.json', 'generated/evidence_quality.generated.json', 'generated/review.generated.json', 'generated/code_shape.generated.json', 'generated/authority.generated.json') }}
`
}
