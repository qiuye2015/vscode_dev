package prometheus_learn

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"testing"

	"github.com/qiuye2015/dev_go/src/fileutil"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func TestHelloWorld(t *testing.T) {
	flag.Parse()
	metricServer := NewGoServer(*addr)
	defer metricServer.GoShutDown(context.Background())
	log.Println("starting server: ", *addr)
	watchStock()
	//TODO:
	// 等待信号
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	for {
		sig := <-sigChan
		log.Println("recv signal exit", sig)
		break
	}
	// t.Fatal("not implemented")
}

var (
	stockPro = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "fjp_stock",
		Help: "fjp_stock",
	}, []string{"company"})
	addr = flag.String("listen-address", ":9999", "The address to listen on for HTTP requests.")
	path = flag.String("watch_file", "./text", "The path to monitor")
)

func watchStock() {
	actionMap := make(map[string]func(s string) error)
	if err := load(*path); err != nil {
		log.Println(err)
	}
	actionMap[*path] = load
	if len(actionMap) > 0 {
		go filetuil.FileWatcher(actionMap)
	}
}

func load(s string) error {
	tmp, err := filetuil.ReadFromFile(s)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_resMap := make(map[string]float64, 0)
	for k, _ := range tmp {

		Vec := strings.Split(k, ". ")
		if len(Vec) != 2 {
			continue
		}
		kv := strings.Split(Vec[1], "：")
		if len(kv) != 2 {
			continue
		}
		marketValue := strings.TrimRight(kv[1], "亿美元")
		marketValueNum, _ := strconv.ParseFloat(marketValue, 64)
		_resMap[kv[0]] = marketValueNum
		//log.Println(kv[0], marketValueNum)
		stockPro.With(prometheus.Labels{"company": kv[0]}).Set(marketValueNum)
	}
	return nil
}
