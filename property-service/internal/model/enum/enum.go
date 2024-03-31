package enum

type BookingStatusValue string
type BookingStatusEnt struct {
	Created     BookingStatusValue
	Payment     BookingStatusValue
	WaitToCheck BookingStatusValue
	Success     BookingStatusValue
	Rejected    BookingStatusValue
	Checkout    BookingStatusValue
}

var BookingStatus = &BookingStatusEnt{
	Created:     "CREATED",
	WaitToCheck: "WAIT_TO_CHECK",
	Success:     "SUCCESS",
	Rejected:    "REJECTED",
	Checkout:    "CHECKOUT",
}

type PropertyTypeValue string
type PropertyTypeEnt struct {
	Apartment PropertyTypeValue
	Home      PropertyTypeValue
	Hotel     PropertyTypeValue
}

var PropertyType = &PropertyTypeEnt{
	Apartment: "APARTMENT",
	Home:      "HOME",
	Hotel:     "HOTEL",
}

type PropertyStatusValue string
type PropertyStatusEnt struct {
	InReview PropertyStatusValue
	Approve  PropertyStatusValue
	Reject   PropertyStatusValue
}

var PropertyStatus = &PropertyStatusEnt{
	InReview: "IN_REVIEW",
	Approve:  "APPROVE",
	Reject:   "REJECT",
}
