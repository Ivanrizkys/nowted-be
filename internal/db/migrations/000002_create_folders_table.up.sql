CREATE TABLE IF NOT EXISTS public.folders(
    folder_id VARCHAR (40) PRIMARY KEY,
    user_id VARCHAR (40) REFERENCES public.users(user_id),
    name VARCHAR (100) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);