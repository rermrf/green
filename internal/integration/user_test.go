package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"green/internal/handler"
	"green/internal/integration/startup"
	"green/internal/repository/dao"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func (s *UserTestSuite) SetupSuite() {
	s.server = startup.InitWebServer()
	s.db = startup.InitDB()
}

func (s *UserTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE IF EXISTS users")
}

func (s *UserTestSuite) TestSignup() {
	testCases := []struct {
		name    string
		reqBody string
		before  func(t *testing.T)
		after   func(t *testing.T)

		wantCode int
		wantRes  handler.Result[string]
	}{
		{
			name:    "注册成功",
			reqBody: `{"email":"test@example.com","password":"wen123...","confirmPassword":"wen123..."}`,
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				var user dao.User
				err := s.db.Where("email = ?", "test@example.com").First(&user).Error
				assert.NoError(s.T(), err)
				assert.True(t, user.Ctime > 0)
				assert.True(t, user.Utime > 0)
				assert.NotEqual(s.T(), dao.User{
					Id: 1,
					Email: sql.NullString{
						String: "test@example.com",
						Valid:  true,
					},
					Password: "wen123...",
				}, user)
			},
			wantCode: http.StatusOK,
			wantRes: handler.Result[string]{
				Code: 0,
				Msg:  "注册成功",
			},
		},
		{
			name:    "邮箱冲突",
			reqBody: `{"email":"test2@example.com","password":"wen123...","confirmPassword":"wen123..."}`,
			before: func(t *testing.T) {
				user := dao.User{
					Email: sql.NullString{
						String: "test2@example.com",
						Valid:  true,
					},
					Password: "wen123...",
				}
				err := s.db.Create(&user).Error
				require.NoError(t, err)
			},
			after: func(t *testing.T) {

			},
			wantCode: http.StatusOK,
			wantRes: handler.Result[string]{
				Code: 4,
				Msg:  "邮箱已经注册",
			},
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			tc.before(t)

			req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(tc.reqBody)))
			require.NoError(t, err)

			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			s.server.ServeHTTP(resp, req)

			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != http.StatusOK {
				return
			}

			var result handler.Result[string]
			err = json.NewDecoder(resp.Body).Decode(&result)
			require.NoError(t, err)
			assert.Equal(t, tc.wantRes, result)

			tc.after(t)
		})
	}
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
