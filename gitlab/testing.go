/*
Copyright 2017 - The GoMiler Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gitlab

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

// MockGitlabAPI populates a []gitlabAPI with mock API data
func MockGitlabAPI(state string) []gitlabAPI {
	currentTime := time.Now()
	gitlabAPImock := []gitlabAPI{}
	for i := 0; i < 10; i++ {
		mock := gitlabAPI{}
		mock.ID = i
		mock.Iid = i
		mock.ProjectID = 1
		mock.Title = "test" + strconv.Itoa(i)
		mock.Description = "test" + strconv.Itoa(i)
		mock.StartDate = "test" + strconv.Itoa(i)
		mock.DueDate = "test" + strconv.Itoa(i)
		mock.UpdatedAt = &currentTime
		mock.CreatedAt = &currentTime
		mock.Name = "test" + strconv.Itoa(i)
		mock.NameSpace.ID = i
		mock.NameSpace.Name = "test" + strconv.Itoa(i)
		mock.NameSpace.Path = "test" + strconv.Itoa(i)
		mock.NameSpace.Kind = "test" + strconv.Itoa(i)
		mock.NameSpace.FullPath = "test" + strconv.Itoa(i)
		if state == "active" {
			mock.State = "active"
		} else {
			mock.State = "closed"
		}
		gitlabAPImock = append(gitlabAPImock, mock)
	}

	return gitlabAPImock
}

// MockGitlabAPIGetRequest creates a mock responder for a specific milestone endpoint and sends back mock JSON data
func MockGitlabAPIGetRequest(URL string, state string) {
	json := MockGitlabAPI(state)
	var strURL []string
	strURL = []string{URL, "/projects/", "1", "/milestones"}
	newURL := strings.Join(strURL, "")
	httpmock.RegisterResponder("GET", newURL,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, json)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

// MockGitlabAPIPostRequest creates a mock responder for a specific milestone endpoint and sends back mock JSON data
func MockGitlabAPIPostRequest(URL string, state string) {
	json := MockGitlabAPI(state)
	var strURL []string
	strURL = []string{URL, "/projects/", "1", "/milestones"}
	newURL := strings.Join(strURL, "")
	httpmock.RegisterResponder("POST", newURL,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, json)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}

// MockGitlabAPIPutRequest creates a mock responder for a specific milestone endpoint and sends back mock JSON data
func MockGitlabAPIPutRequest(URL string, state string, id string) {
	json := MockGitlabAPI(state)
	var strURL []string
	strURL = []string{URL, "/projects/", "1", "/milestones/", id}
	newURL := strings.Join(strURL, "")
	httpmock.RegisterResponder("PUT", newURL,
		func(req *http.Request) (*http.Response, error) {
			resp, err := httpmock.NewJsonResponse(200, json)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		},
	)
}
