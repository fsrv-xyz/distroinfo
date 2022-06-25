package distroinfo

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	parseDistroInfoData()
	if len(distroInfoStore) < 2 {
		t.Error("insufficient distro infos found")
	}
}
