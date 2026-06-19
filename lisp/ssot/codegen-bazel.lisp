(in-package #:myhome-jarvis.ssot)

(defun write-bazel-quality (root)
  (let ((path (backend-artifact-path root "bazel_quality.generated.bzl")))
    (ensure-directories-exist path)
    (with-open-file (stream path
                            :direction :output
                            :if-exists :supersede
                            :if-does-not-exist :create)
      (emit-bazel-quality stream))))

(defun emit-bazel-quality (stream)
  (wf stream "# Generated from lisp/ssot/verification-graph.lisp.")
  (wf stream "def mhj_verification_graph():")
  (dolist (unit (policy-list *verification-graph* :units))
    (emit-bazel-unit stream unit))
  (wf stream "    native.test_suite(")
  (wf stream "        name = \"quality\",")
  (wf stream "        tests = [")
  (dolist (id (unit-ids))
    (wf stream "            \":verify_~A\"," id))
  (wf stream "        ],")
  (wf stream "    )"))

(defun emit-bazel-unit (stream unit)
  (wf stream "    native.sh_test(")
  (wf stream "        name = \"verify_~A\"," (getf unit :id))
  (wf stream "        srcs = [\":run_shell\"],")
  (wf stream "        args = [")
  (dolist (command (bazel-commands unit))
    (wf stream "            ~S," command))
  (wf stream "        ],")
  (wf stream "    )"))

(defun bazel-commands (unit)
  (let ((prefix (getf unit :working_directory)))
    (if prefix
        (mapcar (lambda (command)
                  (format nil "cd ~A && ~A" prefix command))
                (policy-list unit :commands))
        (policy-list unit :commands))))
