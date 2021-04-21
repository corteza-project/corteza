package types

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_toAccess(t *testing.T) {
	var (
		req = require.New(t)
	)

	req.Equal(Private, toAccess("private"))
	req.Equal(Public, toAccess("public"))
	req.Equal(Protected, toAccess("protected"))
	req.Equal(Private|Protected, toAccess("protected|private"))
	req.Equal(Access(0), toAccess(""))
}

func TestAccess_Is(t *testing.T) {
	var (
		req = require.New(t)
	)

	req.True(Private.IsPrivate())
	req.True(Private.Is(Private))
	req.True(Protected.IsProtected())
	req.True(Protected.Is(Protected))
	req.True(Public.IsPublic())
	req.True(Public.Is(Public))
	req.True((Public | Protected).Is(Public))
	req.True((Public | Protected).Is(Protected))
	req.False((Public | Protected).Is(Private))
}

func TestAccess_UnmarshalJSON(t *testing.T) {
	var (
		cc = []struct {
			Data    string
			WantErr bool
			Access  Access
		}{
			{Data: ``, Access: Access(0)},
			{Data: `{}`, Access: Access(0)},
			{Data: `{"private":true}`, Access: Private},
			{Data: `{"private":true,"public":true}`, Access: Private | Public},
			{Data: `true`, WantErr: true},
			{Data: `1`, Access: Private},
			{Data: `3`, Access: Private | Protected},
			{Data: `"private"`, Access: Private},
			{Data: `"private|public"`, Access: Private | Public},
			{Data: `"private,public"`, Access: Private | Public},
		}
	)

	for _, c := range cc {
		t.Run(c.Data, func(t *testing.T) {
			var (
				req = require.New(t)
				a   Access
			)

			err := a.UnmarshalJSON([]byte(c.Data))
			if c.WantErr {
				req.Error(err)
			} else {
				req.NoError(err)
			}

			if err != nil {
				return
			}

			req.Equal(a, c.Access)
		})
	}
}
