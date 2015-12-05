package cue

// #include "CUESDK.h"
import "C"

var (
	CLI_Invalid                  LedId = C.CLI_Invalid
	CLK_Escape                   LedId = C.CLK_Escape
	CLK_F1                       LedId = C.CLK_F1
	CLK_F2                       LedId = C.CLK_F2
	CLK_F3                       LedId = C.CLK_F3
	CLK_F4                       LedId = C.CLK_F4
	CLK_F5                       LedId = C.CLK_F5
	CLK_F6                       LedId = C.CLK_F6
	CLK_F7                       LedId = C.CLK_F7
	CLK_F8                       LedId = C.CLK_F8
	CLK_F9                       LedId = C.CLK_F9
	CLK_F10                      LedId = C.CLK_F10
	CLK_F11                      LedId = C.CLK_F11
	CLK_GraveAccentAndTilde      LedId = C.CLK_GraveAccentAndTilde
	CLK_1                        LedId = C.CLK_1
	CLK_2                        LedId = C.CLK_2
	CLK_3                        LedId = C.CLK_3
	CLK_4                        LedId = C.CLK_4
	CLK_5                        LedId = C.CLK_5
	CLK_6                        LedId = C.CLK_6
	CLK_7                        LedId = C.CLK_7
	CLK_8                        LedId = C.CLK_8
	CLK_9                        LedId = C.CLK_9
	CLK_0                        LedId = C.CLK_0
	CLK_MinusAndUnderscore       LedId = C.CLK_MinusAndUnderscore
	CLK_Tab                      LedId = C.CLK_Tab
	CLK_Q                        LedId = C.CLK_Q
	CLK_W                        LedId = C.CLK_W
	CLK_E                        LedId = C.CLK_E
	CLK_R                        LedId = C.CLK_R
	CLK_T                        LedId = C.CLK_T
	CLK_Y                        LedId = C.CLK_Y
	CLK_U                        LedId = C.CLK_U
	CLK_I                        LedId = C.CLK_I
	CLK_O                        LedId = C.CLK_O
	CLK_P                        LedId = C.CLK_P
	CLK_BracketLeft              LedId = C.CLK_BracketLeft
	CLK_CapsLock                 LedId = C.CLK_CapsLock
	CLK_A                        LedId = C.CLK_A
	CLK_S                        LedId = C.CLK_S
	CLK_D                        LedId = C.CLK_D
	CLK_F                        LedId = C.CLK_F
	CLK_G                        LedId = C.CLK_G
	CLK_H                        LedId = C.CLK_H
	CLK_J                        LedId = C.CLK_J
	CLK_K                        LedId = C.CLK_K
	CLK_L                        LedId = C.CLK_L
	CLK_SemicolonAndColon        LedId = C.CLK_SemicolonAndColon
	CLK_ApostropheAndDoubleQuote LedId = C.CLK_ApostropheAndDoubleQuote
	CLK_LeftShift                LedId = C.CLK_LeftShift
	CLK_NonUsBackslash           LedId = C.CLK_NonUsBackslash
	CLK_Z                        LedId = C.CLK_Z
	CLK_X                        LedId = C.CLK_X
	CLK_C                        LedId = C.CLK_C
	CLK_V                        LedId = C.CLK_V
	CLK_B                        LedId = C.CLK_B
	CLK_N                        LedId = C.CLK_N
	CLK_M                        LedId = C.CLK_M
	CLK_CommaAndLessThan         LedId = C.CLK_CommaAndLessThan
	CLK_PeriodAndBiggerThan      LedId = C.CLK_PeriodAndBiggerThan
	CLK_SlashAndQuestionMark     LedId = C.CLK_SlashAndQuestionMark
	CLK_LeftCtrl                 LedId = C.CLK_LeftCtrl
	CLK_LeftGui                  LedId = C.CLK_LeftGui
	CLK_LeftAlt                  LedId = C.CLK_LeftAlt
	CLK_Lang2                    LedId = C.CLK_Lang2
	CLK_Space                    LedId = C.CLK_Space
	CLK_Lang1                    LedId = C.CLK_Lang1
	CLK_International2           LedId = C.CLK_International2
	CLK_RightAlt                 LedId = C.CLK_RightAlt
	CLK_RightGui                 LedId = C.CLK_RightGui
	CLK_Application              LedId = C.CLK_Application
	CLK_LedProgramming           LedId = C.CLK_LedProgramming
	CLK_Brightness               LedId = C.CLK_Brightness
	CLK_F12                      LedId = C.CLK_F12
	CLK_PrintScreen              LedId = C.CLK_PrintScreen
	CLK_ScrollLock               LedId = C.CLK_ScrollLock
	CLK_PauseBreak               LedId = C.CLK_PauseBreak
	CLK_Insert                   LedId = C.CLK_Insert
	CLK_Home                     LedId = C.CLK_Home
	CLK_PageUp                   LedId = C.CLK_PageUp
	CLK_BracketRight             LedId = C.CLK_BracketRight
	CLK_Backslash                LedId = C.CLK_Backslash
	CLK_NonUsTilde               LedId = C.CLK_NonUsTilde
	CLK_Enter                    LedId = C.CLK_Enter
	CLK_International1           LedId = C.CLK_International1
	CLK_EqualsAndPlus            LedId = C.CLK_EqualsAndPlus
	CLK_International3           LedId = C.CLK_International3
	CLK_Backspace                LedId = C.CLK_Backspace
	CLK_Delete                   LedId = C.CLK_Delete
	CLK_End                      LedId = C.CLK_End
	CLK_PageDown                 LedId = C.CLK_PageDown
	CLK_RightShift               LedId = C.CLK_RightShift
	CLK_RightCtrl                LedId = C.CLK_RightCtrl
	CLK_UpArrow                  LedId = C.CLK_UpArrow
	CLK_LeftArrow                LedId = C.CLK_LeftArrow
	CLK_DownArrow                LedId = C.CLK_DownArrow
	CLK_RightArrow               LedId = C.CLK_RightArrow
	CLK_WinLock                  LedId = C.CLK_WinLock
	CLK_Mute                     LedId = C.CLK_Mute
	CLK_Stop                     LedId = C.CLK_Stop
	CLK_ScanPreviousTrack        LedId = C.CLK_ScanPreviousTrack
	CLK_PlayPause                LedId = C.CLK_PlayPause
	CLK_ScanNextTrack            LedId = C.CLK_ScanNextTrack
	CLK_NumLock                  LedId = C.CLK_NumLock
	CLK_KeypadSlash              LedId = C.CLK_KeypadSlash
	CLK_KeypadAsterisk           LedId = C.CLK_KeypadAsterisk
	CLK_KeypadMinus              LedId = C.CLK_KeypadMinus
	CLK_KeypadPlus               LedId = C.CLK_KeypadPlus
	CLK_KeypadEnter              LedId = C.CLK_KeypadEnter
	CLK_Keypad7                  LedId = C.CLK_Keypad7
	CLK_Keypad8                  LedId = C.CLK_Keypad8
	CLK_Keypad9                  LedId = C.CLK_Keypad9
	CLK_KeypadComma              LedId = C.CLK_KeypadComma
	CLK_Keypad4                  LedId = C.CLK_Keypad4
	CLK_Keypad5                  LedId = C.CLK_Keypad5
	CLK_Keypad6                  LedId = C.CLK_Keypad6
	CLK_Keypad1                  LedId = C.CLK_Keypad1
	CLK_Keypad2                  LedId = C.CLK_Keypad2
	CLK_Keypad3                  LedId = C.CLK_Keypad3
	CLK_Keypad0                  LedId = C.CLK_Keypad0
	CLK_KeypadPeriodAndDelete    LedId = C.CLK_KeypadPeriodAndDelete
	CLK_G1                       LedId = C.CLK_G1
	CLK_G2                       LedId = C.CLK_G2
	CLK_G3                       LedId = C.CLK_G3
	CLK_G4                       LedId = C.CLK_G4
	CLK_G5                       LedId = C.CLK_G5
	CLK_G6                       LedId = C.CLK_G6
	CLK_G7                       LedId = C.CLK_G7
	CLK_G8                       LedId = C.CLK_G8
	CLK_G9                       LedId = C.CLK_G9
	CLK_G10                      LedId = C.CLK_G10
	CLK_VolumeUp                 LedId = C.CLK_VolumeUp
	CLK_VolumeDown               LedId = C.CLK_VolumeDown
	CLK_MR                       LedId = C.CLK_MR
	CLK_M1                       LedId = C.CLK_M1
	CLK_M2                       LedId = C.CLK_M2
	CLK_M3                       LedId = C.CLK_M3
	CLK_G11                      LedId = C.CLK_G11
	CLK_G12                      LedId = C.CLK_G12
	CLK_G13                      LedId = C.CLK_G13
	CLK_G14                      LedId = C.CLK_G14
	CLK_G15                      LedId = C.CLK_G15
	CLK_G16                      LedId = C.CLK_G16
	CLK_G17                      LedId = C.CLK_G17
	CLK_G18                      LedId = C.CLK_G18
	CLK_International5           LedId = C.CLK_International5
	CLK_International4           LedId = C.CLK_International4
	CLK_Fn                       LedId = C.CLK_Fn
	CLM_1                        LedId = C.CLM_1
	CLM_2                        LedId = C.CLM_2
	CLM_3                        LedId = C.CLM_3
	CLM_4                        LedId = C.CLM_4
	CLH_LeftLogo                 LedId = C.CLH_LeftLogo
	CLH_RightLogo                LedId = C.CLH_RightLogo
	CLK_Logo                     LedId = C.CLK_Logo
)

