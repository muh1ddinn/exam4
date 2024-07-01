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

type manger_administrationRepo struct {
	db *pgxpool.Pool
}

func Newmanageradministration(db *pgxpool.Pool) storage.ManageradministrationRepo_RepoRepoI {

	return &manger_administrationRepo{
		db: db,
	}
}
func (s *manger_administrationRepo) Createadministration(ctx context.Context, req *ct.CreateManagerAdministration) (*ct.ManagerAdministration, error) {
	resp := &ct.ManagerAdministration{}
	id := uuid.New().String()

	// Fetch the next value from the sequence
	var count int
	err := s.db.QueryRow(ctx, "SELECT nextval('administration_login_seq')").Scan(&count)
	if err != nil {
		log.Println("error while fetching next login sequence value:", err)
		return resp, err
	}

	// Generate the login value
	login := fmt.Sprintf("A%05d", count)

	// Insert the new administration record into the database
	query := `INSERT INTO administration (
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

	if err != nil {
		log.Println("error while creating administration:", err)
		return resp, err
	}

	// Retrieve the created administration record by ID
	return s.GetadministrationrByID(ctx, &ct.ManagerAdministrationPrimaryKey{Id: id})
}

func (s *manger_administrationRepo) GetadministrationrByID(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (*ct.ManagerAdministration, error) {
	resp := &ct.ManagerAdministration{}

	query := `SELECT 
	id,
	fullname,
	phone,
	password,
	salary,
	ielts_score,
	ielts_attempts_count,
	branch,
	login
	 FROM administration WHERE id=$1 `

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Fullname,
		&resp.Phone,
		&resp.Password,
		&resp.Salary,
		&resp.IeltsScore,
		&resp.IeltsAttemptsCount,
		&resp.Branch,
		&resp.Login)
	if err != nil {
		log.Println("error while getting administration by id:", err)
		return nil, err
	}

	return resp, nil
}
func (s *manger_administrationRepo) GetAlladministration(ctx context.Context, req *ct.GetListManagerAdministrationRequest) (*ct.GetListManagerAdministrationResponse, error) {
	resp := &ct.GetListManagerAdministrationResponse{}

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
	login
	FROM administration
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
		log.Println("error while getting all administrations:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		administration := ct.ManagerAdministration{}
		err := rows.Scan(
			&resp.Count,
			&administration.Id,
			&administration.Fullname,
			&administration.Phone,
			&administration.Password,
			&administration.Salary,
			&administration.IeltsScore,
			&administration.IeltsAttemptsCount,
			&administration.Branch,
			&administration.Login)
		if err != nil {
			log.Println("error while scanning administration:", err)
			return nil, err
		}
		resp.ManagerAdministrations = append(resp.ManagerAdministrations, &administration)
	}
	return resp, nil
}

func (s *manger_administrationRepo) Updateadministration(ctx context.Context, req *ct.UpdateManagerAdministration) (*ct.ManagerAdministration, error) {
	query := `UPDATE administration SET 
        fullname = $2,
        phone = $3,
        password = $4,
        salary = $5,
        ielts_score = $6,
        ielts_attempts_count = $7,
        branch = $8,
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
		log.Println("error while updating administration:", err)
		return nil, err
	}

	return s.GetadministrationrByID(ctx, &ct.ManagerAdministrationPrimaryKey{Id: req.Id})
}

func (s *manger_administrationRepo) Deleteadministration(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (*ct.ManagerAdministrationPrimaryKey, error) {
	query := `DELETE FROM administration WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting administration:", err)
		return nil, err
	}
	return req, nil
}

func (s *manger_administrationRepo) Deleteadministrationsoft(ctx context.Context, req *ct.ManagerAdministrationPrimaryKey) (*ct.ManagerAdministrationPrimaryKey, error) {
	query := `UPDATE administration SET 
	deleted_at=1 WHERE id = $1 AND deleted_at=0`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting administration:", err)
		return nil, err
	}
	return req, nil
}
