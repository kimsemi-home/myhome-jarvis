(in-package #:myhome-jarvis.ssot)

(defun connector (key label category status data allowed forbidden next)
  (list :key key
        :label label
        :category category
        :status status
        :fixture_mode t
        :data_classes data
        :allowed_operations allowed
        :forbidden_operations forbidden
        :next_step next))

(defparameter *connector-policy*
  (list :fixture_only t
        :real_credentials_allowed nil
        :external_api_calls_allowed nil
        :default_status "planned"
        :public_safe_fields #("key" "label" "category" "status" "fixture_mode"
                              "data_classes" "allowed_operations"
                              "forbidden_operations" "next_step")
        :forbidden_fields #("token" "secret" "cookie" "account_id" "card_number"
                            "resident_registration_number" "phone" "email" "local_path")
        :connectors
        (vector
         (connector "mydata" "MyData aggregator" "finance_aggregation" "planned"
                    #("accounts" "cards" "transactions")
                    #("read_fixture" "summarize")
                    #("credential_request" "external_api_call" "transfer" "trade"
                      "card_action")
                    "Define consent and local vault boundaries before any real connector.")
         (connector "banking" "Bank accounts" "finance" "planned"
                    #("accounts" "transactions" "balances")
                    #("read_fixture" "summarize")
                    #("credential_request" "external_api_call" "transfer")
                    "Model read-only account snapshots with fixture replay before credentials.")
         (connector "cards" "Card spending" "finance" "planned"
                    #("cards" "transactions" "benefits")
                    #("read_fixture" "summarize" "recommend_review")
                    #("credential_request" "external_api_call" "card_apply" "card_cancel")
                    "Keep benefit and spend recommendations review-only.")
         (connector "securities" "Securities accounts" "investing" "planned"
                    #("holdings" "transactions")
                    #("read_fixture" "summarize")
                    #("credential_request" "external_api_call" "trade")
                    "Add fixture holdings before any broker adapter.")
         (connector "commerce" "Commerce purchases" "commerce" "planned"
                    #("orders" "items" "recurring_candidates")
                    #("read_fixture" "summarize" "recommend_review")
                    #("credential_request" "cookie_import" "scraping" "purchase")
                    "Extend local purchase fixtures and avoid scraping/cookie capture.")
         (connector "payments" "Payment wallets" "payments" "planned"
                    #("payments" "merchants")
                    #("read_fixture" "summarize")
                    #("credential_request" "external_api_call" "payment" "refund")
                    "Represent wallet activity as local fixture IR first.")
         (connector "external-evidence-lake" "External evidence lake"
                    "public_evidence_boundary" "bootstrap"
                    #("context_pack" "ui_status_metadata" "validation_summary")
                    #("read_public_fixture" "show_status" "link_upstream")
                    #("raw_payload_import" "credential_request" "private_archive"
                      "collector_write")
                    "Render public status only from the evidence-lake UI metadata fixture."))))
