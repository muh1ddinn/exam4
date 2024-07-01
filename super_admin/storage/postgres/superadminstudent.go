package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	ct "superadmin/genproto/superadmin_service"
	"superadmin/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type manger_adminRepo struct {
	db *pgxpool.Pool
}

func Newsuperadminstudent(db *pgxpool.Pool) storage.Superadmin_crmRepoRepoI {

	return &manger_adminRepo{
		db: db,
	}
}

func (s *manger_adminRepo) CreateStudent(ctx context.Context, req *ct.CreateManagerStudent) (*ct.ManagerStudent, error) {
	resp := &ct.ManagerStudent{}
	id := uuid.New().String()

	query := `INSERT INTO students (
        id, fullname, phone, paid_sum, course_count, totalsum
    ) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Exec(ctx, query,
		id, req.Fullname, req.Phone, req.PaidSum, req.CourseCount, req.Totalsum)

	if err != nil {
		log.Println("error while creating student:", err)
		return resp, err
	}

	return s.GetStudentByID(ctx, &ct.ManagerStudentPrimaryKey{Id: id})
}

func (s *manger_adminRepo) GetStudentByID(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (*ct.ManagerStudent, error) {
	resp := &ct.ManagerStudent{}

	query := `SELECT id, fullname, phone, paid_sum, course_count, totalsum FROM students WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id, &resp.Fullname, &resp.Phone, &resp.PaidSum, &resp.CourseCount, &resp.Totalsum)

	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_adminRepo) GetAllStudents(ctx context.Context, req *ct.GetListManagerStudentRequest) (*ct.GetListManagerStudentResponse, error) {
	resp := &ct.GetListManagerStudentResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `SELECT count(id) OVER() AS total_count, id, fullname, phone, paid_sum, course_count, totalsum FROM students ` + filter + ` ORDER BY fullname ASC OFFSET $2 LIMIT $3`

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all students:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		student := ct.ManagerStudent{}
		err := rows.Scan(
			&resp.Count, &student.Id, &student.Fullname, &student.Phone, &student.PaidSum, &student.CourseCount, &student.Totalsum,
		)
		if err != nil {
			log.Println("error while scanning student:", err)
			return nil, err
		}
		resp.ManagerStudents = append(resp.ManagerStudents, &student)
	}
	return resp, nil
}

func (s *manger_adminRepo) UpdateStudent(ctx context.Context, req *ct.UpdateManagerStudent) (*ct.ManagerStudent, error) {
	query := `UPDATE students SET fullname = $1, phone = $2, paid_sum = $3, course_count = $4, totalsum = $5, updated_at = NOW() WHERE id = $6`

	_, err := s.db.Exec(ctx, query, req.Fullname, req.Phone, req.PaidSum, req.CourseCount, req.Totalsum, req.Id)
	if err != nil {
		log.Println("error while updating student:", err)
		return nil, err
	}

	return s.GetStudentByID(ctx, &ct.ManagerStudentPrimaryKey{Id: req.Id})
}

func (s *manger_adminRepo) DeleteStudent(ctx context.Context, req *ct.ManagerStudentPrimaryKey) (*ct.ManagerStudentPrimaryKey, error) {
	query := `DELETE FROM students WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting student:", err)
		return nil, err
	}
	return req, nil
}

func (s *manger_adminRepo) CreateSupportTeacher(ctx context.Context, req *ct.CreateManagerSupportTeacher) (*ct.ManagerSupportTeacher, error) {
	resp := &ct.ManagerSupportTeacher{}
	id := uuid.New().String()

	query := `INSERT INTO support_teachers (
        id, fullname, phone, salary, months_worked, totalsum, ielts_score
    ) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(ctx, query,
		id, req.Fullname, req.Phone, req.Salary, req.MonthsWorked, req.Totalsum, req.IeltsScore)

	if err != nil {
		log.Println("error while creating support teacher:", err)
		return resp, err
	}

	return s.GetSupportTeacherByID(ctx, &ct.ManagerSupportTeacherPrimaryKey{Id: id})
}

