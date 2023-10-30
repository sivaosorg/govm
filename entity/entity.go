package entity

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/sivaosorg/govm/utils"
)

func NewResponseEntity() *responseEntity {
	r := &responseEntity{}
	r.SetHeaders(make(map[string]interface{}))
	r.SetMeta(*NewMetaEntity())
	return r
}

func NewPaginationEntity() *paginationEntity {
	p := &paginationEntity{}
	p.SetIsLast(false)
	return p
}

func NewMetaEntity() *metaEntity {
	m := &metaEntity{}
	m.SetCustomFields(make(map[string]interface{}))
	m.SetRequestedTime(time.Now())
	m.SetApiVersion("v1.0.0")
	return m
}

func (m *metaEntity) SetApiVersion(value string) *metaEntity {
	m.ApiVersion = utils.TrimSpaces(value)
	return m
}

func (m *metaEntity) SetRequestId(value string) *metaEntity {
	m.RequestId = value
	return m
}

func (m *metaEntity) SetRequestedTime(value time.Time) *metaEntity {
	m.RequestedTime = value
	return m
}

func (m *metaEntity) SetCustomFields(value map[string]interface{}) *metaEntity {
	m.CustomFields = value
	return m
}

func (m *metaEntity) AppendCustomField(value string, key interface{}) *metaEntity {
	m.CustomFields[value] = key
	return m
}

func (m *metaEntity) AppendCustomFields(value string, keys ...interface{}) *metaEntity {
	m.CustomFields[value] = keys
	return m
}

func (m *metaEntity) Json() string {
	return utils.ToJson(m)
}

func (p *paginationEntity) SetPage(value int) *paginationEntity {
	if value <= 0 {
		log.Panicf("Invalid page: %v", value)
	}
	p.Page = value
	return p
}

func (p *paginationEntity) SetPerPage(value int) *paginationEntity {
	if value < 0 {
		log.Panicf("Invalid per_page: %v", value)
	}
	p.PerPage = value
	return p
}

func (p *paginationEntity) SetTotalPages(value int) *paginationEntity {
	if value < 0 {
		log.Panicf("Invalid total_pages: %v", value)
	}
	p.TotalPages = value
	return p
}

func (p *paginationEntity) SetTotalItems(value int) *paginationEntity {
	if value < 0 {
		log.Panicf("Invalid total_items: %v", value)
	}
	p.TotalItems = value
	return p
}

func (p *paginationEntity) SetIsLast(value bool) *paginationEntity {
	p.IsLast = value
	return p
}

func (p *paginationEntity) Json() string {
	return utils.ToJson(p)
}

func PaginationEntityValidator(p *paginationEntity) {
	p.SetPage(p.Page).
		SetPerPage(p.PerPage).
		SetTotalPages(p.TotalPages).
		SetTotalItems(p.TotalItems)
}

func (r *responseEntity) SetStatusCode(value int) *responseEntity {
	r.StatusCode = value
	return r
}

func (r *responseEntity) SetTotal(value int) *responseEntity {
	if value < 0 {
		log.Panicf("Invalid total: %v", value)
	}
	r.Total = value
	return r
}

func (r *responseEntity) SetMessage(value string) *responseEntity {
	r.Message = value
	return r
}

func (r *responseEntity) AppendMessage(value ...string) *responseEntity {
	r.Message = strings.Join(value, ",")
	return r
}

func (r *responseEntity) SetData(value interface{}) *responseEntity {
	r.Data = value
	return r
}

func (r *responseEntity) SetErrors(value interface{}) *responseEntity {
	r.Errors = value
	return r
}

func (r *responseEntity) SetError(value error) *responseEntity {
	r.Errors = value.Error()
	return r
}

func (r *responseEntity) SetHeaders(value map[string]interface{}) *responseEntity {
	r.Headers = value
	return r
}

func (r *responseEntity) AppendHeader(key string, value string) *responseEntity {
	r.Headers[key] = value
	return r
}

func (r *responseEntity) AppendHeaders(key string, value ...string) *responseEntity {
	r.Headers[key] = value
	return r
}

func (r *responseEntity) AppendHeaderWith(key string, value interface{}) *responseEntity {
	r.Headers[key] = value
	return r
}

func (r *responseEntity) SetMeta(value metaEntity) *responseEntity {
	r.Meta = value
	return r
}

func (r *responseEntity) SetPagination(value paginationEntity) *responseEntity {
	r.Pagination = value
	return r
}

func (r *responseEntity) Json() string {
	return utils.ToJson(r)
}

func (responseEntity) Ok(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusOK).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (responseEntity) Created(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusCreated).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (responseEntity) BadRequest(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusBadRequest).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) NotFound(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNotFound).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) NotImplemented(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNotImplemented).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) TooManyRequest(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusTooManyRequests).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) Locked(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusLocked).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) NoContent(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusNoContent).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) Processing(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusProcessing).
		SetMessage(message).
		SetData(data).
		SetTotal(1)
	return *r
}

func (responseEntity) UpgradeRequired(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusUpgradeRequired).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) ServiceUnavailable(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusServiceUnavailable).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) InternalServerError(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusInternalServerError).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) GatewayTimeout(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusGatewayTimeout).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) MethodNotAllowed(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusMethodNotAllowed).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) Unauthorized(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusUnauthorized).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) Forbidden(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusForbidden).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) Accepted(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusAccepted).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) RequestTimeout(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusRequestTimeout).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) RequestEntityTooLarge(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusRequestEntityTooLarge).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) UnsupportedMediaType(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusUnsupportedMediaType).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) HTTPVersionNotSupported(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusHTTPVersionNotSupported).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) PaymentRequired(message string, data interface{}) responseEntity {
	r := NewResponseEntity().
		SetStatusCode(http.StatusPaymentRequired).
		SetMessage(message).
		SetData(data).
		SetTotal(0)
	return *r
}

func (responseEntity) BuildX(statusCode int, message string, data interface{}, errors interface{}) responseEntity {
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
