package clientest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
			panic(fmt.Errorf("method doesn't match, want %s, got %s", wantMethod, request.Method))
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

func CheckPayloadHandler(endpoint string, wantPayload interface{}, wantMethod string) Handler {
	return NewHandler(endpoint, func(w http.ResponseWriter, request *http.Request) {
		if err := checkPayload(request, wantPayload); err != nil {
			panic(err)
		}

		if request.Method != wantMethod {
			panic(fmt.Errorf("method doesn't match, want %s, got %s", wantMethod, request.Method))
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	})
}

func checkPayload(request *http.Request, wantPayload interface{}) error {
	if request.Header.Get("Content-Type") == "application/json" {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return err
		}
		if err = checkJson(wantPayload, body); err != nil {
			return err
		}
	}
	if request.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		if err := request.ParseForm(); err != nil {
			return err
		}

		if err := checkForm(wantPayload, request.Form); err != nil {
			return err
		}
	}
	return nil
}

func checkForm(payload interface{}, body url.Values) error {
	want, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	got, err := json.Marshal(body)
	if err != nil {
		return err
	}

	wantJson := string(want)
	gotJson := string(got)
	if wantJson != gotJson {
		return fmt.Errorf("payload doesn't match, want %s, got %s", wantJson, gotJson)
	}
	return nil
}

func checkJson(wantPayload interface{}, body []byte) error {
	want, err := json.Marshal(wantPayload)
	if err != nil {
		return err
	}

	wantJson := string(want)
	bodyJson := string(body)

	if bodyJson != wantJson {
		return fmt.Errorf("payload doesn't match, want %+v, got %+v", wantJson, bodyJson)
	}
	return nil
}

func CheckMethodHandler(endpoint, wantMethod string) Handler {
	return NewHandler(endpoint, func(w http.ResponseWriter, request *http.Request) {
		if request.Method != wantMethod {
			panic(fmt.Errorf("method doesn't match, want %s, got %s", wantMethod, request.Method))
		}
		w.WriteHeader(http.StatusOK)
	})
}
