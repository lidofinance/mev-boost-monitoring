create table if not exists relay_payloads
(
    id                bigserial NOT NULL,
    slot              bigint    NOT NULL,
    block_number      bigint    NOT NULL,
    block_hash        text      NOT NULL,
    fee_recipient     text      NOT NULL,
    transactions_root text      NOT NULL,
    pubkey            text      NOT NULL,
    signature         text      NOT NULL,
    relay_adr         text      NOT NULL,
    relay_timestamp   timestamp not null,
    created_at        timestamp default current_timestamp
);