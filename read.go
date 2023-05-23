package distroinfo

import (
	"errors"
	"sort"
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

func GetDistros() []string {
	var distributions []string
	for distro := range distroInfoStore {
		distributions = append(distributions, distro)
	}
	sort.Strings(distributions)
	return distributions
}

func GetDistroVersions(distro string) ([]string, error) {
	family, familyFound := distroInfoStore[strings.ToLower(distro)]
	if !familyFound {
		return nil, DistroFamilyNotFoundError
	}

	var versions []string
	for _, distroInfoEntry := range family {
		versions = append(versions, distroInfoEntry.Version)
	}
	return versions, nil
}

func GetSupportedDistroVersions(distro string) (map[string]SupportPeriod, error) {
	family, familyFound := distroInfoStore[strings.ToLower(distro)]
	if !familyFound {
		return nil, DistroFamilyNotFoundError
	}

	versions := make(map[string]SupportPeriod)

	for _, distroInfoEntry := range family {
		period := getSupportPeriod(distroInfoEntry)
		if period == Unsupported {
			continue
		}
		if period == Unreleased {
			continue
		}

		versions[distroInfoEntry.Version] = period
	}
	return versions, nil
}
