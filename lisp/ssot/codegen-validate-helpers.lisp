(in-package #:myhome-jarvis.ssot)

(defun policy-list (policy key)
  (coerce (getf policy key) 'list))

(defun require-true (value message &rest args)
  (unless value
    (apply #'error message args)))

(defun require-false (value message &rest args)
  (when value
    (apply #'error message args)))

(defun require-string-value (value message &rest args)
  (unless (and (stringp value) (> (length value) 0))
    (apply #'error message args)))

(defun require-string-equal (value expected message)
  (unless (string= value expected)
    (error message)))

(defun require-member (value values message &rest args)
  (unless (find value values :test #'string=)
    (apply #'error message args)))

(defun require-members (required values message)
  (dolist (value required)
    (require-member value values message value)))

(defun require-command (policy command)
  (require-member command (policy-list policy :commands)
                  "SSOT command missing: ~A" command))

(defun string-prefix-p (prefix value)
  (and (stringp value)
       (>= (length value) (length prefix))
       (string= prefix (subseq value 0 (length prefix)))))

(defun string-suffix-p (suffix value)
  (and (stringp value)
       (>= (length value) (length suffix))
       (string= suffix
                (subseq value (- (length value) (length suffix))))))

(defun private-path-p (value)
  (string-prefix-p "data/private/" value))

(defun private-jsonl-path-p (value)
  (and (private-path-p value)
       (string-suffix-p ".jsonl" value)))

(defun require-private-path (value message)
  (unless (private-path-p value)
    (error message)))

(defun require-private-jsonl (value message)
  (unless (private-jsonl-path-p value)
    (error message)))

(defun require-positive-integer (value message)
  (unless (and (integerp value) (> value 0))
    (error message)))
