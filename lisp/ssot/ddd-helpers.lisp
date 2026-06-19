(in-package #:myhome-jarvis.ssot)

(defun ddd-concept (name context kind description aliases owner targets related)
  (list :canonical_name name
        :bounded_context context
        :ddd_kind kind
        :description description
        :allowed_aliases aliases
        :owner owner
        :generated_targets targets
        :related_concepts related))
