package messaging

import "github.com/google/uuid"

const (
	TopicDeleteMedia = "mediapire.media.delete"
)

type DeleteMediaMessage struct {
	MediaToDelete map[uuid.UUID][]uuid.UUID
}
