(in-package #:myhome-jarvis.ssot)

(defparameter *verification-backends*
  #((:id "github-actions"
     :path "generated/github_quality_workflow.generated.yml"
     :runtime_path ".github/workflows/quality.yml")
    (:id "gitlab-ci" :path "generated/gitlab_quality.generated.yml")
    (:id "local-makefile" :path "generated/local_quality.generated.mk")
    (:id "bazel" :path "generated/bazel_quality.generated.bzl")))

(defparameter *verification-graph*
  (list :schema_version "verification.graph/v1"
        :name "quality"
        :expression "(verify daemon (public-safety) (ssot) (go) (rust) (flutter) (release-check))"
        :source "lisp/ssot/verification-graph.lisp"
        :generated_artifacts #("generated/verification_graph.generated.json"
                               "generated/github_quality_workflow.generated.yml"
                               ".github/workflows/quality.yml"
                               "generated/gitlab_quality.generated.yml"
                               "generated/local_quality.generated.mk"
                               "generated/bazel_quality.generated.bzl"
                               "generated/verification_graph.schema.generated.json"
                               "generated/verification_conformance.generated.json"
                               "generated/verification_tests.generated.json"
                               "generated/release_pipeline.generated.json"
                               "docs/verification-graph.md")
        :backends *verification-backends*
        :evidence #("GitHub job logs" ".github/unit-cache/<unit>/key"
                    "generated backend specs"
                    "generated schema, conformance, and test specs"
                    "redacted local quality run ledger")
        :units *verification-units*))
