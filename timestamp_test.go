package distroinfo

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeStamp_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantTime time.Time
		wantErr  bool
	}{
		{
			name:     "valid date",
			input:    `"2023-05-23"`,
			wantTime: time.Date(2023, 5, 23, 0, 0, 0, 0, time.UTC),
			wantErr:  false,
		},
		{
			name:     "empty date",
			input:    `""`,
			wantTime: time.Time{},
			wantErr:  false,
		},
		{
			name:     "invalid date format",
			input:    `"23-05-2023"`,
			wantTime: time.Time{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var timestamp TimeStamp
			err := json.Unmarshal([]byte(tt.input), &timestamp)

			if (err != nil) != tt.wantErr {
				t.Errorf("TimeStamp.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && !timestamp.Time.Equal(tt.wantTime) {
				t.Errorf("TimeStamp.UnmarshalJSON() got = %v, want = %v", timestamp, tt.wantTime)
			}
		})
	}
}
