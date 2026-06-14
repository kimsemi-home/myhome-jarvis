(load (merge-pathnames "load-ssot.lisp"
                       (make-pathname :directory (pathname-directory *load-truename*))))

(mhj-load-ssot)

(handler-case
    (progn
      (myhome-jarvis.ssot:write-generated-artifacts (mhj-repo-root))
      (format t "Generated artifacts updated~%"))
  (error (condition)
    (format *error-output* "Codegen failed: ~A~%" condition)
    (sb-ext:exit :code 1)))