func (id LedId) String() string {
	switch id {
	case CLI_Invalid:
		return "CLI_Invalid"
	case CLK_Escape:
		return "CLK_Escape"
	case CLK_F1:
		return "CLK_F1"
	case CLK_F2:
		return "CLK_F2"
	case CLK_F3:
		return "CLK_F3"
	case CLK_F4:
		return "CLK_F4"
	case CLK_F5:
		return "CLK_F5"
	case CLK_F6:
		return "CLK_F6"
	case CLK_F7:
		return "CLK_F7"
	case CLK_F8:
		return "CLK_F8"
	case CLK_F9:
		return "CLK_F9"
	case CLK_F10:
		return "CLK_F10"
	case CLK_F11:
		return "CLK_F11"
	case CLK_GraveAccentAndTilde:
		return "CLK_GraveAccentAndTilde"
	case CLK_1:
		return "CLK_1"
	case CLK_2:
		return "CLK_2"
	case CLK_3:
		return "CLK_3"
	case CLK_4:
		return "CLK_4"
	case CLK_5:
		return "CLK_5"
	case CLK_6:
		return "CLK_6"
	case CLK_7:
		return "CLK_7"
	case CLK_8:
		return "CLK_8"
	case CLK_9:
		return "CLK_9"
	case CLK_0:
		return "CLK_0"
	case CLK_MinusAndUnderscore:
		return "CLK_MinusAndUnderscore"
	case CLK_Tab:
		return "CLK_Tab"
	case CLK_Q:
		return "CLK_Q"
	case CLK_W:
		return "CLK_W"
	case CLK_E:
		return "CLK_E"
	case CLK_R:
		return "CLK_R"
	case CLK_T:
		return "CLK_T"
	case CLK_Y:
		return "CLK_Y"
	case CLK_U:
		return "CLK_U"
	case CLK_I:
		return "CLK_I"
	case CLK_O:
		return "CLK_O"
	case CLK_P:
		return "CLK_P"
	case CLK_BracketLeft:
		return "CLK_BracketLeft"
	case CLK_CapsLock:
		return "CLK_CapsLock"
	case CLK_A:
		return "CLK_A"
	case CLK_S:
		return "CLK_S"
	case CLK_D:
		return "CLK_D"
	case CLK_F:
		return "CLK_F"
	case CLK_G:
		return "CLK_G"
	case CLK_H:
		return "CLK_H"
	case CLK_J:
		return "CLK_J"
	case CLK_K:
		return "CLK_K"
	case CLK_L:
		return "CLK_L"
	case CLK_SemicolonAndColon:
		return "CLK_SemicolonAndColon"
	case CLK_ApostropheAndDoubleQuote:
		return "CLK_ApostropheAndDoubleQuote"
	case CLK_LeftShift:
		return "CLK_LeftShift"
	case CLK_NonUsBackslash:
		return "CLK_NonUsBackslash"
	case CLK_Z:
		return "CLK_Z"
	case CLK_X:
		return "CLK_X"
	case CLK_C:
		return "CLK_C"
	case CLK_V:
		return "CLK_V"
	case CLK_B:
		return "CLK_B"
	case CLK_N:
		return "CLK_N"
	case CLK_M:
		return "CLK_M"
	case CLK_CommaAndLessThan:
		return "CLK_CommaAndLessThan"
	case CLK_PeriodAndBiggerThan:
		return "CLK_PeriodAndBiggerThan"
	case CLK_SlashAndQuestionMark:
		return "CLK_SlashAndQuestionMark"
	case CLK_LeftCtrl:
		return "CLK_LeftCtrl"
	case CLK_LeftGui:
		return "CLK_LeftGui"
	case CLK_LeftAlt:
		return "CLK_LeftAlt"
	case CLK_Lang2:
		return "CLK_Lang2"
	case CLK_Space:
		return "CLK_Space"
	case CLK_Lang1:
		return "CLK_Lang1"
	case CLK_International2:
		return "CLK_International2"
	case CLK_RightAlt:
		return "CLK_RightAlt"
	case CLK_RightGui:
		return "CLK_RightGui"
	case CLK_Application:
		return "CLK_Application"
	case CLK_LedProgramming:
		return "CLK_LedProgramming"
	case CLK_Brightness:
		return "CLK_Brightness"
	case CLK_F12:
		return "CLK_F12"
	case CLK_PrintScreen:
		return "CLK_PrintScreen"
	case CLK_ScrollLock:
		return "CLK_ScrollLock"
	case CLK_PauseBreak:
		return "CLK_PauseBreak"
	case CLK_Insert:
		return "CLK_Insert"
	case CLK_Home:
		return "CLK_Home"
	case CLK_PageUp:
		return "CLK_PageUp"
	case CLK_BracketRight:
		return "CLK_BracketRight"
	case CLK_Backslash:
		return "CLK_Backslash"
	case CLK_NonUsTilde:
		return "CLK_NonUsTilde"
	case CLK_Enter:
		return "CLK_Enter"
	case CLK_International1:
		return "CLK_International1"
	case CLK_EqualsAndPlus:
		return "CLK_EqualsAndPlus"
	case CLK_International3:
		return "CLK_International3"
	case CLK_Backspace:
		return "CLK_Backspace"
	case CLK_Delete:
		return "CLK_Delete"
	case CLK_End:
		return "CLK_End"
	case CLK_PageDown:
		return "CLK_PageDown"
	case CLK_RightShift:
		return "CLK_RightShift"
	case CLK_RightCtrl:
		return "CLK_RightCtrl"
	case CLK_UpArrow:
		return "CLK_UpArrow"
	case CLK_LeftArrow:
		return "CLK_LeftArrow"
	case CLK_DownArrow:
		return "CLK_DownArrow"
	case CLK_RightArrow:
		return "CLK_RightArrow"
	case CLK_WinLock:
		return "CLK_WinLock"
	case CLK_Mute:
		return "CLK_Mute"
	case CLK_Stop:
		return "CLK_Stop"
	case CLK_ScanPreviousTrack:
		return "CLK_ScanPreviousTrack"
	case CLK_PlayPause:
		return "CLK_PlayPause"
	case CLK_ScanNextTrack:
		return "CLK_ScanNextTrack"
	case CLK_NumLock:
		return "CLK_NumLock"
	case CLK_KeypadSlash:
		return "CLK_KeypadSlash"
	case CLK_KeypadAsterisk:
		return "CLK_KeypadAsterisk"
	case CLK_KeypadMinus:
		return "CLK_KeypadMinus"
	case CLK_KeypadPlus:
		return "CLK_KeypadPlus"
	case CLK_KeypadEnter:
		return "CLK_KeypadEnter"
	case CLK_Keypad7:
		return "CLK_Keypad7"
	case CLK_Keypad8:
		return "CLK_Keypad8"
	case CLK_Keypad9:
		return "CLK_Keypad9"
	case CLK_KeypadComma:
		return "CLK_KeypadComma"
	case CLK_Keypad4:
		return "CLK_Keypad4"
	case CLK_Keypad5:
		return "CLK_Keypad5"
	case CLK_Keypad6:
		return "CLK_Keypad6"
	case CLK_Keypad1:
		return "CLK_Keypad1"
	case CLK_Keypad2:
		return "CLK_Keypad2"
	case CLK_Keypad3:
		return "CLK_Keypad3"
	case CLK_Keypad0:
		return "CLK_Keypad0"
	case CLK_KeypadPeriodAndDelete:
		return "CLK_KeypadPeriodAndDelete"
	case CLK_G1:
		return "CLK_G1"
	case CLK_G2:
		return "CLK_G2"
	case CLK_G3:
		return "CLK_G3"
	case CLK_G4:
		return "CLK_G4"
	case CLK_G5:
		return "CLK_G5"
	case CLK_G6:
		return "CLK_G6"
	case CLK_G7:
		return "CLK_G7"
	case CLK_G8:
		return "CLK_G8"
	case CLK_G9:
		return "CLK_G9"
	case CLK_G10:
		return "CLK_G10"
	case CLK_VolumeUp:
		return "CLK_VolumeUp"
	case CLK_VolumeDown:
		return "CLK_VolumeDown"
	case CLK_MR:
		return "CLK_MR"
	case CLK_M1:
		return "CLK_M1"
	case CLK_M2:
		return "CLK_M2"
	case CLK_M3:
		return "CLK_M3"
	case CLK_G11:
		return "CLK_G11"
	case CLK_G12:
		return "CLK_G12"
	case CLK_G13:
		return "CLK_G13"
	case CLK_G14:
		return "CLK_G14"
	case CLK_G15:
		return "CLK_G15"
	case CLK_G16:
		return "CLK_G16"
	case CLK_G17:
		return "CLK_G17"
	case CLK_G18:
		return "CLK_G18"
	case CLK_International5:
		return "CLK_International5"
	case CLK_International4:
		return "CLK_International4"
	case CLK_Fn:
		return "CLK_Fn"
	case CLM_1:
		return "CLM_1"
	case CLM_2:
		return "CLM_2"
	case CLM_3:
		return "CLM_3"
	case CLM_4:
		return "CLM_4"
	case CLH_LeftLogo:
		return "CLH_LeftLogo"
	case CLH_RightLogo:
		return "CLH_RightLogo"
	case CLK_Logo:
		return "CLK_Logo"
	default:
		return "Unknown"
	}
}
