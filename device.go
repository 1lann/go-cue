package cue

// #include "CUESDK.h"
import "C"

import (
	"unsafe"
)

type deviceInfo C.CorsairDeviceInfo

type DeviceType int32

var (
	CDT_Unknown  DeviceType = C.CDT_Unknown
	CDT_Mouse    DeviceType = C.CDT_Mouse
	CDT_Keyboard DeviceType = C.CDT_Keyboard
	CDT_Headset  DeviceType = C.CDT_Headset
)

func (t DeviceType) String() string {
	switch t {
	case CDT_Unknown:
		return "CDT_Unknown"
	case CDT_Mouse:
		return "CDT_Mouse"
	case CDT_Keyboard:
		return "CDT_Keyboard"
	case CDT_Headset:
		return "CDT_Headset"
	default:
		return "Unknown"
	}
}

type PhysicalLayoutType int32

var (
	CPL_Invalid PhysicalLayoutType = C.CPL_Invalid

	CPL_US PhysicalLayoutType = C.CPL_US
	CPL_UK PhysicalLayoutType = C.CPL_UK
	CPL_BR PhysicalLayoutType = C.CPL_BR
	CPL_JP PhysicalLayoutType = C.CPL_JP
	CPL_KR PhysicalLayoutType = C.CPL_KR

	CPL_Zones1 PhysicalLayoutType = C.CPL_Zones1
	CPL_Zones2 PhysicalLayoutType = C.CPL_Zones2
	CPL_Zones3 PhysicalLayoutType = C.CPL_Zones3
	CPL_Zones4 PhysicalLayoutType = C.CPL_Zones4
)

func (t PhysicalLayoutType) String() string {
	switch t {
	case CPL_Invalid:
		return "CPL_Invalid"

	case CPL_US:
		return "CPL_US"
	case CPL_UK:
		return "CPL_UK"
	case CPL_BR:
		return "CPL_BR"
	case CPL_JP:
		return "CPL_JP"
	case CPL_KR:
		return "CPL_KR"

	case CPL_Zones1:
		return "CPL_Zones1"
	case CPL_Zones2:
		return "CPL_Zones2"
	case CPL_Zones3:
		return "CPL_Zones3"
	case CPL_Zones4:
		return "CPL_Zones4"
	default:
		return "Unknown"
	}
}

type LogicalLayoutType int32

var (
	CLL_Invalid LogicalLayoutType = C.CLL_Invalid
	CLL_US_Int  LogicalLayoutType = C.CLL_US_Int
	CLL_NA      LogicalLayoutType = C.CLL_NA
	CLL_EU      LogicalLayoutType = C.CLL_EU
	CLL_UK      LogicalLayoutType = C.CLL_UK
	CLL_BE      LogicalLayoutType = C.CLL_BE
	CLL_BR      LogicalLayoutType = C.CLL_BR
	CLL_CH      LogicalLayoutType = C.CLL_CH
	CLL_CN      LogicalLayoutType = C.CLL_CN
	CLL_DE      LogicalLayoutType = C.CLL_DE
	CLL_ES      LogicalLayoutType = C.CLL_ES
	CLL_FR      LogicalLayoutType = C.CLL_FR
	CLL_IT      LogicalLayoutType = C.CLL_IT
	CLL_ND      LogicalLayoutType = C.CLL_ND
	CLL_RU      LogicalLayoutType = C.CLL_RU
	CLL_JP      LogicalLayoutType = C.CLL_JP
	CLL_KR      LogicalLayoutType = C.CLL_KR
	CLL_TW      LogicalLayoutType = C.CLL_TW
	CLL_MEX     LogicalLayoutType = C.CLL_MEX
)

func (t LogicalLayoutType) String() string {
	switch t {
	case CLL_Invalid:
		return "CLL_Invalid"
	case CLL_US_Int:
		return "CLL_US_Int"
	case CLL_NA:
		return "CLL_NA"
	case CLL_EU:
		return "CLL_EU"
	case CLL_UK:
		return "CLL_UK"
	case CLL_BE:
		return "CLL_BE"
	case CLL_BR:
		return "CLL_BR"
	case CLL_CH:
		return "CLL_CH"
	case CLL_CN:
		return "CLL_CN"
	case CLL_DE:
		return "CLL_DE"
	case CLL_ES:
		return "CLL_ES"
	case CLL_FR:
		return "CLL_FR"
	case CLL_IT:
		return "CLL_IT"
	case CLL_ND:
		return "CLL_ND"
	case CLL_RU:
		return "CLL_RU"
	case CLL_JP:
		return "CLL_JP"
	case CLL_KR:
		return "CLL_KR"
	case CLL_TW:
		return "CLL_TW"
	case CLL_MEX:
		return "CLL_MEX"
	default:
		return "Unknown"
	}
}

type DeviceCap int32

var (
	CDC_None     DeviceCap = C.CDC_None
	CDC_Lighting DeviceCap = C.CDC_Lighting
)

func (c DeviceCap) String() string {
	switch c {
	case CDC_None:
		return "CDC_None"
	case CDC_Lighting:
		return "CDC_Lighting"
	default:
		return "Unknown"
	}
}

type Device struct {
	Type           DeviceType
	Model          string
	PhysicalLayout PhysicalLayoutType
	LogicalLayout  LogicalLayoutType
	CapsMask       DeviceCap
}

type internalDevice struct {
	deviceType     int32
	model          *C.char
	physicalLayout int32
	logicalLayout  int32
	capsMask       int32
}

func GetAllDevices() ([]Device, error) {
	count := int(C.CorsairGetDeviceCount())

	if count < 0 {
		return []Device{}, getLastError()
	}

	devices := make([]Device, count)

	for i := 0; i < count; i++ {
		thisDevice := *((*internalDevice)(unsafe.Pointer(C.CorsairGetDeviceInfo(C.int(i)))))
		model := C.GoString(thisDevice.model)

		devices[i] = Device{
			Type:           DeviceType(thisDevice.deviceType),
			Model:          model,
			PhysicalLayout: PhysicalLayoutType(thisDevice.physicalLayout),
			LogicalLayout:  LogicalLayoutType(thisDevice.logicalLayout),
			CapsMask:       DeviceCap(thisDevice.capsMask),
		}
	}

	return devices, nil
}
