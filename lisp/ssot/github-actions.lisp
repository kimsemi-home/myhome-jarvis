(in-package #:myhome-jarvis.ssot)

(defparameter *github-action-refs*
  #((:key "checkout"
     :uses "actions/checkout@v7"
     :runtime "node24"
     :evidence "GitHub tags: actions/checkout v7")
    (:key "setup-go"
     :uses "actions/setup-go@v6"
     :runtime "node24"
     :evidence "GitHub tags: actions/setup-go v6")
    (:key "cache-restore"
     :uses "actions/cache/restore@v5"
     :runtime "node24"
     :evidence "GitHub tags: actions/cache v5")
    (:key "cache-save"
     :uses "actions/cache/save@v5"
     :runtime "node24"
     :evidence "GitHub tags: actions/cache v5")
    (:key "setup-lisp"
     :uses "40ants/setup-lisp@v4"
     :runtime "managed"
     :evidence "maintained Lisp setup action")
    (:key "setup-flutter"
     :uses "subosito/flutter-action@v2"
     :runtime "managed"
     :evidence "maintained Flutter setup action")))

(defun github-action-ref (key)
  (let ((entry (find key *github-action-refs*
                     :key (lambda (item) (getf item :key))
                     :test #'string=)))
    (unless entry
      (error "Missing GitHub action ref for ~A" key))
    (getf entry :uses)))
