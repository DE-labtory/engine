package adapter

import (
	"errors"
	"testing"

	"github.com/it-chain/engine/consensus"
	"github.com/stretchr/testify/assert"
)

func TestPropagateService_BroadcastPrePrepareMsg(t *testing.T) {
	tests := map[string]struct {
		input struct {
			msg consensus.PrePrepareMsg
		}
		err error
	}{
		"success": {
			input: struct {
				msg consensus.PrePrepareMsg
			}{
				msg: consensus.PrePrepareMsg{
					ConsensusId:    consensus.ConsensusId{"c1"},
					SenderId:       "s1",
					Representative: make([]*consensus.Representative, 0),
					ProposedBlock: consensus.ProposedBlock{
						Seal: make([]byte, 0),
						Body: make([]byte, 0),
					},
				},
			},
			err: nil,
		},
		"Consensus ID empty test": {
			input: struct {
				msg consensus.PrePrepareMsg
			}{
				msg: consensus.PrePrepareMsg{
					ConsensusId:    consensus.ConsensusId{""},
					SenderId:       "s1",
					Representative: make([]*consensus.Representative, 0),
					ProposedBlock: consensus.ProposedBlock{
						Seal: make([]byte, 0),
						Body: make([]byte, 0),
					},
				},
			},
			err: errors.New("Consensus ID is empty"),
		},
		"Block empty test": {
			input: struct {
				msg consensus.PrePrepareMsg
			}{
				msg: consensus.PrePrepareMsg{
					ConsensusId:    consensus.ConsensusId{"c1"},
					SenderId:       "s1",
					Representative: make([]*consensus.Representative, 0),
					ProposedBlock: consensus.ProposedBlock{
						Seal: make([]byte, 0),
						Body: nil,
					},
				},
			},
			err: errors.New("Block is empty"),
		},
	}

	publish := func(exchange string, topic string, data interface{}) (e error) {
		assert.Equal(t, "Command", exchange)
		assert.Equal(t, "message.broadcast", topic)

		return nil
	}

	representatives := make([]*consensus.Representative, 0)
	propagateService := NewPropagateService(publish, representatives)

	for testName, test := range tests {
		t.Logf("running test case [%s]", testName)

		err := propagateService.BroadcastPrePrepareMsg(test.input.msg)

		assert.Equal(t, test.err, err)
	}
}

func TestPropagateService_BroadcastPrepareMsg(t *testing.T) {
	tests := map[string]struct {
		input struct {
			msg consensus.PrepareMsg
		}
		err error
	}{
		"success": {
			input: struct {
				msg consensus.PrepareMsg
			}{
				msg: consensus.PrepareMsg{
					ConsensusId: consensus.ConsensusId{"c1"},
					SenderId:    "s1",
					BlockHash:   make([]byte, 0),
				},
			},
			err: nil,
		},
		"Consensus ID empty test": {
			input: struct {
				msg consensus.PrepareMsg
			}{
				msg: consensus.PrepareMsg{
					ConsensusId: consensus.ConsensusId{""},
					SenderId:    "s1",
					BlockHash:   make([]byte, 0),
				},
			},
			err: errors.New("Consensus ID is empty"),
		},
		"Block hash empty test": {
			input: struct {
				msg consensus.PrepareMsg
			}{
				msg: consensus.PrepareMsg{
					ConsensusId: consensus.ConsensusId{"c1"},
					SenderId:    "s1",
					BlockHash:   nil,
				},
			},
			err: errors.New("Block hash is empty"),
		},
	}

	publish := func(exchange string, topic string, data interface{}) (e error) {
		assert.Equal(t, "Command", exchange)
		assert.Equal(t, "message.broadcast", topic)

		return nil
	}

	representatives := make([]*consensus.Representative, 0)
	propagateService := NewPropagateService(publish, representatives)

	for testName, test := range tests {
		t.Logf("running test case [%s]", testName)

		err := propagateService.BroadcastPrepareMsg(test.input.msg)

		assert.Equal(t, test.err, err)
	}
}

func TestPropagateService_BroadcastCommitMsg(t *testing.T) {
	tests := map[string]struct {
		input struct {
			msg consensus.CommitMsg
		}
		err error
	}{
		"success": {
			input: struct {
				msg consensus.CommitMsg
			}{
				msg: consensus.CommitMsg{
					ConsensusId: consensus.ConsensusId{"c1"},
					SenderId:    "s1",
				},
			},
			err: nil,
		},
		"Consensus ID empty test": {
			input: struct {
				msg consensus.CommitMsg
			}{
				msg: consensus.CommitMsg{
					ConsensusId: consensus.ConsensusId{""},
					SenderId:    "s1",
				},
			},
			err: errors.New("Consensus ID is empty"),
		},
	}

	publish := func(exchange string, topic string, data interface{}) (e error) {
		assert.Equal(t, "Command", exchange)
		assert.Equal(t, "message.broadcast", topic)

		return nil
	}

	representatives := make([]*consensus.Representative, 0)
	propagateService := NewPropagateService(publish, representatives)

	for testName, test := range tests {
		t.Logf("running test case [%s]", testName)

		err := propagateService.BroadcastCommitMsg(test.input.msg)

		assert.Equal(t, test.err, err)
	}
}
