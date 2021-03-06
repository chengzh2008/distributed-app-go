package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServicesURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) Add(reg Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock()
	return nil
}

func (r *registry) Remove(url string) error {
	for i := range r.registrations {
		if r.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			r.registrations = append(r.registrations[:i], r.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at url %v not found", url)
}

var reg = registry{registrations: make([]Registration, 0), mutex: new(sync.Mutex)}

type RegistryService struct{}

func RegisterHandlers() {
	http.HandleFunc("/services", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Registration request received")
		switch req.Method {
		case http.MethodPost:
			dec := json.NewDecoder(req.Body)
			var r Registration
			err := dec.Decode(&r)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("Adding service: %v with URL: %v\n", r.ServiceName, r.ServiceURL)
			err = reg.Add(r)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		case http.MethodDelete:
			payload, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			url := string(payload)
			log.Printf("Removing service with URL: %v\n", url)
			err = reg.Remove(url)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})
}
