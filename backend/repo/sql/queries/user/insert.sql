INSERT INTO users (
    created_at,
    email,
    profile_image_url,
    volleynet_user_id,
    volleynet_login,
    role,
    salt,
    hash,
    iterations
)
VALUES (
    :createdAt,
    :email,
    :profileImageUrl,
    :volleynetUserId,
    :volleynetLogin,
    :role,
    :salt,
    :hash,
    :iterations
)