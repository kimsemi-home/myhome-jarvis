(in-package #:myhome-jarvis.ssot)

(defparameter *finance-entities*
  #("Household"
    "Person"
    "Account"
    "Card"
    "Transaction"
    "Merchant"
    "Subscription"
    "Benefit"
    "Recommendation"))

(defparameter *transaction-ir*
  #("transaction_id"
    "source"
    "owner"
    "occurred_at"
    "posted_at"
    "amount"
    "currency"
    "direction"
    "merchant_name"
    "category"
    "account_id"
    "card_id"
    "raw_ref"
    "tags"))
