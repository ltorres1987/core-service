package shipping_order

import (
	"context"
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_SHIPPINGORDER = "SELECT id,idSender,fullNameSender,phoneSender,emailSender,idRecipient,fullNameRecipient,phoneRecipient,emailRecipient,latOrigin,lngOrigin,addressOrigin, " +
		"countryOrigin,zipcodeOrigin,referenceOrigin,latDestination,lngDestination,addressDestination,countryDestination,zipcodeDestination,referenceDestination," +
		"packageSize,quantityProduct,weightProduct,orderStatus,created_user,created_at,updated_user,updated_at,status FROM shipping_order WHERE  id = ? and status = ?"
	QUERY_GET_SHIPPINGORDER_SENDER = "SELECT id,idSender,fullNameSender,phoneSender,emailSender,idRecipient,fullNameRecipient,phoneRecipient,emailRecipient,latOrigin,lngOrigin,addressOrigin, " +
		"countryOrigin,zipcodeOrigin,referenceOrigin,latDestination,lngDestination,addressDestination,countryDestination,zipcodeDestination,referenceDestination," +
		"packageSize,quantityProduct,weightProduct,orderStatus,created_user,created_at,updated_user,updated_at,status FROM shipping_order " +
		"WHERE  id = ? and idSender = ? and status = ?"
	QUERY_CREATE_SHIPPINGORDER = "INSERT INTO shipping_order (idSender,fullNameSender,phoneSender,emailSender,idRecipient,fullNameRecipient,phoneRecipient,emailRecipient,latOrigin,lngOrigin,addressOrigin," +
		"countryOrigin,zipcodeOrigin,referenceOrigin,latDestination,lngDestination,addressDestination,countryDestination,zipcodeDestination,referenceDestination," +
		"packageSize,quantityProduct,weightProduct,orderStatus,created_user,created_at,updated_user,updated_at,status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	QUERY_UPDATE_SHIPPINGORDER = "UPDATE shipping_order SET orderStatus = ? , updated_user = ?, updated_at = ? WHERE id = ?"
)

// Represents that we will use MariaDB in order to implement the methods.
type mariaDBRepository struct {
	mariadb *sql.DB
}

// Create a new repository with MariaDB as the driver.
func NewShippingOrderRepository(mariaDBConnection *sql.DB) ShippingOrderRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

// Gets a single shippingOrder in the database.
func (r *mariaDBRepository) GetShippingOrder(ctx context.Context, shippingOrderID int) (*ShippingOrderOut, error) {
	// Initialize variable.
	shippingOrder := &ShippingOrderOut{}
	shippingOrderSender := &ShippingOrderSender{}
	shippingOrderRecipient := &ShippingOrderRecipient{}
	shippingOrderOrigin := &ShippingOrderOrigin{}
	shippingOrderDestination := &ShippingOrderDestination{}
	shippingOrderPackage := &ShippingOrderPackage{}

	// Prepare SQL to get one shippingOrder.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_SHIPPINGORDER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one shippingOrder and insert it to the 'shippingOrder' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, shippingOrderID, "A").Scan(&shippingOrder.ID,
		&shippingOrderSender.IdSender, &shippingOrderSender.FullNameSender, &shippingOrderSender.PhoneSender, &shippingOrderSender.EmailSender,
		&shippingOrderRecipient.IdRecipient, &shippingOrderRecipient.FullNameRecipient, &shippingOrderRecipient.PhoneRecipient, &shippingOrderRecipient.EmailRecipient,
		&shippingOrderOrigin.LatOrigin, &shippingOrderOrigin.LngOrigin, &shippingOrderOrigin.AddressOrigin, &shippingOrderOrigin.CountryOrigin, &shippingOrderOrigin.ZipcodeOrigin, &shippingOrderOrigin.ReferenceOrigin,
		&shippingOrderDestination.LatDestination, &shippingOrderDestination.LngDestination, &shippingOrderDestination.AddressDestination, &shippingOrderDestination.CountryDestination, &shippingOrderDestination.ZipcodeDestination, &shippingOrderDestination.ReferenceDestination,
		&shippingOrderPackage.PackageSize, &shippingOrderPackage.QuantityProduct, &shippingOrderPackage.WeightProduct,
		&shippingOrder.OrderStatus, &shippingOrder.CreatedUser, &shippingOrder.CreatedAt, &shippingOrder.UpdatedUser, &shippingOrder.UpdatedAt, &shippingOrder.Status)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	shippingOrder.Sender = shippingOrderSender
	shippingOrder.Recipient = shippingOrderRecipient
	shippingOrder.Origin = shippingOrderOrigin
	shippingOrder.Destination = shippingOrderDestination
	shippingOrder.Package = shippingOrderPackage
	return shippingOrder, nil
}

