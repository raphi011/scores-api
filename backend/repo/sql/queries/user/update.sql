UPDATE users
SET 
    profile_image_url = :created_at,
    email = :email,
    volleynet_user_id = :volleynet_user_id,
    volleynet_user = :volleynet_user,
    role = :role,
    salt = :salt,
    hash = :hash,
    iterations = :iterations
WHERE id = :id