(in-package #:myhome-jarvis.ssot)

(defun write-verification-backends (root)
  (write-gitlab-ci root)
  (write-local-makefile root)
  (write-bazel-quality root))

(defun backend-artifact-path (root name)
  (merge-pathnames (format nil "generated/~A" name) root))
