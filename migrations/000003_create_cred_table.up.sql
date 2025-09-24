CREATE TABLE IF NOT EXISTS creds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    cred_name TEXT NOT NULL,
    resume_url TEXT NOT NULL,
    video_url TEXT NOT NULL
)