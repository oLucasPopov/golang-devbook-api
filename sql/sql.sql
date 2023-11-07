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

drop table if exists publicacoes;
create table publicacoes(
  id serial primary key,
  titulo varchar(50) not null,
  conteudo text not null,
  autor_id integer not null,
  curtidas integer not null default 0,
  criadaEm timestamp not null default current_timestamp,
  foreign key (autor_id) references usuarios(id) on delete cascade
)