(in-package #:myhome-jarvis.ssot)

(defparameter *commerce-entities*
  #("ProductPurchase"
    "ProductItem"
    "Order"
    "Merchant"
    "RecurringPurchaseCandidate"
    "PriceTrend"
    "PurchaseRecommendation"))

(defparameter *purchase-ir*
  #("purchase_id"
    "source"
    "owner"
    "purchased_at"
    "merchant_name"
    "order_id"
    "item_name"
    "brand"
    "quantity"
    "unit_price"
    "total_price"
    "category"
    "recurring_candidate"
    "raw_ref"
    "tags"))
