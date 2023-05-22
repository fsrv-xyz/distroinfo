package distroinfo

import (
	"embed"
	"encoding/json"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

// generates json files for distro information
// run go generate to regenerate the files
// see https://sources.debian.org/src/distro-info-data/ for updating the version
//go:generate bash data/generator.sh debian,ubuntu 0.53

//go:embed data/*.json
var distroInfoFS embed.FS

type TimeStamp struct {
	time.Time
}

func (t *TimeStamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
		return nil
	}
	var err error
	t.Time, err = time.Parse("2006-01-02", s)
	return err
}

// DistroInfo holds distro information
type DistroInfo struct {
	Version  string     `json:"version"`
	Codename string     `json:"codename"`
	Series   string     `json:"series"`
	Created  *TimeStamp `json:"created"`
	Release  *TimeStamp `json:"release"`
	Eol      *TimeStamp `json:"eol"`
	EolLts   *TimeStamp `json:"eol-lts"`
	EolElts  *TimeStamp `json:"eol-elts"`
	EolEsm   *TimeStamp `json:"eol-esm"`
}

var distroInfoStore = map[string][]DistroInfo{}

func init() {
	parseDistroInfoData()
}

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
