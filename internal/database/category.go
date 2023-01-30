package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Category struct {
	db *sql.DB
	ID string
	Name string
	Description *string
}

func NewCategory(db *sql.DB) *Category  {
	return &Category{db: db}
}

func (c *Category) Create(name string, description *string) (Category, error) {
	id := uuid.New().String()
	if description != nil {
		_, err := c.db.Exec(
			"INSERT INTO category (id, name, description) VALUES ($1, $2, $3)", 
			id, name, *description,
		)
		if err != nil {
			return Category{}, err
		}
	} else {
		_, err := c.db.Exec(
			"INSERT INTO category (id, name) VALUES ($1, $2)", 
			id, name,
		)		
		if err != nil {
			return Category{}, err
		}
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) GetAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []Category{}
	for rows.Next() {
		category := Category{}
		errRow := rows.Scan(&category.ID, &category.Name, &category.Description)
		if errRow != nil {
			return nil, errRow
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *Category) GetByCourseID(courseID string) (Category, error) {
	var category Category
	err := c.db.QueryRow(
		"SELECT ca.id, ca.name, ca.description " + 
		"FROM category AS ca " + 
		"INNER JOIN course AS co " + 
		"ON ca.id = co.category_id " + 
		"WHERE co.id = $1",
		courseID,
	).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return category, err
	}
	return category, nil
}