package distroinfo

import (
	"errors"
	"time"
)

type SupportPeriod int

const (
	Unsupported SupportPeriod = iota
	Unreleased
	StandardSupport
	LongTermSupport
	ExtendedLongTermSupport
	ExpandedSecurityMaintenance
)

func getSupportPeriod(distroInfoEntry DistroInfo) SupportPeriod {
	var isLts bool
	var isELts bool
	var isStandard bool
	var isEsm bool

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
		isEsm = true
	}

	if distroInfoEntry.EolElts != nil && distroInfoEntry.EolElts.After(time.Now()) {
		isELts = true
	}

	if isStandard {
		return StandardSupport
	}

	if isEsm {
		return ExpandedSecurityMaintenance
	}

	if isLts && !isStandard {
		return LongTermSupport
	}

	if isELts && !isStandard && !isLts {
		return ExtendedLongTermSupport
	}

	return Unsupported
}

func checkPlausibility(distroInfoEntry DistroInfo) error {
	if distroInfoEntry.Eol != nil && distroInfoEntry.EolLts != nil && distroInfoEntry.Eol.After(distroInfoEntry.EolLts.Time) {
		return errors.New("Eol is after EolLts")
	}
	if distroInfoEntry.EolLts != nil && distroInfoEntry.EolElts != nil && distroInfoEntry.EolLts.After(distroInfoEntry.EolElts.Time) {
		return errors.New("EolLts is after EolElts")
	}
	if distroInfoEntry.Eol != nil && distroInfoEntry.EolEsm != nil && distroInfoEntry.Eol.After(distroInfoEntry.EolEsm.Time) {
		return errors.New("Eol is after EolEsm")
	}
	return nil
}
