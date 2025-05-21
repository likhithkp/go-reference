CREATE TABLE student (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    usn TEXT UNIQUE NOT NULL,
    branch TEXT NOT NULL,
    gpa DOUBLE PRECISION NOT NULL,
    birth_place TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
