(in-package #:myhome-jarvis.ssot)

(defparameter *household-policy*
  (list :scopes #("user" "spouse" "household")
        :default_scope "user"
        :aggregate_scope "household"
        :fixture_only t
        :external_actions_allowed nil))
