package messaging

const (
	TopicNodeReady = "mediapire.node.ready"
)

type NodeReadyMessage struct {
	NodeId string
}
