package model

import (
	"encoding/json"
	"testing"

	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func MockUsers(t *testing.T, user User) {
	m := structToMap(t, user)

	mocket.Catcher.Attach([]*mocket.FakeResponse{
		{
			Pattern:  `SELECT * FROM "users"`,
			Response: []map[string]interface{}{m},
		},
		{
			Pattern:  `SELECT * FROM "users"  WHERE (ID = 1) ORDER BY "users"."id" ASC LIMIT 1`,
			Response: []map[string]interface{}{m},
		},
	})
}

func structToMap(t *testing.T, s interface{}) (inInterface map[string]interface{}) {
	inrec, err := json.Marshal(s)
	assert.NoError(t, err)

	err = json.Unmarshal(inrec, &inInterface)
	assert.NoError(t, err)
	return
}
