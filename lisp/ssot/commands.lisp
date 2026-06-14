(in-package #:myhome-jarvis.ssot)

(defparameter *commands*
  (list
   (list :name "display_sleep"
         :summary "Put the display to sleep"
         :payload_fields #()
         :dry_run_default t)
   (list :name "mac_sleep"
         :summary "Put the Mac to sleep"
         :payload_fields #()
         :dry_run_default t)
   (list :name "movie_mode"
         :summary "Apply dry-run movie mode actions"
         :payload_fields #()
         :dry_run_default t)
   (list :name "open_coupang_play"
         :summary "Open Coupang Play"
         :payload_fields #()
         :shortcut_for "open_ott"
         :service "coupangplay"
         :target "https://www.coupangplay.com"
         :dry_run_default t)
   (list :name "open_disney_plus"
         :summary "Open Disney+"
         :payload_fields #()
         :shortcut_for "open_ott"
         :service "disney"
         :target "https://www.disneyplus.com"
         :dry_run_default t)
   (list :name "open_netflix"
         :summary "Open Netflix"
         :payload_fields #()
         :shortcut_for "open_ott"
         :service "netflix"
         :target "https://www.netflix.com"
         :dry_run_default t)
   (list :name "open_ott"
         :summary "Open a supported OTT service"
         :payload_fields #("service")
         :allowed_services #("coupangplay" "disney" "netflix" "tving" "wavve" "youtube")
         :dry_run_default t)
   (list :name "open_url"
         :summary "Open a safe http or https URL"
         :payload_fields #("url")
         :allowed_schemes #("http" "https")
         :dry_run_default t)
   (list :name "open_tving"
         :summary "Open TVING"
         :payload_fields #()
         :shortcut_for "open_ott"
         :service "tving"
         :target "https://www.tving.com"
         :dry_run_default t)
   (list :name "open_wavve"
         :summary "Open Wavve"
         :payload_fields #()
         :shortcut_for "open_ott"
         :service "wavve"
         :target "https://www.wavve.com"
         :dry_run_default t)
   (list :name "open_youtube"
         :summary "Open YouTube"
         :payload_fields #()
         :target "https://www.youtube.com"
         :dry_run_default t)
   (list :name "open_youtube_search"
         :summary "Open a YouTube search"
         :payload_fields #("query")
         :dry_run_default t)
   (list :name "sleep_mode"
         :summary "Apply dry-run sleep mode actions"
         :payload_fields #()
         :dry_run_default t)
   (list :name "volume_down"
         :summary "Lower output volume by a step"
         :payload_fields #("step")
         :dry_run_default t)
   (list :name "volume_mute"
         :summary "Mute output volume"
         :payload_fields #()
         :dry_run_default t)
   (list :name "volume_set"
         :summary "Set output volume to 0..100"
         :payload_fields #("level")
         :min_level 0
         :max_level 100
         :dry_run_default t)
   (list :name "volume_up"
         :summary "Raise output volume by a step"
         :payload_fields #("step")
         :dry_run_default t)))
