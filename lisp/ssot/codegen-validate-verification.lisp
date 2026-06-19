(in-package #:myhome-jarvis.ssot)

(defun validate-verification-graph (graph)
  (require-string-equal (getf graph :schema_version)
                        "verification.graph/v1"
                        "Verification graph schema mismatch")
  (require-members '("generated/verification_graph.generated.json"
                     "generated/github_quality_workflow.generated.yml"
                     ".github/workflows/quality.yml"
                     "generated/gitlab_quality.generated.yml"
                     "generated/local_quality.generated.mk"
                     "generated/bazel_quality.generated.bzl"
                     "docs/verification-graph.md")
                   (policy-list graph :generated_artifacts)
                   "Verification generated artifact missing: ~A")
  (validate-verification-backends graph)
  (require-string-value (getf graph :expression)
                        "Verification graph expression is required")
  (dolist (unit (policy-list graph :units))
    (validate-verification-unit unit))
  t)

(defun validate-verification-unit (unit)
  (require-string-value (getf unit :id)
                        "Verification unit id is required")
  (require-string-value (getf unit :name)
                        "Verification unit name is required")
  (require-string-value (getf unit :kind)
                        "Verification unit kind is required")
  (require-positive-integer (getf unit :timeout)
                            "Verification unit timeout must be positive")
  (require-string-value (getf unit :setup)
                        "Verification unit setup is required")
  (when (getf unit :cache)
    (require-string-value (getf unit :cache)
                          "Verification unit cache key is required")
    (require-true (> (length (policy-list unit :hash_inputs)) 0)
                  "Cached verification unit needs hash inputs"))
  (require-true (> (length (policy-list unit :commands)) 0)
                "Verification unit needs commands"))

(defun validate-verification-backends (graph)
  (dolist (backend (policy-list graph :backends))
    (require-string-value (getf backend :id)
                          "Verification backend id is required")
    (require-string-value (getf backend :path)
                          "Verification backend path is required")))
