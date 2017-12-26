package rundeck

import (
	"testing"

	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses"
	"github.com/lusis/go-rundeck/pkg/rundeck.v21/responses/testdata"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.UsersInfoResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.GetUsers()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Len(t, s, 2)
}

func TestGetUserInfo(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.UserInfoResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.GetUserInfo("admin")
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "admin", s.Login)
	assert.Equal(t, "Admin", s.FirstName)
	assert.Equal(t, "McAdmin", s.LastName)
	assert.Equal(t, "admin@server.com", s.Email)
}

func TestGetCurrentUserInfo(t *testing.T) {
	jsonfile, err := testdata.GetBytes(responses.UserInfoResponseTestFile)
	if err != nil {
		t.Fatal(err.Error())
	}
	client, server, cErr := newTestRundeckClient(jsonfile, "application/json", 200)
	defer server.Close()
	if cErr != nil {
		t.Fatalf(cErr.Error())
	}
	s, err := client.GetCurrentUserInfo()
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, "admin", s.Login)
	assert.Equal(t, "Admin", s.FirstName)
	assert.Equal(t, "McAdmin", s.LastName)
	assert.Equal(t, "admin@server.com", s.Email)
}
