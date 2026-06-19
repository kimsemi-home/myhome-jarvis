(in-package #:myhome-jarvis.ssot)

(defun write-verification-conformance (root)
  (write-json-file
   (artifact-path root "verification_conformance")
   (verification-conformance)))

(defun verification-conformance ()
  (list :schema_version "verification.conformance/v1"
        :source "lisp/ssot/verification-graph.lisp"
        :chain #("Ontology" "Executable SSOT" "Verification Graph"
                 "Codegen" "Generated Artifacts" "Evidence")
        :graph_artifact "generated/verification_graph.generated.json"
        :schema_artifact "generated/verification_graph.schema.generated.json"
        :checks (verification-conformance-checks)
        :backend_artifacts (backend-artifact-list)
        :release_artifact "generated/release_pipeline.generated.json"))

(defun verification-conformance-checks ()
  #((:id "schema-present"
     :evidence "generated/verification_graph.schema.generated.json")
    (:id "graph-present"
     :evidence "generated/verification_graph.generated.json")
    (:id "backends-present"
     :evidence "generated/{github,gitlab,local,bazel}_quality")
    (:id "release-pipeline-present"
     :evidence "generated/release_pipeline.generated.json")
    (:id "drift-protected"
     :evidence "mhj codegen verify and SSOT CI unit")))

(defun backend-artifact-list ()
  (map 'vector
       (lambda (backend)
         (list :id (getf backend :id)
               :path (getf backend :path)))
       (policy-list *verification-graph* :backends)))
