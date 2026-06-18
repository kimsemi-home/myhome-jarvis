package main

import "errors"

func usage() error {
	return errors.New("usage: mhj <version|commands|auth status|auth token create|auth token rotate|audit status|ci verify|code-shape status|security check|security history|command|connectors status|agent-cluster status|learning status|learning record|evidence status|confidence status|translation status|control-plane status|incidents status|evidence-quality status|review status|authority status|harness home|harness finance|harness commerce|toolchain verify|linear status|linear sync|linear pull|linear next|linear comment|linear transition|linear create-from-backlog|linear replay-offline|daemon|daemon status|ddd verify|knowledge verify|knowledge search|repo status|planner status|loop once|loop status|loop worker|benchmark smoke|quality|quality status|codegen|codegen verify>")
}
