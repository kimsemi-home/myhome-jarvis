(defun mhj-script-root ()
  (make-pathname :name nil :type nil :defaults *load-truename*))

(defun mhj-repo-root ()
  (merge-pathnames "../../" (mhj-script-root)))

(defun mhj-load-ssot ()
  (dolist (file '("../ssot/package.lisp"
                  "../ssot/project.lisp"
                  "../ssot/ddd.lisp"
                  "../ssot/commands.lisp"
                  "../ssot/finance.lisp"
                  "../ssot/commerce.lisp"
                  "../ssot/storage.lisp"
                  "../ssot/household.lisp"
                  "../ssot/recommendations.lisp"
                  "../ssot/scheduler.lisp"
                  "../ssot/security.lisp"
                  "../ssot/connectors.lisp"
                  "../ssot/agent-cluster.lisp"
                  "../ssot/learning.lisp"
                  "../ssot/evidence.lisp"
                  "../ssot/confidence.lisp"
                  "../ssot/linear.lisp"
                  "../ssot/planner.lisp"
                  "../ssot/codegen.lisp"))
    (load (merge-pathnames file (mhj-script-root)))))
