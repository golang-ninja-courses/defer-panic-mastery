package envconfig

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDuration(t *testing.T) {
	const envName = "TEST_PARSE_DURATION_ENV"

	cases := []struct {
		name        string
		envValue    string
		expected    time.Duration
		expectError bool
	}{
		{
			name:        "valid",
			envValue:    "300ms",
			expected:    300 * time.Millisecond,
			expectError: false,
		},
		{
			name:        "valid complex",
			envValue:    "23h15m10s373ms939us6ns",
			expected:    23*time.Hour + 15*time.Minute + 10*time.Second + 373*time.Millisecond + 939*time.Microsecond + 6*time.Nanosecond,
			expectError: false,
		},
		{
			name:        "missing units",
			envValue:    "300",
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid units",
			envValue:    "3d20m",
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(envName, tt.envValue)

			duration, err := ParseDuration(envName)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, duration)
		})
	}
}

func TestParseDuration_NoEnv(t *testing.T) {
	const envName = "TEST_PARSE_DURATION_ENV"

	cases := []struct {
		name     string
		envValue string
	}{
		{
			name:     "valid",
			envValue: "300ms",
		},
		{
			name:     "valid complex",
			envValue: "23h15m10s373ms939us6ns",
		},
		{
			name:     "missing units",
			envValue: "300",
		},
		{
			name:     "invalid units",
			envValue: "3d20m",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			duration, err := ParseDuration(envName)
			assert.Error(t, err)
			assert.Empty(t, duration)
		})
	}
}

func TestMustParseDuration(t *testing.T) {
	const envName = "TEST_PARSE_DURATION_ENV"

	t.Run("ok", func(t *testing.T) {
		t.Setenv(envName, "300ms")

		duration := MustParseDuration(envName)
		assert.Equal(t, 300*time.Millisecond, duration)
	})

	t.Run("missing units", func(t *testing.T) {
		t.Setenv(envName, "300")

		assert.PanicsWithError(t, `parse duration: time: missing unit in duration "300"`, func() {
			MustParseDuration(envName)
		})
	})

	t.Run("invalid units", func(t *testing.T) {
		t.Setenv(envName, "3d20m")

		assert.PanicsWithError(t, `parse duration: time: unknown unit "d" in duration "3d20m"`, func() {
			MustParseDuration(envName)
		})
	})

	t.Run("no env", func(t *testing.T) {
		assert.PanicsWithError(t, `parse duration: time: invalid duration ""`, func() {
			MustParseDuration(envName)
		})
	})
}
