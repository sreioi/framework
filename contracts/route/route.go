package route

import (
	contractshttp "github.com/sreioi/framework/contracts/http"
	"net/http"
)

type Route interface {
	Router
	// Fallback 注册一个处理程序，以便在没有其他路径匹配时执行。
	Fallback(handler contractshttp.HandlerFunc)
	// GlobalMiddleware 注册应用于路由器所有路由的全局中间件。
	GlobalMiddleware(middlewares ...contractshttp.Middleware)
	// Run 启动 HTTP 服务器并监听指定主机上的传入连接。
	Run(host ...string) error
	// RunTLS 使用提供的 TLS 配置启动 HTTPS 服务器，并在指定主机上监听。
	RunTLS(host ...string) error
	// RunTLSWithCert 使用提供的证书和密钥文件启动 HTTPS 服务器，并在指定的主机和端口上监听。
	RunTLSWithCert(host, certFile, keyFile string) error
	// ServeHTTP 为 HTTP 请求提供服务。
	ServeHTTP(writer http.ResponseWriter, request *http.Request)
}

type GroupFunc func(router Router)

type Router interface {
	// Group 用指定的处理程序创建一个新的路由器组。
	Group(handler GroupFunc)
	// Prefix 向路由器注册的路由添加一个通用前缀。
	Prefix(addr string) Router
	// Middleware 设置路由器的中间件。
	Middleware(middlewares ...contractshttp.Middleware) Router

	// Any 注册一条响应所有动词的新路线。
	Any(relativePath string, handler contractshttp.HandlerFunc)
	// Get 向路由器注册新的 GET 路由。
	Get(relativePath string, handler contractshttp.HandlerFunc)
	// Post 向路由器注册新的 POST 路由。
	Post(relativePath string, handler contractshttp.HandlerFunc)
	// Delete 向路由器注册新的 Delete 路由。
	Delete(relativePath string, handler contractshttp.HandlerFunc)
	// Patch 向路由器注册新的 PATCH 路由。
	Patch(relativePath string, handler contractshttp.HandlerFunc)
	// Put  向路由器注册新的 Put 路由。
	Put(relativePath string, handler contractshttp.HandlerFunc)
	// Options 向路由器注册新的 Options 路由。
	Options(relativePath string, handler contractshttp.HandlerFunc)
	// Resource registers 向路由器注册新的 Resource 路由。
	Resource(relativePath string, controller contractshttp.ResourceController)

	// Static 注册一个带有路径前缀的新路由，以便从提供的根目录中提供静态文件。
	Static(relativePath, root string)
	// StaticFile 注册一个新路由，该路由具有特定路径，可从文件系统中为静态文件提供服务。
	StaticFile(relativePath, filepath string)
	// StaticFS 注册一个带有路径前缀的新路由，以提供文件系统中的静态文件。
	StaticFS(relativePath string, fs http.FileSystem)
}
