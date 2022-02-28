package handler

import (
	"bit-driver-matching-service/domain/key"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type KeyHandler struct {
	Service *key.Service
}

type KeySetRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (kh *KeyHandler) HandleGet(writer http.ResponseWriter, request *http.Request) {
	var imKey = request.URL.Query().Get("key")
	var imVal, err = kh.Service.Get(imKey)
	if err != nil {
		writer.WriteHeader(http.StatusNoContent)
		return
	}

	_, _ = writer.Write([]byte(fmt.Sprintf(`{"%s":"%s"}`, imKey, imVal)))
}

func (kh *KeyHandler) HandleSet(writer http.ResponseWriter, request *http.Request) {
	var body, err = ioutil.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"status": "error", "error": "%s"}`, err.Error())))
		return
	}

	var req KeySetRequest

	_ = json.Unmarshal(body, &req)

	err = kh.Service.Set(req.Key, req.Value)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		_, _ = writer.Write([]byte(fmt.Sprintf(`{"status": "error", "error": "%s"}`, err.Error())))
		return
	}

	_, _ = writer.Write([]byte(`{"status": "ok"}`))
}

func (kh *KeyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		kh.HandleGet(writer, req)
		logRestRequest(writer, req)
	case http.MethodPost:
		kh.HandleSet(writer, req)
		logRestRequest(writer, req)
	}
}

func logRestRequest(writer http.ResponseWriter, req *http.Request) {
	log.Printf(`{"method": "%s", "host": "%s", "uri": "%s", "response_header": "%s"}\n`,
		req.Method,
		req.Host,
		req.RequestURI,
		writer.Header().Get("Content-Type"),
	)
}
