package externalbootstrap

func checkChildRepoContext(root string, packet Packet, status *ChildRepoStatus) {
	context, ok := readChildJSON[childContextPack](root, ".mhj/context-pack.json")
	if !ok {
		status.addFinding(".mhj/context-pack.json", "missing_file",
			"context pack declaration is missing or invalid")
		return
	}
	expected := packet.ContextHandoff
	checkChildString(status, ".mhj/context-pack.json", context.UpstreamRepo,
		"kimsemi-home/myhome-jarvis", "upstream repo drift")
	checkChildString(status, ".mhj/context-pack.json", context.Context, packet.Context,
		"context drift")
	checkChildString(status, ".mhj/context-pack.json", context.PackID,
		expected.PackID, "pack id drift")
	checkChildString(status, ".mhj/context-pack.json", context.ContextPackVersion,
		expected.ContextPackVersion, "context pack version drift")
	checkChildString(status, ".mhj/context-pack.json", context.OntologyVersion,
		expected.OntologyVersion, "ontology version drift")
	checkChildString(status, ".mhj/context-pack.json", context.AuthorityContractVersion,
		expected.AuthorityContractVersion, "authority contract drift")
	checkChildString(status, ".mhj/context-pack.json", context.SecurityContractVersion,
		expected.SecurityContractVersion, "security contract drift")
	checkChildString(status, ".mhj/context-pack.json", context.CandidateRepo,
		packet.CandidateRepo, "candidate repo drift")
	checkChildBool(status, ".mhj/context-pack.json", context.PrivateLakeStaysPrivate,
		true, "private lake boundary drift")
	checkChildBool(status, ".mhj/context-pack.json", context.RawPayloadPublicAllowed,
		false, "raw payload policy drift")
	checkChildBool(status, ".mhj/context-pack.json", context.ExternalWritesAllowed,
		false, "external write policy drift")
	status.ContextPackValid = countChildFindings(*status, "drift") == 0
}

func checkChildString(status *ChildRepoStatus, path string, got string, want string, message string) {
	if got != want {
		status.addFinding(path, "drift_context_pack", message)
	}
}

func checkChildBool(status *ChildRepoStatus, path string, got bool, want bool, message string) {
	if got != want {
		status.addFinding(path, "drift_context_pack", message)
	}
}
