(in-package #:myhome-jarvis.ssot)

(defparameter *repo-factory-template-files*
  #((:role "generated_ci"
     :path ".github/workflows/quality.yml"
     :source_artifact "generated/github_quality_workflow.generated.yml"
     :purpose "run generated quality and public safety checks")
    (:role "security_scan"
     :path "docs/security.md"
     :source_artifact "generated/security.generated.json"
     :purpose "document required public safety checks")
    (:role "private_data_policy"
     :path "docs/private-data-policy.md"
     :source_artifact "generated/repo_factory.generated.json"
     :purpose "keep household, finance, prompts, and raw assets outside public files")
    (:role "bootstrap_checklist"
     :path "docs/bootstrap-checklist.md"
     :source_artifact "generated/repo_factory.generated.json"
     :purpose "prove the repo was created through the review gate")
    (:role "codex_project"
     :path ".codex/project-goal.md"
     :source_artifact "generated/assistant_vision.generated.json"
     :purpose "seed the public-safe Codex project goal")))

(defparameter *repo-factory-creation-gates*
  #((:key "authority_review" :required t :blocks_repo_creation t
     :evidence "approved review record")
    (:key "public_safety_evidence" :required t :blocks_repo_creation t
     :evidence "mhj security check and mhj security history")
    (:key "generated_ci" :required t :blocks_repo_creation t
     :evidence "generated GitHub Actions workflow")
    (:key "private_data_policy" :required t :blocks_repo_creation t
     :evidence "public private-data policy document")
    (:key "bootstrap_checklist" :required t :blocks_repo_creation t
     :evidence "completed bootstrap checklist")))

(defparameter *repo-factory-bootstrap-checklist*
  #("choose public-safe repository name"
    "generate quality workflow from SSOT"
    "run mhj security check"
    "run mhj security history"
    "record authority review evidence"
    "record public safety evidence"
    "verify generated public files contain no private assets or local paths"
    "open draft pull request before enabling automation"))
