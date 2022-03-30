CREATE TABLE users (
    id SERIAL NOT NULL,
    last_name varchar(100) NOT NULL,
    first_name varchar(100) NOT NULL,
    age smallint,
    gender smallint,
    score smallint,
    time smallint,
    status smallint DEFAULT 0,
    test_datetime timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE questions (
    id SERIAL NOT NULL,
    content varchar(100) NOT NULL,
    image varchar(100) NOT NULL,
    number smallint NOT NULL,
    short_name varchar(100) NOT NULL,
    status smallint DEFAULT 0,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);

CREATE TABLE question_answers (
    id SERIAL NOT NULL,
    form_type smallint NOT NULL,
    question_id bigint NOT NULL,
    answer varchar(100) NOT NULL,
    label_position smallint,
    label varchar(100),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id),
    FOREIGN KEY (question_id)
    REFERENCES questions(id)
);
