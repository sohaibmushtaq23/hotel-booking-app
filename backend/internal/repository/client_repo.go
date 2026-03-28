package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type ClientRepository struct {
	db *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{db: db}
}

var ErrClientNotFound = errors.New("client not found")

func (r *ClientRepository) GetAll(ctx context.Context) ([]models.Client, error) {
	query := `
		SELECT ID, [ClientName], [CNIC], [Phone], [Email], [Discount]
		FROM clients
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
			&c.ClientName,
			&c.CNIC,
			&c.Phone,
			&c.Email,
			&c.Discount,
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
		SELECT ID, [ClientName], [CNIC], [Phone], [Email], [Discount]
		FROM clients
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var c models.Client
	err := row.Scan(
		&c.ID,
		&c.ClientName,
		&c.CNIC,
		&c.Phone,
		&c.Email,
		&c.Discount,
	)

	if err == sql.ErrNoRows {
		return nil, ErrClientNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ClientRepository) Create(ctx context.Context, c *models.Client) error {
	query := `
		INSERT INTO clients 
		([ClientName], [CNIC],[Phone],[Email],[Discount])
		OUTPUT INSERTED.ID
		VALUES (@ClientName,@CNIC,@Phone,@Email,@Discount)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("ClientName", c.ClientName),
		sql.Named("CNIC", c.CNIC),
		sql.Named("Phone", c.Phone),
		sql.Named("Email", c.Email),
		sql.Named("Discount", c.Discount),
	).Scan(&c.ID)
}

func (r *ClientRepository) Update(ctx context.Context, id int, c *models.Client) (*models.Client, error) {
	query := `
        UPDATE clients
        SET ClientName=@ClientName,
			CNIC=@CNIC,
            Phone=@Phone,
            Email=@Email,
            Discount=@Discount
        OUTPUT INSERTED.ID, INSERTED.ClientName, INSERTED.CNIC, INSERTED.Phone,
               INSERTED.Email, INSERTED.Discount
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("ClientName", c.ClientName),
		sql.Named("CNIC", c.CNIC),
		sql.Named("Phone", c.Phone),
		sql.Named("Email", c.Email),
		sql.Named("Discount", c.Discount),
		sql.Named("id", id),
	)

	var updated models.Client
	err := row.Scan(
		&updated.ID,
		&updated.ClientName,
		&updated.CNIC,
		&updated.Phone,
		&updated.Email,
		&updated.Discount,
	)
	if err == sql.ErrNoRows {
		return nil, ErrClientNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *ClientRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM clients WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrClientNotFound
	}

	return nil
}
