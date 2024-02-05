package query

import "time"

type timestampQuery struct {
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type paginationQuery struct {
	enabled    bool
	page       int
	size       int
	CreatedAt  timestampQuery `json:"created_at,omitempty"`
	ModifiedAt timestampQuery `json:"modified_at,omitempty"`
}

type IPageQuery struct {
	paginationQuery
	TargetAt   timestampQuery `json:"target_at,omitempty"`
	ReceivedAt timestampQuery `json:"received_at,omitempty"`
	EventAt    timestampQuery `json:"event_at,omitempty"`
	IndexAt    timestampQuery `json:"index_at,omitempty"`
}

type decision struct {
	IsEnabled bool        `json:"enabled"`
	Value     interface{} `json:"value,omitempty"`
	on        time.Time
}

type Modify map[string]decision
