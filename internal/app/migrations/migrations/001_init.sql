-- +goose Up

CREATE TABLE IF NOT EXISTS tgusers (
       id bigint not null primary key,
       username       char(50),
       reg_date   date default CURRENT_DATE
);

-- +goose Down

DROP TABLE tgusers
