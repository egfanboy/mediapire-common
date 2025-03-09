package messaging

const (
	TopicNodeReady = "mediapire.node.ready"
)

type NodeReadyMessage struct {
	Id, Name string
}
