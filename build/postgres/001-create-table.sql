CREATE TABLE Users (
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

CREATE TABLE Questions (
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
