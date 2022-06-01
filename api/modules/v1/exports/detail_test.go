package exports

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/protsack-stephan/dev-toolkit/pkg/storage"
	"github.com/protsack-stephan/gin-toolkit/httpmw"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const detailTestDbName = "enwiki"
const detailTestErrDbName = "e"
const detailTestNs = "0"
const detailTestErrNs = "10"
const detailTestData = `{"name":"Earth","identifier":9228,"version":12,"dateModified":"0001-01-01T00:00:00Z","url":"http://en.wikipedia.org/wiki/Earth"}`
const detailTestErrMsg = "key does not exist"
const detailTestGroup = "group_1"

type detailMockStorage struct {
	mock.Mock
}

func (ms *detailMockStorage) Get(path string) (io.ReadCloser, error) {
	args := ms.Called(path)

	return args.Get(0).(io.ReadCloser), args.Error(1)
}

func setupDetailRBACMW(group string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := new(httpmw.CognitoUser)
		user.SetUsername("user")
		user.SetGroups([]string{group})

		c.Set("user", user)
	}
}

func createDetailTestServer(middleware gin.HandlerFunc, storage storage.Getter, group string) http.Handler {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(middleware)
	router.Handle(http.MethodGet, "/:namespace/:project", Detail(storage, group))

	return router
}

func TestDetail(t *testing.T) {
	assert := assert.New(t)

	t.Run("detail success", func(t *testing.T) {
		path := fmt.Sprintf("export/%s/%s_%s.json", detailTestDbName, detailTestDbName, detailTestNs)
		store := new(detailMockStorage)
		mw := setupDetailRBACMW("unlimited")
		srv := httptest.NewServer(createDetailTestServer(mw, store, detailTestGroup))
		defer srv.Close()
		store.
			On("Get", path).
			Return(ioutil.NopCloser(strings.NewReader(detailTestData)), nil)

		res, err := http.Get(fmt.Sprintf("%s/%s/%s", srv.URL, detailTestNs, detailTestDbName))
		assert.NoError(err)
		defer res.Body.Close()
		assert.Equal(http.StatusOK, res.StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		assert.Equal(detailTestData, string(data))
	})

	t.Run("detail success for custom group", func(t *testing.T) {
		path := fmt.Sprintf("export/%s/%s_group_1_%s.json", detailTestDbName, detailTestDbName, detailTestNs)
		store := new(detailMockStorage)
		mw := setupDetailRBACMW("group_1")
		srv := httptest.NewServer(createDetailTestServer(mw, store, detailTestGroup))
		defer srv.Close()
		store.
			On("Get", path).
			Return(ioutil.NopCloser(strings.NewReader(detailTestData)), nil)

		res, err := http.Get(fmt.Sprintf("%s/%s/%s", srv.URL, detailTestNs, detailTestDbName))
		assert.NoError(err)
		defer res.Body.Close()
		assert.Equal(http.StatusOK, res.StatusCode)

		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		assert.Equal(detailTestData, string(data))
	})

	t.Run("detail ns error", func(t *testing.T) {
		store := new(detailMockStorage)
		mw := setupDetailRBACMW("group_2")
		srv := httptest.NewServer(createDetailTestServer(mw, store, detailTestGroup))
		defer srv.Close()

		res, err := http.Get(fmt.Sprintf("%s/%s/%s", srv.URL, detailTestErrNs, detailTestDbName))
		assert.NoError(err)

		defer res.Body.Close()
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		assert.Contains(string(data), detailTestErrNs)
	})

	t.Run("detail dbName error", func(t *testing.T) {
		store := new(detailMockStorage)
		mw := setupDetailRBACMW("unlimited")
		srv := httptest.NewServer(createDetailTestServer(mw, store, detailTestGroup))
		defer srv.Close()

		res, err := http.Get(fmt.Sprintf("%s/%s/%s", srv.URL, detailTestNs, detailTestErrDbName))
		assert.NoError(err)

		defer res.Body.Close()
		assert.Equal(http.StatusBadRequest, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		assert.Contains(string(data), detailTestErrDbName)
	})

	t.Run("detail storage error", func(t *testing.T) {
		store := new(detailMockStorage)
		mw := setupDetailRBACMW("group_2")
		srv := httptest.NewServer(createDetailTestServer(mw, store, detailTestGroup))
		defer srv.Close()
		store.
			On("Get", fmt.Sprintf("export/%s/%s_%s.json", detailTestDbName, detailTestDbName, detailTestNs)).
			Return(ioutil.NopCloser(strings.NewReader("")), errors.New(detailTestErrMsg))

		res, err := http.Get(fmt.Sprintf("%s/%s/%s", srv.URL, detailTestNs, detailTestDbName))
		assert.NoError(err)

		defer res.Body.Close()
		assert.Equal(http.StatusNotFound, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.NoError(err)
		assert.Contains(string(data), detailTestErrMsg)
	})
}
