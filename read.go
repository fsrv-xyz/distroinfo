package distroinfo

import (
	"errors"
	"strings"
)

var (
	DistroFamilyNotFoundError  = errors.New("distro family not found")
	DistroVersionNotFoundError = errors.New("distro version not found")
)

func GetDistroInfoByVersion(distro string, version string) (*DistroInfo, error) {
	family, familyFound := distroInfoStore[strings.ToLower(distro)]
	if !familyFound {
		return nil, DistroFamilyNotFoundError
	}

	for _, distroInfoEntry := range family {
		if distroInfoEntry.Version == version {
			return &distroInfoEntry, nil
		}
	}
	return nil, DistroVersionNotFoundError
}
