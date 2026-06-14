(in-package #:myhome-jarvis.ssot)

(defparameter *security-policy*
  (list :allowed_languages *allowed-languages*
        :forbidden_languages *forbidden-languages*
        :private_paths #("data/private" "data/lake" "secrets")
        :forbidden_file_markers #("token" "secret" "credential" "cookie")
        :dry_run_default t
        :default_bind_host "127.0.0.1"))
