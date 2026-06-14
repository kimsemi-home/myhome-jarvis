(defpackage #:myhome-jarvis.ssot
  (:use #:cl)
  (:export
   #:*project*
   #:*commands*
   #:*finance-entities*
   #:*commerce-entities*
   #:*storage-policy*
   #:*household-policy*
   #:*recommendation-policy*
   #:*security-policy*
   #:*linear-policy*
   #:*planner-policy*
   #:validate-ssot
   #:write-generated-artifacts))

(in-package #:myhome-jarvis.ssot)
