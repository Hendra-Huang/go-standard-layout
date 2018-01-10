CREATE SEQUENCE "users_id_seq";
CREATE TABLE users (
    id integer DEFAULT nextval('users_id_seq'::regclass) NOT NULL,
    name character varying(100),
    email character varying(100)
);
