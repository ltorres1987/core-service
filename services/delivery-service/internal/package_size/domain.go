package package_size

import (
	"context"
	"time"
)

// PackageSize struct to describe PackageSize object.
type PackageSize struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Nemo        string    `db:"nemo"`
	Limitvalue  int       `db:"limitvalue"`
	CreatedUser string    `db:"created_user"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedUser string    `db:"updated_user"`
	UpdatedAt   time.Time `db:"updated_at"`
	Status      string    `db:"status"`
}

type PackageSizeOut struct {
	ID          int       `json:"id" `
	Name        string    `json:"name"`
	Nemo        string    `json:"nemo"`
	Limitvalue  int       `json:"limitvalue"`
	CreatedUser string    `json:"created_user"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedUser string    `json:"updated_user"`
	UpdatedAt   time.Time `json:"updated_at"`
	Status      string    `json:"status"`
}

// Our repository will implement these methods.
type PackageSizeRepository interface {
	GetPackageSize(ctx context.Context, packageSizeLimitValue int) (*[]PackageSizeOut, error)
}
