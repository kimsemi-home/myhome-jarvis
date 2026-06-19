(in-package #:myhome-jarvis.ssot)

(defun write-local-makefile (root)
  (let ((path (backend-artifact-path root "local_quality.generated.mk")))
    (ensure-directories-exist path)
    (with-open-file (stream path
                            :direction :output
                            :if-exists :supersede
                            :if-does-not-exist :create)
      (emit-local-makefile stream))))

(defun emit-local-makefile (stream)
  (wf stream "# Generated from lisp/ssot/verification-graph.lisp.")
  (wf stream "SHELL := /bin/bash")
  (wf stream ".SHELLFLAGS := -euo pipefail -c")
  (wf stream ".ONESHELL:")
  (wf stream ".PHONY: verify ~{verify-~A~^ ~}"
      (mapcar (lambda (unit) (getf unit :id))
              (policy-list *verification-graph* :units)))
  (wf stream "")
  (wf stream "verify: ~{verify-~A~^ ~}" (unit-ids))
  (wf stream "")
  (loop for unit in (policy-list *verification-graph* :units)
        for first = t then nil
        do (progn
             (unless first (wf stream ""))
             (emit-make-unit stream unit))))

(defun unit-ids ()
  (mapcar (lambda (unit) (getf unit :id))
          (policy-list *verification-graph* :units)))

(defun emit-make-unit (stream unit)
  (wf stream "verify-~A:" (getf unit :id))
  (when (getf unit :working_directory)
    (wf stream "~Acd ~A" #\Tab (getf unit :working_directory)))
  (emit-make-commands stream (policy-list unit :commands)))

(defun emit-make-commands (stream commands)
  (dolist (command commands)
    (wf stream "~A~A" #\Tab (make-escape (local-command command)))))

(defun local-command (command)
  (let ((prefix "ros -Q run -- --script "))
    (if (string-prefix-p prefix command)
        (format nil "sbcl --script ~A" (subseq command (length prefix)))
        command)))

(defun make-escape (command)
  (with-output-to-string (stream)
    (loop for char across command
          do (progn
               (when (char= char #\$)
                 (write-char #\$ stream))
               (write-char char stream)))))
