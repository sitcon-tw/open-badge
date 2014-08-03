package server

import (
	"os"
	"sync"
	"strings"
	"runtime"
	"net"
	"fmt"
	"net/http"
	"bitbucket.org/kardianos/service"
)

type handler struct {
	w sync.WaitGroup
	ServeMux *http.ServeMux
	IssuerMux *http.ServeMux
	BadgeMux *http.ServeMux
	AssertionMux *http.ServeMux
}

var (
	log service.Logger
	serviceFlag chan bool
	apiListener net.Listener
)

func (h *handler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	h.w.Add(1)
	defer h.w.Done()
	if strings.HasPrefix(r.URL.Path, "/api/issuer") {
		h.IssuerMux.ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/badge") {
		h.BadgeMux.ServeHTTP(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/api/assertion") {
		h.AssertionMux.ServeHTTP(w, r)
	} else {
		h.ServeMux.ServeHTTP(w, r)
	}
	//important to flush before decrementing the wait group.
	//we won't get a chance to once main() ends.
	w.(http.Flusher).Flush()
}

func startService() {
	var err error
	var apiHandler *handler
	if apiListener, err = net.Listen("tcp",":8010"); err == nil {
		apiHandler = &handler{
			ServeMux: http.NewServeMux(),
			IssuerMux: issuerServeMux(),
			BadgeMux: badgeServeMux(),
			AssertionMux: assertionServeMux(),
		}
		apiHandler.ServeMux.HandleFunc("/", notFound)
		go http.Serve(apiListener, apiHandler)
	} else {
		panic(err)
	}
	apiHandler.w.Wait()
	serviceFlag <- true
}

func notFound(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
	return
}

func stopService() {
	apiListener.Close()
	<- serviceFlag
}

func Run() {
	cpu_num := runtime.NumCPU()
	fmt.Println("Cpu num: ", cpu_num)
	runtime.GOMAXPROCS(cpu_num)

	var name = "sitconBadge"
	var displayName = "SITCON Open Badge"
	var desc = "SITCON Open badge system."

	var s, err = service.NewService(name, displayName, desc)
	log = s

	if err != nil {
		fmt.Printf("%s unable to start: %s", displayName, err)
		return
	}

	if len(os.Args) > 1 {
		var err error
		verb := os.Args[1]
		switch verb {
		case "install":
			err = s.Install()
			if err != nil {
				fmt.Printf("Failed to install: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" installed.\n", displayName)
		case "remove":
			err = s.Remove()
			if err != nil {
				fmt.Printf("Failed to remove: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" removed.\n", displayName)
		case "run":
			startService()
		case "start":
			err = s.Start()
			if err != nil {
				fmt.Printf("Failed to start: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" started.\n", displayName)
		case "stop":
			err = s.Stop()
			if err != nil {
				fmt.Printf("Failed to stop: %s\n", err)
				return
			}
			fmt.Printf("Service \"%s\" stopped.\n", displayName)
		}
		return
	}
	err = s.Run(func() error {
		// start
		go startService()
		return nil
	}, func() error {
		// stop
		stopService()
		return nil
	})
	if err != nil {
		s.Error(err.Error())
	}
}