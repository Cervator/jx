package cmd

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDevPodLabel(t *testing.T) {
	labels := []string{"go", "nodejs", "maven"}

	fileToValues := map[string]string{
		"Jenkinsfile.nodejs": "nodejs",
	}

	for fileName, label := range fileToValues {
		testFile := path.Join("test_data", "jenkinsfiles", fileName)

		answer, err := findDevPodLabelFromJenkinsfile(testFile, labels)
		assert.NoError(t, err, "Failed to find label for file %s", testFile)
		if err == nil {
			assert.Equal(t, label, answer, "Failed to find label for file %s", testFile)
		}
	}
}
