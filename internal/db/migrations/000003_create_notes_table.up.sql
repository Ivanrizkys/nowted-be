CREATE TABLE IF NOT EXISTS public.notes(
    note_id VARCHAR (40) PRIMARY KEY,
    user_id VARCHAR (40) REFERENCES public.users(user_id),
    folder_id VARCHAR (40) REFERENCES public.folders(folder_id),
    title VARCHAR (255) NOT NULL,
    content JSON,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);