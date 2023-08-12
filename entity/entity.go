package entity

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewResponseEntity() *ResponseEntity {
	r := &ResponseEntity{}
	r.SetHeaders(make(map[string]interface{}))
	r.SetMeta(*NewMetaEntity())
	return r
}

func NewPaginationEntity() *PaginationEntity {
	p := &PaginationEntity{}
	p.SetIsLast(false)
	return p
}

func NewMetaEntity() *MetaEntity {
	m := &MetaEntity{}
	m.SetCustomFields(make(map[string]interface{}))
	m.SetRequestedTime(time.Now())
	m.SetApiVersion("v1.0.0")
	return m
}

func (m *MetaEntity) SetApiVersion(value string) *MetaEntity {
	m.ApiVersion = utils.TrimSpaces(value)
	return m
}

func (m *MetaEntity) SetRequestId(value string) *MetaEntity {
	m.RequestId = value
	return m
}

func (m *MetaEntity) SetRequestedTime(value time.Time) *MetaEntity {
	m.RequestedTime = value
	return m
}

func (m *MetaEntity) SetCustomFields(value map[string]interface{}) *MetaEntity {
	m.CustomFields = value
	return m
}

func (m *MetaEntity) AppendCustomField(value string, key interface{}) *MetaEntity {
	m.CustomFields[value] = key
	return m
}

func (m *MetaEntity) AppendCustomFields(value string, keys ...interface{}) *MetaEntity {
	m.CustomFields[value] = keys
	return m
}

func (m *MetaEntity) Json() string {
	return utils.ToJson(m)
}

func (p *PaginationEntity) SetPage(value int) *PaginationEntity {
	if value <= 0 {
		log.Panicf("Invalid page: %v", value)
	}
	p.Page = value
	return p
}

func (p *PaginationEntity) SetPerPage(value int) *PaginationEntity {
	if value < 0 {
		log.Panicf("Invalid per_page: %v", value)
	}
	p.PerPage = value
	return p
}

func (p *PaginationEntity) SetTotalPages(value int) *PaginationEntity {
	if value < 0 {
		log.Panicf("Invalid total_pages: %v", value)
	}
	p.TotalPages = value
	return p
}

func (p *PaginationEntity) SetTotalItems(value int) *PaginationEntity {
	if value < 0 {
		log.Panicf("Invalid total_items: %v", value)
	}
	p.TotalItems = value
	return p
}

func (p *PaginationEntity) SetIsLast(value bool) *PaginationEntity {
	p.IsLast = value
	return p
}

func (p *PaginationEntity) Json() string {
	return utils.ToJson(p)
}

func PaginationEntityValidator(p *PaginationEntity) {
	p.SetPage(p.Page).
		SetPerPage(p.PerPage).
		SetTotalPages(p.TotalPages).
		SetTotalItems(p.TotalItems)
}

func (r *ResponseEntity) SetStatusCode(value int) *ResponseEntity {
	r.StatusCode = value
	return r
}

func (r *ResponseEntity) SetTotal(value int) *ResponseEntity {
	if value < 0 {
		log.Panicf("Invalid total: %v", value)
	}
	r.Total = value
	return r
}

func (r *ResponseEntity) SetMessage(value string) *ResponseEntity {
	r.Message = value
	return r
}

func (r *ResponseEntity) AppendMessage(value ...string) *ResponseEntity {
	r.Message = strings.Join(value, ",")
	return r
}

func (r *ResponseEntity) SetData(value interface{}) *ResponseEntity {
	r.Data = value
	return r
}

func (r *ResponseEntity) SetErrors(value interface{}) *ResponseEntity {
	r.Errors = value
	return r
}

func (r *ResponseEntity) SetError(value error) *ResponseEntity {
	r.Errors = value.Error()
	return r
}

func (r *ResponseEntity) SetHeaders(value map[string]interface{}) *ResponseEntity {
	r.Headers = value
	return r
}

func (r *ResponseEntity) AppendHeader(key string, value string) *ResponseEntity {
	r.Headers[key] = value
	return r
}

func (r *ResponseEntity) AppendHeaders(key string, value ...string) *ResponseEntity {
	r.Headers[key] = value
	return r
}

func (r *ResponseEntity) AppendHeaderWith(key string, value interface{}) *ResponseEntity {
	r.Headers[key] = value
	return r
}

func (r *ResponseEntity) SetMeta(value MetaEntity) *ResponseEntity {
	r.Meta = value
	return r
}

func (r *ResponseEntity) SetPagination(value PaginationEntity) *ResponseEntity {
	r.Pagination = value
	return r
}

func (r *ResponseEntity) Json() string {
	return utils.ToJson(r)
}

func (ResponseEntity) Ok(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusOK).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (ResponseEntity) Created(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusCreated).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (ResponseEntity) BadRequest(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusBadRequest).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) NotFound(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNotFound).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) NotImplemented(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNotImplemented).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) TooManyRequest(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusTooManyRequests).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) Locked(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusLocked).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) NoContent(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNoContent).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) Processing(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusProcessing).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (ResponseEntity) UpgradeRequired(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusUpgradeRequired).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) ServiceUnavailable(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusServiceUnavailable).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) InternalServerError(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusInternalServerError).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) GatewayTimeout(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusGatewayTimeout).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) MethodNotAllowed(message string, data interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusMethodNotAllowed).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (ResponseEntity) BuildX(statusCode int, message string, data interface{}, errors interface{}) ResponseEntity {
	r := NewResponseEntity().
		SetStatusCode(statusCode).
		SetMessage(message).
		SetData(data).
		SetErrors(errors).
		SetTotal(1)
	return *r
}

func IsStatusCodeSuccess(statusCode int) bool {
	return statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices
}

func IsStatusCodeFailure(statusCode int) bool {
	return statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices
}

func DescribeHttpMessage(statusCode int) string {
	return http.StatusText(statusCode)
}
