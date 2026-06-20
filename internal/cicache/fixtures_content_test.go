package cicache

func graphFixture() string {
	return `{"generated_artifacts":["generated/a.generated.json","generated/b.generated.yml",".github/workflows/quality.yml"],"units":[{"id":"public-safety","name":"Public safety","kind":"release-check","commands":["go run ./cmd/mhj security check"]},{"id":"ssot","name":"SSOT and generated artifacts","kind":"conformance","cache":"ssot","hash_inputs":["generated/**","lisp/ssot/**"],"commands":["test -s generated/a.generated.json"]},{"id":"go","name":"Go daemon and CLI","kind":"unit-test","cache":"go","hash_inputs":["cmd/**/*.go","internal/**/*.go","generated/*.json"],"commands":["go test ./..."]}]}`
}

func workflowFixture() string {
	return `jobs:
  ssot:
    steps:
      - name: Restore SSOT and generated artifacts unit cache
        id: unit-cache
        with:
          path: .github/unit-cache/ssot
          key: ssot-${{ runner.os }}-${{ env.UNIT_CACHE_VERSION }}-${{ hashFiles('generated/**') }}
      - name: Report SSOT and generated artifacts cache hit
        if: steps.unit-cache.outputs.cache-hit == 'true'
      - name: Run SSOT and generated artifacts verification
        if: steps.unit-cache.outputs.cache-hit != 'true'
      - name: Save SSOT and generated artifacts unit cache
        if: steps.unit-cache.outputs.cache-hit != 'true' && github.event_name == 'push'
  go:
    steps:
      - name: Restore Go daemon and CLI unit cache
        id: unit-cache
        with:
          path: .github/unit-cache/go
          key: go-${{ runner.os }}-${{ env.UNIT_CACHE_VERSION }}-${{ hashFiles('cmd/**/*.go') }}
      - name: Report Go daemon and CLI cache hit
        if: steps.unit-cache.outputs.cache-hit == 'true'
      - name: Run Go daemon and CLI verification
        if: steps.unit-cache.outputs.cache-hit != 'true'
      - name: Save Go daemon and CLI unit cache
        if: steps.unit-cache.outputs.cache-hit != 'true' && github.event_name == 'push'
`
}
