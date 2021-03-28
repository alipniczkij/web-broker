package handler

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"
// )

// func TestHandler_Put(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		queue              string
// 		timeout            string
// 		putData            map[string]string
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name:               "Ok",
// 			queue:              "color",
// 			timeout:            "10",
// 			putData:            map[string]string{"color": "green"},
// 			expectedStatusCode: 200,
// 			expectedResponse:   "green",
// 		},
// 		{
// 			name:               "Not found",
// 			queue:              "color",
// 			timeout:            "10",
// 			putData:            map[string]string{},
// 			expectedStatusCode: 404,
// 			expectedResponse:   "404 not found.",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			storageName := "storage.json"
// 			handler := NewHandler(storageName)

// 			for queue, value := range test.putData {
// 				req, err := http.NewRequest("PUT", fmt.Sprintf("/%s", value), nil)
// 				if err != nil {
// 					t.Fatal(err)
// 				}
// 			}

// 			req, err := http.NewRequest("GET", "/", nil)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 		})
// 	}
// }

// func TestLongPoll(t *testing.T) {

// }
