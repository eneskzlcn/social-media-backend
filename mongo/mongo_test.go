package mongo

import (
	"github.com/eneskzlcn/hackernews/config"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMongoDBInitializesSuccessfullyWithGivenValidConfig(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	testDBConfig := config.DB{
		Host: "localhost",
		Port: 27017,
	}
	mongoDB := NewDB(testDBConfig)
	assert.NotNil(t, mongoDB)
}

func TestMongoDBNotInitializesWithInvalidConfig(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip()
	}
	invalidDBConfigCases := []config.DB {
		{
			Host: "",
			Port: 27017,
		},
		{
			Host: "test",
			Port: 0,
		},
		{
			Host: "",
			Port: 0,
		},
	}

	for _, item := range invalidDBConfigCases {
		mongoDB := NewDB(item)
		assert.Nil(t, mongoDB)
	}
}