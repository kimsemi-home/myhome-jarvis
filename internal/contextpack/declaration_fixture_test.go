package contextpack

import "encoding/json"

func validDeclaration() Declaration {
	return Declaration{PackID: "myhome-jarvis/context-pack",
		ContextPackVersion:           "v1",
		UpstreamCompatibilityVersion: "myhome-jarvis/context-pack/v1",
		OntologyVersion:              "concept-registry/v1",
		AuthorityContractVersion:     "authority/v1",
		SecurityContractVersion:      "security/v1",
		VerificationProfile:          "quality",
		SSOTArtifactVersions:         testArtifacts()}
}

func declarationJSON(declaration Declaration) string {
	body, err := json.Marshal(declaration)
	if err != nil {
		panic(err)
	}
	return string(body) + "\n"
}
