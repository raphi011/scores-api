INSERT INTO users (
    created_at,
    email,
    profile_image_url,
    volleynet_user_id,
    volleynet_login,
    role
)
VALUES (CURRENT_TIMESTAMP, ?, ?, ?, ?, ?)