package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type ReservationRepository struct {
	db *sql.DB
}

func NewReservationRepository(db *sql.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

var ErrReservationNotFound = errors.New("reservation not found")

func (r *ReservationRepository) GetAll(ctx context.Context) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) GetByID(ctx context.Context, id int) (*models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var rv models.Reservation
	err := row.Scan(
		&rv.ID,
		&rv.IDCustomer,
		&rv.IDRoom,
		&rv.BookingStart,
		&rv.BookingEnd,
		&rv.ExtraCharges,
		&rv.AmountPaid,
		&rv.ReservedAt,
		&rv.IDReservedBy,
		&rv.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrReservationNotFound
	}
	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func (r *ReservationRepository) GetByIDClient(ctx context.Context, idClient int) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE IDCustomer=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idClient)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) GetByIDRoom(ctx context.Context, idRoom int) ([]models.Reservation, error) {
	query := `
		SELECT ID, [IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [ReservedAt], [IDReservedBy], [Status]
		FROM reservations
		WHERE IDRoom=@p1
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query, idRoom)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation

	for rows.Next() {
		var rv models.Reservation

		err := rows.Scan(
			&rv.ID,
			&rv.IDCustomer,
			&rv.IDRoom,
			&rv.BookingStart,
			&rv.BookingEnd,
			&rv.ExtraCharges,
			&rv.AmountPaid,
			&rv.ReservedAt,
			&rv.IDReservedBy,
			&rv.Status,
		)
		if err != nil {
			return nil, err
		}

		reservations = append(reservations, rv)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationRepository) Create(ctx context.Context, rv *models.Reservation) (int, error) {
	query := `
		INSERT INTO reservations 
		([IDCustomer], [IDRoom], [BookingStart], [BookingEnd], [ExtraCharges], [AmountPaid], [IDReservedBy], [Status])
		OUTPUT INSERTED.ID, INSERTED.IDRoom
		VALUES (@IDCustomer, @IDRoom, @BookingStart, @BookingEnd, @ExtraCharges, @AmountPaid, @IDReservedBy, @Status)
	`
	var newID, roomID int
	err := r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("IDCustomer", rv.IDCustomer),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("AmountPaid", rv.AmountPaid),
		sql.Named("IDReservedBy", rv.IDReservedBy),
		sql.Named("Status", rv.Status),
	).Scan(&newID, &roomID)

	if err != nil {
		return 0, err
	}
	rv.ID = newID
	return roomID, nil

}

func (r *ReservationRepository) Update(ctx context.Context, id int, rv *models.Reservation) (*models.Reservation, error) {
	query := `
        UPDATE reservations
        SET IDCustomer=@IDCustomer,
			IDRoom=@IDRoom,
			BookingStart=@BookingStart,
			BookingEnd=@BookingEnd,
			ExtraCharges=@ExtraCharges,
            AmountPaid=@AmountPaid,
            IDReservedBy=@IDReservedBy,
			Status=@Status
        OUTPUT INSERTED.ID, INSERTED.IDCustomer, INSERTED.IDRoom, INSERTED.BookingStart
			,INSERTED.BookingEnd, INSERTED.ExtraCharges, INSERTED.AmountPaid
            ,INSERTED.IDReservedBy, INSERTED.Status
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("IDCustomer", rv.IDCustomer),
		sql.Named("IDRoom", rv.IDRoom),
		sql.Named("BookingStart", rv.BookingStart),
		sql.Named("BookingEnd", rv.BookingEnd),
		sql.Named("ExtraCharges", rv.ExtraCharges),
		sql.Named("AmountPaid", rv.AmountPaid),
		sql.Named("IDReservedBy", rv.IDReservedBy),
		sql.Named("Status", rv.Status),
		sql.Named("id", id),
	)

	var updated models.Reservation
	err := row.Scan(
		&updated.ID,
		&updated.IDCustomer,
		&updated.IDRoom,
		&updated.BookingStart,
		&updated.BookingEnd,
		&updated.ExtraCharges,
		&updated.AmountPaid,
		&updated.IDReservedBy,
		&updated.Status,
	)
	if err == sql.ErrNoRows {
		return nil, ErrReservationNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *ReservationRepository) Delete(ctx context.Context, id int) (int, error) {
	var roomID int

	queryGet := `SELECT IDRoom FROM reservations WHERE id=@p1`

	err := r.db.QueryRowContext(ctx, queryGet, id).Scan(&roomID)
	if err != nil {
		return 0, err
	}

	query := `DELETE FROM reservations WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return 0, ErrReservationNotFound
	}

	return roomID, nil
}

func (r *ReservationRepository) GetAllWithDetails(ctx context.Context) ([]models.BookingDetails, error) {
	query := `
		SELECT 
			b.id,
			c.clientName AS customerName,
			r.roomNo,
			b.bookingStart,
			b.bookingEnd,
			b.extraCharges,
			b.amountPaid,
			b.reservedAt,
			u.userName AS reservedBy,
			b.status,
			b.IDCustomer,
			b.IDRoom,
			b.IDReservedBy
		FROM reservations b
		JOIN clients c ON b.idCustomer = c.id
		JOIN rooms r ON b.idRoom = r.id
		JOIN users u ON b.idReservedBy = u.id
		ORDER BY b.bookingStart DESC
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.BookingDetails

	for rows.Next() {
		var b models.BookingDetails
		var bookingStart, bookingEnd, reservedAt sql.NullString // handle NULLs

		err := rows.Scan(
			&b.ID,
			&b.CustomerName,
			&b.RoomNo,
			&bookingStart,
			&bookingEnd,
			&b.ExtraCharges,
			&b.AmountPaid,
			&reservedAt,
			&b.ReservedBy,
			&b.Status,
			&b.IDCustomer,
			&b.IDRoom,
			&b.IDReservedBy,
		)
		if err != nil {
			return nil, err
		}

		// Convert NULL to nil pointers
		if bookingStart.Valid {
			b.BookingStart = &bookingStart.String
		} else {
			b.BookingStart = nil
		}
		if bookingEnd.Valid {
			b.BookingEnd = &bookingEnd.String
		} else {
			b.BookingEnd = nil
		}
		if reservedAt.Valid {
			b.ReservedAt = &reservedAt.String
		} else {
			b.ReservedAt = nil
		}

		bookings = append(bookings, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookings, nil
}
