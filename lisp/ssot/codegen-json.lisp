(in-package #:myhome-jarvis.ssot)

(defun write-json-file (path value)
  (ensure-directories-exist path)
  (with-open-file (stream path
                          :direction :output
                          :if-exists :supersede
                          :if-does-not-exist :create)
    (write-json-value value stream)
    (terpri stream)))

(defun write-json-value (value stream)
  (cond
    ((stringp value) (write-json-string value stream))
    ((integerp value) (princ value stream))
    ((eq value t) (princ "true" stream))
    ((null value) (princ "false" stream))
    ((vectorp value) (write-json-array value stream))
    ((plist-object-p value) (write-json-object value stream))
    ((listp value) (write-json-array (coerce value 'vector) stream))
    (t (write-json-string (princ-to-string value) stream))))

(defun write-json-object (plist stream)
  (princ "{" stream)
  (loop for (key value) on plist by #'cddr
        for first = t then nil
        do (progn
             (unless first (princ "," stream))
             (write-json-string (json-key key) stream)
             (princ ":" stream)
             (write-json-value value stream)))
  (princ "}" stream))

(defun write-json-array (vector stream)
  (princ "[" stream)
  (loop for index from 0 below (length vector)
        do (progn
             (unless (= index 0) (princ "," stream))
             (write-json-value (aref vector index) stream)))
  (princ "]" stream))

(defun write-json-string (value stream)
  (princ "\"" stream)
  (loop for char across value
        do (case char
             (#\" (princ "\\\"" stream))
             (#\\ (princ "\\\\" stream))
             (#\Newline (princ "\\n" stream))
             (#\Return (princ "\\r" stream))
             (#\Tab (princ "\\t" stream))
             (otherwise (princ char stream))))
  (princ "\"" stream))

(defun plist-object-p (value)
  (and (listp value)
       (evenp (length value))
       (loop for tail on value by #'cddr
             always (keywordp (first tail)))))

(defun json-key (key)
  (let ((name (symbol-name key)))
    (if (mixed-case-symbol-name-p name)
        name
        (string-downcase name))))

(defun mixed-case-symbol-name-p (name)
  (loop for char across name
        thereis (lower-case-p char)))
