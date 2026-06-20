package main

import "errors"

func usage() error {
	return errors.New("usage: mhj <version|commands|assistant status|work-item status|auth status|auth token create|auth token rotate|audit status|ci verify|ci-cache status|verification verify|verification evidence|code-shape status|security status|security check|security history|command|connectors status|agent-cluster status|learning status|learning record|evidence status|evidence-integrity status|confidence status|translation status|control-plane status|control-plane verify|context-pack status|context-pack verify [path]|incidents status|evidence-quality status|finance-consent status|codex-cost status|codex-cost record|codex-sustainability status|codex-sustainability record-quality|media-readiness status|merge-evidence status|storage-archive status|storage-archive run|monetization status|repo-factory status|review status|authority status|authority-review status|authority-review request|authority-review evidence|authority-review queue|pdca status|harness home|harness finance|harness commerce|toolchain verify|linear status|linear sync|linear pull|linear next|linear comment|linear transition|linear create-from-backlog|linear replay-offline|daemon|daemon status|ddd verify|knowledge verify|knowledge search|repo status|planner status|loop once|loop status|loop worker|benchmark smoke|quality|quality status|codegen|codegen verify>")
}
