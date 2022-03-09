CREATE TABLE Users (
    id SERIAL NOT NULL,
    last_name varchar(100) NOT NULL,
    first_name varchar(100) NOT NULL,
    age smallint,
    gender smallint,
    score smallint,
    time smallint,
    status smallint,
    test_datetime timestamp,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    PRIMARY KEY (id)
);
