package repositories

import (
	"projeto_turismo_jp/db"
	"projeto_turismo_jp/models"
)


func (r *tripRepositoryImpl) DeleteTrip(t *models.Trip) error {
	query := "DELETE FROM trip WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(t.ID)
	return err
}

