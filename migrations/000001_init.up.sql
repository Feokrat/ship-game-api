CREATE TABLE ship
(
    id UUID NOT NULL ,
    name TEXT NOT NULL ,
    level INT NOT NULL ,
    health FLOAT NOT NULL ,
    attack FLOAT NOT NULL ,
    luck FLOAT NOT NULL ,
    fuel FLOAT NOT NULL ,

    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE commander
(
    id           UUID NOT NULL ,
    name         TEXT NOT NULL ,
    age INT NOT NULL ,

    PRIMARY KEY (id),
    UNIQUE (name)
);

CREATE TABLE fleet
(
    id           UUID NOT NULL ,
    name         TEXT NOT NULL ,
    commander_id UUID REFERENCES commander (id),
    ship_ids     UUID[],

    PRIMARY KEY (id),
    UNIQUE (name, commander_id)
);