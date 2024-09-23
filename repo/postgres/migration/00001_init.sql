-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto with schema public;
CREATE TABLE POST
(
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- UUID primary key with default generation
    title       VARCHAR(255)                   NOT NULL,    -- Title of the post (maximum 255 characters)
    description TEXT,                                       -- Description of the post (unlimited text)
    created_at  TIMESTAMP        DEFAULT NOW() NOT NULL,    -- Automatically set creation time
    created_by  VARCHAR(255)                   NOT NULL,    -- ID of the user who created the post (UUID)
    updated_at  TIMESTAMP        DEFAULT NOW() not null,    -- Automatically set update time (on creation as well)
    updated_by  VARCHAR(255)                   NOT NULL,    -- ID of the user who last updated the post (UUID)
    thumbnail   TEXT,                                       -- URL or path to the thumbnail (unlimited text)
    content     TEXT,                                       -- Content of the post (unlimited text)
    deleted_at  timestamp
);

CREATE TABLE TAG
(
    ID          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    NAME        VARCHAR(255) NOT NULL UNIQUE,
    DESCRIPTION TEXT         NOT NULL,
    CREATED_AT  TIMESTAMP        DEFAULT NOW(),
    deleted_at  timestamp
);

CREATE TABLE post_tag
(
    post_id    UUID      NOT NULL, -- Foreign key to the posts table (UUID)
    tag_id     UUID      NOT NULL, -- Foreign key to the tags table (UUID)
    created_at timestamp not null,
    PRIMARY KEY (post_id, tag_id), -- Composite primary key (post_id, tag_id)

    CONSTRAINT fk_post
        FOREIGN KEY (post_id)
            REFERENCES post (id)
            ON DELETE CASCADE,

    CONSTRAINT fk_tag
        FOREIGN KEY (tag_id)
            REFERENCES tag (id)
            ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Drop the post_tag table
DROP TABLE IF EXISTS post_tag;

-- Drop the tag table
DROP TABLE IF EXISTS tag;

-- Drop the post table
DROP TABLE IF EXISTS post;

-- Optionally, drop the pgcrypto extension if it was created specifically for this migration
DROP EXTENSION IF EXISTS "pgcrypto";
-- +goose StatementEnd
