package messaging

import (
	"github.com/google/uuid"
)

const (
	TopicTransfer            = "mediapire.transfer"
	TopicTransferUpdate      = "mediapire.transfer.update"
	TopicTransferReady       = "mediapire.transfer.ready"
	TopicTransferReadyUpdate = "mediapire.transfer.ready.update"
)

type TransferMessage struct {
	Id       string                    `json:"id"`
	TargetId uuid.UUID                 `json:"targetId"`
	Inputs   map[uuid.UUID][]uuid.UUID `json:"inputs"`
}

type TransferUpdateMessage struct {
	Success       bool      `json:"success"`
	FailureReason string    `json:"failureReason"`
	NodeId        uuid.UUID `json:"nodeId"`
	TransferId    string    `json:"transferId"`
}

type TransferReadyMessage struct {
	Content    []byte    `json:"content"`
	TargetId   uuid.UUID `json:"targetId"`
	TransferId string    `json:"transferId"`
}

type TransferReadyUpdateMessage struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failureReason"`
	TransferId    string `json:"transferId"`
}
