DROP TABLE IF EXISTS Names;
CREATE TABLE IF NOT EXISTS Names (
    Id INTEGER NOT NULL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL
);
INSERT INTO Names (Id, Name)
VALUES (1, 'Sundaram'),
    (2, 'Aman'),
    (3, 'Rajesh'),
    (4, 'Ramesh'),
    (5, 'Rakesh');