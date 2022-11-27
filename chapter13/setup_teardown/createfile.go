package setup_teardown

import (
	"os"
	"testing"
)

func createFile(t *testing.T) (string, error) {
	f, err := os.Create("tempFile")
	if err != nil {
		return "", err
	}
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}
func TestFileProcessing(t *testing.T) {
	_, err := createFile(t)
	if err != nil {
		t.Fatal(err)
	}
}
