package primitive

import (
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
	"github.ibm.com/dash/dash_utils/dashtest"
)

func TestMain(m *testing.M) {

	dashtest.ControlCoverage(m)
}
func TestTempFile(t *testing.T) {
	_, err := tempfile("/invalid/invalid", "txt")
	if err == nil {
		t.Error("Expected error but got no error")
	}
}

func TestWithMode(t *testing.T) {
	result := WithMode(ModeCombo)
	if result == nil {
		t.Error("Expected string but got no result")
	}
}

func TestTransform(t *testing.T) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/birds.png")
	f, _ := os.Open(imgPath)
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "png", 1, opts)
	if err != nil {
		t.Errorf("Expected no error but got error:: %v", err)
	}
}
func TestTransform_PrimitiveError(t *testing.T) {
	h, _ := homedir.Dir()
	imgPath := filepath.Join(h, "img/birds.png")
	f, _ := os.Open(imgPath)
	opts := WithMode(ModeCombo)
	_, err := Transform(f, "", 1, opts)
	if err == nil {
		t.Errorf("Expected error but got no error ")
	}
}
