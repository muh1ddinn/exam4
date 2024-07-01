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

type manger_teacherRepo struct {
	db *pgxpool.Pool
}

func Newmanagerteacher(db *pgxpool.Pool) storage.Managerteacher_RepoRepoI {

	return &manger_teacherRepo{
		db: db,
	}
}
func (s *manger_teacherRepo) Createteacher(ctx context.Context, req *ct.CreateManagerTeacher) (*ct.ManagerTeacher, error) {
	resp := &ct.ManagerTeacher{}
	id := uuid.New().String()

	// Fetch the next value from the sequence
	var count int
	err := s.db.QueryRow(ctx, "SELECT nextval('teacher_login_seq')").Scan(&count)
	if err != nil {
		log.Println("error while fetching next login sequence value:", err)
		return resp, err
	}

	// Generate the login value
	login := fmt.Sprintf("T%05d", count)

	// Insert the new teacher record into the database
	query := `INSERT INTO teacher (
        id,
        fullname,
        phone,
        password,
        salary,
        ielts_score,
        ielts_attempts_count,
        branch,
        login,
        created_at,
		group_id
     ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW(),$10)`
	_, err = s.db.Exec(ctx, query,
		id,
		req.Fullname,
		req.Phone,
		req.Password,
		req.Salary,
		req.IeltsScore,
		req.IeltsAttemptsCount,
		req.Branch,
		login,
		req.GroupId) // Ensure RoleID is included in req

	if err != nil {
		log.Println("error while creating teacher:", err)
		return resp, err
	}

	// Retrieve the created teacher record by ID
	return s.GetteacherByID(ctx, &ct.ManagerTeacherPrimaryKey{Id: id})
}

func (s *manger_teacherRepo) GetteacherByID(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (*ct.ManagerTeacher, error) {
	resp := &ct.ManagerTeacher{}

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
	 FROM teacher WHERE id=$1 `

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
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}

	return resp, nil
}
func (s *manger_teacherRepo) GetAllteacher(ctx context.Context, req *ct.GetListManagerTeacherRequest) (*ct.GetListManagerTeacherResponse, error) {
	resp := &ct.GetListManagerTeacherResponse{}

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
	FROM teacher
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
		teacher := ct.ManagerTeacher{}
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
			log.Println("error while scanning teacher:", err)
			return nil, err
		}
		resp.ManagerTeachers = append(resp.ManagerTeachers, &teacher)
	}
	return resp, nil
}

func (s *manger_teacherRepo) Updateteacher(ctx context.Context, req *ct.UpdateManagerTeacher) (*ct.ManagerTeacher, error) {
	query := `UPDATE teacher SET 
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
		req.Branch,
		req.RoleId)

	if err != nil {
		log.Println("error while updating teacher:", err)
		return nil, err
	}

	return s.GetteacherByID(ctx, &ct.ManagerTeacherPrimaryKey{Id: req.Id})
}

func (s *manger_teacherRepo) Deleteteacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (*ct.ManagerTeacherPrimaryKey, error) {
	query := `DELETE FROM teacher WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting teacher:", err)
		return nil, err
	}
	return req, nil
}

func (s *manager_crmRepo) Deleteteachersoft(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (*ct.ManagerTeacherPrimaryKey, error) {
	query := `UPDATE teacher SET 
	deleted_at=1 WHERE id = $1 AND deleted_at=0`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting teacher:", err)
		return nil, err
	}
	return req, nil
}
