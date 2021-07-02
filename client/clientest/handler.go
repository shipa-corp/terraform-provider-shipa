package clientest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

type Handler struct {
	endpoint    string
	handlerFunc func(http.ResponseWriter, *http.Request)
}

func (h Handler) GetHandlerFunc() func(http.ResponseWriter, *http.Request) {
	return h.handlerFunc
}

func (h Handler) GetEndpoint() string {
	return h.endpoint
}

func NewHandler(endpoint string, handlerFunc func(http.ResponseWriter, *http.Request)) Handler {
	return Handler{endpoint: endpoint, handlerFunc: handlerFunc}
}

func PrintJsonHandler(endpoint string, payload interface{}, wantMethod string) Handler {
	return NewHandler(endpoint, func(w http.ResponseWriter, request *http.Request) {
		if request.Method != wantMethod {
			panic(fmt.Errorf("method doesn't metach, want %s, got %s", wantMethod, request.Method))
		}

		data, err := json.Marshal(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if _, err := w.Write(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

type ComparablePayload struct {
	Want interface{}
	Got  interface{}
}

func (p ComparablePayload) Check() error {
	if !reflect.DeepEqual(p.Want, p.Got) {
		return fmt.Errorf("payload doesn't match, want %+v, got %+v", p.Want, p.Got)
	}
	return nil
}

func CheckPayloadHandler(endpoint string, payload ComparablePayload, wantMethod string) Handler {
	return NewHandler(endpoint, func(w http.ResponseWriter, request *http.Request) {
		if err := payload.Check(); err != nil {
			panic(err)
		}
		if request.Method != wantMethod {
			panic(fmt.Errorf("method doesn't metach, want %s, got %s", wantMethod, request.Method))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func CheckMethodHandler(endpoint, wantMethod string) Handler {
	return NewHandler(endpoint, func(w http.ResponseWriter, request *http.Request) {
		if request.Method != wantMethod {
			panic(fmt.Errorf("method doesn't metach, want %s, got %s", wantMethod, request.Method))
		}
		w.WriteHeader(http.StatusOK)
	})
}
