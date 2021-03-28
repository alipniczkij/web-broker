package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alipniczkij/web-broker/pkg/repository"
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
			fileData:           map[string][]string{"color": []string{"green"}}, //"{\"color\": \"green\"}",
			fileName:           "temp.json",
			expectedStatusCode: 200,
			expectedResponse:   "green",
		},
		{
			name:               "Empty file",
			key:                "color",
			timeout:            "1",
			fileData:           map[string][]string{}, // "{}",
			fileName:           "temp.json",
			expectedStatusCode: 404,
			expectedResponse:   "404 not found.",
		},
		{
			name:               "Not found",
			key:                "color",
			timeout:            "1",
			fileData:           map[string][]string{"game": []string{"football"}}, // "{}",
			fileName:           "temp.json",
			expectedStatusCode: 404,
			expectedResponse:   "404 not found.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpfile, err := ioutil.TempFile("../../", test.fileName)
			defer os.Remove(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			log.Printf("file %v was created with %v", tmpfile.Name(), test.fileData)
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

			if status := rr.Code; status != test.expectedStatusCode {
				t.Errorf("handler returned wrong status code:\ngot\n%v\nwant\n%v",
					status, http.StatusOK)
			}

			if rr.Body.String() != test.expectedResponse {
				t.Errorf("handler returned unexpected body:\ngot\n%v\nwant\n%v",
					rr.Body.String(), test.expectedResponse)
			}
		})
	}
}

func TestHandler_Put(t *testing.T) {

}

func TestLongPoll(t *testing.T) {

}
