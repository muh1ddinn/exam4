package postgres

import (
	"context"
	ct "crmservice/genproto/manager_service"
	"crmservice/storage"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type manager_crmRepo struct {
	db *pgxpool.Pool
}

func Newmanager_crmRepo(db *pgxpool.Pool) storage.Manager_crmRepoRepoI {

	return &manager_crmRepo{
		db: db,
	}
}
func (s *manager_crmRepo) CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (*ct.ManagerStudent, error) {
	resp := &ct.ManagerStudent{}
	id := uuid.New().String()

	// Fetch the next value from the sequence
	var count int
	err := s.db.QueryRow(ctx, "SELECT nextval('student_login_seq')").Scan(&count)
	if err != nil {
		log.Println("error while fetching next login sequence value:", err)
		return resp, err
	}

	// Generate the login value
	login := fmt.Sprintf("ST%05d", count)

	query := `INSERT INTO student (
        id,
        first_name,
        last_name,
        phone,
        password,
        created_at,
        group_id,
        login,
		journal_id
    ) VALUES ($1, $2, $3, $4, $5, NOW(), $6, $7,$8)`
	_, err = s.db.Exec(ctx, query,
		id,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Password,
		req.GroupId,
		login,
		req.Journal)

	if err != nil {
		log.Println("error while creating student:", err)
		return resp, err
	}

	return s.GetStudentByID(ctx, &ct.ManagerStudentPrimaryKey{Id: id})
}

func (s *manager_crmRepo) GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (*ct.ManagerStudent, error) {
	resp := &ct.ManagerStudent{}

	query := `SELECT 
        id,
        first_name,
        last_name,
        phone,
        password,
        group_id,
        role_id,
		login
    FROM student WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.Phone,
		&resp.Password,
		&resp.GroupId,
		&resp.RoleId,
		&resp.Login)

	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}

	return resp, nil
}
func (s *manager_crmRepo) GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (*ct.GetListManagerStudentResponse, error) {
	resp := &ct.GetListManagerStudentResponse{}

	baseFilter := "WHERE deleted_at = 0"
	searchFilter := ""
	args := []interface{}{}

	// Handle search term
	if req.Search != "" {
		searchTerm := fmt.Sprintf("%%%s%%", req.Search)
		searchFilter = ` AND (
			first_name ILIKE $1 OR 
			last_name ILIKE $1 
		)`
		args = append(args, searchTerm)
	}

	// Construct the SQL query
	query := `
	SELECT
		count(id) OVER() AS total_count,
		id,
		first_name,
		last_name,
		phone,
		password,
		group_id,
		role_id,
		created_at
	FROM student
	` + baseFilter + searchFilter + `
	ORDER BY last_name ASC
`
	fmt.Println("query: ", query)
	fmt.Println("args: ", args)

	// Execute the query
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all students:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		student := ct.ManagerStudent{}
		var createdAt time.Time

		err := rows.Scan(
			&resp.Count,
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&student.Phone,
			&student.Password,
			&student.GroupId,
			&student.RoleId,
			&createdAt,
		)
		if err != nil {
			log.Println("error while scanning student:", err)
			return nil, err
		}

		student.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
		resp.ManagerStudents = append(resp.ManagerStudents, &student)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Println("error while iterating over rows:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manager_crmRepo) UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (*ct.ManagerStudent, error) {
	query := `UPDATE student SET 
        first_name = $1,
        last_name = $2,
        phone = $3,
        password = $4,
        group_id = $5,
        updated_at = NOW()
    WHERE id = $6`

	_, err := s.db.Exec(ctx, query,
		req.FirstName,
		req.LastName,
		req.Phone,
		req.Password,
		req.GroupId,
		req.Id)

	if err != nil {
		log.Println("error while updating student:", err)
		return nil, err
	}

	return s.GetStudentByID(ctx, &ct.ManagerStudentPrimaryKey{Id: req.Id})
}

func (s *manager_crmRepo) DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (*ct.ManagerStudentPrimaryKey, error) {
	query := `DELETE FROM student WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting student:", err)
		return nil, err
	}
	return req, nil
}
