package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPut(t *testing.T) {
	s := storage{}

	t.Run("valid json", func(t *testing.T) {
		in := []byte(`{"Time":"2020-08-20T01:45:02.6393092+03:00","Action":"run","Package":"github.com/KolesnikDmitriy/gotesting/storage","Test":"TestPut"}`)

		err := s.Put(in)

		require.NoError(t, err)
		require.NotEmpty(t, s.events)
		assert.Equal(t, time.Date(2020, time.August, 20, 01, 45, 02, 0, time.FixedZone("MSK", 3*60*60)).Unix(), s.events[0].Time.Unix())
		assert.Equal(t, run, s.events[0].Action)
		assert.Equal(t, "github.com/KolesnikDmitriy/gotesting/storage", s.events[0].Package)
		assert.Equal(t, "TestPut", s.events[0].Test)
		assert.Empty(t, s.events[0].Output)
		assert.Zero(t, s.events[0].Elapsed)
	})

	t.Run("not json", func(t *testing.T) {
		in := []byte(`not json`)

		err := s.Put(in)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid character")
	})

	t.Run("empty string", func(t *testing.T) {
		in := []byte(``)

		err := s.Put(in)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "unexpected end of JSON input")
	})
}