func (s *manger_adminRepo) GetSupportTeacherByID(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (*ct.ManagerSupportTeacher, error) {
	resp := &ct.ManagerSupportTeacher{}

	query := `SELECT id, fullname, phone, salary, months_worked, totalsum, ielts_score FROM support_teachers WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id, &resp.Fullname, &resp.Phone, &resp.Salary, &resp.MonthsWorked, &resp.Totalsum, &resp.IeltsScore)

	if err != nil {
		log.Println("error while getting support teacher by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_adminRepo) GetAllSupportTeachers(ctx context.Context, req *ct.GetListManagerSupportTeacherRequest) (*ct.GetListManagerSupportTeacherResponse, error) {
	resp := &ct.GetListManagerSupportTeacherResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `SELECT count(id) OVER() AS total_count, id, fullname, phone, salary, months_worked, totalsum, ielts_score FROM support_teachers ` + filter + ` ORDER BY fullname ASC OFFSET $2 LIMIT $3`

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all support teachers:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		supportTeacher := ct.ManagerSupportTeacher{}
		err := rows.Scan(
			&resp.Count, &supportTeacher.Id, &supportTeacher.Fullname, &supportTeacher.Phone, &supportTeacher.Salary, &supportTeacher.MonthsWorked, &supportTeacher.Totalsum, &supportTeacher.IeltsScore,
		)
		if err != nil {
			log.Println("error while scanning support teacher:", err)
			return nil, err
		}
		resp.ManagerSupportTeachers = append(resp.ManagerSupportTeachers, &supportTeacher)
	}
	return resp, nil
}

func (s *manger_adminRepo) UpdateSupportTeacher(ctx context.Context, req *ct.UpdateManagerSupportTeacher) (*ct.ManagerSupportTeacher, error) {
	query := `UPDATE support_teachers SET fullname = $1, phone = $2, salary = $3, months_worked = $4, totalsum = $5, ielts_score = $6, updated_at = NOW() WHERE id = $7`

	_, err := s.db.Exec(ctx, query, req.Fullname, req.Phone, req.Salary, req.MonthsWorked, req.Totalsum, req.IeltsScore, req.Id)
	if err != nil {
		log.Println("error while updating support teacher:", err)
		return nil, err
	}

	return s.GetSupportTeacherByID(ctx, &ct.ManagerSupportTeacherPrimaryKey{Id: req.Id})
}
func (s *manger_adminRepo) DeleteSupportTeacher(ctx context.Context, req *ct.ManagerSupportTeacherPrimaryKey) (*ct.ManagerSupportTeacherPrimaryKey, error) {
	query := `DELETE FROM support_teachers WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting support teacher:", err)
		return nil, err
	}
	return req, nil
}

