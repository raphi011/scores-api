UPDATE users
SET 
    profile_image_url = :profile_image_url,
    email = :email,
    player_id = :player_id,
    player_login = :player_login,
    role = :role,
    pw_salt = :pw_salt,
    pw_hash = :pw_hash,
    pw_iterations = :pw_iterations
WHERE id = :id