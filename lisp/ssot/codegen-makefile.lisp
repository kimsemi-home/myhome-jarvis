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
  (emit-make-commands stream
                      (policy-list unit :commands)
                      (getf unit :working_directory)))

(defun emit-make-commands (stream commands working-directory)
  (dolist (command commands)
    (wf stream "~A~A" #\Tab
        (make-escape (local-command command working-directory)))))

(defun local-command (command &optional working-directory)
  (let ((prefix "ros -Q run -- --script "))
    (let ((local (if (string-prefix-p prefix command)
                     (format nil "sbcl --script ~A" (subseq command (length prefix)))
                     command)))
      (if working-directory
          (format nil "cd ~A && ~A" working-directory local)
          local))))

(defun make-escape (command)
  (with-output-to-string (stream)
    (loop for char across command
          do (progn
               (when (char= char #\$)
                 (write-char #\$ stream))
               (write-char char stream)))))
