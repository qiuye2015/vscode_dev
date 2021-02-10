package prometheus_learn

//go14 test -v --listen-address :8888 --watch_file ./test.txt
import (
	"context"

	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var serverMux = http.NewServeMux()

type GoServer struct {
	server *http.Server
	wg     sync.WaitGroup
}

func NewGoServer(addr string) *GoServer {
	var goServer GoServer
	if addr != "" {
		goServer.server = &http.Server{
			Addr:         addr,
			Handler:      serverMux,
			ReadTimeout:  time.Second * 1, //设置3秒的读超时
			WriteTimeout: time.Second * 3, //设置3秒的写超时
		}
		goServer.wg.Add(1)
		go func() {
			if err := goServer.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Println("server listen error", addr, err)
			}
			goServer.wg.Done()
		}()
	}
	return &goServer
}

func (s *GoServer) GoShutDown(ctx context.Context) error {
	if s.server == nil {
		return nil
	}
	if err := s.server.Shutdown(ctx); err != nil { //Shutdown-->Close
		return err
	}
	s.wg.Wait()
	return nil
}

//定义两个对外注册函数
func GoHandle(pattern string, handler http.Handler) {
	serverMux.Handle(pattern, handler)
}

//func GoHandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
//	serverMux.HandleFunc(pattern, handler)
//}

func init() {
	GoHandle("/metrics", promhttp.Handler())
}
