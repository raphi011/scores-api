SELECT
    u.id,
    u.created_at,
    u.updated_at,
    u.deleted_at,
    u.email,
    u.pw_hash,
    COALESCE(u.pw_iterations, 0) as pw_iterations,
    COALESCE(u.profile_image_url, '') as profile_image_url,
    u.role,
    u.pw_salt,
    u.player_id,
    u.player_login
FROM users u