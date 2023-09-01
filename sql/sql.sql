--POSTGRESQL
create database if not exists devbook;

DROP table if exists usuarios;

CREATE TABLE usuarios (
  id SERIAL PRIMARY KEY,
  nome VARCHAR(50) NOT NULL,
  nick VARCHAR(50) unique NOT NULL,
  email VARCHAR(50) NOT NULL,
  senha VARCHAR(20) NOT NULL
  criado_em TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);