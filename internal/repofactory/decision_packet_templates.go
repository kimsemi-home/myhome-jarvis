package repofactory

func packetTemplates(
	files []TemplateFile,
	forbidden []string,
) []TemplateEvidence {
	templates := make([]TemplateEvidence, 0, len(files))
	for _, file := range files {
		templates = append(templates, packetTemplate(file, forbidden))
	}
	return templates
}

func packetTemplate(file TemplateFile, forbidden []string) TemplateEvidence {
	state := "ready"
	path := file.Path
	if forbiddenTemplateValue(file, forbidden) {
		state = "invalid_forbidden_template_value"
		path = "redacted_forbidden_template_path"
	}
	return TemplateEvidence{
		Role:           file.Role,
		PublicPath:     path,
		SourceArtifact: file.SourceArtifact,
		State:          state,
	}
}

func packetTemplateReadyCount(templates []TemplateEvidence) int {
	count := 0
	for _, item := range templates {
		if item.State == "ready" {
			count++
		}
	}
	return count
}
