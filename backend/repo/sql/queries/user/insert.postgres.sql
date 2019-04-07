INSERT INTO users (
    created_at,
    email,
    profile_image_url,
    player_id,
    player_login,
    role,
    pw_salt,
    pw_hash,
    pw_iterations
)
VALUES (
    :created_at,
    :email,
    :profile_image_url,
    :player_id,
    :player_login,
    :role,
    :pw_salt,
    :pw_hash,
    :pw_iterations
)
RETURNING id