CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user' CHECK (role IN ('user', 'admin', 'employer')),
    is_active BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE job_post (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE job_post_application (
    id SERIAL PRIMARY KEY,
    job_post_id INTEGER NOT NULL REFERENCES job_post(id) ON DELETE CASCADE,
    applicant_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    application_text TEXT,
    applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (job_post_id, applicant_id)
);

CREATE TABLE account_action_token (
    id SERIAL PRIMARY KEY,
    account_id INTEGER NOT NULL REFERENCES account(id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    action VARCHAR(50) NOT NULL  CHECK (action IN ('activation', 'recovery')),
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
