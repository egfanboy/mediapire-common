package messaging

const (
	TopicDeleteMedia = "mediapire.media.delete"
)

type DeleteMediaMessage struct {
	MediaToDelete map[string][]string
}
