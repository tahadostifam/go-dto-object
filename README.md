# Dto Object (Validator)

A straightforward dto object validator written in GoLang and on top of the [Go-Validator](https://github.com/go-playground/validator) library.
Good for both restful and gRPC handlers.

## Install

```bash
go get github.com/tahadostifam/go-dto-object@latest
```

## Usage

```go
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	Gender         string     `validate:"oneof=male female prefer_not_to_say"`
	FavoriteColor string      `validate:"iscolor"`                 // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"`  // a person can have a home and cottage...
}

err := dto_object.Validate(&User{
    FirsName: "",
    ...
})

errs := Divide(err) // returns []string{}
```
