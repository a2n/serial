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

type WebService struct {
	id *IdService
}

func NewWebService() *WebService {
	flag.Parse()
	return &WebService{}
}

func (this *WebService) Start() {
	go this.listenToSignal()

	c, e := NewConfigService().Get()
	if e != nil {
		glog.Fatalf("%+v", e)
	}

	glog.Infof("Starting serial service %s, with value %d.", c.Port, c.Value)
	this.initId(c.Value)
	e = http.ListenAndServe(c.Port, this.initRouters())
	if e != nil {
		glog.Fatalf("%+v", e)
	}
}

func (this *WebService) initId(n uint64) {
	this.id = NewIdService()
	this.id.Set(n)
}

func (this *WebService) initRouters() *mux.Router {
	r := mux.NewRouter()

	// Get.
	r.HandleFunc("/no", this.get).Methods("GET")

	// Increase.
	r.HandleFunc("/no", this.increase).Methods("POST")

	return r
}

func (this *WebService) listenToSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	glog.Info("Listening signals...")
	<-c

	e := this.saveConfig()
	if e != nil {
		glog.Errorf("%+v", e)
	}

	os.Exit(0)
}

func (this *WebService) saveConfig() error {
	glog.Info("Receives stop signals.")

	cs := NewConfigService()
	c, e := cs.Get()
	if e != nil {
		return errors.Wrap(e, "")
	}

	glog.Infof("Setting value to %d.", this.id.Get())
	c.Value = this.id.Get()
	e = cs.Save(c)
	if e != nil {
		return errors.Wrap(e, "")
	}

	return nil
}

func (this *WebService) get(w http.ResponseWriter, r *http.Request) {
	s := strconv.FormatUint(this.id.Get(), 10)
	w.Write([]byte(s))
}

func (this *WebService) increase(w http.ResponseWriter, r *http.Request) {
	s := strconv.FormatUint(this.id.Increase(), 10)
	w.Write([]byte(s))
}
