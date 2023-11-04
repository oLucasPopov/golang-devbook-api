insert into usuarios(nome, nick, email, senha)
values ('Usuário 1', 'usuario1', 'usuario1@usuario.com', '$2a$12$BwbB85f6P6NEc3WsYbqqvu0I.PT0NnqLRcDGZrpa6PBUcoPpv9SeO')
      ,('Usuário 2', 'usuario2', 'usuario2@usuario.com', '$2a$12$BwbB85f6P6NEc3WsYbqqvu0I.PT0NnqLRcDGZrpa6PBUcoPpv9SeO')
      ,('Usuário 3', 'usuario3', 'usuario3@usuario.com', '$2a$12$BwbB85f6P6NEc3WsYbqqvu0I.PT0NnqLRcDGZrpa6PBUcoPpv9SeO')
      ,('Usuário 4', 'usuario4', 'usuario4@usuario.com', '$2a$12$BwbB85f6P6NEc3WsYbqqvu0I.PT0NnqLRcDGZrpa6PBUcoPpv9SeO')      
      ,('Usuário 5', 'usuario5', 'usuario5@usuario.com', '$2a$12$BwbB85f6P6NEc3WsYbqqvu0I.PT0NnqLRcDGZrpa6PBUcoPpv9SeO');
      
insert into seguidores(usuario_id, seguidor_id)     
values (1, 2)
      ,(3,1)
      ,(1,3)
      ,(1,5);