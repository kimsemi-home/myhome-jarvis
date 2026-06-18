(in-package #:myhome-jarvis.ssot)

(defun validate-ddd-related-concepts (concepts)
  (dolist (concept (coerce *concept-registry* 'list))
    (dolist (related (coerce (getf concept :related_concepts) 'list))
      (require-true (gethash related concepts)
                    "Concept ~A references unknown related concept ~A"
                    (getf concept :canonical_name) related))))

(defun validate-ddd-used-patterns (used-patterns)
  (dolist (pattern (coerce *ddd-patterns* 'list))
    (require-true (gethash pattern used-patterns)
                  "DDD pattern is not used by any concept: ~A" pattern)))

(defun validate-domain-events (contexts concepts)
  (require-true (> (length *domain-events*) 0)
                "Domain events must be defined in SSOT")
  (dolist (event (coerce *domain-events* 'list))
    (let ((name (getf event :name))
          (context (getf event :bounded_context))
          (emitted-by (getf event :emitted_by))
          (payload-fields (getf event :payload_fields)))
      (require-string-value name "Domain event name is required")
      (require-true (gethash context contexts)
                    "Domain event ~A references unknown bounded context ~A"
                    name context)
      (require-true (gethash emitted-by concepts)
                    "Domain event ~A references unknown emitter concept ~A"
                    name emitted-by)
      (require-true (> (length payload-fields) 0)
                    "Domain event ~A must declare payload fields" name))))

(defun validate-harness-contracts (contexts)
  (require-true (> (length *harness-case-contracts*) 0)
                "Harness case contracts must be defined in SSOT")
  (dolist (harness (coerce *harness-case-contracts* 'list))
    (let ((name (getf harness :name))
          (context (getf harness :bounded_context))
          (command (getf harness :command))
          (target (getf harness :evidence_target)))
      (require-string-value name "Harness case name is required")
      (require-true (gethash context contexts)
                    "Harness case ~A references unknown bounded context ~A"
                    name context)
      (require-string-value command
                            "Harness case must declare a command")
      (require-string-value target
                            "Harness case must declare an evidence target"))))
