(in-package #:myhome-jarvis.ssot)

(defun validate-context-pack-policy (policy)
  (require-string-equal (getf policy :context) "CrossRepoContextPack"
                        "Context pack policy context mismatch")
  (require-string-equal (getf policy :generated_artifact)
                        "generated/context_pack.generated.json"
                        "Context pack artifact mismatch")
  (require-string-value (getf policy :pack_id)
                        "Context pack id required")
  (require-string-value (getf policy :upstream_compatibility_version)
                        "Context pack compatibility version required")
  (require-string-value (getf policy :ontology_version)
                        "Context pack ontology version required")
  (require-true (getf policy :public_status_redacted)
                "Context pack must be public redacted")
  (require-false (getf policy :raw_private_context_public_allowed)
                 "Context pack must not expose private context")
  (validate-context-pack-lists policy)
  (require-command policy "mhj context-pack status")
  (require-command policy "mhj context-pack verify"))

(defun validate-context-pack-lists (policy)
  (require-members '("responsibility_overload" "ownership_boundary"
                     "independent_release_cadence" "private_data_boundary"
                     "ci_cost_cache_impact")
                   (mapcar (lambda (item) (getf item :key))
                           (policy-list policy :split_criteria))
                   "Context pack split criterion missing: ~A")
  (require-members '("mission" "ontology" "authority" "security"
                     "verification" "repo_factory")
                   (mapcar (lambda (item) (getf item :role))
                           (policy-list policy :exported_artifacts))
                   "Context pack export role missing: ~A")
  (require-members '("pack_id" "context_pack_version"
                     "upstream_compatibility_version" "ontology_version"
                     "authority_contract_version" "security_contract_version")
                   (policy-list policy :required_declaration_fields)
                   "Context pack declaration field missing: ~A"))
