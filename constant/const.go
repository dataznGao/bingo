package constant

type FaultType int8

const (
	ValueFault = iota
	NullFault
	ExceptionShortcircuitFault
	ExceptionUncaughtFault
	ExceptionUnhandledFault
	AttributeShadowedFault
	AttributeReversoFault
	SwitchMissDefaultFault
	ConditionBorderFault
	ConditionInversedFault
	SyncFault
)

const (
	StrValueFault                 = "ValueFault"
	StrNullFault                  = "NullFault"
	StrExceptionShortcircuitFault = "ExceptionShortcircuitFault"
	StrExceptionUncaughtFault     = "ExceptionUncaughtFault"
	StrExceptionUnhandledFault    = "ExceptionUnhandledFault"
	StrAttributeShadowedFault     = "AttributeShadowedFault"
	StrAttributeReversoFault      = "AttributeReversoFault"
	StrSwitchMissDefaultFault     = "SwitchMissDefaultFault"
	StrConditionBorderFault       = "ConditionBorderFault"
	StrConditionInversedFault     = "ConditionInversedFault"
	StrSyncFault                  = "SyncFault"
)

const (
	ConfigFile = "/Users/misery/GolandProjects/code_fault/config.yaml"
)

const (
	Always = "ALWAYS"
	Random = "RANDOM"
)

var FaultTypeMap map[string]FaultType

func InitFaultTypeMap() {
	FaultTypeMap = make(map[string]FaultType, 0)
	FaultTypeMap[StrValueFault] = ValueFault
	FaultTypeMap[StrNullFault] = NullFault
	FaultTypeMap[StrExceptionUncaughtFault] = ExceptionUncaughtFault
	FaultTypeMap[StrExceptionShortcircuitFault] = ExceptionShortcircuitFault
	FaultTypeMap[StrExceptionUnhandledFault] = ExceptionUnhandledFault
	FaultTypeMap[StrAttributeShadowedFault] = AttributeShadowedFault
	FaultTypeMap[StrAttributeReversoFault] = AttributeReversoFault
	FaultTypeMap[StrSwitchMissDefaultFault] = SwitchMissDefaultFault
	FaultTypeMap[StrConditionBorderFault] = ConditionBorderFault
	FaultTypeMap[StrConditionInversedFault] = ConditionInversedFault
	FaultTypeMap[StrSyncFault] = SyncFault
}
