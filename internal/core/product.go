package core

var (
	ErrInvalid = NewError("", "invalid parameter", nil)
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

func (p Product) Validate() *Error {
	if p.ID == "" {
		return ErrInvalid.SetCause("id is required")
	}

	if p.Name == "" {
		return ErrInvalid.SetCause("name is required")
	}

	if p.Price == 0 {
		return ErrInvalid.SetCause("price required")
	}

	if p.Price < 0 {
		return ErrInvalid.SetCause("price cannot be negative")
	}

	return nil
}
