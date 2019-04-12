package history

import (
	"time"

	"github.com/over-engineering/msa/common/models/types"
)

type EventTimeline map[time.Time]struct {
	Events []Event `json:"events"`
}

type Event struct {
	ID          types.UID `json:"id"`
	Description string    `json:"description"`
	Effects     []Effect  `json:"effects"`
}

type Effect struct {
	TargetID        types.UID `json:"target_id"`
	TransitionBlock `json:"transition_block"`
}

// effect의 target id를 보고 해당하는 로컬 체인에 transition block을 끼워 넣는다.
