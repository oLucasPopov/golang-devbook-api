--POSTGRESQL
create database if not exists devbook;

DROP table if exists usuarios;

CREATE TABLE usuarios (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(50) NOT NULL,
  nick VARCHAR(50) unique NOT NULL,
  email VARCHAR(50) NOT NULL,
  senha VARCHAR(64) NOT NULL,
  criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

DROP table if exists seguidores;
create table seguidores(
  usuario_id integer not null,
  seguidor_id integer not null,
  foreign key (usuario_id) references usuarios(id) on delete cascade,
  foreign key (seguidor_id) references usuarios(id) on delete cascade,
  primary key (usuario_id, seguidor_id)
)