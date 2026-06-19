(defun mhj-script-root ()
  (make-pathname :name nil :type nil :defaults *load-truename*))

(defun mhj-repo-root ()
  (merge-pathnames "../../" (mhj-script-root)))

(defun mhj-load-ssot ()
  (dolist (file (mhj-ssot-files))
    (load (merge-pathnames file (mhj-script-root)))))

(defun mhj-ssot-files ()
  (append '("../ssot/package.lisp" "../ssot/project.lisp")
          (mhj-ddd-files)
          '("../ssot/commands.lisp" "../ssot/finance.lisp" "../ssot/commerce.lisp"
            "../ssot/storage.lisp" "../ssot/household.lisp" "../ssot/recommendations.lisp"
            "../ssot/scheduler.lisp" "../ssot/security.lisp" "../ssot/connectors.lisp"
            "../ssot/agent-cluster.lisp" "../ssot/learning.lisp" "../ssot/evidence.lisp"
            "../ssot/confidence.lisp" "../ssot/translation.lisp" "../ssot/control-plane.lisp"
            "../ssot/incidents.lisp" "../ssot/evidence-quality.lisp" "../ssot/review.lisp"
            "../ssot/assistant-vision.lisp")
          (mhj-authority-files)
          '("../ssot/pdca.lisp" "../ssot/code-shape.lisp"
            "../ssot/linear.lisp" "../ssot/planner.lisp" "../ssot/verification-units.lisp"
            "../ssot/github-actions.lisp" "../ssot/verification-graph.lisp"
            "../ssot/codegen-validate-helpers.lisp"
            "../ssot/codegen-validate-basics.lisp" "../ssot/codegen-validate-agent.lisp"
            "../ssot/codegen-validate-assistant-vision.lisp"
            "../ssot/codegen-validate-shape-learning.lisp" "../ssot/codegen-validate-evidence.lisp"
            "../ssot/codegen-validate-confidence.lisp" "../ssot/codegen-validate-ddd-core.lisp"
            "../ssot/codegen-validate-ddd-links.lisp" "../ssot/codegen-validate-control.lisp"
            "../ssot/codegen-validate-incidents.lisp" "../ssot/codegen-validate-evidence-quality.lisp"
            "../ssot/codegen-validate-review.lisp" "../ssot/codegen-validate-review-details.lisp"
            "../ssot/codegen-validate-authority.lisp" "../ssot/codegen-validate-authority-details.lisp"
            "../ssot/codegen-validate-pdca.lisp" "../ssot/codegen-validate-translation.lisp"
            "../ssot/codegen-validate-linear-planner.lisp" "../ssot/codegen-validate-verification.lisp"
            "../ssot/codegen-json.lisp" "../ssot/codegen-workflow-helpers.lisp"
            "../ssot/codegen-workflow-steps.lisp" "../ssot/codegen-workflow-unit.lisp"
            "../ssot/codegen-workflow.lisp" "../ssot/codegen-gitlab.lisp"
            "../ssot/codegen-makefile.lisp" "../ssot/codegen-bazel.lisp"
            "../ssot/codegen-backends.lisp" "../ssot/codegen-control-plane-verification.lisp"
            "../ssot/codegen-verification-schema.lisp" "../ssot/codegen-verification-evidence.lisp"
            "../ssot/codegen-verification-conformance.lisp" "../ssot/codegen-verification-tests.lisp"
            "../ssot/codegen-release-pipeline.lisp" "../ssot/codegen-verification-doc.lisp"
            "../ssot/codegen-artifacts.lisp" "../ssot/codegen.lisp")))

(defun mhj-ddd-files ()
  '("../ssot/ddd-helpers.lisp" "../ssot/ddd-contexts.lisp"
    "../ssot/ddd-concepts-core.lisp" "../ssot/ddd-concepts-agent-observation.lisp"
    "../ssot/ddd-concepts-agent-governance.lisp" "../ssot/ddd-concepts-ops.lisp"
    "../ssot/ddd-concepts-ops-adapters.lisp" "../ssot/ddd.lisp"
    "../ssot/ddd-events.lisp" "../ssot/ddd-artifacts.lisp"
    "../ssot/ddd-planning.lisp"))

(defun mhj-authority-files ()
  '("../ssot/authority-core.lisp" "../ssot/authority-reasoning.lisp"
    "../ssot/authority-roles.lisp" "../ssot/authority-decisions.lisp"
    "../ssot/authority-boundary.lisp" "../ssot/authority.lisp"))
