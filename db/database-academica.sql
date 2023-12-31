-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 1.0.0-beta1
-- PostgreSQL version: 15.0
-- Project Site: pgmodeler.io
-- Model Author: Andres Felipe Martin Rodriguez

-- object: agenda | type: SCHEMA --
DROP SCHEMA IF EXISTS agenda CASCADE;
CREATE SCHEMA agenda;
-- ddl-end --
ALTER SCHEMA agenda OWNER TO postgres;
-- ddl-end --

SET search_path TO pg_catalog,public,agenda;
-- ddl-end --

-- object: agenda.contacto | type: TABLE --
-- DROP TABLE IF EXISTS agenda.contacto CASCADE;
CREATE TABLE agenda.contacto (
	id serial NOT NULL,
	nombre_completo varchar(60) NOT NULL,
	numero_documento varchar(60),
	direccion varchar(60),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_contacto PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda.contacto OWNER TO postgres;
-- ddl-end --

-- object: agenda.numero_telefonico | type: TABLE --
-- DROP TABLE IF EXISTS agenda.numero_telefonico CASCADE;
CREATE TABLE agenda.numero_telefonico (
	id serial NOT NULL,
	contacto_id integer NOT NULL,
	numero_contacto varchar(10) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_numero_telefonico PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda.numero_telefonico OWNER TO postgres;
-- ddl-end --

-- object: agenda.email | type: TABLE --
-- DROP TABLE IF EXISTS agenda.email CASCADE;
CREATE TABLE agenda.email (
	id serial NOT NULL,
	contacto_id integer NOT NULL,
	email varchar(60) NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_email PRIMARY KEY (id)
);
-- ddl-end --
ALTER TABLE agenda.email OWNER TO postgres;
-- ddl-end --

-- object: fk_numero_telefono_contacto | type: CONSTRAINT --
-- ALTER TABLE agenda.numero_telefonico DROP CONSTRAINT IF EXISTS fk_numero_telefono_contacto CASCADE;
ALTER TABLE agenda.numero_telefonico ADD CONSTRAINT fk_numero_telefono_contacto FOREIGN KEY (contacto_id)
REFERENCES agenda.contacto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_email_contacto | type: CONSTRAINT --
-- ALTER TABLE agenda.email DROP CONSTRAINT IF EXISTS fk_email_contacto CASCADE;
ALTER TABLE agenda.email ADD CONSTRAINT fk_email_contacto FOREIGN KEY (contacto_id)
REFERENCES agenda.contacto (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --


