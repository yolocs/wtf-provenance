package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	log.Println("Starting...")
	projectID := os.Getenv("PROJECT_ID")
	if projectID == "" {
		log.Fatal("no PROJECT_ID specified")
	}

	v := viper.New()
	v.AddRemoteProvider("firestore", projectID, "configs/wtf-provenance")
	v.SetConfigType("yaml")
	if err := v.ReadRemoteConfig(); err != nil {
		log.Fatal(err)
	}
	h := &handler{v: v}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cv := h.load(r.URL.Path)
		w.Write(cv)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
	v *viper.Viper
}

func (h *handler) load(path string) []byte {
	ck := strings.Trim(path, "/")
	ck = strings.ReplaceAll(ck, "/", ".")
	cv := h.v.Get(ck)
	return []byte(fmt.Sprintf("%v", cv))
}
