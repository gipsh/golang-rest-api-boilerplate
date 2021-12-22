
-- +migrate Up

CREATE TABLE articles (
        id bigserial NOT NULL,
        created_at timestamptz NULL,
        updated_at timestamptz NULL,
        deleted_at timestamptz NULL,        
        "title" text NULL,
        "body" text NULL,
        CONSTRAINT articles_pkey PRIMARY KEY (id)
);

-- +migrate Down

DROP TABLE articles;
