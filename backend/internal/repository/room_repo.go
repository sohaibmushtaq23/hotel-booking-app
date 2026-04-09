package repository

import (
	"context"
	"database/sql"
	"errors"
	"hotel-booking-backend/internal/models"
)

type RoomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) *RoomRepository {
	return &RoomRepository{db: db}
}

var ErrRoomNotFound = errors.New("room not found")

func (r *RoomRepository) GetAll(ctx context.Context) ([]models.Room, error) {
	query := `
		SELECT ID, [RoomNo], [RoomWidth],[RoomLength],[DoubleBeds],[SingleBeds]
		,[Windows],[AC],[Wifi],[HotWater],[Balcony],[Location],[RoomCharges]
		,[RoomImage],[Remarks],[Status]
		FROM rooms
		ORDER BY ID
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []models.Room

	for rows.Next() {
		var rm models.Room

		err := rows.Scan(
			&rm.ID,
			&rm.RoomNo,
			&rm.RoomWidth,
			&rm.RoomLength,
			&rm.DoubleBeds,
			&rm.SingleBeds,
			&rm.Windows,
			&rm.AC,
			&rm.Wifi,
			&rm.HotWater,
			&rm.Balcony,
			&rm.Location,
			&rm.RoomCharges,
			&rm.RoomImage,
			&rm.Remarks,
			&rm.Status,
		)
		if err != nil {
			return nil, err
		}

		rooms = append(rooms, rm)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *RoomRepository) GetByID(ctx context.Context, id int) (*models.Room, error) {
	query := `
		SELECT ID, [RoomNo], [RoomWidth],[RoomLength],[DoubleBeds],[SingleBeds]
		,[Windows],[AC],[Wifi],[HotWater],[Balcony],[Location],[RoomCharges]
		,[RoomImage],[Remarks],[Status]
		FROM rooms
		WHERE id = @p1
	`

	row := r.db.QueryRowContext(ctx, query, id)
	var rm models.Room
	err := row.Scan(
		&rm.ID,
		&rm.RoomNo,
		&rm.RoomWidth,
		&rm.RoomLength,
		&rm.DoubleBeds,
		&rm.SingleBeds,
		&rm.Windows,
		&rm.AC,
		&rm.Wifi,
		&rm.HotWater,
		&rm.Balcony,
		&rm.Location,
		&rm.RoomCharges,
		&rm.RoomImage,
		&rm.Remarks,
		&rm.Status,
	)

	if err == sql.ErrNoRows {
		return nil, ErrRoomNotFound
	}
	if err != nil {
		return nil, err
	}

	return &rm, nil
}

func (r *RoomRepository) Create(ctx context.Context, rm *models.Room) error {
	query := `
		INSERT INTO rooms 
		([RoomNo], [RoomWidth],[RoomLength],[DoubleBeds],[SingleBeds]
		,[Windows],[AC],[Wifi],[HotWater],[Balcony],[Location],[RoomCharges]
		,[RoomImage],[Remarks],[Status])
		OUTPUT INSERTED.ID
		VALUES (@RoomNo,@RoomWidth,@RoomLength,@DoubleBeds,@SingleBeds,@Windows,@AC, @Wifi, @HotWater, @Balcony, @Location, @RoomCharges, @RoomImage, @Remarks,  @Status)
	`

	return r.db.QueryRowContext(
		ctx,
		query,
		sql.Named("RoomNo", rm.RoomNo),
		sql.Named("RoomWidth", rm.RoomWidth),
		sql.Named("RoomLength", rm.RoomLength),
		sql.Named("DoubleBeds", rm.DoubleBeds),
		sql.Named("SingleBeds", rm.SingleBeds),
		sql.Named("Windows", rm.Windows),
		sql.Named("AC", rm.AC),
		sql.Named("Wifi", rm.Wifi),
		sql.Named("HotWater", rm.HotWater),
		sql.Named("Balcony", rm.Balcony),
		sql.Named("Location", rm.Location),
		sql.Named("RoomCharges", rm.RoomCharges),
		sql.Named("RoomImage", rm.RoomImage),
		sql.Named("Remarks", rm.Remarks),
		sql.Named("Status", rm.Status),
	).Scan(&rm.ID)
}

func (r *RoomRepository) Update(ctx context.Context, id int, rm *models.Room) (*models.Room, error) {
	query := `
        UPDATE rooms
        SET RoomNo=@RoomNo,
			RoomWidth=@RoomWidth,
            RoomLength=@RoomLength,
            DoubleBeds=@DoubleBeds,
            SingleBeds=@SingleBeds,
            Windows=@Windows,
            AC=@AC,
            Wifi=@Wifi,
            HotWater=@HotWater,
            Balcony=@Balcony,
            Location=@Location,
			RoomCharges=@RoomCharges,
			RoomImage=@RoomImage,
			Remarks=@Remarks,
			Status=@Status
        OUTPUT INSERTED.ID, INSERTED.RoomNo, INSERTED.RoomWidth, INSERTED.RoomLength,
               INSERTED.DoubleBeds, INSERTED.SingleBeds, INSERTED.Windows, INSERTED.AC,
               INSERTED.Wifi, INSERTED.HotWater, INSERTED.Balcony, INSERTED.Location,
               INSERTED.RoomCharges, INSERTED.RoomImage, INSERTED.Remarks, INSERTED.Status
        WHERE ID=@id
    `

	row := r.db.QueryRowContext(ctx, query,
		sql.Named("RoomNo", rm.RoomNo),
		sql.Named("RoomWidth", rm.RoomWidth),
		sql.Named("RoomLength", rm.RoomLength),
		sql.Named("DoubleBeds", rm.DoubleBeds),
		sql.Named("SingleBeds", rm.SingleBeds),
		sql.Named("Windows", rm.Windows),
		sql.Named("AC", rm.AC),
		sql.Named("Wifi", rm.Wifi),
		sql.Named("HotWater", rm.HotWater),
		sql.Named("Balcony", rm.Balcony),
		sql.Named("Location", rm.Location),
		sql.Named("RoomCharges", rm.RoomCharges),
		sql.Named("RoomImage", rm.RoomImage),
		sql.Named("Remarks", rm.Remarks),
		sql.Named("Status", rm.Status),
		sql.Named("id", id),
	)

	var updated models.Room
	err := row.Scan(
		&updated.ID,
		&updated.RoomNo,
		&updated.RoomWidth,
		&updated.RoomLength,
		&updated.DoubleBeds,
		&updated.SingleBeds,
		&updated.Windows,
		&updated.AC,
		&updated.Wifi,
		&updated.HotWater,
		&updated.Balcony,
		&updated.Location,
		&updated.RoomCharges,
		&updated.RoomImage,
		&updated.Remarks,
		&updated.Status,
	)
	if err == sql.ErrNoRows {
		return nil, ErrRoomNotFound
	}
	if err != nil {
		return nil, err
	}
	return &updated, nil
}

func (r *RoomRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM rooms WHERE id=@p1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return ErrRoomNotFound
	}

	return nil
}
