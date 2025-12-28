package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"job4j.ru/go-lang-base/internal/base"
)

func Test_Validate(t *testing.T) {
	t.Parallel()

	t.Run("Request status: Empty", func(t *testing.T) {
		t.Parallel()

		res := base.Validate(nil)

		assert.Equal(t, []string{"Empty request!"}, res)
	})

	t.Run("Request status: UserID is required", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{UserID: "", Title: "Big Boss", Description: "Take all money"}
		res := base.Validate(&req)
		assert.Equal(t, []string{"UserID is required"}, res)
	})

	t.Run("Request status: Title is required", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{UserID: "1", Title: "", Description: "Take all money"}
		res := base.Validate(&req)

		assert.Equal(t, []string{"Title is required"}, res)
	})

	t.Run("Request status: Description is required", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{UserID: "1", Title: "Big Boss", Description: ""}
		res := base.Validate(&req)

		assert.Equal(t, []string{"Description is required"}, res)
	})

	t.Run("Request status: Ok", func(t *testing.T) {
		t.Parallel()

		req := base.ValidateRequest{UserID: "1", Title: "Big Boss", Description: "Take all money"}
		res := base.Validate(&req)

		assert.Equal(t, []string{}, res)
	})

}
