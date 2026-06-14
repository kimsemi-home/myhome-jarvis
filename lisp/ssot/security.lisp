(in-package #:myhome-jarvis.ssot)

(defparameter *security-policy*
  (list :allowed_languages *allowed-languages*
        :forbidden_languages *forbidden-languages*
        :private_paths #("data/private" "data/lake" "secrets")
        :forbidden_file_markers #("token" "secret" "credential" "cookie")
        :dry_run_default t
        :default_bind_host "127.0.0.1"
        :local_token_file "data/private/local-token.txt"
        :lan_requires_bearer_token t))
