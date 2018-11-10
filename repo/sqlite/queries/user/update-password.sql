UPDATE users
SET salt = ?, hash = ?, iterations = ?
WHERE id = ?