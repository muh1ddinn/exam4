CREATE TABLE teachers (
    id UUID PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    months_worked INT NOT NULL,
    totalsum DECIMAL(10, 2) NOT NULL,
    ielts_score DECIMAL(4, 2) NOT NULL
);

CREATE TABLE support_teachers (
    id UUID PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    months_worked INT NOT NULL,
    totalsum DECIMAL(10, 2) NOT NULL,
    ielts_score DECIMAL(4, 2) NOT NULL
);

CREATE TABLE administrators (
    id UUID PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    salary DECIMAL(10, 2) NOT NULL,
    months_worked INT NOT NULL,
    totalsum DECIMAL(10, 2) NOT NULL,
    ielts_score DECIMAL(4, 2) NOT NULL
);

CREATE TABLE students (
    id UUID PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    paid_sum DECIMAL(10, 2) NOT NULL,
    course_count INT NOT NULL,
    totalsum DECIMAL(10, 2) NOT NULL
);

CREATE TABLE manager (
    id UUID PRIMARY KEY,
    fullname VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    password VARCHAR(100) NOT NULL,
    salary NUMERIC NOT NULL 
);   