func (s *manger_adminRepo) CreateTeacher(ctx context.Context, req *ct.CreateManagerTeacher) (*ct.ManagerTeacher, error) {
	resp := &ct.ManagerTeacher{}
	id := uuid.New().String()

	query := `INSERT INTO teachers (
        id, fullname, phone, salary, months_worked, totalsum, ielts_score
    ) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(ctx, query,
		id, req.Fullname, req.Phone, req.Salary, req.MonthsWorked, req.Totalsum, req.IeltsScore)

	if err != nil {
		log.Println("error while creating teacher:", err)
		return resp, err
	}

	return s.GetTeacherByID(ctx, &ct.ManagerTeacherPrimaryKey{Id: id})
}

func (s *manger_adminRepo) GetTeacherByID(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (*ct.ManagerTeacher, error) {
	resp := &ct.ManagerTeacher{}

	query := `SELECT id, fullname, phone, salary, months_worked, totalsum, ielts_score FROM teachers WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id, &resp.Fullname, &resp.Phone, &resp.Salary, &resp.MonthsWorked, &resp.Totalsum, &resp.IeltsScore)

	if err != nil {
		log.Println("error while getting teacher by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_adminRepo) GetAllTeachers(ctx context.Context, req *ct.GetListManagerTeacherRequest) (*ct.GetListManagerTeacherResponse, error) {
	resp := &ct.GetListManagerTeacherResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `SELECT count(id) OVER() AS total_count, id, fullname, phone, salary, months_worked, totalsum, ielts_score FROM teachers ` + filter + ` ORDER BY fullname ASC OFFSET $2 LIMIT $3`

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
			&resp.Count, &teacher.Id, &teacher.Fullname, &teacher.Phone, &teacher.Salary, &teacher.MonthsWorked, &teacher.Totalsum, &teacher.IeltsScore,
		)
		if err != nil {
			log.Println("error while scanning teacher:", err)
			return nil, err
		}
		resp.ManagerTeachers = append(resp.ManagerTeachers, &teacher)
	}
	return resp, nil
}
func (s *manger_adminRepo) UpdateTeacher(ctx context.Context, req *ct.UpdateManagerTeacher) (*ct.ManagerTeacher, error) {
	query := `UPDATE teachers SET fullname = $1, phone = $2, salary = $3, months_worked = $4, totalsum = $5, ielts_score = $6, updated_at = NOW() WHERE id = $7`

	_, err := s.db.Exec(ctx, query, req.Fullname, req.Phone, req.Salary, req.MonthsWorked, req.Totalsum, req.IeltsScore, req.Id)
	if err != nil {
		log.Println("error while updating teacher:", err)
		return nil, err
	}

	return s.GetTeacherByID(ctx, &ct.ManagerTeacherPrimaryKey{Id: req.Id})
}

func (s *manger_adminRepo) DeleteTeacher(ctx context.Context, req *ct.ManagerTeacherPrimaryKey) (*ct.ManagerTeacherPrimaryKey, error) {
	query := `DELETE FROM teachers WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting teacher:", err)
		return nil, err
	}
	return req, nil
}

func (s *manger_adminRepo) CreateManager(ctx context.Context, req *ct.CreateManagerRequest) (*ct.Manager, error) {
	id := uuid.New().String()

	query := `INSERT INTO manager (
        id, fullname, phone, password, salary
    ) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(ctx, query, id, req.Fullname, req.Phone, req.Password, req.Salary)
	if err != nil {
		log.Println("error while creating manager:", err)
		return nil, err
	}

	return s.GetManagerByID(ctx, &ct.ManagerPrimaryKey{Id: id})
}

