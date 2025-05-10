-- +goose Up
-- +goose StatementBegin
CREATE TYPE property_type AS ENUM(
    'vehicle',
    'land',
    'residental_house',
    'investment_property'
);

CREATE TABLE documents (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    property_type property_type NOT NULL,
    name VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    metadata JSONB
);

CREATE INDEX idx_property_type ON documents(property_type);
CREATE INDEX idx_property_created_at ON documents(created_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_property_type;
DROP INDEX idx_property_created_at;
DROP TABLE documents;
DROP TYPE property_type;
-- +goose StatementEnd
