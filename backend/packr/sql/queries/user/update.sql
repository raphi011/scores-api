UPDATE users
SET 
    profile_image_url = :createdAt,
    email = :email,
    volleynet_user_id = :volleynetuserid,
    volleynet_user = :volleynetuser,
    role = :role,
    salt = :salt,
    hash = :hash,
    iterations = :iterations
WHERE id = :id