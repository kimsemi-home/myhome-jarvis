(in-package #:myhome-jarvis.ssot)

(defparameter *planning-rules*
  (list :knowledge_index_required_before_planning t
        :default_knowledge_query "planner KnowledgeIndex Linear closed loop"
        :semantic_changes_require_ssot_first t
        :ssot_change_requires_codegen t
        :small_cohesive_change_required t
        :validation_steps #("focused test" "harness" "codegen verify"
                            "ddd verify" "security check")))

(defparameter *knowledge-index-schema*
  (list :kind "local-lexical"
        :external_vector_db_allowed nil
        :cloud_rag_allowed nil
        :index_roots #("lisp/ssot" "generated" "cmd" "internal"
                       "apps/flutter" "docs" "fixtures" "harness/golden"
                       "data/private/linear-offline-queue.jsonl")
        :query_types #("concept definition location"
                       "bounded context owner"
                       "related implementation files"
                       "related tests and generated files"
                       "related Linear issue"
                       "semantic duplication"
                       "must-read files before change")
        :evidence_fields #("canonical_name" "bounded_context" "owner"
                           "matched_terms" "paths" "linear_issues"
                           "duplicate_suspicions" "must_read_files")))