func (s *manger_adminRepo) GetManagerByID(ctx context.Context, req *ct.ManagerPrimaryKey) (resp *ct.Manager, err error) {
	resp = &ct.Manager{}

	query := `SELECT id, fullname, phone, password, salary FROM manager WHERE id=$1`
	err = s.db.QueryRow(ctx, query, req.Id).Scan(&resp.Id, &resp.Fullname, &resp.Phone, &resp.Password, &resp.Salary)
	if err != nil {
		log.Println("error while getting manager by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_adminRepo) GetAllManagers(ctx context.Context, req *ct.GetAllManagersRequest) (*ct.GetAllManagersResponse, error) {
	resp := &ct.GetAllManagersResponse{}

	filter := ""
	args := []interface{}{}
	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%' OR phone ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `
        SELECT count(id) OVER() AS total_count, id, fullname, phone, password, salary
        FROM manager
        ` + filter + `
        ORDER BY fullname ASC
        OFFSET $2 LIMIT $3`

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all managers:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		manager := ct.Manager{}
		err := rows.Scan(&resp.TotalCount, &manager.Id, &manager.Fullname, &manager.Phone, &manager.Password, &manager.Salary)
		if err != nil {
			log.Println("error while scanning manager:", err)
			return nil, err
		}
		resp.Managers = append(resp.Managers, &manager)
	}
	return resp, nil
}

func (s *manger_adminRepo) UpdateManager(ctx context.Context, req *ct.UpdateManagerRequest) (*ct.Manager, error) {
	query := `UPDATE manager SET 
        fullname = $1, phone = $2, password = $3, salary = $4, updated_at = NOW()
    WHERE id = $5`

	_, err := s.db.Exec(ctx, query, req.Fullname, req.Phone, req.Password, req.Salary, req.Id)
	if err != nil {
		log.Println("error while updating manager:", err)
		return nil, err
	}

	return s.GetManagerByID(ctx, &ct.ManagerPrimaryKey{Id: req.Id})
}

func (s *manger_adminRepo) DeleteManager(ctx context.Context, req *ct.ManagerPrimaryKey) (*ct.ManagerPrimaryKey, error) {
	query := `DELETE FROM manager WHERE id = $1`
	_, err := s.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting manager:", err)
		return nil, err
	}
	return req, nil
}

func (r *manger_adminRepo) CreateAdministrator(ctx context.Context, req *ct.CreateManagerAdministrator) (resp *ct.ManagerAdministrator, err error) {
	resp = &ct.ManagerAdministrator{}
	id := uuid.New().String()

	query := `INSERT INTO administrators (
        id, fullname, phone, salary, months_worked, totalsum, ielts_score
    ) VALUES ($1, $2, $3, $4, $5,$6,$7)`
	_, err = r.db.Exec(ctx, query,
		id, req.Fullname, req.Phone, req.Salary, req.MonthsWorked, req.Totalsum, req.IeltsScore)

	if err != nil {
		log.Println("error while creating administrator:", err)
		return resp, err
	}

	return r.GetAdministratorByID(ctx, &ct.ManagerAdministratorPrimaryKey{Id: id})
}

func (r *manger_adminRepo) GetAdministratorByID(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (resp *ct.ManagerAdministrator, err error) {
	resp = &ct.ManagerAdministrator{}

	query := `SELECT id, fullname, phone, salary, months_worked, totalsum, ielts_score
	FROM administrators WHERE id=$1`

	err = r.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id, &resp.Fullname, &resp.Phone, &resp.Salary, &resp.MonthsWorked, &resp.Totalsum, &resp.IeltsScore)

	if err != nil {
		log.Println("error while getting administrator by id:", err)
		return nil, err
	}

	return resp, nil
}
func (r *manger_adminRepo) GetAllAdministrators(ctx context.Context, req *ct.GetListManagerAdministratorRequest) (*ct.GetListManagerAdministratorResponse, error) {
	resp := &ct.GetListManagerAdministratorResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE fullname ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `SELECT count(id) OVER() AS total_count, id, fullname, phone, salary, months_worked, totalsum, ielts_score 
	FROM administrators ` + filter + ` ORDER BY fullname ASC OFFSET $2 LIMIT $3`

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all administrators:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		admin := ct.ManagerAdministrator{}
		err := rows.Scan(
			&admin.Id, &admin.Fullname, &admin.Phone, &admin.Salary, &admin.MonthsWorked, &admin.Totalsum, &admin.IeltsScore,
		)
		if err != nil {
			log.Println("error while scanning administrator:", err)
			return nil, err
		}
		resp.ManagerAdministrators = append(resp.ManagerAdministrators, &admin)
	}

	if err := rows.Err(); err != nil {
		log.Println("error with rows:", err)
		return nil, err
	}

	resp.Count = int32(len(resp.ManagerAdministrators))
	return resp, nil
}

func (r *manger_adminRepo) DeleteAdministrator(ctx context.Context, req *ct.ManagerAdministratorPrimaryKey) (*ct.ManagerAdministratorPrimaryKey, error) {
	query := `DELETE FROM administrators WHERE id=$1`

	_, err := r.db.Exec(ctx, query, req.Id)
	if err != nil {
		log.Println("error while deleting administrator:", err)
		return nil, err
	}

	return req, nil
}
func (s *manger_adminRepo) ExportToCSV(ctx context.Context, req *ct.TableName) (resp *ct.Empty, err error) {
	conn, err := s.db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	file, err := os.Create(req.Outfilename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	exportQuery := fmt.Sprintf(`COPY (SELECT * FROM %s) TO STDOUT WITH CSV HEADER`, req.TableName)
	_, err = conn.Conn().PgConn().CopyTo(ctx, file, exportQuery)
	if err != nil {
		return nil, err
	}

	return &ct.Empty{}, nil
}
