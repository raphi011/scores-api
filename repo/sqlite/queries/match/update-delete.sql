UPDATE matches
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?