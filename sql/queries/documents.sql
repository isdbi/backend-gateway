-- name: CreateDocument :one
INSERT INTO documents (
    property_type,
    name,
    content,
    metadata
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetDocument :one
SELECT * FROM documents
WHERE id = $1 LIMIT 1;

-- name: ListDocuments :many
SELECT * FROM documents
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;

-- name: ListDocumentsByPropertyType :many
SELECT * FROM documents
WHERE property_type = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateDocument :one
UPDATE documents
SET 
    property_type = $2,
    name = $3,
    content = $4,
    metadata = $5,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDocument :exec
DELETE FROM documents
WHERE id = $1;

-- name: SearchDocuments :many
SELECT * FROM documents
WHERE name ILIKE $1 OR content ILIKE $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;
