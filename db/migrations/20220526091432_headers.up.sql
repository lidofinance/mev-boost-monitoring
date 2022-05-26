create table if not exists headers (
    id bigserial NOT NULL,
    version varchar,
    data jsonb default null
);