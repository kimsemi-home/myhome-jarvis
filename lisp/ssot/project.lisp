(in-package #:myhome-jarvis.ssot)

(defparameter *project*
  (list :name "myhome-jarvis"
        :version "0.1.0-bootstrap"
        :go_version "1.26.5"
        :default_bind_host "127.0.0.1"
        :dry_run_default t
        :local_first t))

(defparameter *allowed-languages*
  #("Go" "Rust" "Common Lisp" "Flutter" "Dart for Flutter"))

(defparameter *forbidden-languages*
  #("Python" "Node.js" "TypeScript"))
