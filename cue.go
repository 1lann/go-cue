// +build windows

package cue

/*
#cgo CFLAGS: -Iinclude
#cgo amd64 LDFLAGS: -L${SRCDIR}/lib -lCUESDK.x64_2013
#cgo 386 LDFLAGS: -L${SRCDIR}/lib -lCUESDK_2013


#include "CUESDK.h"
*/
import "C"

import (
	"errors"
)

type ProtocolDetails struct {
	SDKVersion            string
	ServerVersion         string
	SDKProtocolVersion    int
	ServerProtocolVersion int
	BreakingChanges       bool
}

var (
	internalCE_Success                  C.CorsairError = C.CE_Success
	internalCE_ServerNotFound           C.CorsairError = C.CE_ServerNotFound
	internalCE_NoControl                C.CorsairError = C.CE_NoControl
	internalCE_ProtocolHandshakeMissing C.CorsairError = C.CE_ProtocolHandshakeMissing
	internalCE_IncompatibleProtocol     C.CorsairError = C.CE_IncompatibleProtocol
	internalCE_InvalidArguments         C.CorsairError = C.CE_InvalidArguments
)

var (
	CE_Success                  error = nil
	CE_ServerNotFound                 = errors.New("cue: could not find CUE server")
	CE_NoControl                      = errors.New("cue: cannot take control")
	CE_ProtocolHandshakeMissing       = errors.New("cue: protocol handshake missing (did you call cue.Initialize()?)")
	CE_IncompatibleProtocol           = errors.New("cue: incompatible protocol (is the server and client up to date?)")
	CE_InvalidArguments               = errors.New("cue: invalid arguments were supplied")
	ErrUnknown                        = errors.New("cue: unknown error")
)

func getLastError() error {
	lastError := C.CorsairGetLastError()
	switch lastError {
	case internalCE_Success:
		return CE_Success
	case internalCE_ServerNotFound:
		return CE_ServerNotFound
	case internalCE_NoControl:
		return CE_NoControl
	case internalCE_ProtocolHandshakeMissing:
		return CE_ProtocolHandshakeMissing
	case internalCE_IncompatibleProtocol:
		return CE_IncompatibleProtocol
	case internalCE_InvalidArguments:
		return CE_InvalidArguments
	default:
		return ErrUnknown
	}
}

func RequestExclusiveControl() error {
	if C.CorsairRequestControl(C.CorsairAccessMode(0)) == 1 {
		return nil
	} else {
		return getLastError()
	}
}

func Initialize() (ProtocolDetails, error) {
	rawDetails := C.CorsairPerformProtocolHandshake()

	if rawDetails.sdkVersion == nil || rawDetails.serverVersion == nil {
		return ProtocolDetails{}, getLastError()
	}

	details := ProtocolDetails{
		SDKVersion:            C.GoString(rawDetails.sdkVersion),
		ServerVersion:         C.GoString(rawDetails.serverVersion),
		SDKProtocolVersion:    int(rawDetails.sdkProtocolVersion),
		ServerProtocolVersion: int(rawDetails.serverProtocolVersion),
		BreakingChanges:       rawDetails.breakingChanges == 1,
	}

	return details, nil
}
