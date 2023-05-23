package distroinfo

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
