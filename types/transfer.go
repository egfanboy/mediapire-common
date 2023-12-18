package types

import "time"

type Transfer struct {
	Id            string    `json:"id"`
	Status        string    `json:"status"`
	FailureReason *string   `json:"failureReason"`
	TargetId      string    `json:"targetId"`
	Expiry        time.Time `json:"expiry"`
}
