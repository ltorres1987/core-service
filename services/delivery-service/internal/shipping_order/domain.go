package shipping_order

import (
	"context"
	"database/sql"
	"time"
)

// ShippingOrder struct to describe ShippingOrder object.
type ShippingOrder struct {
	ID                   int       `db:"id"`
	IdSender             string    `db:"idSender"`
	FullNameSender       string    `db:"fullNameSender"`
	PhoneSender          string    `db:"phoneSender"`
	EmailSender          string    `db:"emailSender"`
	IdRecipient          string    `db:"idRecipient"`
	FullNameRecipient    string    `db:"fullNameRecipient"`
	PhoneRecipient       string    `db:"phoneRecipient"`
	EmailRecipient       string    `db:"emailRecipient"`
	LatOrigin            string    `db:"latOrigin"`
	LngOrigin            string    `db:"lngOrigin"`
	AddressOrigin        string    `db:"addressOrigin"`
	CountryOrigin        string    `db:"countryOrigin"`
	ZipcodeOrigin        string    `db:"zipcodeOrigin"`
	ReferenceOrigin      string    `db:"referenceOrigin"`
	LatDestination       string    `db:"latDestination"`
	LngDestination       string    `db:"lngDestination"`
	AddressDestination   string    `db:"addressDestination"`
	CountryDestination   string    `db:"countryDestination"`
	ZipcodeDestination   string    `db:"zipcodeDestination"`
	ReferenceDestination string    `db:"referenceDestination"`
	PackageSize          string    `db:"packageSize"`
	QuantityProduct      int       `db:"quantityProduct"`
	WeightProduct        int       `db:"weightProduct"`
	OrderStatus          string    `db:"orderStatus"`
	CreatedUser          string    `db:"created_user"`
	CreatedAt            time.Time `db:"created_at"`
	UpdatedUser          string    `db:"updated_user"`
	UpdatedAt            time.Time `db:"updated_at"`
	Status               string    `db:"status"`
}

// struct to describe register a new shipping_order.
type ShippingOrderInsert struct {
	Sender      *ShippingOrderSender      `json:"sender"`
	Recipient   *ShippingOrderRecipient   `json:"recipient"`
	Origin      *ShippingOrderOrigin      `json:"origin"`
	Destination *ShippingOrderDestination `json:"destination"`
	Package     *ShippingOrderPackage     `json:"package"`
	CreatedUser string                    `json:"createdUser" validate:"required,lte=200"`
}

type ShippingOrderSender struct {
	IdSender       string `json:"idSender" validate:"required,lte=200"`
	FullNameSender string `json:"fullNameSender" validate:"required,lte=200"`
	PhoneSender    string `json:"phoneSender" validate:"required,lte=200"`
	EmailSender    string `json:"emailSender" validate:"required,lte=200,email"`
}

type ShippingOrderRecipient struct {
	IdRecipient       string `json:"idRecipient" validate:"required,lte=200"`
	FullNameRecipient string `json:"fullNameRecipient" validate:"required,lte=200"`
	PhoneRecipient    string `json:"phoneRecipient" validate:"required,lte=200"`
	EmailRecipient    string `json:"emailRecipient" validate:"required,lte=200,email"`
}

type ShippingOrderOrigin struct {
	LatOrigin       string `json:"latOrigin" validate:"required,latitude"`
	LngOrigin       string `json:"lngOrigin" validate:"required,longitude"`
	AddressOrigin   string `json:"addressOrigin" validate:"required,lte=200"`
	CountryOrigin   string `json:"countryOrigin" validate:"required,lte=200"`
	ZipcodeOrigin   string `json:"zipcodeOrigin" validate:"required,lte=200"`
	ReferenceOrigin string `json:"referenceOrigin" validate:"required,lte=200"`
}

type ShippingOrderDestination struct {
	LatDestination       string `json:"latDestination" validate:"required,latitude"`
	LngDestination       string `json:"lngDestination" validate:"required,longitude"`
	AddressDestination   string `json:"addressDestination" validate:"required,lte=200"`
	CountryDestination   string `json:"countryDestination" validate:"required,lte=200"`
	ZipcodeDestination   string `json:"zipcodeDestination" validate:"required,lte=200"`
	ReferenceDestination string `json:"referenceDestination" validate:"required,lte=200"`
}

type ShippingOrderPackage struct {
	PackageSize     string `json:"packageSize" validate:"required,lte=1,eq=S|eq=M|eq=L"`
	QuantityProduct int    `json:"quantityProduct" validate:"required,numeric,gt=0"`
	WeightProduct   int    `json:"weightProduct" validate:"required,numeric,gt=0"`
}

// ShippingOrderUpdate struct to describe update shipping_order.
type ShippingOrderUpdate struct {
	OrderStatus string `json:"orderStatus" validate:"required,lte=200,eq=recolectado|eq=en_estacion|eq=en_ruta|eq=entregado"`
	UpdatedUser string `json:"updatedUser" validate:"required,lte=200"`
}

// ShippingOrderCancel struct to describe update shipping_order.
type ShippingOrderCancel struct {
	OrderStatus string `json:"orderStatus"`
	Refund      string `json:"refund" validate:"required,lte=1,eq=S|eq=N"`
	UpdatedUser string `json:"updatedUser" validate:"required,lte=200"`
}

type ShippingOrderOut struct {
	ID          int                       `json:"id" `
	Sender      *ShippingOrderSender      `json:"sender"`
	Recipient   *ShippingOrderRecipient   `json:"recipient"`
	Origin      *ShippingOrderOrigin      `json:"origin"`
	Destination *ShippingOrderDestination `json:"destination"`
	Package     *ShippingOrderPackage     `json:"package"`
	OrderStatus string                    `json:"orderStatus"`
	CreatedUser string                    `json:"created_user"`
	CreatedAt   time.Time                 `json:"created_at"`
	UpdatedUser string                    `json:"updated_user"`
	UpdatedAt   time.Time                 `json:"updated_at"`
	Status      string                    `json:"status"`
}

// Our repository will implement these methods.
type ShippingOrderRepository interface {
	GetShippingOrder(ctx context.Context, shippingOrderID int) (*ShippingOrderOut, error)
	GetSenderShippingOrder(ctx context.Context, shippingOrderID int, idSender string) (*ShippingOrderOut, error)
	CreateShippingOrder(ctx context.Context, shipping_order *ShippingOrder) (sql.Result, error)
	UpdateShippingOrder(ctx context.Context, shippingOrderID int, shipping_order *ShippingOrder) error
}

// Our use-case or service will implement these methods.
type ShippingOrderService interface {
	GetShippingOrder(ctx context.Context, shippingOrderID int) (*ShippingOrderOut, error)
	GetSenderShippingOrder(ctx context.Context, shippingOrderID int, idSender string) (*ShippingOrderOut, error)
	CreateShippingOrder(ctx context.Context, shippingOrderInsert *ShippingOrderInsert) (*ShippingOrderOut, error)
	UpdateShippingOrder(ctx context.Context, shippingOrderID int, shippingOrderUpdate *ShippingOrderUpdate) (*ShippingOrderOut, error)
	CancelShippingOrder(ctx context.Context, shippingOrderID int, shippingOrderCancel *ShippingOrderCancel) error
}
