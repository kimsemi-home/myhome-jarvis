(in-package #:myhome-jarvis.ssot)

(defparameter *planner-policy*
  (list :loop_mode "closed-loop"
        :max_task_scope "one file or tightly connected files"
        :checkpoint_root "data/private/checkpoints"
        :quality_required t
        :linear_offline_fallback t
        :default_next "choose the highest-priority ready task"
        :task_graph
        (vector
         (list :id "repo_safety"
               :title "Inspect repository safety state"
               :owner "go"
               :status "completed"
               :depends_on #())
         (list :id "ssot_verify"
               :title "Validate SSOT and generated artifacts"
               :owner "lisp"
               :status "completed"
               :depends_on #("repo_safety"))
         (list :id "quality_gate"
               :title "Run local quality gate"
               :owner "go"
               :status "completed"
               :depends_on #("ssot_verify"))
         (list :id "daemon_surface"
               :title "Expose local daemon status surfaces"
               :owner "go"
               :status "completed"
               :depends_on #("quality_gate"))
         (list :id "flutter_surface"
               :title "Reflect daemon status in local client"
               :owner "flutter"
               :status "completed"
               :depends_on #("daemon_surface"))
         (list :id "linear_sync"
               :title "Sync Linear only after explicit external-write approval"
               :owner "go"
               :status "blocked_external_write"
               :depends_on #("quality_gate")))
        :linear_templates
        (vector
         (list :name "implementation_task"
               :title_prefix "[myhome-jarvis]"
               :labels #("local-first" "quality-gated"))
         (list :name "safety_review"
               :title_prefix "[safety]"
               :labels #("public-safe" "no-secrets")))))
