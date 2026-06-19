package evidence

import "fmt"

func inspectSource(
	root string,
	policy Policy,
	source PrivateSource,
	status *Status,
	artifactRefs map[string]bool,
) (SourceStatus, error) {
	switch source.Format {
	case "jsonl":
		if source.Key == "learning" {
			return inspectLearningSource(root, policy, source, status, artifactRefs)
		}
		return inspectJSONLSource(root, source, status)
	case "directory":
		return inspectDirectorySource(root, source, status)
	default:
		return newSourceStatus(source), fmt.Errorf("evidence source %q has unsupported format", source.Key)
	}
}

func inspectJSONLSource(root string, source PrivateSource, status *Status) (SourceStatus, error) {
	sourceStatus := newSourceStatus(source)
	count, present, last, err := countJSONL(root, source.Path)
	if err != nil {
		return sourceStatus, err
	}
	sourceStatus.Present = present
	sourceStatus.Count = count
	status.NodeCount += count
	status.ByNodeKind[source.NodeKind] += count
	updateLastObserved(status, last)
	return sourceStatus, nil
}

func inspectDirectorySource(root string, source PrivateSource, status *Status) (SourceStatus, error) {
	sourceStatus := newSourceStatus(source)
	count, present, err := countDirectoryFiles(root, source.Path)
	if err != nil {
		return sourceStatus, err
	}
	sourceStatus.Present = present
	sourceStatus.Count = count
	status.NodeCount += count
	status.ByNodeKind[source.NodeKind] += count
	return sourceStatus, nil
}
