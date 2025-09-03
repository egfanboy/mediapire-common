package messaging

const (
	TopicNodeReady        = "mediapire.node.ready"
	TopicNodeMediaChanged = "mediapire.node.media.change"
)

type NodeReadyMessage struct {
	Id, Name string
}
