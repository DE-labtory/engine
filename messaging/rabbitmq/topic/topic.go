package topic

type Event int

const (
	MessageCreated Event = iota
	ConsensusMessagePublishEvent
	BlockConfirmEvent
	MessageDeliverEvent
	MessageReceiveEvent
	NewConnEvent
	ConnCreateCmd
	ConnCreateEvent
	ConsensusCreateCmd

	//txpool Event
	TransactionReceiveEvent
	TransactionSendEvent
	BlockProposeEvent

	//peer Event
	//todo mq 관련 메세지 아닌것 ( rpc 관련) 분리하기
	LeaderInfoPublishEvent
	LeaderInfoRequestCmd
	LeaderChangeEvent
	PeerConnectEvent
	PeerDisconnectEvent
)

func (e Event) String() string {
	switch e {
	case MessageCreated:
		return "MessageCreated"
	case ConsensusMessagePublishEvent:
		return "ConsensusMessagePublishEvent"
	case BlockConfirmEvent:
		return "BlockConfirmEvent"
	case MessageDeliverEvent:
		return "MessageDeliverEvent"
	case MessageReceiveEvent:
		return "MessageReceiveEvent"
	case NewConnEvent:
		return "NewConnEvent"
	case TransactionReceiveEvent:
		return "TransactionReceiveEvent"
	case TransactionSendEvent:
		return "TransactionSendEvent"
	case BlockProposeEvent:
		return "BlockProposeEvent"
	case ConnCreateCmd:
		return "ConnCreateCmd"
	case ConsensusCreateCmd:
		return "ConsensusCreateCmd"
	//peer
	case LeaderInfoPublishEvent:
		return "LeaderInfoPublishEvent"
	case LeaderInfoRequestCmd:
		return "LeaderInfoRequestCmd"
	case LeaderChangeEvent:
		return "LeaderChangeEvent"
	case PeerConnectEvent:
		return "PeerConnectEvent"
	case PeerDisconnectEvent:
		return "PeerDisconnectEvent"
	}

	return "error"
}
