package main

import (
	"net/http"
	"strconv"
	"fmt"
)

func handle(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	a, _ := strconv.Atoi(v.Get("dividend"))
	b, _ := strconv.Atoi(v.Get("divisor"))
	fmt.Fprintf(w, "%d / %d = %d", a, b, a/b)
}

func errorHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()
		fn(w, r)
	}
}

/**
 * 실행 후 url 호출
 * - 정상 건 : http://127.0.0.1:8080/divide?dividend=9&divisor=3
 * - 오류 건 : http://127.0.0.1:8080/divide?dividend=9&divisor=0
 */
func main() {
	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		errorHandler(handle)(w, r)
	})
	http.ListenAndServe(":8080", nil)
}