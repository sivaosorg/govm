package query

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewTimestampQuery() *timestampQuery {
	return &timestampQuery{}
}

func (t *timestampQuery) SetFrom(value time.Time) *timestampQuery {
	t.From = value
	return t
}

func (t *timestampQuery) SetTo(value time.Time) *timestampQuery {
	t.To = value
	return t
}

func (t *timestampQuery) Json() string {
	return utils.ToJson(t)
}

func GetTimestampQuerySample() *timestampQuery {
	t := NewTimestampQuery().SetFrom(time.Now()).SetTo(time.Now())
	return t
}

func NewPaginationQuery() *paginationQuery {
	return &paginationQuery{
		enabled: true,
	}
}

func (p *paginationQuery) SetPage(value int) *paginationQuery {
	p.page = value
	return p
}

func (p *paginationQuery) SetSize(value int) *paginationQuery {
	p.size = value
	return p
}

func (p *paginationQuery) SetEnabled(value bool) *paginationQuery {
	p.enabled = value
	return p
}

func (p *paginationQuery) SetCreatedAt(value timestampQuery) *paginationQuery {
	p.CreatedAt = value
	return p
}

func (p *paginationQuery) SetModifiedAt(value timestampQuery) *paginationQuery {
	p.ModifiedAt = value
	return p
}

func (p *paginationQuery) IsEnabled() bool {
	return p.enabled
}

func (p *paginationQuery) GetOffset() int {
	if p.page == 0 {
		return 0
	}
	return (p.page - 1) * p.size
}

func (p *paginationQuery) GetPage() int {
	return p.page
}

func (p *paginationQuery) GetSize() int {
	return p.size
}

func (p *paginationQuery) AvailableOffset() bool {
	return p.GetOffset() >= 0
}

func (p *paginationQuery) AvailableLimit() bool {
	return p.GetPage() > 0
}

func (p *paginationQuery) GetQueryString() string {
	return fmt.Sprintf("page=%v&size=%v", p.page, p.size)
}

func (p *paginationQuery) Json() string {
	return utils.ToJson(p)
}

func NewIPageQuery() *IPageQuery {
	return &IPageQuery{
		paginationQuery: *NewPaginationQuery(),
	}
}

func (i *IPageQuery) SetPaginator(value paginationQuery) *IPageQuery {
	i.paginationQuery = value
	return i
}

func (i *IPageQuery) SetTargetAt(value timestampQuery) *IPageQuery {
	i.TargetAt = value
	return i
}

func (i *IPageQuery) SetReceivedAt(value timestampQuery) *IPageQuery {
	i.ReceivedAt = value
	return i
}

func (i *IPageQuery) SetEventAt(value timestampQuery) *IPageQuery {
	i.EventAt = value
	return i
}

func (i *IPageQuery) SetIndexAt(value timestampQuery) *IPageQuery {
	i.IndexAt = value
	return i
}

// GetPaginator
// Return suffix sql string including: LIMIT and OFFSET
func (i *IPageQuery) GetPaginator() string {
	if !i.IsEnabled() {
		return ""
	}
	var q strings.Builder
	if i.AvailableLimit() {
		q.WriteString(fmt.Sprintf(" LIMIT %d ", i.GetSize()))
		if i.AvailableOffset() {
			q.WriteString(fmt.Sprintf(" OFFSET %d ", i.GetOffset()))
		}
	}
	return q.String()
}

func (i *IPageQuery) Json() string {
	return utils.ToJson(i)
}

func GetTotalPages(totalCount int, size int) int {
	d := float64(totalCount) / float64(size)
	return int(math.Ceil(d))
}

func GetHasMore(currentPage, totalCount, size int) bool {
	return currentPage < totalCount/size
}