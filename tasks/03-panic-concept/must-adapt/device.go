package devices

type DeviceType int

const (
	DeviceTypeUnspecified DeviceType = iota
	DeviceTypeArduinoMega2560
	DeviceTypeRaspberryPi4
	DeviceTypeLogiBone2
)
