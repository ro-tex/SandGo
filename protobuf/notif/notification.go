package notif

func Hi() {
	println("hello")
}

type LinkMethod string

type Link struct {
	Href    string       `json:"href"`
	Methods []LinkMethod `json:"methods"`
	Rel     string       `json:"rel"`
}

type OrderLinks struct {
	Self Link `json:"link"`
}

type Order struct {
	OrderID string     `json:"orderId"`
	Links   OrderLinks `json:"links"`
}

type FulfillerLinks struct {
	Self Link `json:"link"`
}

type Fulfiller struct {
	Links FulfillerLinks `json:"links"`
}

type NotificationItemStatus string

type NotificationItemLinks struct {
	Self   Link `json:"link"`
	Accept Link `json:"accept"`
	Reject Link `json:"reject"`
}

type NotificationItem struct {
	ItemID string                 `json:"itemId"`
	Status NotificationItemStatus `json:"status"`
	Links  NotificationItemLinks  `json:"links"`
}

type AddressChange struct {
	Country         string `json:"country"`
	PostalCode      string `json:"postalCode"`
	StateOrProvince string `json:"stateOrProvince"`
	City            string `json:"city"`
	Company         string `json:"company"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	MiddleName      string `json:"middleName"`
	Phone           string `json:"phone"`
	PhoneExt        string `json:"phoneExt"`
	Street1         string `json:"street1"`
	Street2         string `json:"street2"`
	Email           string `json:"email"`
	DoorCode        string `json:"doorCode"`
}

type DeliveryChange struct {
	DestinationAddress AddressChange `json:"destinationAddress"`
}

type ChangeRequest struct {
	DeliveryChangeDetails DeliveryChange `json:"deliveryChangeDetails"`
}

type RetryDetails struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

type NotificationStatus string

type NotificationType string

type NotificationLinks struct {
	Self   Link `json:"link"`
	Accept Link `json:"accept"`
	Reject Link `json:"reject"`
}

type Notification struct {
	NotificationID    string             `json:"notificationId"`
	Type              NotificationType   `json:"type"`
	Status            NotificationStatus `json:"status"`
	CreatedDate       string             `json:"createdDate"`
	Order             Order              `json:"order"`
	Fulfiller         Fulfiller          `json:"fulfiller"`
	Items             []NotificationItem `json:"items"`
	Links             NotificationLinks  `json:"links"`
	ChangeRequest     ChangeRequest      `json:"changeRequest"`
	RetryOrderRequest RetryDetails       `json:"retryOrderRequest"`
}

type Notifications []Notification
