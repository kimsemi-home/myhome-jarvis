(in-package #:myhome-jarvis.ssot)

(defparameter *linear-policy*
  (list :token_env "LINEAR_API_KEY"
        :team_key_env "LINEAR_TEAM_KEY"
        :team_id_env "LINEAR_TEAM_ID"
        :token_file "data/private/linear-token.txt"
        :endpoint "https://api.linear.app/graphql"
        :personal_api_key_authorization "Authorization: <API_KEY>"
        :oauth_authorization "Authorization: Bearer <ACCESS_TOKEN>"
        :offline_queue "data/private/linear-offline-queue.jsonl"
        :default_poll_seconds 60
        :sync_claim_requires_api_success t
        :pull_active_only t
        :team_scope_optional t
        :team_scope_private t
        :project_issue_title_prefix "[myhome-jarvis]"
        :next_prefers_project_issues t
        :next_requires_project_issue t
        :backlog_seed_current_project_only t
        :backlog_seed_dedupes_by_title t
        :backlog_seed_queries_existing_titles t
        :commands #("mhj linear status"
                    "mhj linear sync"
                    "mhj linear pull"
                    "mhj linear next"
                    "mhj linear comment <issue-id> <message>"
                    "mhj linear transition <issue-id> <state>"
                    "mhj linear create-from-backlog")
        :offline_action_kinds #("linear_sync"
                                "linear_pull"
                                "linear_next"
                                "linear_comment"
                                "linear_transition"
                                "linear_create_from_backlog")))
