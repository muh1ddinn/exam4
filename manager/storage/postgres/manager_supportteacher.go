package postgres

import (
	"context"
	ct "crmservice/genproto/manager_service"
	"crmservice/storage"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type manger_supportteacherRepo struct {
	db *pgxpool.Pool
}

func Newmanagersupportteacher(db *pgxpool.Pool) storage.Managersupportteacher_RepoRepoI {

	return &manger_supportteacherRepo{
		db: db,
	}
}
func (s *manger_supportteacherRepo) Createsupportteacher(ctx context.Context, req *ct.CreateManagerSupportTeacher) (*ct.ManagerSupportTeacher, error) {
	resp := &ct.ManagerSupportTeacher{}
	id := uuid.New().String()

	// Fetch the next value from the sequence
	var count int
	err := s.db.QueryRow(ctx, "SELECT nextval('supportteacher_login_seq')").Scan(&count)
	if err != nil {
		log.Println("error while fetching next login sequence value:", err)
		return resp, err
	}

	// Generate the login value
	login := fmt.Sprintf("S%05d", count)

	// Insert the new teacher record into the database
	query := `INSERT INTO supportteacher (
        id,
        fullname,
        phone,
        password,
        salary,
        ielts_score,
        ielts_attempts_count,
        branch,
        login,
        created_at
        
    ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())`
	_, err = s.db.Exec(ctx, query,
		id,
		req.Fullname,
		req.Phone,
		req.Password,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.Branch,
		login) // Ensure RoleID is included in req

	fmt.Println(err)

	if err != nil {
		log.Println("error while creating supportteacher:", err)
		return resp, err
	}

	// Retrieve the created teacher record by ID
	return s.GetsupportteacherByID(ctx, &ct.ManagerSupportTeacherPrimaryKey{Id: id})
}

func (s *manger_supportteacherRepo) GetsupportteacherByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (*ct.ManagerSupportTeacher, error) {
	resp := &ct.ManagerSupportTeacher{}

	query := `SELECT 
	id,
	fullname,
	phone,
	password,
	salary,
	ielts_score,
	ielts_attempts_count,
	branch,
	role_id,
	login
	 FROM supportteacher WHERE id=$1 `

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Password,
		&resp.Salary,
		&resp.IeltsScore,
		&resp.IeltsAttemptsCount,
		&resp.Branch,
		&resp.RoleId,
		&resp.Login)
	if err != nil {
		log.Println("error while getting supportteacher by id:", err)
		return nil, err
	}

	return resp, nil
}
func (s *manger_supportteacherRepo) GetAllsupportteacher(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (*ct.GetListManagerSupportTeacherResponse, error) {
	resp := &ct.GetListManagerSupportTeacherResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%' OR login ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `
	SELECT
	count(id) OVER() AS total_count,
	id,
	fullname,
	phone,
	password,
	salary,
	ielts_score,
	ielts_attempts_count,
	branch,
	role_id,
	login
	FROM supportteacher
	` + filter + `
	ORDER BY fullname ASC
	OFFSET $2 LIMIT $3
	`

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all teachers:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		teacher := ct.ManagerSupportTeacher{}
		err := rows.Scan(
			&resp.Count,
			&teacher.Id,
			&teacher.Fullname,
			&teacher.Phone,
			&teacher.Password,
			&teacher.Salary,
			&teacher.IeltsScore,
			&teacher.IeltsAttemptsCount,
			&teacher.Branch,
			&teacher.RoleId,
			&teacher.Login)
		if err != nil {
			log.Println("error while scanning supportteacher:", err)
			return nil, err
		}
		resp.ManagerSupportTeachers = append(resp.ManagerSupportTeachers, &teacher)
	}
	return resp, nil
}

func (s *manger_supportteacherRepo) Updatesupportteacher(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (*ct.ManagerSupportTeacher, error) {
	query := `UPDATE supportteacher SET 
        fullname = $2,
        phone = $3,
        password = $4,
        salary = $5,
        ielts_score = $6,
        ielts_attempts_count = $7,
        branch = $8,
        role_id = $9,
        updated_at = NOW()
    WHERE id = $1`

	_, err := s.db.Exec(ctx, query,
		req.Id,
		req.Fullname,
		req.Phone,
		req.Password,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.Branch)

	if err != nil {
		log.Println("error while updating supportteacher:", err)
		return nil, err
	}

	return s.GetsupportteacherByID(ctx, &ct.ManagerSupportTeacherPrimaryKey{Id: req.Id})
}

func (s *manger_supportteacherRepo) Deletesupportteacher(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (*ct.ManagerSupportTeacherPrimaryKey, error) {
	query := `DELETE FROM supportteacher WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting supportteacher:", err)
		return nil, err
	}
	return req, nil
}

func (s *manger_supportteacherRepo) Deletetsupporteachersoft(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (*ct.ManagerSupportTeacherPrimaryKey, error) {
	query := `UPDATE supportteacher SET 
	deleted_at=1 WHERE id = $1 AND deleted_at=0`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting supportteacher:", err)
		return nil, err
	}
	return req, nil
}
