CREATE DATABASE eloweb;
\c eloweb;

DROP TABLE IF EXISTS Results;

CREATE TABLE Results(
	Winner varchar(100),
	Loser  varchar(100)
);

INSERT INTO Results VALUES ('Tom','Steve');
INSERT INTO Results VALUES ('Tom','Steve');
INSERT INTO Results VALUES ('Tom','John');
INSERT INTO Results VALUES ('John','Steve');
INSERT INTO Results VALUES ('John','Steve');
