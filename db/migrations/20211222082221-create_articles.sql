
-- +migrate Up

CREATE TABLE `articles` (
    `id` integer,
    `created_at` datetime,
    `updated_at` datetime,
    `deleted_at` datetime,
    `title` text,
    `body` text,
    PRIMARY KEY (`id`));


-- +migrate Down

DROP TABLE articles;
