package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GO Gotchas:
// 1 nil in Go actually has a type! and it assumes the type of pointer it needs when assigned (internally).
// nil in go can be a "concretely typed value" var i *int, for instance. You can imagine it as if every
// value in the language instead of being just the value is actually the tuple (Type, Value).

// 2 calling methods on nil pointers is fine, since they get promoted (they carry a pointer to the type)

// 3 types based on other types and values of those types passed to funcs/meths as base and derived types

// 4 implementing interfaces is only about the implementers having the right methods on them. compare vs 3.

// 5 returning concrete error types from func will create bugs when checking err vs nil!

// 6 we get npds when doing dot access on nil pointers because go automatically dereferences these pointers to values,
// and they do not point to any valid memory.

// 7 limits of immutability with value parameters: if we pass a value struct that contains a ptr to a struct
// to a func, we can mutate that inner struct thru the ptr. maybe we need ptr semantics on the outer too
// to make mutability clear.

func main() {
	http.HandleFunc("/", GetRoot)
	http.HandleFunc("/hello", GetHello)
	http.ListenAndServe(":666", nil)
}

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
