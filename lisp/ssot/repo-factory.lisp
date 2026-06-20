(in-package #:myhome-jarvis.ssot)

(defparameter *repo-factory-policy*
  (list :context "PublicSafeRepoFactory"
        :version "v1"
        :generated_artifact "generated/repo_factory.generated.json"
        :target_owner "kimsemi-home"
        :public_repo_default t
        :codex_project_required t
        :repo_creation_allowed_without_review nil
        :public_safety_evidence_required t
        :authority_review_required t
        :private_assets_public_allowed nil
        :local_paths_public_allowed nil
        :template_files *repo-factory-template-files*
        :creation_gates *repo-factory-creation-gates*
        :bootstrap_checklist *repo-factory-bootstrap-checklist*
        :allowed_public_path_prefixes #(".github/" ".codex/" "cmd/" "docs/"
                                        "generated/" "internal/" "README.md"
                                        "LICENSE")
        :forbidden_public_fragments #("absolute_home_path" "old_private_owner"
                                      "private_team_slug" "private_storage_prefix"
                                      ".env" "id_rsa"
                                      "private_key" "access_token"
                                      "client_secret")
        :public_summary_fields #("policy_path" "template_file_count"
                                 "creation_gate_count"
                                 "bootstrap_check_count"
                                 "authority_review_required"
                                 "public_safety_evidence_required"
                                 "public_safe" "missing_template_role_count"
                                 "missing_creation_gate_count"
                                 "forbidden_template_value_count"
                                 "checked_at")
        :commands #("mhj repo-factory status")))
