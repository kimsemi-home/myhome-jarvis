(in-package #:myhome-jarvis.ssot)

(defun wf-checkout (stream unit)
  (wf stream "      - name: Check out repository")
  (wf stream "        uses: ~A" (github-action-ref "checkout"))
  (when (getf unit :history)
    (wf stream "        with:")
    (wf stream "          fetch-depth: 0")))

(defun wf-setup (stream unit)
  (let ((condition (wf-cache-miss-if unit))
        (setup (getf unit :setup)))
    (cond
      ((string= setup "go") (wf-setup-go stream condition))
      ((string= setup "lisp") (wf-setup-lisp stream condition))
      ((string= setup "rust") (wf-setup-rust stream condition))
      ((string= setup "flutter") (wf-setup-flutter stream condition)))))

(defun wf-setup-go (stream condition)
  (wf stream "      - name: Set up Go")
  (wf-if stream condition)
  (wf stream "        uses: ~A" (github-action-ref "setup-go"))
  (wf stream "        with:")
  (wf stream "          go-version: ${{ env.GO_VERSION }}")
  (wf stream "          cache: true")
  (wf stream "          cache-dependency-path: go.mod"))

(defun wf-setup-lisp (stream condition)
  (wf stream "      - name: Set up Common Lisp")
  (wf-if stream condition)
  (wf stream "        uses: ~A" (github-action-ref "setup-lisp")))

(defun wf-setup-rust (stream condition)
  (wf stream "      - name: Set up Rust")
  (wf-if stream condition)
  (wf-run-block stream
                '("rustup toolchain install \"${RUST_TOOLCHAIN}\" --profile minimal --component rustfmt --component clippy"
                  "rustup default \"${RUST_TOOLCHAIN}\"")
                "        "))

(defun wf-setup-flutter (stream condition)
  (wf stream "      - name: Set up Flutter")
  (wf-if stream condition)
  (wf stream "        uses: ~A" (github-action-ref "setup-flutter"))
  (wf stream "        with:")
  (wf stream "          flutter-version: ${{ env.FLUTTER_VERSION }}")
  (wf stream "          channel: stable")
  (wf stream "          cache: true"))

(defun wf-command-step (stream unit)
  (wf stream "      - name: Run ~A verification" (getf unit :name))
  (wf-if stream (wf-cache-miss-if unit))
  (when (getf unit :working_directory)
    (wf stream "        working-directory: ~A"
        (getf unit :working_directory)))
  (wf-run-block stream (policy-list unit :commands) "        "))
