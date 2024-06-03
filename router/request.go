package router

type Request struct {
	Method      string
	QueryParams map[string]string
	UrlParams   map[string]string
	Headers     map[string]string
	Body        []byte
}
