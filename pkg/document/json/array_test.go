package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hackerwins/yorkie/pkg/document/json"
	"github.com/hackerwins/yorkie/pkg/document/time"
)

func TestArray(t *testing.T) {
	t.Run("marshal test", func(t *testing.T) {
		a := json.NewArray(json.NewRGA(), time.InitialTicket)

		a.Add(json.NewPrimitive("1", time.InitialTicket))
		assert.Equal(t, `["1"]`, a.Marshal())
		a.Add(json.NewPrimitive("2", time.InitialTicket))
		assert.Equal(t, `["1","2"]`, a.Marshal())
		a.Add(json.NewPrimitive("3", time.InitialTicket))
		assert.Equal(t, `["1","2","3"]`, a.Marshal())
	})
}
