package charge

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

// GetQueryParam returns the value of a query parameter from the request.
func GetQueryParam(req *http.Request, param string) string {
	values := req.URL.Query()
	return values.Get(param)
}

// GetHeader returns the value of a header from the request.
func GetHeader(req *http.Request, header string) string {
	return req.Header.Get(header)
}

// GetCookie returns the value of a cookie from the request.
func GetCookie(req *http.Request, cookieName string) string {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return ""
	}
	return cookie.Value
}

// GetSessionID extracts a session ID from a request's cookies.
func GetSessionID(req *http.Request) string {
	return GetCookie(req, "session_id")
}

// IsAjaxRequest checks if the request is an AJAX request.
func IsAjaxRequest(req *http.Request) bool {
	return req.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// IsJSONRequest checks if the request has a JSON content type.
func IsJSONRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	return contentType == "application/json"
}

// IsGET checks if the request method is GET.
func IsGET(req *http.Request) bool {
	return req.Method == "GET"
}

// IsPOST checks if the request method is POST.
func IsPOST(req *http.Request) bool {
	return req.Method == "POST"
}

// IsPUT checks if the request method is PUT.
func IsPUT(req *http.Request) bool {
	return req.Method == "PUT"
}

// IsDELETE checks if the request method is DELETE.
func IsDELETE(req *http.Request) bool {
	return req.Method == "DELETE"
}

// IsPATCH checks if the request method is PATCH.
func IsPATCH(req *http.Request) bool {
	return req.Method == "PATCH"
}

// GetClientIP returns the client's IP address from the request.
func GetClientIP(req *http.Request) string {
	ip := req.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = req.Header.Get("X-Real-IP")
	}
	if ip == "" {
		ip = req.RemoteAddr
	}
	return ip
}

// IsTLS checks if the request is using TLS (HTTPS).
func IsTLS(req *http.Request) bool {
	return req.TLS != nil
}

// GetRequestURL returns the full URL of the request.
func GetRequestURL(req *http.Request) string {
	return req.URL.String()
}

// GetRequestPath returns the path of the request URL.
func GetRequestPath(req *http.Request) string {
	return req.URL.Path
}

// GetRequestUserAgent returns the user agent from the request headers.
func GetRequestUserAgent(req *http.Request) string {
	return req.UserAgent()
}

// GetRequestReferer returns the referer (referrer) URL from the request headers.
func GetRequestReferer(req *http.Request) string {
	return req.Referer()
}

// GetRequestContentType returns the content type of the request.
func GetRequestContentType(req *http.Request) string {
	return req.Header.Get("Content-Type")
}

// GetRequestContentLength returns the content length of the request, if provided.
func GetRequestContentLength(req *http.Request) int64 {
	length := req.Header.Get("Content-Length")
	if length == "" {
		return 0
	}
	// Parse the content length, or return 0 if it fails.
	contentLength, err := strconv.ParseInt(length, 10, 64)
	if err != nil {
		return 0
	}
	return contentLength
}

// GetRequestAcceptHeader returns the value of the "Accept" header from the request.
func GetRequestAcceptHeader(req *http.Request) string {
	return req.Header.Get("Accept")
}

// GetRequestLanguage returns the preferred language from the "Accept-Language" header.
func GetRequestLanguage(req *http.Request) string {
	acceptLanguage := req.Header.Get("Accept-Language")
	languages := strings.Split(acceptLanguage, ",")
	if len(languages) > 0 {
		return strings.TrimSpace(strings.Split(languages[0], ";")[0])
	}
	return ""
}

// GetRequestAuthBearerToken extracts the Bearer token from the "Authorization" header.
func GetRequestAuthBearerToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return ""
}

// ParseFormValues parses form values from the request body and returns a map.
func ParseFormValues(req *http.Request) (map[string][]string, error) {
	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	return req.Form, nil
}

