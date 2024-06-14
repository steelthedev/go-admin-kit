package admin

type HandlerFunc func(Context)

type Router interface {
	GET(string, HandlerFunc)
	POST(string, HandlerFunc)
}

type Context interface {
	Param(string) string
	Bind(interface{}) error
	HTML(int, string, interface{})
	Redirect(int, string)
}
