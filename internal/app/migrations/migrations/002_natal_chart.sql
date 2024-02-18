-- +goose Up

CREATE TABLE IF NOT EXISTS natal_chart_result (
       key char(200) not null primary key,
       json_result jsonb not null,
       date_created timestamp not null
);

-- +goose Down

DROP TABLE natal_chart_result
