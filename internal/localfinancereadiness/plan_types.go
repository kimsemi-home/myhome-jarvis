package localfinancereadiness

type Handle struct {
	Service string `json:"service"`
	Account string `json:"account"`
}

type Schedule struct {
	Day      int    `json:"day"`
	Hour     int    `json:"hour"`
	Minute   int    `json:"minute"`
	Timezone string `json:"timezone"`
}

type Launchd struct {
	TemplatePath     string   `json:"template_path"`
	Label            string   `json:"label"`
	ProgramArguments []string `json:"program_arguments"`
}
