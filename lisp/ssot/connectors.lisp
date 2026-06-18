(in-package #:myhome-jarvis.ssot)

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
        :connectors #((:key "mydata"
                       :label "MyData aggregator"
                       :category "finance_aggregation"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("accounts" "cards" "transactions")
                       :allowed_operations #("read_fixture" "summarize")
                       :forbidden_operations #("credential_request" "external_api_call"
                                               "transfer" "trade" "card_action")
                       :next_step "Define consent and local vault boundaries before any real connector.")
                      (:key "banking"
                       :label "Bank accounts"
                       :category "finance"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("accounts" "transactions" "balances")
                       :allowed_operations #("read_fixture" "summarize")
                       :forbidden_operations #("credential_request" "external_api_call"
                                               "transfer")
                       :next_step "Model read-only account snapshots with fixture replay before credentials.")
                      (:key "cards"
                       :label "Card spending"
                       :category "finance"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("cards" "transactions" "benefits")
                       :allowed_operations #("read_fixture" "summarize" "recommend_review")
                       :forbidden_operations #("credential_request" "external_api_call"
                                               "card_apply" "card_cancel")
                       :next_step "Keep benefit and spend recommendations review-only.")
                      (:key "securities"
                       :label "Securities accounts"
                       :category "investing"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("holdings" "transactions")
                       :allowed_operations #("read_fixture" "summarize")
                       :forbidden_operations #("credential_request" "external_api_call" "trade")
                       :next_step "Add fixture holdings before any broker adapter.")
                      (:key "commerce"
                       :label "Commerce purchases"
                       :category "commerce"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("orders" "items" "recurring_candidates")
                       :allowed_operations #("read_fixture" "summarize" "recommend_review")
                       :forbidden_operations #("credential_request" "cookie_import"
                                               "scraping" "purchase")
                       :next_step "Extend local purchase fixtures and avoid scraping/cookie capture.")
                      (:key "payments"
                       :label "Payment wallets"
                       :category "payments"
                       :status "planned"
                       :fixture_mode t
                       :data_classes #("payments" "merchants")
                       :allowed_operations #("read_fixture" "summarize")
                       :forbidden_operations #("credential_request" "external_api_call"
                                               "payment" "refund")
                       :next_step "Represent wallet activity as local fixture IR first."))))
