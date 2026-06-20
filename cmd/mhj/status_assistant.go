package main

import "github.com/kimsemi-home/myhome-jarvis/internal/commandcenter"

func runAssistant(root string, args []string) error {
	if len(args) != 1 {
		return usage()
	}
	switch args[0] {
	case "status":
		return assistantStatus(root)
	case "vision-audit":
		return assistantVisionAudit(root)
	default:
		return usage()
	}
}

func assistantStatus(root string) error {
	status, err := commandcenter.StatusForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}

func assistantVisionAudit(root string) error {
	audit, err := commandcenter.VisionAuditForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(audit)
}

func workItemStatus(root string) error {
	status, err := commandcenter.WorkItemForRoot(root)
	if err != nil {
		return err
	}
	return writeJSON(status)
}
