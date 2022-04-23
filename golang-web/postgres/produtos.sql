DROP TABLE IF EXISTS "produtos";
DROP SEQUENCE IF EXISTS produtos_id_seq;
CREATE SEQUENCE produtos_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."produtos" (
    "id" integer DEFAULT nextval('produtos_id_seq') NOT NULL,
    "nome" character varying(100) NOT NULL,
    "descricao" text NOT NULL,
    "preco" real NOT NULL,
    "quantidade" integer NOT NULL,
    CONSTRAINT "produtos_pkey" PRIMARY KEY ("id")
) WITH (oids = false);