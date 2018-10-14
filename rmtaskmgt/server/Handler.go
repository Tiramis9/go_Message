/***********************************************************
Handle to the HTTP
GrpcHandlerFunc函数是用于判断请求是来源于Rpc客户端还是Restful Api的请求，
根据不同的请求注册不同的ServeHTTP服务；r.ProtoMajor == 2也代表着请求必须基于HTTP/2


************************************************************************/

package server

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"runtime"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"huapu.info/rmcloud.exp/rmtaskmgt/pkg/ui/data/swagger"
)

var (
	//	SwaggerDir = "huapu.info/rmcloud.exp/rmtaskmgt/api/gRPC-Gateway"
	SwaggerDir = "huapu.info/rmcloud.exp/rmtaskmgt/pkg/ui/data/swagger"
)

func serveSwaggerUI(mux *http.ServeMux) {
	rlog := logrus.WithField("server", "serveSwaggerUI")
	rlog.Info("serveSwaggerUI")
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
func serveSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}
	SwaggerPAath := strings.TrimPrefix(r.URL.Path, "/swagger/")
	SwaggerPAath = path.Join(SwaggerDir, SwaggerPAath)
	http.ServeFile(w, r, SwaggerPAath)
}

func GrpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	if otherHandler == nil {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			grpcServer.ServeHTTP(w, r)
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}
func FindCallerLine(skip int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCallerLine(skip + i)
		if !strings.HasPrefix(file, "logrus") {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}
func getCallerLine(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= 2 {
				file = file[i+1:]
				break
			}
		}
	}
	return file, line
}
