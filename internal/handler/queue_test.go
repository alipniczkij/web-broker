package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alipniczkij/web-broker/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Get(t *testing.T) {
	tests := []struct {
		name               string
		key                string
		timeout            string
		fileName           string
		fileData           map[string][]string
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name:               "Ok",
			key:                "color",
			timeout:            "1",
			fileData:           map[string][]string{"color": []string{"green"}},
			fileName:           "temp.json",
			expectedStatusCode: 200,
			expectedResponse:   "green",
		},
		{
			name:               "Empty file",
			key:                "color",
			timeout:            "1",
			fileData:           map[string][]string{},
			fileName:           "temp.json",
			expectedStatusCode: 404,
			expectedResponse:   "404 not found.\n",
		},
		{
			name:               "Not found",
			key:                "color",
			timeout:            "1",
			fileData:           map[string][]string{"game": []string{"football"}},
			fileName:           "temp.json",
			expectedStatusCode: 404,
			expectedResponse:   "404 not found.\n",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("../../", test.fileName)
			defer os.Remove(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			jsonString, _ := json.Marshal(test.fileData)
			_, err = tmpfile.Write(jsonString)
			if err != nil {
				t.Fatal(err)
			}

			repo := &repository.Repository{Queue: repository.NewQueueRepo(tmpfile.Name())}
			handler := &Handler{repo: repo}

			req, err := http.NewRequest("GET", fmt.Sprintf("/%v?timeout=%v", test.key, test.timeout), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			testHandler := http.HandlerFunc(handler.GetHandler)

			testHandler.ServeHTTP(rr, req)

			assert.Equal(t, test.expectedStatusCode, rr.Code)
			assert.Equal(t, test.expectedResponse, rr.Body.String())
		})
	}
}

// func TestHandler_Put(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		key                string
// 		body               string
// 		fileName           string
// 		fileData           map[string][]string
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name:               "Ok",
// 			key:                "queue",
// 			body:               `{"v": "1"}`,
// 			fileData:           map[string][]string{},
// 			fileName:           "temp.json",
// 			expectedStatusCode: 200,
// 			expectedResponse:   "OK",
// 		},
// 		{
// 			name:               "Empty body",
// 			key:                "queue",
// 			body:               "",
// 			fileData:           map[string][]string{},
// 			fileName:           "temp.json",
// 			expectedStatusCode: 200,
// 			expectedResponse:   "OK",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			tmpfile, err := ioutil.TempFile("../../", test.fileName)
// 			defer os.Remove(tmpfile.Name())
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			jsonString, _ := json.Marshal(test.fileData)
// 			_, err = tmpfile.Write(jsonString)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			repo := &repository.Repository{Queue: repository.NewQueueRepo(tmpfile.Name())}
// 			handler := &Handler{repo: repo}
// 			req, err := http.NewRequest("PUT", fmt.Sprintf("/%v", test.key), bytes.NewBufferString(test.body))
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			rr := httptest.NewRecorder()
// 			testHandler := http.HandlerFunc(handler.PutHandler)

// 			testHandler.ServeHTTP(rr, req)

// 			assert.Equal(t, test.expectedStatusCode, rr.Code)
// 			assert.Equal(t, test.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
