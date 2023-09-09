package testhandlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetRoot(w http.ResponseWriter, req *http.Request) {
	if strings.Contains(req.URL.Path, "roots") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if strings.ToLower(req.Method) == strings.ToLower(http.MethodPut) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("body: %v\n", string(body))

	if len(body) > 0 {
		if strings.Contains(string(body), "unbelievable") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("request body must not contain word 'unbelievable'"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "root")
}

func GetHello(w http.ResponseWriter, req *http.Request) {
	_ = req
	io.WriteString(w, "hello")
}

