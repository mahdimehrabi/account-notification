package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDTOValidation(t *testing.T) {
	t.Run("Valid DTO", func(t *testing.T) {
		dto := CreateDTO{
			AccountID: 123,
			Balance:   45.67,
		}

		err := dto.Validate()
		assert.NoError(t, err, "Validation should pass for a valid DTO")
	})

	t.Run("Missing AccountID", func(t *testing.T) {
		dto := CreateDTO{
			// AccountID: 123, // Commenting this line will cause validation error
			Balance: 45.67,
		}

		err := dto.Validate()
		assert.Error(t, err, "Validation should fail for a DTO with missing AccountID")
	})

	t.Run("Missing Balance", func(t *testing.T) {
		dto := CreateDTO{
			AccountID: 123,
			// Balance:   45.67, // Commenting this line will cause validation error
		}

		err := dto.Validate()
		assert.Error(t, err, "Validation should fail for a DTO with missing Balance")
	})

	t.Run("Missing Both AccountID and Balance", func(t *testing.T) {
		dto := CreateDTO{
			// AccountID: 123, // Commenting both lines will cause validation error
			// Balance:   45.67,
		}

		err := dto.Validate()
		assert.Error(t, err, "Validation should fail for a DTO with missing AccountID and Balance")
	})
}
