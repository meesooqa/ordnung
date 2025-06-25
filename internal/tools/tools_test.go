package tools

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseYtDuration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{
			name:     "only seconds",
			input:    "PT37S",
			expected: 37 * time.Second,
			wantErr:  false,
		},
		{
			name:     "minutes and seconds",
			input:    "PT2M48S",
			expected: 2*time.Minute + 48*time.Second,
			wantErr:  false,
		},
		{
			name:     "hours, minutes and seconds",
			input:    "PT1H5M30S",
			expected: time.Hour + 5*time.Minute + 30*time.Second,
			wantErr:  false,
		},
		{
			name:     "only minutes",
			input:    "PT15M",
			expected: 15 * time.Minute,
			wantErr:  false,
		},
		{
			name:     "fractional seconds",
			input:    "PT0.5S",
			expected: 500 * time.Millisecond,
			wantErr:  false,
		},
		{
			name:     "invalid format - missing PT",
			input:    "30S",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "invalid format - wrong order",
			input:    "PT30S5M",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseYtDuration(tt.input)

			if tt.wantErr {
				assert.Error(t, err, "expected error for input %s", tt.input)
				return
			}

			assert.NoError(t, err, "unexpected error for input %s", tt.input)
			assert.Equal(t, tt.expected, got, "duration mismatch for input %s", tt.input)
		})
	}
}
