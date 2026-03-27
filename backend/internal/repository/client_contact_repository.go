package repository

import (
	"context"
	"database/sql"
	"errors"

	"clientmanager/internal/models"
)

var ErrContactNotFound = errors.New("contact not found")

type ClientContactRepositoryInterface interface {
	GetByID(ctx context.Context, id int) (*models.ClientContact, error)
	GetByClientID(ctx context.Context, idClient int) ([]models.ClientContact, error)
	Create(ctx context.Context, contact *models.ClientContact) error
	Update(ctx context.Context, id int, contact *models.ClientContact) error
	Delete(ctx context.Context, id int) error
}

type ClientContactRepository struct {
	db *sql.DB
}

func NewClientContactRepository(db *sql.DB) *ClientContactRepository {
	return &ClientContactRepository{db: db}
}

func (r *ClientContactRepository) GetByClientID(ctx context.Context, idClient int) ([]models.ClientContact, error) {
	query := `
        SELECT ID, IDClient, FirstName, LastName, Designation,
               Email, Mobile, Gender, BirthDate, IsPrimary,
               Notes, CreatedAt
        FROM tbl_client_contacts
        WHERE IDClient = @p1
        ORDER BY IsPrimary DESC, FirstName, LastName
    `

	rows, err := r.db.QueryContext(ctx, query, sql.Named("p1", idClient))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.ClientContact
	for rows.Next() {
		var c models.ClientContact
		err := rows.Scan(
			&c.ID,
			&c.IDClient,
			&c.FirstName,
			&c.LastName,
			&c.Designation,
			&c.Email,
			&c.Mobile,
			&c.Gender,
			&c.BirthDate,
			&c.IsPrimary,
			&c.Notes,
			&c.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r *ClientContactRepository) GetByID(ctx context.Context, id int) (*models.ClientContact, error) {
	query := `
        SELECT ID, IDClient, FirstName, LastName, Designation,
               Email, Mobile, Gender, BirthDate, IsPrimary,
               Notes, CreatedAt
        FROM tbl_client_contacts
        WHERE ID = @p1
        ORDER BY IsPrimary DESC, FirstName, LastName
    `
	row := r.db.QueryRowContext(ctx, query, id)
	var c models.ClientContact
	err := row.Scan(
		&c.ID,
		&c.IDClient,
		&c.FirstName,
		&c.LastName,
		&c.Designation,
		&c.Email,
		&c.Mobile,
		&c.Gender,
		&c.BirthDate,
		&c.IsPrimary,
		&c.Notes,
		&c.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrContactNotFound
	}
	if err != nil {
		return nil, err
	}

	return &c, nil

}

func (r *ClientContactRepository) Create(ctx context.Context, c *models.ClientContact) error {
	query := `
		INSERT INTO tbl_client_contacts 
		(IDClient, FirstName, LastName, Designation,
			Email, Mobile, Gender, BirthDate, IsPrimary,
			Notes)
		OUTPUT INSERTED.ID, INSERTED.CreatedAt
		VALUES (@IDClient,@FName,@LName,@Designation,@Email,@Mobile,@Gender, @BirthDate, @IsPrimary, @Notes)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("IDClient", c.IDClient),
		sql.Named("FName", c.FirstName),
		sql.Named("LName", c.LastName),
		sql.Named("Designation", c.Designation),
		sql.Named("Email", c.Email),
		sql.Named("Mobile", c.Mobile),
		sql.Named("Gender", c.Gender),
		sql.Named("BirthDate", c.BirthDate),
		sql.Named("IsPrimary", c.IsPrimary),
		sql.Named("Notes", c.Notes),
	).Scan(&c.ID, &c.CreatedAt)
}

func (r *ClientContactRepository) Update(ctx context.Context, id int, c *models.ClientContact) error {
	query := `
		UPDATE tbl_client_contacts
		SET FirstName=@FName,
		    LastName=@LName,
			Designation=@Designation,
		    Email=@Email,
			Mobile=@Mobile,
		    Gender=@Gender,
			BirthDate=@BirthDate, 
			IsPrimary=@IsPrimary,
			Notes=@Notes
		WHERE ID=@id
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		sql.Named("FName", c.FirstName),
		sql.Named("LName", c.LastName),
		sql.Named("Designation", c.Designation),
		sql.Named("Email", c.Email),
		sql.Named("Mobile", c.Mobile),
		sql.Named("Gender", c.Gender),
		sql.Named("BirthDate", c.BirthDate),
		sql.Named("IsPrimary", c.IsPrimary),
		sql.Named("Notes", c.Notes),
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrContactNotFound
	}

	return nil
}

func (r *ClientContactRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM tbl_client_contacts WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrContactNotFound
	}

	return nil
}
