UPDATE users
SET 
    profile_image_url = ?,
    email = ?,
    volleynet_user_id = ?,
    volleynet_login = ?,
    role = ?,
    salt = ?,
    hash = ?,
    iterations = ?
WHERE id = ?