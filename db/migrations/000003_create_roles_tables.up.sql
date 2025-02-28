CREATE TABLE IF NOT EXISTS teachers (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    nip VARCHAR(100) NOT NULL,
    phone_number VARCHAR(50),
    specialization VARCHAR(50),
    status VARCHAR(20),
    profile_picture_url TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS students (
	id SERIAL PRIMARY KEY,
	name VARCHAR(50) NOT NULL,
	nis VARCHAR(50) NOT NULL,
	phone_number VARCHAR(50),
	grade INT,
	curr_score INT,
	status VARCHAR(20),
	profile_picture_url TEXT,
	created_at TIMESTAMPTZ DEFAULT NOW(),
	updated_at TIMESTAMPTZ DEFAULT NOW()
);
