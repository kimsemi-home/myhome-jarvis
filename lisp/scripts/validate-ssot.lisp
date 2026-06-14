(load (merge-pathnames "load-ssot.lisp"
                       (make-pathname :directory (pathname-directory *load-truename*))))

(mhj-load-ssot)

(handler-case
    (progn
      (myhome-jarvis.ssot:validate-ssot)
      (format t "SSOT valid~%"))
  (error (condition)
    (format *error-output* "SSOT invalid: ~A~%" condition)
    (sb-ext:exit :code 1)))
