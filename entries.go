package solidgate

type FormInitDTO struct {
	PaymentIntent string
	Merchant      string
	Signature     string
}

type FormUpdateDTO struct {
	PartialIntent string
	Signature     string
}

type FormResignDTO struct {
	ResignIntent string
	Merchant     string
	Signature    string
}
