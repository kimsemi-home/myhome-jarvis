(in-package #:myhome-jarvis.ssot)

(defun write-gitlab-ci (root)
  (let ((path (backend-artifact-path root "gitlab_quality.generated.yml")))
    (ensure-directories-exist path)
    (with-open-file (stream path
                            :direction :output
                            :if-exists :supersede
                            :if-does-not-exist :create)
      (emit-gitlab-ci stream))))

(defun emit-gitlab-ci (stream)
  (wf stream "# Generated from lisp/ssot/verification-graph.lisp.")
  (wf stream "stages:")
  (dolist (stage (gitlab-stages))
    (wf stream "  - ~A" stage))
  (wf stream "")
  (loop for unit in (policy-list *verification-graph* :units)
        for first = t then nil
        do (progn
             (unless first (wf stream ""))
             (emit-gitlab-unit stream unit))))

(defun gitlab-stages ()
  (remove-duplicates
   (mapcar (lambda (unit) (getf unit :kind))
           (policy-list *verification-graph* :units))
   :test #'string= :from-end t))

(defun emit-gitlab-unit (stream unit)
  (wf stream "~A:" (getf unit :id))
  (wf stream "  stage: ~A" (getf unit :kind))
  (wf stream "  timeout: ~Am" (getf unit :timeout))
  (wf stream "  script:")
  (dolist (command (policy-list unit :commands))
    (wf stream "    - ~A" (gitlab-command command)))
  (wf stream "  artifacts:")
  (wf stream "    when: always")
  (wf stream "    paths:")
  (wf stream "      - generated/")
  (wf stream "  rules:")
  (wf stream "    - when: on_success"))

(defun gitlab-command (command)
  (if (find #\: command)
      (format nil "'~A'" command)
      command))
