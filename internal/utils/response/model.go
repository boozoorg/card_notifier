package response

// knownErrors it is a type of all known errors that, this service catch and show to client
type knownErrors int

// Error method needs to knownErrors implement error interface
//
//	type error interface {
//		Error() string
//	}
func (kw knownErrors) Error() string {
	return err[kw]
}

// errors const variable
const (
	EmptyBody knownErrors = iota + 1
	WrongCard
	WrongWebUrl
	WrongOrderType
)

// errors const description
var err = map[knownErrors]string{
	EmptyBody:      "Body is empty",
	WrongCard:      "Card number is empty or invalid.",
	WrongWebUrl:    "Website url is empty or invalid.",
	WrongOrderType: "Order type is empty or invalid.",
}
