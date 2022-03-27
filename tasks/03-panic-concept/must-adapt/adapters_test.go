package devices

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func TestAdaptProtoDeviceTypeToDeviceType(t *testing.T) {
	cases := []struct {
		name    string
		in      ProtoDeviceType
		out     DeviceType
		wantErr bool
	}{
		{
			name: "unspecified",
			in:   DeviceType_DEVICE_TYPE_UNSPECIFIED,
			out:  DeviceTypeUnspecified,
		},
		{
			name: "arduino",
			in:   DeviceType_DEVICE_TYPE_ARDUINO_MEGA_2560,
			out:  DeviceTypeArduinoMega2560,
		},
		{
			name: "raspberry",
			in:   DeviceType_DEVICE_TYPE_RASPBERRY_PI_4,
			out:  DeviceTypeRaspberryPi4,
		},
		{
			name: "bone",
			in:   DeviceType_DEVICE_TYPE_LOGI_BONE_2,
			out:  DeviceTypeLogiBone2,
		},
		{
			name:    "unknown 1",
			in:      -1,
			wantErr: true,
		},
		{
			name:    "unknown 2",
			in:      ProtoDeviceType(4 + rand.Int31n(math.MaxInt32-4)),
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("adapt", func(t *testing.T) {
				d, err := AdaptProtoDeviceTypeToDeviceType(tt.in)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.out, d)
				}
			})

			t.Run("must adapt", func(t *testing.T) {
				if tt.wantErr {
					assert.Panics(t, func() {
						_ = MustAdaptProtoDeviceTypeToDeviceType(tt.in)
					})
				} else {
					assert.NotPanics(t, func() {
						d := MustAdaptProtoDeviceTypeToDeviceType(tt.in)
						assert.Equal(t, tt.out, d)
					})
				}
			})
		})
	}
}

func TestAdaptDeviceTypeToProtoDeviceType(t *testing.T) {
	cases := []struct {
		name    string
		in      DeviceType
		out     ProtoDeviceType
		wantErr bool
	}{
		{
			name: "unspecified",
			in:   DeviceTypeUnspecified,
			out:  DeviceType_DEVICE_TYPE_UNSPECIFIED,
		},
		{
			name: "arduino",
			in:   DeviceTypeArduinoMega2560,
			out:  DeviceType_DEVICE_TYPE_ARDUINO_MEGA_2560,
		},
		{
			name: "raspberry",
			in:   DeviceTypeRaspberryPi4,
			out:  DeviceType_DEVICE_TYPE_RASPBERRY_PI_4,
		},
		{
			name: "bone",
			in:   DeviceTypeLogiBone2,
			out:  DeviceType_DEVICE_TYPE_LOGI_BONE_2,
		},
		{
			name:    "unknown 1",
			in:      -1,
			wantErr: true,
		},
		{
			name:    "unknown 2",
			in:      DeviceType(4 + rand.Intn(math.MaxInt-4)),
			wantErr: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("adapt", func(t *testing.T) {
				d, err := AdaptDeviceTypeToProtoDeviceType(tt.in)
				if tt.wantErr {
					require.Error(t, err)
				} else {
					require.NoError(t, err)
					assert.Equal(t, tt.out, d)
				}
			})

			t.Run("must adapt", func(t *testing.T) {
				if tt.wantErr {
					assert.Panics(t, func() {
						_ = MustAdaptDeviceTypeToProtoDeviceType(tt.in)
					})
				} else {
					assert.NotPanics(t, func() {
						d := MustAdaptDeviceTypeToProtoDeviceType(tt.in)
						assert.Equal(t, tt.out, d)
					})
				}
			})
		})
	}
}
