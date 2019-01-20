INSERT INTO users (
    created_at,
    updated_at,
    email,
    profile_image_url,
    volleynet_user_id,
    volleynet_user,
    role,
    salt,
    hash,
    iterations
)
VALUES (
    :created_at,
    :updated_at,
    :email,
    :profile_image_url,
    :volleynet_user_id,
    :volleynet_user,
    :role,
    :salt,
    :hash,
    :iterations
)
RETURNING id