(in-package #:myhome-jarvis.ssot)

(defparameter *recommendation-policy*
  (list :mode "fixture_only"
        :external_actions_allowed nil
        :kinds #("cash_buffer"
                 "recurring_purchase_review"
                 "subscription_review")
        :score_min 0
        :score_max 100
        :default_currency "KRW"
        :recommendation_only_actions #("subscription_review"
                                       "recurring_purchase_review")))
