CREATE USER 'wot'@'172.17.0.0/24' IDENTIFIED WITH caching_sha2_password BY 'wot';
GRANT USAGE ON *.* TO 'wot'@'172.17.0.0/24';
ALTER USER 'wot'@'172.17.0.0/24' REQUIRE NONE WITH MAX_QUERIES_PER_HOUR 0 MAX_CONNECTIONS_PER_HOUR 0 MAX_UPDATES_PER_HOUR 0 MAX_USER_CONNECTIONS 0;

GRANT SELECT, INSERT, UPDATE, DELETE, LOCK TABLES ON `wot`.* TO 'wot'@'172.17.0.0/24';
ALTER USER 'wot'@'172.17.0.0/24' ;
