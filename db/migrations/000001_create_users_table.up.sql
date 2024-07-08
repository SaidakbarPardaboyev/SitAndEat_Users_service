create table users(
    id uuid primary key DEFAULT gen_random_uuid() not null,
    username varchar UNIQUE not null,
    password varchar not null,
    email varchar UNIQUE NOT NULL,
    phone_number varchar UNIQUE not null
);

create table refresh_token(
    id uuid primary key DEFAULT gen_random_uuid() not null,
    user_id uuid references users(id),
    token text UNIQUE NOT NULL,
    expires_at bigint,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP
);