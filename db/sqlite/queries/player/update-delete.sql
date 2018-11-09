UPDATE players
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = ?