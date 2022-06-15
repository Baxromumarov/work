CREATE TABLE IF NOT EXISTS datas (
        id Integer NOT NULL,
        user_id Integer NOT NULL,
        title text,
        body text
);

CREATE TABLE IF NOT EXISTS paginations (
		total INTEGER NOT NULL,
		pages INTEGER NOT NULL,
		page INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS links (
		previous TEXT,
		current TEXT,
		next TEXT
);