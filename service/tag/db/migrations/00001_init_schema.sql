-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tag
(
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name            text     NOT NULL CHECK ( name <> '' ),
    name_idx        tsvector NOT NULL GENERATED ALWAYS AS ( TO_TSVECTOR('english', name) ) STORED,
    description     text CHECK ( description IS NULL OR description <> '' ),
    description_idx tsvector GENERATED ALWAYS AS ( TO_TSVECTOR('english', description) ) STORED
);
COMMENT ON TABLE tag IS 'Tags that can be applied to entities.';
COMMENT ON COLUMN tag.name IS 'The primary english name of the tag. Must not be empty.';
COMMENT ON COLUMN tag.description IS 'The primary english description of the tag. Must not be empty if set.';

CREATE INDEX tag_name_idx ON tag USING gin (name_idx);
CREATE INDEX tag_description_idx ON tag USING gin (description_idx);


CREATE TABLE IF NOT EXISTS tag_i8n
(
    tag_id          uuid NOT NULL REFERENCES tag (id) ON DELETE CASCADE,
    lang            text NOT NULL,
    name            text NOT NULL CHECK ( name <> '' ),
    name_idx        tsvector GENERATED ALWAYS AS ( TO_TSVECTOR(lang, name) ) STORED,
    description     text CHECK ( description IS NULL OR description <> '' ),
    description_idx tsvector GENERATED ALWAYS AS ( TO_TSVECTOR(lang, description) ) STORED,
    PRIMARY KEY (tag_id, lang)
);
COMMENT ON TABLE tag_i8n IS 'Internationalized names and descriptions for tags.';

CREATE INDEX tag_18n_tag_idx ON tag_i8n (tag_id);
CREATE INDEX tag_i8n_lang_idx ON tag_i8n (lang);
CREATE INDEX tag_i8n_name_idx ON tag_i8n USING gin (name_idx);
CREATE INDEX tag_i8n_description_idx ON tag_i8n USING gin (description_idx);

CREATE TABLE IF NOT EXISTS tag_assignment
(
    id          uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    tag_id      uuid NOT NULL REFERENCES tag (id),
    entity_type text NOT NULL CHECK ( entity_type <> '' ),
    entity_id   uuid NOT NULL,
    score       real NOT NULL    DEFAULT 0.0 CHECK ( score >= 0.0 AND score <= 1.0 ),
    UNIQUE (tag_id, entity_type, entity_id)
);
COMMENT ON TABLE tag_assignment IS 'Assignments of tags to entities.';
COMMENT ON COLUMN tag_assignment.tag_id IS 'The tag that is assigned.';
COMMENT ON COLUMN tag_assignment.entity_type IS 'The type of the entity that is assigned to the tag.';
COMMENT ON COLUMN tag_assignment.entity_id IS 'The id of the entity that is assigned to the tag.';
COMMENT ON COLUMN tag_assignment.score IS 'The score of the assignment. Higher is better.';

CREATE INDEX tag_assignment_entity_idx ON tag_assignment (entity_type, entity_id);
CREATE INDEX tag_assignment_tag_idx ON tag_assignment (tag_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tag_assignment;
DROP TABLE IF EXISTS tag_i8n;
DROP TABLE IF EXISTS tag;
-- +goose StatementEnd
