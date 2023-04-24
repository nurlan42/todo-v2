package todo

import (
	"fmt"
	"github.com/nurlan42/todo/internal/domain/entity"

	"database/sql"
	"runtime/debug"

	_ "github.com/lib/pq"
)

type TODO struct {
	db *sql.DB
}

func NewTODO(db *sql.DB) *TODO {
	return &TODO{
		db: db,
	}
}

func (t *TODO) Create(td entity.TODO) error {
	q := `INSERT INTO adapter (id, public_id, title, description, due_date, completed) 
			VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := t.db.Exec(q, td.ID, td.PublicID, td.Title, td.Description, td.DueDate, td.Completed)
	if err != nil {
		return fmt.Errorf("create(): %v; %s%", err, debug.Stack())
	}

	return nil
}

func (t *TODO) GetByID(ID string) (entity.TODO, error) {
	row, err := t.db.Query("SELECT * from adapter WHERE id = ?;", ID)
	if err != nil {
		return entity.TODO{}, err
	}

	var td entity.TODO
	if err := row.Scan(td); err != nil {
		return entity.TODO{}, fmt.Errorf("GetByID: %v; %v", err, debug.Stack())
	}

	return td, nil
}

func (t *TODO) GetAll() ([]entity.TODO, error) {
	return nil, nil
}

// UpdateByID takes ID, td as parameter
func (t *TODO) UpdateByID(_ string, _ entity.TODO) error {
	return nil
}

// DeleteByID takes ID
func (t *TODO) DeleteByID(_ string) error {
	return nil
}
