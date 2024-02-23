-- +goose Up

ALTER TABLE natal_chart_result
    ALTER COLUMN key TYPE character(500);

-- +goose Down

DROP TABLE natal_chart_result
