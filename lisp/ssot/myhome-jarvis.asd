(asdf:defsystem #:myhome-jarvis-ssot
  :description "Executable SSOT for myhome-jarvis."
  :serial t
  :components ((:file "package")
               (:file "project")
               (:file "commands")
               (:file "finance")
               (:file "commerce")
               (:file "storage")
               (:file "household")
               (:file "recommendations")
               (:file "scheduler")
               (:file "security")
               (:file "linear")
               (:file "planner")
               (:file "codegen")))
