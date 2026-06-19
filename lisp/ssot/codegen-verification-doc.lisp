(in-package #:myhome-jarvis.ssot)

(defun write-verification-doc (root)
  (let ((path (merge-pathnames "docs/verification-graph.md" root)))
    (ensure-directories-exist path)
    (with-open-file (stream path
                            :direction :output
                            :if-exists :supersede
                            :if-does-not-exist :create)
      (emit-verification-doc stream))))

(defun emit-verification-doc (stream)
  (format stream "# Verification Graph~%~%")
  (format stream "Source: `~A`~%~%" (getf *verification-graph* :source))
  (format stream "Expression: `~A`~%~%" (getf *verification-graph* :expression))
  (format stream "Generated artifacts:~%")
  (dolist (artifact (policy-list *verification-graph* :generated_artifacts))
    (format stream "- `~A`~%" artifact))
  (format stream "~%| Unit | Kind | Cache | Evidence |~%")
  (format stream "| --- | --- | --- | --- |~%")
  (dolist (unit (policy-list *verification-graph* :units))
    (format stream "| `~A` | ~A | ~A | GitHub log + command exit code |~%"
            (getf unit :id)
            (getf unit :kind)
            (or (getf unit :cache) "none"))))
