package postgres

import (
	"context"
	ct "crmservice/genproto/manager_service"
	"database/sql"

	"crmservice/storage"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type manger_groupRepo struct {
	db *pgxpool.Pool
}

func Newmanagergroup(db *pgxpool.Pool) storage.ManagerGroupService_RepoRepoI {

	return &manger_groupRepo{
		db: db,
	}
}
func (s *manger_groupRepo) CreateJournal(ctx context.Context, req *ct.CreateJournalRequest) (*ct.JournalID, error) {

	id := "9"
	query := `INSERT INTO journal (
        id,
        student_name,
        from_date,
        to_date,
        schedules_id
       -- created_at
    ) VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Exec(ctx, query,
		id,
		req.StudentName,
		req.FromDate,
		req.ToDate,
		req.SchedulesId)

	if err != nil {
		log.Println("error while creating journal:", err)
		return &ct.JournalID{Id: id}, err
	}

	return &ct.JournalID{Id: id}, nil
}
func (s *manger_groupRepo) GetJournalByID(ctx context.Context, req *ct.JournalID) (*ct.Journal, error) {
	resp := &ct.Journal{}
	schedule := &ct.Schedules{}
	var students []*ct.Student

	// Fetch journal details along with schedules
	journalQuery := `
        SELECT 
            j.id,
            j.student_name,
            j.from_date::text,
            j.to_date::text,
            j.schedules_id,
            s.id,
            s.start_time::text,
            s.end_time::text,
            s.lesson_id
        FROM journal j
        LEFT JOIN schedules s ON j.schedules_id = s.id
        WHERE j.id = $1
    `

	err := s.db.QueryRow(ctx, journalQuery, req.Id).Scan(
		&resp.Id,
		&resp.StudentName,
		&resp.FromDate,
		&resp.ToDate,
		&resp.SchedulesId,
		&schedule.Id,
		&schedule.StartTime,
		&schedule.EndTime,
		&schedule.LessonId,
	)
	if err != nil {
		log.Println("error while getting journal by id:", err)
		return nil, err
	}

	resp.Schedules = schedule

	// Fetch students associated with the journal
	studentsQuery := `SELECT id, first_name, last_name FROM student WHERE journal_id = $1`
	rows, err := s.db.Query(ctx, studentsQuery, req.Id)
	if err != nil {
		log.Println("error while getting students by journal id:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		student := &ct.Student{}
		err := rows.Scan(&student.Id, &student.FirstName, &student.LastName)
		if err != nil {
			log.Println("error while scanning student:", err)
			return nil, err
		}
		students = append(students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println("error after scanning students:", err)
		return nil, err
	}

	resp.Students = students

	log.Println("journal fetched successfully with students:", resp)
	return resp, nil
}

func (s *manger_groupRepo) CreateGroup(ctx context.Context, req *ct.CreateGroupRequest) (*ct.GroupID, error) {
	id := uuid.New().String()

	query := `INSERT INTO groups (
         id,
         group_name,
         branch,
         teacher,
         support_teacher,
         journal_id,
         type_id
         --created_at 
		 ) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(ctx, query,
		id,
		req.GroupName,
		req.Branch,
		req.Teacher,
		req.SupportTeacher,
		req.JournalId,
		req.TypeId)

	if err != nil {
		log.Println("error while creating teacher:", err)
		return &ct.GroupID{Id: id}, err
	}

	// Retrieve the created teacher record by ID
	return &ct.GroupID{Id: id}, err
}

func (s *manger_groupRepo) GetGroupByID(ctx context.Context, groupId *ct.GroupID) (*ct.Group, error) {
	group := &ct.Group{}
	journal := &ct.Journal{}
	groupType := &ct.Type{}
	lesson := &ct.Lesson{}
	schedules := &ct.Schedules{}

	query := `
        SELECT 
            g.id,
            g.group_name,
            g.branch,
            g.teacher,
            g.support_teacher,
            j.id,
            j.student_name,
            j.from_date::text,
            j.to_date::text,
            j.schedules_id,
            t.id,
            t.course_name,
            t.description,
            schedules.id,
            schedules.start_time::text,
            schedules.end_time::text,
            schedules.lesson_id,
            lesson.id,
            lesson.lesson_name,
            lesson.task_id
        FROM groups g
        INNER JOIN journal j ON g.journal_id = j.id
        INNER JOIN type t ON g.type_id = t.id
        INNER JOIN schedules ON j.schedules_id = schedules.id
        INNER JOIN lesson ON schedules.lesson_id = lesson.id
        WHERE g.id = $1 AND g.deleted_at = 0 AND j.deleted_at = 0 AND t.deleted_at = 0 AND lesson.deleted_at = 0
    `

	err := s.db.QueryRow(ctx, query, groupId.Id).Scan(
		&group.Id,
		&group.GroupName,
		&group.Branch,
		&group.Teacher,
		&group.SupportTeacher,
		&journal.Id,
		&journal.StudentName,
		&journal.FromDate,
		&journal.ToDate,
		&journal.SchedulesId,
		&groupType.Id,
		&groupType.CourseName,
		&groupType.Description,
		&schedules.Id,
		&schedules.StartTime,
		&schedules.EndTime,
		&schedules.LessonId,
		&lesson.Id,
		&lesson.LessonName,
		&lesson.TaskId,
	)
	if err != nil {
		log.Println("error while getting group by id:", err)
		return nil, err
	}

	group.Journal = journal
	group.Type = groupType
	journal.Schedules = schedules
	schedules.Lesoon = lesson

	return group, nil
}
func (s *manger_groupRepo) CreateTask(ctx context.Context, req *ct.CreateTaskRequest) (*ct.TaskID, error) {
	query := `
		INSERT INTO task (
			label,
			deadline,
			score,
			created_by
		) VALUES ($1, $2, $3, $4)
		RETURNING id`

	var Id string
	err := s.db.QueryRow(ctx, query,
		req.Label,
		req.Deadline,
		req.Score,
		req.CreatedBy).Scan(&Id)

	if err != nil {
		log.Println("error while creating task:", err)
		return nil, err
	}

	return &ct.TaskID{Id: Id}, nil
}

func (s *manger_groupRepo) GetTaskByID(ctx context.Context, req *ct.TaskID) (*ct.Task, error) {
	resp := &ct.Task{}

	query := `SELECT 
        id,
        label,
		deadline::text,
		score,
		created_by
    FROM task WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.Label,
		&resp.Deadline,
		&resp.Score,
		&resp.CreatedBy)

	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_groupRepo) GetAllTasks(ctx context.Context, req *ct.GetAllTasksRequest) (*ct.GetAllTasksResponse, error) {
	resp := &ct.GetAllTasksResponse{}

	filter := ""
	args := []interface{}{}

	if req.Search != "" {
		filter = `WHERE label ILIKE '%' || $1 || '%' OR created_by ILIKE '%' || $1 || '%'`
		args = append(args, req.Search)
	}

	query := `
    SELECT
        count(id) OVER() AS total_count,
        id,
        label,
        deadline::text,
        score,
        created_by
    FROM task
    ` + filter + `
    ORDER BY deadline ASC
	OFFSET $2 LIMIT $3

    `

	offset := (req.Page - 1) * req.Limit
	if offset <= 0 {
		offset = 0
	}

	args = append(args, offset, req.Limit)

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Println("error while getting all tasks:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		task := ct.Task{}
		err := rows.Scan(
			&resp.Count,
			&task.Id,
			&task.Label,
			&task.Deadline,
			&task.Score,
			&task.CreatedBy,
		)
		if err != nil {
			log.Println("error while scanning task:", err)
			return nil, err
		}
		resp.Tasks = append(resp.Tasks, &task)
	}
	return resp, nil
}

func (s *manger_groupRepo) CreateLesson(ctx context.Context, req *ct.CreateLessonRequest) (*ct.LessonID, error) {
	query := `
		INSERT INTO lesson (
			lesson_name,
			task_id
			) VALUES ($1, $2)
		RETURNING id`

	var Id string
	err := s.db.QueryRow(ctx, query,
		req.LessonName,
		req.TaskId).Scan(&Id)

	if err != nil {
		log.Println("error while creating task:", err)
		return nil, err
	}

	return &ct.LessonID{Id: Id}, nil
}

func (s *manger_groupRepo) UpdateTask(ctx context.Context, req *ct.UpdateTaskRequest) (*ct.Task, error) {
	query := `UPDATE task SET 
        label = $1,
        deadline = $2,
        score = $3,
        created_by = $4,
        updated_at = NOW()
    WHERE id = $5`

	_, err := s.db.Exec(ctx, query,
		req.Label,
		req.Deadline,
		req.Score,
		req.CreatedBy,
		req.Id)

	if err != nil {
		log.Println("error while updating task:", err)
		return nil, err
	}

	return s.GetTaskByID(ctx, &ct.TaskID{Id: req.Id})
}

func (s *manger_groupRepo) GetLessonByID(ctx context.Context, req *ct.LessonID) (*ct.Lesson, error) {
	resp := &ct.Lesson{}
	task := &ct.Task{}

	query := `SELECT 
	lesson.id,
	lesson.lesson_name,
	lesson.task_id,
	task.id,
	task.label,
	task.deadline::text,
	task.score ,
	task.created_by 
    FROM lesson    
	JOIN task ON lesson.task_id=task.id   
	WHERE  lesson.id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.Id,
		&resp.LessonName,
		&resp.TaskId,
		&task.Id,
		&task.Label,
		&task.Deadline,
		&task.Score,
		&task.CreatedBy,
	)

	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}

	resp.Task = task
	return resp, nil
}

