package evidence

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func inspectLearningIntegrity(root string, policy Policy, status *IntegrityStatus) error {
	source, ok := learningSource(policy)
	if !ok {
		return nil
	}
	file, err := os.Open(filepath.Join(root, filepath.FromSlash(source.Path)))
	if errors.Is(err, os.ErrNotExist) {
		return nil
	}
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		err := inspectIntegrityLine(root, policy, status, scanner.Text())
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}
