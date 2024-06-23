package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var uh userHandler

func TestFetchUserList(t *testing.T) {
	// param map[string]struct{
	// 	want
	// }
	t.Run("", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		body := bytes.NewBufferString("{\"foo\":\"bar\"}")
		c.Request, _ = http.NewRequest("POST", "/v1/users", body)
		c.Request.Header.Add("application/json", binding.MIMEJSON)
		uh.FetchUserList(c)
		// assert.Equal(t, 200, w.Code)
	})

}
