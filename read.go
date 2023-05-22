package distroinfo

import (
	"errors"
	"sort"
	"strings"
	"time"
)

var (
	DistroFamilyNotFoundError  = errors.New("distro family not found")
	DistroVersionNotFoundError = errors.New("distro version not found")
)

type SupportPeriod int

const (
	Unsupported SupportPeriod = iota
	Unreleased
	StandardSupport
	LongTermSupport
	ExtendedLongTermSupport
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

func getSupportPeriod(distroInfoEntry DistroInfo) SupportPeriod {
	var isLts bool
	var isELts bool
	var isStandard bool

	if distroInfoEntry.Release != nil && distroInfoEntry.Release.After(time.Now()) {
		return Unreleased
	}

	if distroInfoEntry.Eol != nil && distroInfoEntry.Eol.After(time.Now()) {
		isStandard = true
	}

	if distroInfoEntry.EolLts != nil && distroInfoEntry.EolLts.After(time.Now()) {
		isLts = true
	}

	if distroInfoEntry.EolEsm != nil && distroInfoEntry.EolEsm.After(time.Now()) {
		isLts = true
	}

	if distroInfoEntry.EolElts != nil && distroInfoEntry.EolElts.After(time.Now()) {
		isELts = true
	}

	if isStandard {
		return StandardSupport
	}

	if isLts && !isStandard {
		return LongTermSupport
	}

	if isELts && !isStandard && !isLts {
		return ExtendedLongTermSupport
	}

	return Unsupported
}
