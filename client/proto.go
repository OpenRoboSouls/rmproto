package client

import (
	"fmt"

	"github.com/wintbiit/tsf4g/tdrcom"
)

type NewMessage func() tdrcom.Message

func newMessage[T any]() NewMessage {
	var msg T
	if _, ok := any(&msg).(tdrcom.Message); !ok {
		panic(fmt.Errorf("type %T does not implement tdrcom.Message", any(&msg)))
	}

	return func() tdrcom.Message {
		var msg T
		return any(&msg).(tdrcom.Message)
	}
}

var MessageTypeMap = map[uint16]NewMessage{
	uint16(CmdGameState):  newMessage[GameStateData](),
	uint16(CmdGameResult): newMessage[GameResultData](),
}
