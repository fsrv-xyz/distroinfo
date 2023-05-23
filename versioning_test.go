package distroinfo

import (
	"testing"
	"time"
)

func TestGetSupportPeriod(t *testing.T) {
	now := TimeStamp{time.Now()}
	future := TimeStamp{now.AddDate(1, 0, 0)}
	past := TimeStamp{now.AddDate(-1, 0, 0)}

	tests := []struct {
		name        string
		distroInfo  DistroInfo
		wantSupport SupportPeriod
	}{
		{
			name:        "Standard Support - Upcoming LTS",
			distroInfo:  DistroInfo{Eol: &future, EolLts: &future},
			wantSupport: StandardSupport,
		},
		{
			name:        "Standard Support - Past LTS",
			distroInfo:  DistroInfo{Eol: &future},
			wantSupport: StandardSupport,
		},
		{
			name:        "LTS - Upcoming ELTS",
			distroInfo:  DistroInfo{Eol: &past, EolLts: &future, EolElts: &future},
			wantSupport: LongTermSupport,
		},
		{
			name:        "LTS - Past ELTS",
			distroInfo:  DistroInfo{EolLts: &future, EolElts: &past},
			wantSupport: LongTermSupport,
		},
		{
			name:        "Standard & LTS - Upcoming ELTS",
			distroInfo:  DistroInfo{Eol: &future, EolLts: &future, EolElts: &future},
			wantSupport: StandardSupport,
		},
		{
			name:        "Standard Support - Past LTS, Upcoming ELTS",
			distroInfo:  DistroInfo{Eol: &future, EolLts: &past, EolElts: &future},
			wantSupport: StandardSupport,
		},
		{
			name:        "Unsupported - Past Standard, LTS, ELTS",
			distroInfo:  DistroInfo{Eol: &past, EolLts: &past, EolElts: &past},
			wantSupport: Unsupported,
		},
		{
			name:        "Unreleased",
			distroInfo:  DistroInfo{Release: &future},
			wantSupport: Unreleased,
		},
		{
			name:        "Unreleased - Past EOL",
			distroInfo:  DistroInfo{Release: &future, Eol: &past},
			wantSupport: Unreleased,
		},
		{
			name:        "Ubuntu Special Case",
			distroInfo:  DistroInfo{Eol: &past, EolEsm: &future},
			wantSupport: ExpandedSecurityMaintenance,
		},
		{
			name:        "ELTS",
			distroInfo:  DistroInfo{Eol: &past, EolLts: &past, EolElts: &future},
			wantSupport: ExtendedLongTermSupport,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSupport := getSupportPeriod(tt.distroInfo)
			if gotSupport != tt.wantSupport {
				t.Errorf("getSupportPeriod() got = %v, want = %v", gotSupport, tt.wantSupport)
			}
		})
	}
}

func TestCheckPlausibility(t *testing.T) {
	now := TimeStamp{time.Now()}
	past := TimeStamp{now.Add(-time.Hour)}
	future := TimeStamp{now.Add(time.Hour)}

	tests := []struct {
		name    string
		d       DistroInfo
		wantErr bool
	}{
		{
			name:    "Eol is after EolLts",
			d:       DistroInfo{Eol: &future, EolLts: &now},
			wantErr: true,
		},
		{
			name:    "EolLts is after EolElts",
			d:       DistroInfo{EolLts: &future, EolElts: &now},
			wantErr: true,
		},
		{
			name:    "Eol is after EolEsm",
			d:       DistroInfo{Eol: &future, EolEsm: &now},
			wantErr: true,
		},
		{
			name:    "Plausible dates",
			d:       DistroInfo{Eol: &past, EolLts: &now, EolElts: &now, EolEsm: &future},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := checkPlausibility(tt.d); (err != nil) != tt.wantErr {
				t.Errorf("checkPlausibility() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTestdataPlausibility(t *testing.T) {
	t.Parallel()
	for _, distro := range GetDistros() {
		for _, d := range distroInfoStore[distro] {
			t.Run(d.Version+d.Codename, func(t *testing.T) {
				if err := checkPlausibility(d); err != nil {
					t.Errorf("checkPlausibility() error = %v", err)
				}
			})
		}
	}
}
