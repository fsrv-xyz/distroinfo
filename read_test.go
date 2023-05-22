package distroinfo

import (
	"reflect"
	"testing"
	"time"
)

func mustParseTime(layout string, value string) *TimeStamp {
	result, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &TimeStamp{result}
}

const timeFormat = "2006-01-02"

func TestGetDistroInfoByVersion(t *testing.T) {
	type args struct {
		distro  string
		version string
	}
	tests := []struct {
		name    string
		args    args
		want    *DistroInfo
		wantErr bool
	}{
		{
			name: "existing version",
			args: args{
				distro:  "debian",
				version: "1.2",
			},
			want: &DistroInfo{
				Version:  "1.2",
				Codename: "Rex",
				Series:   "rex",
				Created:  mustParseTime(timeFormat, "1996-06-17"),
				Release:  mustParseTime(timeFormat, "1996-12-12"),
				Eol:      mustParseTime(timeFormat, "1998-06-05"),
			},
			wantErr: false,
		},
		{
			name: "non-existing version",
			args: args{
				distro:  "debian",
				version: "non-existing",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "non-existing distro",
			args: args{
				distro:  "non-existing",
				version: "1.2",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDistroInfoByVersion(tt.args.distro, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDistroInfoByVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDistroInfoByVersion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDistros(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "get distros",
			want: []string{"debian", "ubuntu"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDistros(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDistros() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDistroVersions(t *testing.T) {
	tests := []struct {
		name    string
		distro  string
		wantErr bool
	}{
		{
			name:    "get debian versions",
			distro:  "debian",
			wantErr: false,
		},
		{
			name:    "get ubuntu versions",
			distro:  "ubuntu",
			wantErr: false,
		},
		{
			name:    "get non-existing distro versions",
			distro:  "non-existing",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := GetDistroVersions(tt.distro); (err != nil) != tt.wantErr {
				t.Errorf("GetDistroVersions() error = %v, wantErr %v", err, tt.wantErr)
			} else if !tt.wantErr && len(got) == 0 {
				t.Errorf("GetDistroVersions() got = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestGetSupportedDistroVersions(t *testing.T) {
	for _, distro := range GetDistros() {
		t.Run(distro, func(t *testing.T) {
			versions, err := GetSupportedDistroVersions(distro)
			if err != nil {
				t.Error(err)
			}

			if len(versions) == 0 {
				t.Error("No supported versions found")
			}

			t.Log(versions)
		})
	}
}
