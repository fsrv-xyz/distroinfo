package distroinfo

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"
)

// generates json files for distro information
// run go generate to regenerate the files
// see https://sources.debian.org/src/distro-info-data/ for updating the version
//go:generate bash data/generator.sh debian,ubuntu 0.58

//go:embed data/*.json
var distroInfoFS embed.FS

var distroInfoStore = map[string][]DistroInfo{}

func init() {
	parseDistroInfoData()
}

// parseDistroInfoData parses the embedded distro info data
func parseDistroInfoData() {
	// get file list from embedded data directory
	dirEntries, distroInfoFSReadError := distroInfoFS.ReadDir("data")
	if distroInfoFSReadError != nil {
		panic(distroInfoFSReadError)
	}

	// itarate over the file list and parse each file
	for _, entry := range dirEntries {
		if entry.IsDir() {
			continue
		}
		parseJsonEntryFile(entry)
	}
}

// open and parse a single json file to distroInfoStore map
func parseJsonEntryFile(entry fs.DirEntry) {
	file, fileOpenError := distroInfoFS.Open("data/" + entry.Name())
	if fileOpenError != nil {
		return
	}
	defer file.Close()

	var info []DistroInfo
	json.NewDecoder(file).Decode(&info)

	distroInfoStore[fileNameWithoutExtTrimSuffix(entry.Name())] = info
}

// helper function for getting the file name without the extension
func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
