(in-package #:myhome-jarvis.ssot)

(defun validate-repo-factory-policy (policy)
  (require-string-equal (getf policy :context)
                        "PublicSafeRepoFactory"
                        "Repo factory policy context mismatch")
  (require-true (getf policy :public_repo_default)
                "Repo factory must default to public repos")
  (require-true (and (getf policy :codex_project_required)
                     (getf policy :public_safety_evidence_required)
                     (getf policy :authority_review_required))
                "Repo factory must require Codex project, public safety evidence, and authority review")
  (require-false (getf policy :repo_creation_allowed_without_review)
                 "Repo creation must not bypass review")
  (require-false (getf policy :private_assets_public_allowed)
                 "Repo factory must not allow private assets in public output")
  (require-false (getf policy :local_paths_public_allowed)
                 "Repo factory must not allow local paths in public output")
  (validate-repo-factory-template-files (policy-list policy :template_files))
  (validate-repo-factory-gates (policy-list policy :creation_gates))
  (require-members '("choose public-safe repository name"
                     "generate quality workflow from SSOT"
                     "run mhj security check"
                     "run mhj security history"
                     "record authority review evidence"
                     "record public safety evidence"
                     "declare consumed context pack version")
                   (policy-list policy :bootstrap_checklist)
                   "Repo factory bootstrap checklist missing: ~A")
  (require-command policy "mhj repo-factory status"))

(defun validate-repo-factory-template-files (files)
  (let ((roles (mapcar (lambda (file) (getf file :role)) files)))
    (require-members '("generated_ci" "security_scan" "private_data_policy"
                       "bootstrap_checklist" "codex_project"
                       "context_pack_declaration")
                     roles
                     "Repo factory template role missing: ~A"))
  (dolist (file files)
    (require-string-value (getf file :path)
                          "Repo factory template path must be present")
    (require-string-value (getf file :source_artifact)
                          "Repo factory template source artifact must be present")
    (require-string-value (getf file :purpose)
                          "Repo factory template purpose must be present")))

(defun validate-repo-factory-gates (gates)
  (let ((keys (mapcar (lambda (gate) (getf gate :key)) gates)))
    (require-members '("authority_review" "public_safety_evidence"
                       "generated_ci" "private_data_policy"
                       "bootstrap_checklist")
                     keys
                     "Repo factory creation gate missing: ~A"))
  (dolist (gate gates)
    (require-true (and (getf gate :required)
                       (getf gate :blocks_repo_creation))
                  "Repo factory creation gate must be required and blocking")
    (require-string-value (getf gate :evidence)
                          "Repo factory creation gate evidence must be present")))
