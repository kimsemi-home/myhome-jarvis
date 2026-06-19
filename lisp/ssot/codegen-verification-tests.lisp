(in-package #:myhome-jarvis.ssot)

(defun write-verification-tests (root)
  (write-json-file
   (artifact-path root "verification_tests")
   (verification-tests)))

(defun verification-tests ()
  (list :schema_version "verification.tests/v1"
        :source "lisp/ssot/verification-graph.lisp"
        :command "mhj verification verify"
        :target_graph "generated/verification_graph.generated.json"
        :tests (verification-test-cases)))

(defun verification-test-cases ()
  #((:id "graph-artifacts-exist"
     :kind "artifact"
     :evidence "all generated_artifacts paths exist")
    (:id "backend-artifacts-exist"
     :kind "backend"
     :evidence "all backend paths exist")
    (:id "schema-json-valid"
     :kind "schema"
     :evidence "schema artifact parses as JSON")
    (:id "conformance-manifest-linked"
     :kind "conformance"
     :evidence "conformance manifest links graph, schema, tests, release")
    (:id "release-gates-cover-units"
     :kind "release"
     :evidence "release gates cover every verification unit")
    (:id "control-plane-verifier-linked"
     :kind "sidecar"
     :evidence "control-plane verifier artifact and command are in graph")
    (:id "verification-evidence-linked"
     :kind "evidence"
     :evidence "verification evidence artifact and command are in graph")
    (:id "pdca-linked"
     :kind "learning-loop"
     :evidence "PDCA artifact and status command are in graph")
    (:id "local-makefile-ssot-drift-check"
     :kind "local"
     :evidence "generated Makefile verify-ssot command")))