func (s *manger_groupRepo) CreateType(ctx context.Context, req *ct.CreateTypeRequest) (*ct.TypeID, error) {

	query := `INSERT INTO type (
        course_name,
        description   
    ) VALUES ($1, $2)
	RETURNING id`
	var Id string
	err := s.db.QueryRow(ctx, query,
		req.CourseName,
		req.Description).Scan(&Id)

	if err != nil {
		log.Println("error while creating task:", err)
		return nil, err
	}

	return &ct.TypeID{Id: Id}, nil
}

func (s *manger_groupRepo) GetTypeByID(ctx context.Context, req *ct.TypeID) (*ct.Type, error) {
	resp := &ct.Type{}

	query := `SELECT 
	course_name,
	description,
    FROM type WHERE id=$1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&resp.CourseName,
		&resp.Description)

	if err != nil {
		log.Println("error while getting student by id:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_groupRepo) CreateSchedules(ctx context.Context, req *ct.CreateSchedulesRequest) (*ct.SchedulesID, error) {
	id := uuid.New().String()

	query := `INSERT INTO schedules (
        id,
        start_time,
        end_time,
        lesson_id
    ) VALUES ($1, $2, $3, $4)`

	// Ensure the time format is correct for the database
	_, err := s.db.Exec(ctx, query,
		id,
		req.StartTime,
		req.EndTime,
		req.LessonId)

	if err != nil {
		log.Println("error while creating schedule:", err)
		return nil, err
	}

	return &ct.SchedulesID{Id: id}, nil
}

func (s *manger_groupRepo) GetSchedulesByID(ctx context.Context, req *ct.SchedulesID) (*ct.Schedules, error) {
	lesson := &ct.Lesson{}
	schedules := &ct.Schedules{}

	query := `SELECT 
		schedules.id,
		schedules.start_time::text,
		schedules.end_time::text,
		schedules.lesson_id,
		lesson.id,
		lesson.lesson_name,
		lesson.task_id
	FROM schedules  
	JOIN lesson ON schedules.lesson_id = lesson.id   
	WHERE schedules.id = $1`

	err := s.db.QueryRow(ctx, query, req.Id).Scan(
		&schedules.Id,
		&schedules.StartTime,
		&schedules.EndTime,
		&schedules.LessonId,
		&lesson.Id,
		&lesson.LessonName,
		&lesson.TaskId,
	)

	if err != nil {
		log.Println("error while getting schedule by id:", err)
		return nil, err
	}

	schedules.Lesoon = lesson
	return schedules, nil
}
func (s *manger_groupRepo) GetTeacherInfo(ctx context.Context, req *ct.TeacherGetTeacherInfoRequest) (*ct.TeacherGetTeacherInfoResponse, error) {
	resp := &ct.TeacherGetTeacherInfoResponse{
		Teacher:   &ct.TeacherTeacher{},
		Journals:  []*ct.Journal{},
		Schedules: []*ct.Schedules{},
		Students:  []*ct.Student{},
		TaslScore: []*ct.TaskScore{},
	}

	// Query to fetch teacher, journal, schedules, and student details
	query := `
        SELECT
            t.id AS teacher_id,
            t.fullname AS teacher_fullname,
            j.id AS journal_id,
            s.id AS schedule_id,
            s.start_time::text,
            s.end_time::text,
            st.id AS student_id,
            st.first_name,
            st.last_name,
            ts.score
        FROM teacher t
        LEFT JOIN groups g ON t.id = g.teacher_id
        LEFT JOIN journal j ON j.id = g.journal_id
        LEFT JOIN schedules s ON j.schedules_id = s.id
        LEFT JOIN student st ON j.id = st.journal_id
        LEFT JOIN task_score ts ON st.id = ts.student_id
        WHERE t.id = $1
    `

	rows, err := s.db.Query(ctx, query, req.Id)
	if err != nil {
		log.Println("error while querying teacher information:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		journal := &ct.Journal{}
		schedule := &ct.Schedules{}
		student := &ct.Student{}
		taskScores := &ct.TaskScore{}
		var score sql.NullString

		err := rows.Scan(
			&resp.Teacher.Id,
			&resp.Teacher.Fullname,
			&journal.Id,
			&schedule.Id,
			&schedule.StartTime,
			&schedule.EndTime,
			&student.Id,
			&student.FirstName,
			&student.LastName,
			&score,
		)
		if err != nil {
			log.Println("error while scanning row:", err)
			return nil, err
		}

		// Assign score value to taskScores.Score if it is not NULL
		if score.Valid {
			taskScores.Score = score.String
		} else {
			taskScores.Score = ""
		}

		// Append the scanned entities to their respective slices
		resp.Journals = append(resp.Journals, journal)
		resp.Schedules = append(resp.Schedules, schedule)
		resp.Students = append(resp.Students, student)
		resp.TaslScore = append(resp.TaslScore, taskScores)
	}

	if err := rows.Err(); err != nil {
		log.Println("error after scanning rows:", err)
		return nil, err
	}

	return resp, nil
}

func (s *manger_groupRepo) GetSupporTeacherInfo(ctx context.Context, req *ct.TeacherGetTeacherInfoRequest) (*ct.TeacherGetSupportInfoResponse, error) {
	resp := &ct.TeacherGetSupportInfoResponse{
		Support:   &ct.SupportTeacher{},
		Journals:  []*ct.Journal{},
		Schedules: []*ct.Schedules{},
		Students:  []*ct.Student{},
	}

	// Query to fetch teacher, journal, schedules, and student details
	query := `
        SELECT
            t.id AS teacher_id,
            t.fullname AS teacher_fullname,
            j.id AS journal_id,
            s.id AS schedule_id,
            s.start_time::text,
            s.end_time::text,
            st.id AS student_id,
            st.first_name,
            st.last_name
        FROM teacher t
        LEFT JOIN groups g ON t.id = g.teacher_id
        LEFT JOIN journal j ON j.id = g.journal_id
        LEFT JOIN schedules s ON j.schedules_id = s.id
        LEFT JOIN student st ON j.id = st.journal_id
        WHERE t.id = $1
    `

	rows, err := s.db.Query(ctx, query, req.Id)
	if err != nil {
		log.Println("error while querying teacher information:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		journal := &ct.Journal{}
		schedule := &ct.Schedules{}
		student := &ct.Student{}

		err := rows.Scan(
			&resp.Support.Id,
			&resp.Support.Fullname,
			&journal.Id,
			&schedule.Id,
			&schedule.StartTime,
			&schedule.EndTime,
			&student.Id,
			&student.FirstName,
			&student.LastName,
		)
		if err != nil {
			log.Println("error while scanning row:", err)
			return nil, err
		}

		// Append the scanned entities to their respective slices
		resp.Journals = append(resp.Journals, journal)
		resp.Schedules = append(resp.Schedules, schedule)
		resp.Students = append(resp.Students, student)
	}

	if err := rows.Err(); err != nil {
		log.Println("error after scanning rows:", err)
		return nil, err
	}

	return resp, nil
}
