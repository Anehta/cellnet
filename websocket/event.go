package websocket

import (
	"fmt"

	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/proto/gamedef"
)

var (
	Event_SessionConnected = uint32(cellnet.MessageMetaByName("gamedef.SessionConnected").ID)
	Event_SessionClosed    = uint32(cellnet.MessageMetaByName("gamedef.SessionClosed").ID)
	Event_SessionAccepted  = uint32(cellnet.MessageMetaByName("gamedef.SessionAccepted").ID)
)

// 会话事件
type SessionEvent struct {
	*cellnet.Packet
	Ses cellnet.Session
}

func (self SessionEvent) String() string {
	return fmt.Sprintf("SessionEvent msgid: %d data: %v", self.MsgID, self.Data)
}

func NewSessionEvent(msgid uint32, s cellnet.Session, data []byte) *SessionEvent {
	return &SessionEvent{
		Packet: &cellnet.Packet{MsgID: msgid, Data: data},
		Ses:    s,
	}
}
