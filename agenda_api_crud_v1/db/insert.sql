insert into	agenda.contacto (
	nombre_completo,
	numero_documento,
	direccion,
	activo,
	fecha_creacion,
	fecha_modificacion
) values (
'Juan Perez',
'1234567890',
'cll 40',
true,
now(),
now()
);

INSERT INTO agenda.email
(contacto_id , email, activo, fecha_creacion, fecha_modificacion)
VALUES(1, 'juan.perez@gmail.com', true, now(), now());

INSERT INTO agenda.email
(contacto_id, email, activo, fecha_creacion, fecha_modificacion)
VALUES(1, 'juan.perez@outlook.com', true, now(), now());

INSERT INTO agenda.numero_telefonico
(contacto_id, numero_contacto, activo, fecha_creacion, fecha_modificacion)
VALUES(1, '3120000000', true, now(), now());

INSERT INTO agenda.numero_telefonico
(contacto_id, numero_contacto, activo, fecha_creacion, fecha_modificacion)
VALUES(1, '3200000000', true, now(), now());
