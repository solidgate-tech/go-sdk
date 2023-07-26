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
