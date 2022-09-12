package package_size

import (
	"context"
	"database/sql"
)

// Queries that we will use.
const (
	QUERY_GET_PACKAGESIZE = "SELECT id, name, nemo, limitvalue, created_user, created_at, updated_user, updated_at, status " +
		"FROM package_size " +
		"WHERE  limitvalue >= ? and status = ? " +
		"order by limitvalue asc"
)

// Represents that we will use MariaDB in order to implement the methods.
type mariaDBRepository struct {
	mariadb *sql.DB
}

// Create a new repository with MariaDB as the driver.
func NewPackageSizeRepository(mariaDBConnection *sql.DB) PackageSizeRepository {
	return &mariaDBRepository{
		mariadb: mariaDBConnection,
	}
}

// Gets a single packageSize in the database.
func (r *mariaDBRepository) GetPackageSize(ctx context.Context, packageSizeLimitValue int) (*[]PackageSizeOut, error) {
	// Initialize variables.
	var packageSizes []PackageSizeOut

	// Get all packageSizes.
	res, err := r.mariadb.QueryContext(ctx, QUERY_GET_PACKAGESIZE, packageSizeLimitValue, "A")
	if err != nil {
		return nil, err
	}
	defer res.Close()

	// Scan all of the results to the 'packageSizes' array.
	// If it's empty, return null.
	for res.Next() {
		packageSize := &PackageSizeOut{}
		err = res.Scan(&packageSize.ID, &packageSize.Name, &packageSize.Nemo, &packageSize.Limitvalue, &packageSize.CreatedUser, &packageSize.CreatedAt, &packageSize.UpdatedUser, &packageSize.UpdatedAt, &packageSize.Status)
		if err != nil && err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, err
		}
		packageSizes = append(packageSizes, *packageSize)
	}

	// Return all of our packageSizes.
	return &packageSizes, nil
}
