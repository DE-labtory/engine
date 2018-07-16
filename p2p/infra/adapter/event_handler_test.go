package adapter_test

import (
	"testing"

	"github.com/it-chain/it-chain-Engine/p2p"
	"github.com/it-chain/it-chain-Engine/p2p/infra/adapter"
	"github.com/it-chain/midgard"
	"github.com/magiconair/properties/assert"
)

func TestEventHandler_HandleConnCreatedEvent(t *testing.T) {

	tests := map[string]struct {
		input struct {
			nodeId  string
			address string
		}
		err error
	}{
		"success": {
			input: struct {
				nodeId  string
				address string
			}{nodeId: string("123"), address: string("123")},
			err: nil,
		},
		"empty address test": {
			input: struct {
				nodeId  string
				address string
			}{nodeId: string("123"), address: string("")},
			err: adapter.ErrEmptyAddress,
		},
	}
	eventHandler := adapter.NewEventHandler(EventHandlerMockPeerApi{})

	for testName, test := range tests {
		t.Logf("running test case %s", testName)
		err := eventHandler.HandleConnCreatedEvent(p2p.ConnectionCreatedEvent{EventModel: midgard.EventModel{ID: test.input.nodeId}, Address: test.input.address})
		assert.Equal(t, err, test.err)
	}

}

func TestEventHandler_HandleConnDisconnectedEvent(t *testing.T) {

	tests := map[string]struct {
		input struct {
			id string
		}
		err error
	}{
		"success": {
			input: struct {
				id string
			}{id: string(123)},
			err: nil,
		},
		"empty node id test": {
			input: struct {
				id string
			}{id: string("")},
			err: adapter.ErrEmptyPeerId,
		},
	}

	eventHandler := adapter.NewEventHandler(EventHandlerMockPeerApi{})

	for testName, test := range tests {

		t.Logf("running test case %s", testName)

		event := p2p.ConnectionDisconnectedEvent{
			EventModel: midgard.EventModel{
				ID: test.input.id,
			},
		}

		err := eventHandler.HandleConnDisconnectedEvent(event)

		assert.Equal(t, err, test.err)

	}
}

type MockPeerRepository struct{}

func (nr MockPeerRepository) Save(data p2p.Peer) error { return nil }

type MockLeaderRepository struct{}

func (lr MockLeaderRepository) SetLeader(leader p2p.Leader) {}

func TestRepositoryProjector_HandleLeaderUpdatedEvent(t *testing.T) {
	repositoryProjector := SetupRepositoryProjector()

	tests := map[string]struct {
		input struct {
			id string
		}
		err error
	}{
		"success": {
			input: struct {
				id string
			}{id: "123"},
			err: nil,
		},
		"empty node id test": {
			input: struct {
				id string
			}{id: string("")},
			err: adapter.ErrEmptyPeerId,
		},
	}

	for testName, test := range tests {
		t.Logf("running test case %s", testName)
		event := p2p.LeaderUpdatedEvent{
			EventModel: midgard.EventModel{
				ID: test.input.id,
			},
		}
		err := repositoryProjector.HandleLeaderUpdatedEvent(event)
		assert.Equal(t, err, test.err)
	}
}

func SetupRepositoryProjector() *adapter.RepositoryProjector {

	repositoryProjector := adapter.NewRepositoryProjector(MockPeerRepository{}, MockLeaderRepository{})

	return repositoryProjector
}