// ParseJSONRequest parses a JSON request body into a target interface.
func ParseJSONRequest(req *http.Request, target interface{}) error {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// IsMultipartForm checks if the request is a multipart form request.
func IsMultipartForm(req *http.Request) bool {
	return strings.Contains(req.Header.Get("Content-Type"), "multipart/form-data") ||
		strings.Contains(req.Header.Get("Content-Type"), "application/x-www-form-urlencoded")
}

// GetMultipartForm parses the multipart form data from the request.
func GetMultipartForm(req *http.Request, maxMemory int64) (*multipart.Form, error) {
	if !IsMultipartForm(req) {
		return nil, errors.New("Request is not a multipart form")
	}
	err := req.ParseMultipartForm(maxMemory)
	if err != nil {
		return nil, err
	}
	return req.MultipartForm, nil
}

// IsWebSocketUpgrade checks if the request is an upgrade to a WebSocket connection.
func IsWebSocketUpgrade(req *http.Request) bool {
	return req.Header.Get("Upgrade") == "websocket"
}

// IsGzipAccepted checks if the request accepts gzip encoding for response compression.
func IsGzipAccepted(req *http.Request) bool {
	return strings.Contains(req.Header.Get("Accept-Encoding"), "gzip")
}

// IsRequestFromLocalhost checks if the request originated from the localhost.
func IsRequestFromLocalhost(req *http.Request) bool {
	remoteAddr := req.RemoteAddr
	// Replace square brackets for IPv6 addresses.
	remoteAddr = strings.TrimPrefix(remoteAddr, "[::1]:")
	remoteAddr = strings.TrimPrefix(remoteAddr, "127.0.0.1:")
	return remoteAddr == "localhost" || remoteAddr == "127.0.0.1" || remoteAddr == "::1"
}

// GetRequestQueryParams returns a map of query parameters from the request URL.
func GetRequestQueryParams(req *http.Request) map[string]string {
	queryParams := make(map[string]string)
	values := req.URL.Query()
	for key, values := range values {
		queryParams[key] = values[0]
	}
	return queryParams
}

// GetRequestRawBody reads and returns the raw request body as a string.
func GetRequestRawBody(req *http.Request) (string, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// IsRequestContentTypeJSON checks if the request's content type is JSON.
func IsRequestContentTypeJSON(req *http.Request) bool {
	contentType := GetRequestContentType(req)
	return contentType == "application/json" || strings.HasSuffix(contentType, "+json")
}

// IsRequestContentTypeXML checks if the request's content type is XML.
func IsRequestContentTypeXML(req *http.Request) bool {
	contentType := GetRequestContentType(req)
	return contentType == "application/xml" || strings.HasSuffix(contentType, "+xml")
}

// IsHTTPSRedirect checks if the request should be redirected to HTTPS.
func IsHTTPSRedirect(req *http.Request) bool {
	return req.Header.Get("X-Forwarded-Proto") != "https" && req.Host != "localhost"
}

// IsWebSocketRequest checks if the request is a WebSocket upgrade request.
func IsWebSocketRequest(req *http.Request) bool {
	return strings.ToLower(req.Header.Get("Upgrade")) == "websocket" &&
		strings.ToLower(req.Header.Get("Connection")) == "upgrade"
}

// IsGRPCRequest checks if the request is for gRPC communication.
func IsGRPCRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/grpc")
}

// IsSOAPRequest checks if the request is for SOAP communication.
func IsSOAPRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/soap+xml")
}

// IsGraphQLRequest checks if the request is for GraphQL communication.
func IsGraphQLRequest(req *http.Request) bool {
	contentType := req.Header.Get("Content-Type")
	return strings.Contains(contentType, "application/graphql")
}

// IsJSONPRequest checks if the request is a JSONP request.
func IsJSONPRequest(req *http.Request) bool {
	callback := req.URL.Query().Get("callback")
	return callback != ""
}

// IsOpenIDConnectRequest checks if the request is for OpenID Connect authentication.
func IsOpenIDConnectRequest(req *http.Request) bool {
	acceptHeader := req.Header.Get("Accept")
	return strings.Contains(acceptHeader, "application/openid-connect")
}

// IsFileUploadRequest checks if the request is for handling file uploads.
func IsFileUploadRequest(req *http.Request) bool {
	return strings.HasPrefix(req.Header.Get("Content-Type"), "multipart/form-data") && req.Method == "POST"
}

// IsMobileRequest checks if the request is from a mobile device.
func IsMobileRequest(req *http.Request) bool {
	userAgent := req.Header.Get("User-Agent")
	return strings.Contains(userAgent, "Mobile") || strings.Contains(userAgent, "Android") || strings.Contains(userAgent, "iPhone")
}

// IsStreamingRequest checks if the request is for streaming content.
func IsStreamingRequest(req *http.Request) bool {
	return strings.Contains(req.Header.Get("Accept"), "text/event-stream")
}
