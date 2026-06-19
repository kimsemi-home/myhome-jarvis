(in-package #:myhome-jarvis.ssot)

(defparameter *concept-registry*
  (concatenate 'vector
               *concept-registry-core*
               *concept-registry-agent-observation*
               *concept-registry-agent-governance*
               *concept-registry-ops*
               *concept-registry-ops-adapters*))
