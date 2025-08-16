package rmproto

import (
	"fmt"

	"github.com/openrobosouls/rm-proto/client"
	"github.com/openrobosouls/rm-proto/core"
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
	uint16(core.CmdGameState):  newMessage[client.GameStateData](),
	uint16(core.CmdGameResult): newMessage[client.GameResultData](),
}
