package myapp2

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// NewServer starts and returns a new Server.
	// The caller should call Close when finished, to shut it down
	// Create a mockup testing server
	ts := httptest.NewServer(NewHandler())

	// Close the server
	defer ts.Close()

	resp, err := http.Get(ts.URL)

	// assert.NoError returns bool
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// ReadAll reads from r until an error or EOF and returns the data it read.
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	data, _ := ioutil.ReadAll(resp.Body)

	// Contains asserts that the specified string, list(array, slice...)
	// or map contains the specified substring or element.
	// assert.Contains returns bool value.
	assert.Contains(string(data), "Get UserInfo")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	// if /users/89 url has not been defined, it will redirect to its upper directory.
	// which is, in this case, /users, connected with usersHandler function.
	resp, err := http.Get(ts.URL + "/users/110")

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// ReadAll reads from r until an error or EOF and returns the data it read.
	data, _ := ioutil.ReadAll(resp.Body)

	assert.Contains(string(data), "No User ID: 110")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	// A Server is an HTTP server listening on a system-chosen port on the
	// local loopback interface, for use in end-to-end HTTP tests.
	ts := httptest.NewServer(NewHandler())

	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"jonghyun", "last_name":"sung", "email":"nellow1102@gmail.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)

	assert.NoError(err)

	// NotEqual asserts that the specified values are NOT equal: a.NotEqual(obj1, obj2)
	assert.NotEqual(0, user.ID)

	id := user.ID

	// Get issues a GET to the specified URL. If the response is one of the following redirect codes, Get follows the redirect, up to a maximum of 10 redirects:
	// 301 (Moved Permanently)	// 302 (Found)	// 303 (See Other)
	// 307 (Temporary Redirect)	// 308 (Permanent Redirect)
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))

	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)

	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
	// assert.Equal(user, user2)
}
