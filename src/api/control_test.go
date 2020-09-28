package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/theamrendrasingh/boolipi/db"
	mock "github.com/theamrendrasingh/boolipi/mocks"
)

type responseFrame struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value bool   `json:"value"`
}

func TestGetting200(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Fetch("someuuid").Return(db.Entry{
		Uuid:  "someuuid",
		Value: true,
		Key:   "name"}, nil)

	db.SetRepo(mockRepo)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/:id", Getting)

	req, err := http.NewRequest(http.MethodGet, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusOK)

	e := responseFrame{}
	body, err := ioutil.ReadAll(w.Body)
	err = json.Unmarshal([]byte(string(body)), &e)

	assert.Equal(t, e.Id, "someuuid")
	assert.Equal(t, e.Key, "name")
	assert.Equal(t, e.Value, true)
}

func TestGetting404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Fetch("someuuid").Return(db.Entry{}, errors.New("record not found"))

	db.SetRepo(mockRepo)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/:id", Getting)

	req, err := http.NewRequest(http.MethodGet, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusNotFound)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	assert.Equal(t, string(body), "")
}

func TestGetting500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Fetch("someuuid").Return(db.Entry{}, errors.New("some random error"))

	db.SetRepo(mockRepo)

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/:id", Getting)

	req, err := http.NewRequest(http.MethodGet, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, http.StatusInternalServerError)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	assert.Equal(t, string(body), "")
}

///////////////////////////////////////////////////////

func TestPosting200(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	mockRepo.EXPECT().Create(gomock.Any(), true, "name").Return(db.Entry{
		Uuid:  "someuuid",
		Value: true,
		Key:   "name"}, nil)

	var jsonStr = []byte(`{"value":true, "key" : "name"}`)

	r := gin.Default()
	r.POST("/", Posting)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	e := responseFrame{}
	body, err := ioutil.ReadAll(w.Body)
	err = json.Unmarshal([]byte(string(body)), &e)

	assert.Equal(t, "name", e.Key)
	assert.Equal(t, true, e.Value)
}

func TestPosting400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	var jsonStr = []byte(`{"cat":true, "dog" : "name"}`)

	r := gin.Default()
	r.POST("/", Posting)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "", string(body))
}

func TestPosting500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	mockRepo.EXPECT().Create(gomock.Any(), true, "name").Return(db.Entry{}, errors.New("some error"))

	var jsonStr = []byte(`{"value":true, "key" : "name"}`)

	r := gin.Default()
	r.POST("/", Posting)

	req, err := http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "", string(body))
}

////////////////////////////////////////////////////////

func TestPatching200(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	mockRepo.EXPECT().Patch(gomock.Any(), true, "name").Return(db.Entry{
		Uuid:  "someuuid",
		Value: true,
		Key:   "name"}, nil)

	var jsonStr = []byte(`{"value":true, "key" : "name"}`)

	r := gin.Default()
	r.PATCH("/:id", Patching)

	req, err := http.NewRequest(http.MethodPatch, "/someuuid", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	e := responseFrame{}
	body, err := ioutil.ReadAll(w.Body)
	err = json.Unmarshal([]byte(string(body)), &e)

	assert.Equal(t, "name", e.Key)
	assert.Equal(t, true, e.Value)
}

func TestPatching400(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	var jsonStr = []byte(`{"value":"some", "key" : "name"}`)

	r := gin.Default()
	r.PATCH("/:id", Patching)

	req, err := http.NewRequest(http.MethodPatch, "/someuuid", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "", string(body))
}

func TestPatching404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	mockRepo.EXPECT().Patch(gomock.Any(), true, "name").Return(db.Entry{},
		errors.New("record not found"))

	var jsonStr = []byte(`{"value":true, "key" : "name"}`)

	r := gin.Default()
	r.PATCH("/:id", Patching)

	req, err := http.NewRequest(http.MethodPatch, "/someuuid", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "", string(body))
}

func TestPatching500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	mockRepo.EXPECT().Patch(gomock.Any(), true, "name").Return(db.Entry{},
		errors.New("some random error"))

	var jsonStr = []byte(`{"value":true, "key" : "name"}`)

	r := gin.Default()
	r.PATCH("/:id", Patching)

	req, err := http.NewRequest(http.MethodPatch, "/someuuid", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "", string(body))
}

//////////////////////////////////////////////////////////////

func TestDeleting204(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Delete("someuuid").Return(nil)
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/:id", Deleting)

	req, err := http.NewRequest(http.MethodDelete, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	assert.Equal(t, "", string(body))
}

func TestDeleting404(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Delete("someuuid").Return(errors.New("record not found"))
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/:id", Deleting)

	req, err := http.NewRequest(http.MethodDelete, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	assert.Equal(t, "", string(body))
}

func TestDeleting500(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepo := mock.NewMockRepo(ctrl)

	mockRepo.EXPECT().Delete("someuuid").Return(errors.New("some random error"))
	db.SetRepo(mockRepo)
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.DELETE("/:id", Deleting)

	req, err := http.NewRequest(http.MethodDelete, "/someuuid", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	body, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}
	assert.Equal(t, "", string(body))
}
