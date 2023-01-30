package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db *sql.DB
	ID string
	Name string
	Description *string
	CategoryID string
}

func NewCourse(db *sql.DB) *Course  {
	return &Course{db: db}
}

func (c *Course) Create(name string, description *string, categoryID string) (Course, error) {
	id := uuid.New().String()
	if description != nil {
		_, err := c.db.Exec(
			"INSERT INTO course (id, name, description, category_id) VALUES ($1, $2, $3, $4)", 
			id, name, *description, categoryID,
		)
		if err != nil {
			return Course{}, err
		}
	} else {
		_, err := c.db.Exec(
			"INSERT INTO course (id, name, category_id) VALUES ($1, $2, $3)", 
			id, name, categoryID,
		)		
		if err != nil {
			return Course{}, err
		}
	}
	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil
}

func (c *Course) GetAll() ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		course := Course{}
		errRow := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if errRow != nil {
			return nil, errRow
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *Course) GetByCategoryID(categoryID string) ([]Course, error) {
	rows, err := c.db.Query("SELECT * FROM course WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	courses := []Course{}
	for rows.Next() {
		course := Course{}
		errRow := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID)
		if errRow != nil {
			return nil, errRow
		}
		courses = append(courses, course)
	}
	return courses, nil
}