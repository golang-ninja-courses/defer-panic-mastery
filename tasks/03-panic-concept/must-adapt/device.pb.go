//nolint:stylecheck,revive
package devices

type ProtoDeviceType int32

const (
	DeviceType_DEVICE_TYPE_UNSPECIFIED       ProtoDeviceType = 0
	DeviceType_DEVICE_TYPE_ARDUINO_MEGA_2560 ProtoDeviceType = 1
	DeviceType_DEVICE_TYPE_RASPBERRY_PI_4    ProtoDeviceType = 2
	DeviceType_DEVICE_TYPE_LOGI_BONE_2       ProtoDeviceType = 3
)
