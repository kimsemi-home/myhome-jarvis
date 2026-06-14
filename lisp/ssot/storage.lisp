(in-package #:myhome-jarvis.ssot)

(defparameter *storage-policy*
  (list :fixture_format "jsonl"
        :lake_layers #("raw" "bronze" "silver" "gold")
        :datasets #("finance_transactions" "commerce_purchases")
        :long_term_format "parquet"
        :compression "zstd"
        :private_root "data/lake"))
