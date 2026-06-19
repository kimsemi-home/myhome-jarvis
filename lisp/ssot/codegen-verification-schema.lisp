(in-package #:myhome-jarvis.ssot)

(defun write-verification-schema (root)
  (write-json-file
   (artifact-path root "verification_graph.schema")
   (verification-schema)))

(defun verification-schema ()
  (list :|$schema| "https://json-schema.org/draft/2020-12/schema"
        :|$id| "generated/verification_graph.schema.generated.json"
        :title "Verification Graph"
        :type "object"
        :required #("schema_version" "name" "expression" "source"
                    "generated_artifacts" "backends" "evidence" "units")
        :properties
        (list :schema_version (schema-string "verification.graph/v1")
              :name (schema-string nil)
              :expression (schema-string nil)
              :source (schema-string nil)
              :generated_artifacts (schema-string-array)
              :backends (schema-object-array #("id" "path"))
              :evidence (schema-string-array)
              :units (schema-object-array #("id" "name" "kind" "timeout"
                                            "setup" "commands")))
        :|additionalProperties| t))

(defun schema-string (const)
  (if const
      (list :type "string" :const const)
      (list :type "string" :|minLength| 1)))

(defun schema-string-array ()
  (list :type "array"
        :items (list :type "string" :|minLength| 1)
        :|minItems| 1))

(defun schema-object-array (required)
  (list :type "array"
        :items (list :type "object"
                     :required required
                     :|additionalProperties| t)
        :|minItems| 1))
