CREATE TABLE IF NOT EXISTS users (
    id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    balance DECIMAL(10,5),
    name STRING NOT NULL UNIQUE,
    lat FLOAT NOT NULL,
    lng FLOAT NOT NULL,
    PRIMARY KEY (id, name)       
);

CREATE TABLE IF NOT EXISTS places (
    id UUID UNIQUE NOT NULL DEFAULT gen_random_uuid(),
    name STRING NOT NULL UNIQUE,
    category STRING NOT NULL,
    lat FLOAT NOT NULL,
    lng FLOAT NOT NULL,
    PRIMARY KEY (id, name, category)
);

CREATE TABLE IF NOT EXISTS comments (
    id UUID    NOT NULL DEFAULT gen_random_uuid(),
    user_id UUID  NOT NULL REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE,
    place_id UUID NOT NULL REFERENCES places (id) ON UPDATE CASCADE ON DELETE CASCADE,
    context STRING NOT NULL,
    is_pay BOOL NOT NULL,
    PRIMARY KEY (id, place_id, user_id)
);