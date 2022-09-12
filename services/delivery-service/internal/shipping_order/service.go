package shipping_order

import (
	"context"
	"delivery-service/internal/package_size"
	"delivery-service/internal/utils"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Implementation of the repository in this service.
type shippingOrderService struct {
	shippingOrderRepository ShippingOrderRepository
	packageSizeRepository   package_size.PackageSizeRepository
}

// Create a new 'service' or 'use-case' for 'ShippingOrder' entity.
func NewShippingOrderService(r ShippingOrderRepository, p package_size.PackageSizeRepository) ShippingOrderService {
	return &shippingOrderService{
		shippingOrderRepository: r,
		packageSizeRepository:   p,
	}
}

// Implementation of 'GetShippingOrder'.
func (s *shippingOrderService) GetShippingOrder(ctx context.Context, shippingOrderID int) (*ShippingOrderOut, error) {
	return s.shippingOrderRepository.GetShippingOrder(ctx, shippingOrderID)
}

// Implementation of 'GetSenderShippingOrder'.
func (s *shippingOrderService) GetSenderShippingOrder(ctx context.Context, shippingOrderID int, idSender string) (*ShippingOrderOut, error) {
	return s.shippingOrderRepository.GetSenderShippingOrder(ctx, shippingOrderID, idSender)
}

// Implementation of 'CreateShippingOrder'.
func (s *shippingOrderService) CreateShippingOrder(ctx context.Context, shippingOrderInsert *ShippingOrderInsert) (*ShippingOrderOut, error) {
	// Create a new shippingOrder struct.
	shippingOrder := &ShippingOrder{}

	// Set initialized default data for shippingOrder:
	shippingOrder.IdSender = shippingOrderInsert.Sender.IdSender
	shippingOrder.FullNameSender = shippingOrderInsert.Sender.FullNameSender
	shippingOrder.PhoneSender = shippingOrderInsert.Sender.PhoneSender
	shippingOrder.EmailSender = shippingOrderInsert.Sender.EmailSender

	shippingOrder.IdRecipient = shippingOrderInsert.Recipient.IdRecipient
	shippingOrder.FullNameRecipient = shippingOrderInsert.Recipient.FullNameRecipient
	shippingOrder.PhoneRecipient = shippingOrderInsert.Recipient.PhoneRecipient
	shippingOrder.EmailRecipient = shippingOrderInsert.Recipient.EmailRecipient

	shippingOrder.LatOrigin = shippingOrderInsert.Origin.LatOrigin
	shippingOrder.LngOrigin = shippingOrderInsert.Origin.LngOrigin
	shippingOrder.AddressOrigin = shippingOrderInsert.Origin.AddressOrigin
	shippingOrder.CountryOrigin = shippingOrderInsert.Origin.CountryOrigin
	shippingOrder.ZipcodeOrigin = shippingOrderInsert.Origin.ZipcodeOrigin
	shippingOrder.ReferenceOrigin = shippingOrderInsert.Origin.ReferenceOrigin

	shippingOrder.LatDestination = shippingOrderInsert.Destination.LatDestination
	shippingOrder.LngDestination = shippingOrderInsert.Destination.LngDestination
	shippingOrder.AddressDestination = shippingOrderInsert.Destination.AddressDestination
	shippingOrder.CountryDestination = shippingOrderInsert.Destination.CountryDestination
	shippingOrder.ZipcodeDestination = shippingOrderInsert.Destination.ZipcodeDestination
	shippingOrder.ReferenceDestination = shippingOrderInsert.Destination.ReferenceDestination

	shippingOrder.PackageSize = shippingOrderInsert.Package.PackageSize
	shippingOrder.QuantityProduct = shippingOrderInsert.Package.QuantityProduct
	shippingOrder.WeightProduct = shippingOrderInsert.Package.WeightProduct

	shippingOrder.OrderStatus = "creado"
	shippingOrder.CreatedUser = shippingOrderInsert.CreatedUser
	shippingOrder.CreatedAt = time.Now()
	shippingOrder.Status = "A"

	//valid package size
	maxSize, _ := strconv.Atoi(os.Getenv("MAX_SIZE"))

	if shippingOrder.WeightProduct > maxSize {
		return nil, fmt.Errorf("orders greater than 25KG must contact the company to make a special agreement")
	}

	packageSize, err := s.packageSizeRepository.GetPackageSize(ctx, shippingOrder.WeightProduct)
	if err != nil {
		return nil, utils.FailOnError(err, "packet size information could not be retrieved")
	}

	sizeNemo := ""
	for _, s := range *packageSize {

		sizeNemo = s.Nemo
		break
	}

	if shippingOrder.PackageSize != sizeNemo {
		return nil, fmt.Errorf("the type of the package has no relation to size")
	}

	// Pass to the repository layer.
	result, err := s.shippingOrderRepository.CreateShippingOrder(ctx, shippingOrder)

	if err != nil {
		return nil, utils.FailOnError(err, "problems creating the record")
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, utils.FailOnError(err, "it is not possible to retrieve the id from the record")
	}

	shippingOrderSenderOut := &ShippingOrderSender{
		IdSender:       shippingOrder.IdSender,
		FullNameSender: shippingOrder.FullNameSender,
		PhoneSender:    shippingOrder.PhoneSender,
		EmailSender:    shippingOrder.EmailSender,
	}

	shippingOrderRecipientOut := &ShippingOrderRecipient{
		IdRecipient:       shippingOrder.IdRecipient,
		FullNameRecipient: shippingOrder.FullNameRecipient,
		PhoneRecipient:    shippingOrder.PhoneRecipient,
		EmailRecipient:    shippingOrder.EmailRecipient,
	}

	shippingOrderOriginOut := &ShippingOrderOrigin{
		LatOrigin:       shippingOrder.LatOrigin,
		LngOrigin:       shippingOrder.LngOrigin,
		AddressOrigin:   shippingOrder.AddressOrigin,
		CountryOrigin:   shippingOrder.CountryOrigin,
		ZipcodeOrigin:   shippingOrder.ZipcodeOrigin,
		ReferenceOrigin: shippingOrder.ReferenceOrigin,
	}

	shippingOrderDestinationOut := &ShippingOrderDestination{
		LatDestination:       shippingOrder.LatDestination,
		LngDestination:       shippingOrder.LngDestination,
		AddressDestination:   shippingOrder.AddressDestination,
		CountryDestination:   shippingOrder.CountryDestination,
		ZipcodeDestination:   shippingOrder.ZipcodeDestination,
		ReferenceDestination: shippingOrder.ReferenceDestination,
	}

	shippingOrderPackageOut := &ShippingOrderPackage{
		PackageSize:     shippingOrder.PackageSize,
		QuantityProduct: shippingOrder.QuantityProduct,
		WeightProduct:   shippingOrder.WeightProduct,
	}

	ShippingOrderOut := &ShippingOrderOut{
		ID:          int(insertedID),
		Sender:      shippingOrderSenderOut,
		Recipient:   shippingOrderRecipientOut,
		Origin:      shippingOrderOriginOut,
		Destination: shippingOrderDestinationOut,
		Package:     shippingOrderPackageOut,
		OrderStatus: shippingOrder.OrderStatus,
		CreatedUser: shippingOrder.CreatedUser,
		CreatedAt:   shippingOrder.CreatedAt,
		UpdatedUser: shippingOrder.UpdatedUser,
		UpdatedAt:   shippingOrder.UpdatedAt,
		Status:      shippingOrder.Status,
	}
	return ShippingOrderOut, err
}

// Implementation of 'UpdateShippingOrder'.
func (s *shippingOrderService) UpdateShippingOrder(ctx context.Context, shippingOrderID int, shippingOrderUpdate *ShippingOrderUpdate) (*ShippingOrderOut, error) {

	// Create a new shippingOrder struct.
	shippingOrder := &ShippingOrder{}

	// Check if shippingOrder exists.
	searchedShippingOrder, err := s.shippingOrderRepository.GetShippingOrder(ctx, shippingOrderID)
	if err != nil {
		return nil, utils.FailOnError(err, "information could not be retrieved")
	}
	if searchedShippingOrder == nil {
		return nil, fmt.Errorf("There is no shippingOrder with this ID")
	}

	if shippingOrderUpdate.OrderStatus == "recolectado" && searchedShippingOrder.OrderStatus != "creado" {
		return nil, fmt.Errorf("to make this change the order status must be created")
	}

	if shippingOrderUpdate.OrderStatus == "en_estacion" && searchedShippingOrder.OrderStatus != "recolectado" {
		return nil, fmt.Errorf("to make this change the status of the order must be collected")
	}

	if shippingOrderUpdate.OrderStatus == "en_ruta" && searchedShippingOrder.OrderStatus != "en_estacion" {
		return nil, fmt.Errorf("to make this change the status of the order must be in station")
	}

	if shippingOrderUpdate.OrderStatus == "entregado" && searchedShippingOrder.OrderStatus != "en_ruta" {
		return nil, fmt.Errorf("to make this change the status of the order must be en route")
	}

	// Set value for 'Modified' attribute.
	shippingOrder.OrderStatus = shippingOrderUpdate.OrderStatus
	shippingOrder.UpdatedUser = shippingOrderUpdate.UpdatedUser
	shippingOrder.UpdatedAt = time.Now()

	// Pass to the repository layer.
	err = s.shippingOrderRepository.UpdateShippingOrder(ctx, shippingOrderID, shippingOrder)

	if err != nil {
		return nil, utils.FailOnError(err, "could not update record")
	}

	return s.shippingOrderRepository.GetShippingOrder(ctx, shippingOrderID)
}

// Implementation of 'CancelShippingOrder'.
func (s *shippingOrderService) CancelShippingOrder(ctx context.Context, shippingOrderID int, shippingOrderCancel *ShippingOrderCancel) error {

	// Create a new shippingOrder struct.
	shippingOrder := &ShippingOrder{}

	// Check if shippingOrder exists.
	searchedShippingOrder, err := s.shippingOrderRepository.GetShippingOrder(ctx, shippingOrderID)
	if err != nil {
		return utils.FailOnError(err, "information could not be retrieved")
	}
	if searchedShippingOrder == nil {
		return fmt.Errorf("There is no shippingOrder with this ID")
	}

	if searchedShippingOrder.OrderStatus == "en_ruta" || searchedShippingOrder.OrderStatus == "entregado" {
		return fmt.Errorf("the order must not have the status of en route or delivered")
	}

	//Refund
	if shippingOrderCancel.Refund == "S" {
		tm1 := searchedShippingOrder.CreatedAt
		tm2 := time.Now()
		dif := tm2.Sub(tm1).Minutes()

		minCancel, _ := strconv.ParseFloat(os.Getenv("MIN_CANCEL"), 64)

		if dif > minCancel {
			return fmt.Errorf("cancellation with refund only applies within 2 minutes")
		}
	}

	// Set value for 'Modified' attribute.
	shippingOrder.OrderStatus = "cancelado"
	shippingOrder.UpdatedUser = shippingOrderCancel.UpdatedUser
	shippingOrder.UpdatedAt = time.Now()

	// Pass to the repository layer.
	err = s.shippingOrderRepository.UpdateShippingOrder(ctx, shippingOrderID, shippingOrder)

	if err != nil {
		return utils.FailOnError(err, "could not update record")
	}

	return nil
}
