(in-package #:myhome-jarvis.ssot)

(defparameter *scheduler-policy*
  (list :name "closed_loop"
        :interval_seconds 60
        :heartbeat_interval_seconds 15
        :min_backoff_seconds 5
        :max_backoff_seconds 300
        :checkpoint_every 1
        :state_root "data/private/scheduler"
        :crash_recovery t
        :rate_limited t))
