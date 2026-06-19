(in-package #:myhome-jarvis.ssot)

(defun write-control-plane-verification (root)
  (write-json-file
   (artifact-path root "control_plane_verification")
   (control-plane-verification)))

(defun control-plane-verification ()
  (list :schema_version "control-plane.verification/v1"
        :source "lisp/ssot/control-plane.lisp"
        :policy_artifact (getf *control-plane-policy* :generated_artifact)
        :command (getf *control-plane-policy* :verification_command)
        :status_command "mhj control-plane status"
        :verifier_separation_required
        (getf *control-plane-policy* :verifier_separation_required)
        :checks (control-plane-verification-checks)))

(defun control-plane-verification-checks ()
  (map 'vector
       (lambda (check)
         (list :id check
               :evidence (control-plane-check-evidence check)))
       (policy-list *control-plane-policy* :verifier_checks)))

(defun control-plane-check-evidence (check)
  (cond ((string= check "policy-json-valid")
         "generated control-plane policy parses and validates")
        ((string= check "status-public-redacted")
         "public status JSON excludes forbidden control-plane fields")
        ((string= check "lease-bounds-valid")
         "lease bounds are positive and ordered")
        ((string= check "verifier-separation-required")
         "reviewer and verifier roles must differ")
        ((string= check "manifest-debt-evaluated")
         "invalid manifests and verifier violations are counted as debt")
        (t "control-plane verification check")))
