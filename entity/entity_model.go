package entity

import (
	"time"
)

type ResponseEntity struct {
	StatusCode int                    `json:"status_code" binding:"required"`
	Total      int                    `json:"total"`
	Message    string                 `json:"message"`
	Data       interface{}            `json:"data"`
	Errors     interface{}            `json:"errors"`
	Headers    map[string]interface{} `json:"headers"`
	Meta       MetaEntity             `json:"meta,omitempty"`
	Pagination PaginationEntity       `json:"pagination,omitempty"`
}

type PaginationEntity struct {
	Page       int  `json:"page,omitempty"`
	PerPage    int  `json:"per_page,omitempty"`
	TotalPages int  `json:"total_pages,omitempty"`
	TotalItems int  `json:"total_items,omitempty"`
	IsLast     bool `json:"is_last,omitempty"`
}

type MetaEntity struct {
	ApiVersion    string                 `json:"api_version,omitempty"`
	RequestId     string                 `json:"request_id,omitempty"`
	Locale        string                 `json:"locale,omitempty"`
	RequestedTime time.Time              `json:"requested_time,omitempty"`
	CustomFields  map[string]interface{} `json:"custom_fields,omitempty"`
}
