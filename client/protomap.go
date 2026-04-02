package client

import (
	"strings"

	"google.golang.org/protobuf/proto"
)

// ProtoMessageFactories maps protobuf message names to constructors.
// Supported keys include both the short message name and the full protobuf name.
var ProtoMessageFactories = map[string]func() proto.Message{}

func init() {
	registerProtoMessage(&KeyboardMouseControl{})
	registerProtoMessage(&CustomControl{})
	registerProtoMessage(&GameStatus{})
	registerProtoMessage(&GlobalUnitStatus{})
	registerProtoMessage(&GlobalLogisticsStatus{})
	registerProtoMessage(&GlobalSpecialMechanism{})
	registerProtoMessage(&Event{})
	registerProtoMessage(&RobotInjuryStat{})
	registerProtoMessage(&RobotRespawnStatus{})
	registerProtoMessage(&RobotStaticStatus{})
	registerProtoMessage(&RobotDynamicStatus{})
	registerProtoMessage(&RobotModuleStatus{})
	registerProtoMessage(&RobotPosition{})
	registerProtoMessage(&Buff{})
	registerProtoMessage(&PenaltyInfo{})
	registerProtoMessage(&RobotPathPlanInfo{})
	registerProtoMessage(&MapClickInfoNotify{})
	registerProtoMessage(&RadarInfoToClient{})
	registerProtoMessage(&RadarSingleRobotInfo{})
	registerProtoMessage(&CustomByteBlock{})
	registerProtoMessage(&AssemblyCommand{})
	registerProtoMessage(&TechCoreMotionStateSync{})
	registerProtoMessage(&RobotPerformanceSelectionCommand{})
	registerProtoMessage(&RobotPerformanceSelectionSync{})
	registerProtoMessage(&CommonCommand{})
	registerProtoMessage(&HeroDeployModeEventCommand{})
	registerProtoMessage(&DeployModeStatusSync{})
	registerProtoMessage(&RuneActivateCommand{})
	registerProtoMessage(&RuneStatusSync{})
	registerProtoMessage(&SentryStatusSync{})
	registerProtoMessage(&DartCommand{})
	registerProtoMessage(&DartSelectTargetStatusSync{})
	registerProtoMessage(&SentryCtrlCommand{})
	registerProtoMessage(&SentryCtrlResult{})
	registerProtoMessage(&AirSupportCommand{})
	registerProtoMessage(&AirSupportStatusSync{})
}

func registerProtoMessage(msg proto.Message) {
	fullName := string(msg.ProtoReflect().Descriptor().FullName())
	shortName := string(msg.ProtoReflect().Descriptor().Name())
	typ := msg.ProtoReflect().Type()

	factory := func() proto.Message {
		return typ.New().Interface()
	}

	ProtoMessageFactories[fullName] = factory
	ProtoMessageFactories[shortName] = factory
	ProtoMessageFactories[strings.ToLower(fullName)] = factory
	ProtoMessageFactories[strings.ToLower(shortName)] = factory
}

// NewMessageByName returns a new protobuf message instance by name.
func NewMessageByName(name string) (proto.Message, bool) {
	factory, ok := ProtoMessageFactories[name]
	if !ok {
		factory, ok = ProtoMessageFactories[strings.ToLower(name)]
		if !ok {
			return nil, false
		}
	}

	return factory(), true
}
