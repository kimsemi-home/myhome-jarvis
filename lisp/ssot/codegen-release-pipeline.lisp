(in-package #:myhome-jarvis.ssot)

(defun write-release-pipeline (root)
  (write-json-file
   (artifact-path root "release_pipeline")
   (release-pipeline)))

(defun release-pipeline ()
  (list :schema_version "release.pipeline/v1"
        :name "quality-release"
        :source "lisp/ssot/verification-graph.lisp"
        :trigger "after verification graph success"
        :release_gate "all verification units pass"
        :public_repository t
        :private_data_allowed nil
        :gates (release-gates)
        :evidence #("GitHub quality run" "same-SHA cache-hit rerun"
                    "generated verification conformance manifest"
                    "generated verification test manifest"
                    "generated control-plane verifier manifest"
                    "generated verification evidence manifest")))

(defun release-gates ()
  (map 'vector
       (lambda (unit)
         (list :id (getf unit :id)
               :kind (getf unit :kind)
               :required t
               :evidence (release-gate-evidence unit)))
       (policy-list *verification-graph* :units)))

(defun release-gate-evidence (unit)
  (if (getf unit :cache)
      (format nil ".github/unit-cache/~A/key or command log"
              (getf unit :cache))
      "command log"))
