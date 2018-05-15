package serial

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

// WebService Web 服務
type WebService struct {
	id *IDService
}

// NewWebService 創建
func NewWebService() *WebService {
	flag.Parse()
	return &WebService{}
}

// Start 啟動
func (ws *WebService) Start() {
	go ws.listenToSignal()

	c, e := NewConfigService().Get()
	if e != nil {
		glog.Fatalf("%+v", e)
	}

	glog.Infof("Starting serial service %s, with value %d.", c.Port, c.Value)
	ws.initID(c.Value)
	e = http.ListenAndServe(c.Port, ws.initRouters())
	if e != nil {
		glog.Fatalf("%+v", e)
	}
}

// initId 初始識別
func (ws *WebService) initID(n uint64) {
	ws.id = NewIDService()
	ws.id.Set(n)
}

// initRouters 初始路由
func (ws *WebService) initRouters() *mux.Router {
	r := mux.NewRouter()

	// Get.
	r.HandleFunc("/no", ws.get).Methods("GET")

	// Increase.
	r.HandleFunc("/no", ws.increase).Methods("POST")

	return r
}

// listenToSignal 聆聽訊號
func (ws *WebService) listenToSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	glog.Info("Listening signals...")
	<-c

	e := ws.saveConfig()
	if e != nil {
		glog.Errorf("%+v", e)
	}

	os.Exit(0)
}

// saveConfig 保存設定
func (ws *WebService) saveConfig() error {
	glog.Info("Receives stop signals.")

	cs := NewConfigService()
	c, e := cs.Get()
	if e != nil {
		return errors.Wrap(e, "")
	}

	glog.Infof("Setting value to %d.", ws.id.Get())
	c.Value = ws.id.Get()
	e = cs.Save(c)
	if e != nil {
		return errors.Wrap(e, "")
	}

	return nil
}

// get 取得
func (ws *WebService) get(w http.ResponseWriter, r *http.Request) {
	s := strconv.FormatUint(ws.id.Get(), 10)
	w.Write([]byte(s))
}

// increase 遞增
func (ws *WebService) increase(w http.ResponseWriter, r *http.Request) {
	s := strconv.FormatUint(ws.id.Increase(), 10)
	w.Write([]byte(s))
}
