package records

import (
	"errors"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func NewMockRepo() *redismock.ClientMock {
	mr, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}

func IsAvailable(c redis.Cmdable) bool {
	return c.Ping().Err() == nil
}

func TestNewRepository(t *testing.T) {
	mock := NewMockRepo()

	mock.On("Ping").Return(redis.NewStatusResult("Pong", errors.New("server not available")))

	assert.False(t, IsAvailable(mock))
}
