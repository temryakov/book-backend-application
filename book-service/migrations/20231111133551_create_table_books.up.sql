create table public.books
(
    title           varchar(30)                                            not null,
    author          varchar(30)                                            not null,
    created_at      timestamp default now()                                not null,
    updated_at      timestamp,
    deleted_at      timestamp,
    id              smallint  default nextval('snippets_id_seq'::regclass) not null
        constraint snippets_pkey
            primary key,
    chapters_amount integer                                                not null
);

alter table public.books
    owner to admin;

