package dto_object

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Dto1 struct {
	Field1 string `validate:"required"`
}

type Dto2 struct {
	Field1 int `validate:"gte=18,lte=130"`
}

type Dto3 struct {
	Field1 string `validate:"oneof=male female prefer_not_to_say"`
}

type Dto4 struct {
	Field1 string `validate:"iscolor"`
}

type Dto5 struct {
	Field1 string `validate:"iscolor"`
	Field2 string `validate:"required"`
}

func TestValidate(t *testing.T) {
	testCases := []struct {
		Dto interface{}
		Err error
	}{
		{
			Dto: Dto1{
				Field1: "hello_world",
			},
			Err: nil,
		},
		{
			Dto: Dto1{
				Field1: "",
			},
			Err: errors.New("Field1 is required"),
		},
		{
			Dto: Dto2{
				Field1: 19,
			},
			Err: nil,
		},
		{
			Dto: Dto2{
				Field1: 17,
			},
			Err: errors.New("Field1 must be greater than or equal to 18"),
		},
		{
			Dto: Dto3{
				Field1: "male",
			},
			Err: nil,
		},
		{
			Dto: Dto3{
				Field1: "shalgam",
			},
			Err: errors.New("Field1 is not one of the list"),
		},
		{
			Dto: Dto4{
				Field1: "#000000",
			},
			Err: nil,
		},
		{
			Dto: Dto4{
				Field1: "#00 00 00",
			},
			Err: errors.New("Field1 got an invalid color format"),
		},
	}

	for _, v := range testCases {
		err := Validate(v)
		if err != nil {
			assert.Equal(t, err.Error(), v.Err.Error())
		}
	}
}

func TestDivide(t *testing.T) {
	err := Validate(Dto5{
		Field1: "#invalid_hex",
	})

	errs := Divide(err)
	assert.Equal(t, errs[0], "Field1 got an invalid color format")
	assert.Equal(t, errs[1], "Field2 is required")
}

func BenchmarkValidate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Validate(Dto1{
			Field1: "required",
		})
	}
}
