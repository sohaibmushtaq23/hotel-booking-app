package repository

import (
	"context"
	"database/sql"
	"errors"

	"clientmanager/internal/models"
)

var ErrNotFound = errors.New("client not found")

type ClientRepositoryInterface interface {
	GetByID(ctx context.Context, id int) (*models.Client, error)
	GetAll(ctx context.Context) ([]models.Client, error)
	Create(ctx context.Context, client *models.Client) error
	Update(ctx context.Context, id int, client *models.Client) error
	Delete(ctx context.Context, id int) error
}

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

func (r *ClientRepository) GetAll(ctx context.Context) ([]models.Client, error) {
	query := `
		SELECT ID, ClientCode, CompanyName, Industry,
			Email, Phone, [Website], [Country],[City],
			[Address], IsActive, CreditLimit, CreatedAt
		FROM tbl_clients
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []models.Client

	for rows.Next() {
		var c models.Client

		err := rows.Scan(
			&c.ID,
			&c.ClientCode,
			&c.CompanyName,
			&c.Industry,
			&c.Email,
			&c.Phone,
			&c.Website,
			&c.Country,
			&c.City,
			&c.Address,
			&c.IsActive,
			&c.CreditLimit,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		clients = append(clients, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *ClientRepository) GetByID(ctx context.Context, id int) (*models.Client, error) {
	query := `
		SELECT ID, ClientCode, CompanyName, Industry,
			Email, Phone, [Website], [Country],[City],
			[Address], IsActive, CreditLimit, CreatedAt
		FROM tbl_clients
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var c models.Client
	err := row.Scan(
		&c.ID,
		&c.ClientCode,
		&c.CompanyName,
		&c.Industry,
		&c.Email,
		&c.Phone,
		&c.Website,
		&c.Country,
		&c.City,
		&c.Address,
		&c.IsActive,
		&c.CreditLimit,
		&c.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ClientRepository) Create(ctx context.Context, c *models.Client) error {
	query := `
		INSERT INTO tbl_clients 
		(ClientCode, CompanyName, Industry, Email, Phone, [Website], [Country], [City], [Address], IsActive, CreditLimit)
		OUTPUT INSERTED.ID, INSERTED.CreatedAt
		VALUES (@Code,@Name,@Industry,@Email,@Phone,@Web,@Country, @City, @Address, @Active, @CreditLimit)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("Code", c.ClientCode),
		sql.Named("Name", c.CompanyName),
		sql.Named("Industry", c.Industry),
		sql.Named("Email", c.Email),
		sql.Named("Phone", c.Phone),
		sql.Named("Web", c.Website),
		sql.Named("Country", c.Country),
		sql.Named("City", c.City),
		sql.Named("Address", c.Address),
		sql.Named("Active", c.IsActive),
		sql.Named("CreditLimit", c.CreditLimit),
	).Scan(&c.ID, &c.CreatedAt)
}

func (r *ClientRepository) Update(ctx context.Context, id int, c *models.Client) (*models.Client, error) {
	query := `
        UPDATE tbl_clients
        SET ClientCode=@Code,
            CompanyName=@Name,
            Industry=@Industry,
            Email=@Email,
            Phone=@Phone,
            Website=@Web,
            Country=@Country,
            City=@City,
            Address=@Address,
            IsActive=@Active,
            CreditLimit=@CreditLimit
        OUTPUT INSERTED.ID, INSERTED.ClientCode, INSERTED.CompanyName, INSERTED.Industry,
               INSERTED.Email, INSERTED.Phone, INSERTED.Website, INSERTED.Country,
               INSERTED.City, INSERTED.Address, INSERTED.IsActive, INSERTED.CreditLimit,
               INSERTED.CreatedAt
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("Code", c.ClientCode),
		sql.Named("Name", c.CompanyName),
		sql.Named("Industry", c.Industry),
		sql.Named("Email", c.Email),
		sql.Named("Phone", c.Phone),
		sql.Named("Web", c.Website),
		sql.Named("Country", c.Country),
		sql.Named("City", c.City),
		sql.Named("Address", c.Address),
		sql.Named("Active", c.IsActive),
		sql.Named("CreditLimit", c.CreditLimit),
		sql.Named("id", id),
	)

	var updated models.Client
	err := row.Scan(
		&updated.ID,
		&updated.ClientCode,
		&updated.CompanyName,
		&updated.Industry,
		&updated.Email,
		&updated.Phone,
		&updated.Website,
		&updated.Country,
		&updated.City,
		&updated.Address,
		&updated.IsActive,
		&updated.CreditLimit,
		&updated.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *ClientRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM tbl_clients WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
