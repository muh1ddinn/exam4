-- Create the custom type 'role_names'
CREATE TYPE role_names AS ENUM ('student', 'teacher', 'supportteacher', 'administration');
CREATE SEQUENCE student_login_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE teacher_login_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE supportteacher_login_seq START WITH 1 INCREMENT BY 1;
CREATE SEQUENCE administration_login_seq START WITH 1 INCREMENT BY 1;




-- Create the 'roles' table
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    role_name role_names UNIQUE NOT NULL
);

-- Insert roles into the 'roles' table
INSERT INTO roles(role_name) VALUES ('student'), ('teacher'), ('supportteacher'), ('administration');

-- Create the 'journal' table


-- Create the 'task' table
CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    label VARCHAR(50),
    deadline TIMESTAMP,
    score FLOAT,
    created_by VARCHAR(50), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'type' table
CREATE TABLE type (
    id SERIAL PRIMARY KEY,
    course_name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'lesson' table
CREATE TABLE lesson (
    id SERIAL PRIMARY KEY,
    lesson_name VARCHAR(100) NOT NULL,
    task_id INT NOT NULL REFERENCES task (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'schedules' table
CREATE TABLE schedules (
    id UUID PRIMARY KEY,
    lesson_id INT NOT NULL REFERENCES lesson (id),
    start_time TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP,
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE journal (
    id SERIAL PRIMARY KEY,
    student_name VARCHAR(50),
    from_date TIMESTAMP,
    to_date TIMESTAMP,
    schedules_id UUID NOT NULL REFERENCES schedules (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);


-- Create the 'groups' table (renamed from 'group')
CREATE TABLE groups (
    id UUID PRIMARY KEY,
    group_name VARCHAR(100) NOT NULL,
    branch VARCHAR(50) NOT NULL,
    teacher_id UUID NOT NULL REFERENCES teacher (id),
     support_teacher VARCHAR(50) NOT NULL,
    journal_id INT NOT NULL REFERENCES journal (id),
    type_id INT NOT NULL REFERENCES type (id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'student' table
CREATE TABLE student (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50)NOT NULL,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    group_id UUID NOT NULL REFERENCES groups (id),
    journal_id INT REFERENCES journal (id),
    role_id INT REFERENCES roles (id) DEFAULT 1,
    login VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

CREATE TABLE task_score (
    id SERIAL NOT NULL,
    student_id UUID NOT NULL REFERENCES student (id),
    score FLOAT,
    task_id INT NOT NULL REFERENCES task (id),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);


-- Create the 'teacher' table
CREATE TABLE teacher (
    id UUID PRIMARY KEY,
    fullname VARCHAR(50) NOT NULL,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    salary FLOAT NOT NULL,
    ielts_score FLOAT,
    ielts_attempts_count VARCHAR(70),
    branch VARCHAR(60),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id INT REFERENCES roles (id) DEFAULT 2,
    group_id UUID NOT NULL REFERENCES groups (id),
    login VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'supportteacher' table
CREATE TABLE supportteacher (
    id UUID PRIMARY KEY,
    fullname VARCHAR(50) NOT NULL,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    salary FLOAT NOT NULL,
    ielts_score FLOAT,
    ielts_attempts_count VARCHAR(70),
    branch VARCHAR(60),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id INT REFERENCES roles (id) DEFAULT 3,
    login VARCHAR(50) NOT NULL,
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'administration' table
CREATE TABLE administration (
    id UUID PRIMARY KEY,
    fullname VARCHAR(50) NOT NULL,
    phone VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    salary FLOAT NOT NULL,
    ielts_score FLOAT,
    ielts_attempts_count VARCHAR(70),
    branch VARCHAR(60),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    role_id INT REFERENCES roles (id) DEFAULT 4,
    login VARCHAR(50) NOT NULL,
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0

);



-- Create the 'student_roles' table
CREATE TABLE student_roles (
    student_id UUID NOT NULL REFERENCES student (id),
    role_id INT NOT NULL REFERENCES roles (id),
    PRIMARY KEY (student_id, role_id),
    updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'teacher_roles' table
CREATE TABLE teacher_roles (
    teacher_id UUID NOT NULL REFERENCES teacher (id),
    role_id INT NOT NULL REFERENCES roles (id),
    PRIMARY KEY (teacher_id, role_id),
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'supportteacher_roles' table
CREATE TABLE supportteacher_roles (
    supportteacher_id UUID NOT NULL REFERENCES supportteacher (id),
    role_id INT NOT NULL REFERENCES roles (id),
    PRIMARY KEY (supportteacher_id, role_id),
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'administration_roles' table
CREATE TABLE administration_roles (
    administration_id UUID NOT NULL REFERENCES administration (id),
    role_id INT NOT NULL REFERENCES roles (id),
    PRIMARY KEY (administration_id, role_id),
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'event' table
CREATE TABLE event (
    id UUID PRIMARY KEY,
    student_name VARCHAR(50) NOT NULL,
    topic VARCHAR(50),
    day VARCHAR(10),
    branch VARCHAR(20),
    start_time TIMESTAMP,
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);

-- Create the 'payment' table
CREATE TABLE payment (
    student_name VARCHAR(50) NOT NULL,
    type VARCHAR(50),
    type_id INT NOT NULL REFERENCES type (id),
    paid_sum FLOAT,
    branch VARCHAR(20),
    from_date TIMESTAMP,
    to_date TIMESTAMP,
     updated_at TIMESTAMP,
    deleted_at INT DEFAULT 0
);
