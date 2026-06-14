(defpackage #:myhome-jarvis.ssot
  (:use #:cl)
  (:export
   #:*project*
   #:*commands*
   #:*bounded-contexts*
   #:*ddd-patterns*
   #:*concept-registry*
   #:*generated-artifact-contracts*
   #:*planning-rules*
   #:*knowledge-index-schema*
   #:*finance-entities*
   #:*commerce-entities*
   #:*storage-policy*
   #:*household-policy*
   #:*recommendation-policy*
   #:*scheduler-policy*
   #:*security-policy*
   #:*linear-policy*
   #:*planner-policy*
   #:validate-ssot
   #:write-generated-artifacts))

(in-package #:myhome-jarvis.ssot)
