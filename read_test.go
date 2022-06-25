package distroinfo

import (
	"reflect"
	"testing"
)

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
				Created:  "1996-06-17",
				Release:  "1996-12-12",
				Eol:      "1998-06-05",
			},
			wantErr: false,
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
