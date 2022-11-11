package unittests

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetEnvVars(t *testing.T) {
	var MyEnvVars EnvVars
	MyEnvVars = GetEnvValues()
	t.Log(MyEnvVars)
	require.NotNil(t, MyEnvVars.ip)
	require.NotNil(t, MyEnvVars.user)
	require.NotNil(t, MyEnvVars.pass)
	require.NotNil(t, MyEnvVars.cmp)
	require.NotNil(t, MyEnvVars.csc)
	require.NotNil(t, MyEnvVars.cni)
	require.NotEmpty(t, MyEnvVars.ip)
}

func TestNew(t *testing.T) {
	cam := New(&http.Client{})
	require.Nil(t, cam.d)
	require.NotNil(t, cam.httpClient)
}

func TestInit(t *testing.T) {
	cam := New(&http.Client{})

	err := cam.Init(context.TODO())
	require.NoError(t, err)
	require.NotNil(t, cam.d)
}
