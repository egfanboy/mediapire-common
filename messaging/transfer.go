package messaging

const (
	TopicTransfer            = "mediapire.transfer"
	TopicTransferUpdate      = "mediapire.transfer.update"
	TopicTransferReady       = "mediapire.transfer.ready"
	TopicTransferReadyUpdate = "mediapire.transfer.ready.update"
)

type TransferMessage struct {
	Id       string              `json:"id"`
	TargetId string              `json:"targetId"`
	Inputs   map[string][]string `json:"inputs"`
}

type TransferUpdateMessage struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failureReason"`
	NodeId        string `json:"nodeId"`
	TransferId    string `json:"transferId"`
}

type TransferReadyMessage struct {
	Content    []byte `json:"content"`
	TargetId   string `json:"targetId"`
	TransferId string `json:"transferId"`
}

type TransferReadyUpdateMessage struct {
	Success       bool   `json:"success"`
	FailureReason string `json:"failureReason"`
	TransferId    string `json:"transferId"`
}
