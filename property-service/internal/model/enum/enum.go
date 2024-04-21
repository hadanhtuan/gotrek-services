package enum

type BookingStatusValue string
type BookingStatusEnt struct {
	WaitToCheck BookingStatusValue
	Success     BookingStatusValue
	Rejected    BookingStatusValue
	Checkout    BookingStatusValue
}

var BookingStatus = &BookingStatusEnt{
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
	InReview        PropertyStatusValue //wait admin confirm property
	AdminReject     PropertyStatusValue //admin reject property, cannot booking
	InBooking       PropertyStatusValue //in booking
	WaitHostApprove PropertyStatusValue //user booking and wait host approve
	Available       PropertyStatusValue //property available for booking
}

var PropertyStatus = &PropertyStatusEnt{
	Available:       "AVAILABLE",
	InBooking:       "IN_BOOKING",
	WaitHostApprove: "WAIT_HOST_APPROVE",
	InReview:        "IN_REVIEW",
	AdminReject:     "ADMIN_REJECT",
}

type UserRoleValue string
type UserRoleEnt struct {
	User  UserRoleValue
	Host  UserRoleValue
	Admin UserRoleValue
}

var UserRole = &UserRoleEnt{
	User:  "USER",
	Host:  "HOST",
	Admin: "ADMIN",
}