// Gets a single shippingOrder in the database.
func (r *mariaDBRepository) GetSenderShippingOrder(ctx context.Context, shippingOrderID int, idSender string) (*ShippingOrderOut, error) {
	// Initialize variable.
	shippingOrder := &ShippingOrderOut{}
	shippingOrderSender := &ShippingOrderSender{}
	shippingOrderRecipient := &ShippingOrderRecipient{}
	shippingOrderOrigin := &ShippingOrderOrigin{}
	shippingOrderDestination := &ShippingOrderDestination{}
	shippingOrderPackage := &ShippingOrderPackage{}

	// Prepare SQL to get one shippingOrder.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_GET_SHIPPINGORDER_SENDER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Get one shippingOrder and insert it to the 'shippingOrder' struct.
	// If it's empty, return null.
	err = stmt.QueryRowContext(ctx, shippingOrderID, idSender, "A").Scan(&shippingOrder.ID,
		&shippingOrderSender.IdSender, &shippingOrderSender.FullNameSender, &shippingOrderSender.PhoneSender, &shippingOrderSender.EmailSender,
		&shippingOrderRecipient.IdRecipient, &shippingOrderRecipient.FullNameRecipient, &shippingOrderRecipient.PhoneRecipient, &shippingOrderRecipient.EmailRecipient,
		&shippingOrderOrigin.LatOrigin, &shippingOrderOrigin.LngOrigin, &shippingOrderOrigin.AddressOrigin, &shippingOrderOrigin.CountryOrigin, &shippingOrderOrigin.ZipcodeOrigin, &shippingOrderOrigin.ReferenceOrigin,
		&shippingOrderDestination.LatDestination, &shippingOrderDestination.LngDestination, &shippingOrderDestination.AddressDestination, &shippingOrderDestination.CountryDestination, &shippingOrderDestination.ZipcodeDestination, &shippingOrderDestination.ReferenceDestination,
		&shippingOrderPackage.PackageSize, &shippingOrderPackage.QuantityProduct, &shippingOrderPackage.WeightProduct,
		&shippingOrder.OrderStatus, &shippingOrder.CreatedUser, &shippingOrder.CreatedAt, &shippingOrder.UpdatedUser, &shippingOrder.UpdatedAt, &shippingOrder.Status)
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	// Return result.
	shippingOrder.Sender = shippingOrderSender
	shippingOrder.Recipient = shippingOrderRecipient
	shippingOrder.Origin = shippingOrderOrigin
	shippingOrder.Destination = shippingOrderDestination
	shippingOrder.Package = shippingOrderPackage
	return shippingOrder, nil
}

// Creates a single shippingOrder in the database.
func (r *mariaDBRepository) CreateShippingOrder(ctx context.Context, shippingOrder *ShippingOrder) (sql.Result, error) {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_CREATE_SHIPPINGORDER)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Insert one shippingOrder.
	result, err := stmt.ExecContext(ctx, shippingOrder.IdSender, shippingOrder.FullNameSender, shippingOrder.PhoneSender, shippingOrder.EmailSender,
		shippingOrder.IdRecipient, shippingOrder.FullNameRecipient, shippingOrder.PhoneRecipient, shippingOrder.EmailRecipient,
		shippingOrder.LatOrigin, shippingOrder.LngOrigin, shippingOrder.AddressOrigin, shippingOrder.CountryOrigin, shippingOrder.ZipcodeOrigin, shippingOrder.ReferenceOrigin,
		shippingOrder.LatDestination, shippingOrder.LngDestination, shippingOrder.AddressDestination, shippingOrder.CountryDestination, shippingOrder.ZipcodeDestination, shippingOrder.ReferenceDestination,
		shippingOrder.PackageSize, shippingOrder.QuantityProduct, shippingOrder.WeightProduct,
		shippingOrder.OrderStatus, shippingOrder.CreatedUser, shippingOrder.CreatedAt, shippingOrder.UpdatedUser, shippingOrder.UpdatedAt, shippingOrder.Status)
	if err != nil {
		return nil, err
	}

	// Return empty.
	return result, nil
}

// Updates a single shippingOrder in the database.
func (r *mariaDBRepository) UpdateShippingOrder(ctx context.Context, shippingOrderID int, shippingOrder *ShippingOrder) error {
	// Prepare context to be used.
	stmt, err := r.mariadb.PrepareContext(ctx, QUERY_UPDATE_SHIPPINGORDER)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Update one shippingOrder.
	_, err = stmt.ExecContext(ctx, shippingOrder.OrderStatus, shippingOrder.UpdatedUser, shippingOrder.UpdatedAt, shippingOrderID)
	if err != nil {
		return err
	}

	// Return empty.
	return nil
}
