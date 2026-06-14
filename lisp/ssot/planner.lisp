(in-package #:myhome-jarvis.ssot)

(defparameter *planner-policy*
  (list :loop_mode "closed-loop"
        :max_task_scope "one file or tightly connected files"
        :checkpoint_root "data/private/checkpoints"
        :quality_required t
        :linear_offline_fallback t))
