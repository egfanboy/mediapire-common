package messaging

const (
	TopicUpdateMedia  = "mediapire.media.update"
	TopicMediaUpdated = "mediapire.media.updated"
)

type UpdatedItem struct {
	MediaId string `json:"mediaId"`
	Content []byte `json:"content"`
}

type UpdateMediaMessage struct {
	ChangesetId string                   `json:"changesetId"`
	Items       map[string][]UpdatedItem `json:"items"`
}

type MediaUpdatedMessage struct {
	Success       bool              `json:"success"`
	FailureReason *string           `json:"failureReason"`
	NodeId        string            `json:"nodeId"`
	ChangesetId   string            `json:"transferId"`
	IdChanges     map[string]string `json:"idChanges"`
}
