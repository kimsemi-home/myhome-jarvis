(in-package #:myhome-jarvis.ssot)

(defun wf (stream control &rest args)
  (apply #'format stream control args)
  (terpri stream))

(defun wf-join (values)
  (format nil "~{~A~^, ~}" values))

(defun wf-single-quote (value)
  (format nil "'~A'" value))

(defun wf-hash-files (unit)
  (format nil "${{ hashFiles(~{~A~^, ~}) }}"
          (mapcar #'wf-single-quote (policy-list unit :hash_inputs))))

(defun wf-cache-miss-if (unit)
  (when (getf unit :cache)
    "steps.unit-cache.outputs.cache-hit != 'true'"))

(defun wf-if (stream condition)
  (when condition
    (wf stream "        if: ~A" condition)))

(defun wf-run-block (stream lines indent)
  (wf stream "~Arun: |" indent)
  (dolist (line lines)
    (wf stream "~A  ~A" indent line)))
