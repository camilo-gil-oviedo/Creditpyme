--
-- PostgreSQL database dump
--

\restrict 7CSwmnCTc4t5UHY9WPDgnt2PiHQHfUvVmPzdOt132gHWwa9EiJd7Ybf7jXCjgR1

-- Dumped from database version 18.0
-- Dumped by pg_dump version 18.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: asignacion; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.asignacion (
    id integer NOT NULL,
    solicitud_id integer,
    operador_id uuid,
    estado character varying(50) DEFAULT 'pendiente'::character varying,
    fecha_asignacion timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.asignacion OWNER TO postgres;

--
-- Name: asignacion_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.asignacion_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.asignacion_id_seq OWNER TO postgres;

--
-- Name: asignacion_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.asignacion_id_seq OWNED BY public.asignacion.id;


--
-- Name: cliente; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.cliente (
    id integer NOT NULL,
    nombre character varying(100) NOT NULL,
    apellido character varying(100) NOT NULL,
    correo character varying(100) NOT NULL,
    contrasena character varying(255) NOT NULL,
    fecha_registro timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);


ALTER TABLE public.cliente OWNER TO postgres;

--
-- Name: cliente_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.cliente_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.cliente_id_seq OWNER TO postgres;

--
-- Name: cliente_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.cliente_id_seq OWNED BY public.cliente.id;


--
-- Name: clientes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.clientes (
    id bigint NOT NULL,
    nombre character varying(100) NOT NULL,
    apellido character varying(100) NOT NULL,
    correo character varying(100) NOT NULL,
    contrasena character varying(255) NOT NULL,
    fecha_registro timestamp with time zone
);


ALTER TABLE public.clientes OWNER TO postgres;

--
-- Name: clientes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.clientes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.clientes_id_seq OWNER TO postgres;

--
-- Name: clientes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.clientes_id_seq OWNED BY public.clientes.id;


--
-- Name: empresa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.empresa (
    id integer NOT NULL,
    cliente_id integer,
    nombre character varying(100) NOT NULL,
    direccion character varying(255) NOT NULL,
    ciudad character varying(100) NOT NULL
);


ALTER TABLE public.empresa OWNER TO postgres;

--
-- Name: empresa_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.empresa_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.empresa_id_seq OWNER TO postgres;

--
-- Name: empresa_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.empresa_id_seq OWNED BY public.empresa.id;


--
-- Name: empresas; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.empresas (
    id bigint NOT NULL,
    cliente_id bigint,
    nombre character varying(100) NOT NULL,
    direccion character varying(255) NOT NULL,
    ciudad character varying(100) NOT NULL
);


ALTER TABLE public.empresas OWNER TO postgres;

--
-- Name: empresas_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.empresas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.empresas_id_seq OWNER TO postgres;

--
-- Name: empresas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.empresas_id_seq OWNED BY public.empresas.id;


--
-- Name: operadors; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.operadors (
    id bigint NOT NULL,
    nombre character varying(100) NOT NULL,
    apellido character varying(100) NOT NULL,
    correo character varying(100) NOT NULL,
    total_asignados bigint
);


ALTER TABLE public.operadors OWNER TO postgres;

--
-- Name: operadors_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.operadors_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.operadors_id_seq OWNER TO postgres;

--
-- Name: operadors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.operadors_id_seq OWNED BY public.operadors.id;


--
-- Name: solicitud_credito; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.solicitud_credito (
    id integer NOT NULL,
    cliente_id integer NOT NULL,
    monto_solicitado numeric(12,2) NOT NULL,
    plazo_meses integer NOT NULL,
    destino_credito character varying(255) NOT NULL
);


ALTER TABLE public.solicitud_credito OWNER TO postgres;

--
-- Name: solicitud_credito_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.solicitud_credito_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.solicitud_credito_id_seq OWNER TO postgres;

--
-- Name: solicitud_credito_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.solicitud_credito_id_seq OWNED BY public.solicitud_credito.id;


--
-- Name: solicitud_creditos; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.solicitud_creditos (
    id bigint NOT NULL,
    cliente_id bigint NOT NULL,
    monto_solicitado numeric NOT NULL,
    plazo_meses bigint NOT NULL,
    destino_credito character varying(255) NOT NULL,
    estado character varying(50) NOT NULL,
    operador_id bigint
);


ALTER TABLE public.solicitud_creditos OWNER TO postgres;

--
-- Name: solicitud_creditos_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.solicitud_creditos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.solicitud_creditos_id_seq OWNER TO postgres;

--
-- Name: solicitud_creditos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.solicitud_creditos_id_seq OWNED BY public.solicitud_creditos.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    rol character varying(50) DEFAULT 'cliente'::character varying NOT NULL,
    activo boolean DEFAULT true
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: asignacion id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.asignacion ALTER COLUMN id SET DEFAULT nextval('public.asignacion_id_seq'::regclass);


--
-- Name: cliente id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente ALTER COLUMN id SET DEFAULT nextval('public.cliente_id_seq'::regclass);


--
-- Name: clientes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clientes ALTER COLUMN id SET DEFAULT nextval('public.clientes_id_seq'::regclass);


--
-- Name: empresa id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresa ALTER COLUMN id SET DEFAULT nextval('public.empresa_id_seq'::regclass);


--
-- Name: empresas id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresas ALTER COLUMN id SET DEFAULT nextval('public.empresas_id_seq'::regclass);


--
-- Name: operadors id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.operadors ALTER COLUMN id SET DEFAULT nextval('public.operadors_id_seq'::regclass);


--
-- Name: solicitud_credito id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.solicitud_credito ALTER COLUMN id SET DEFAULT nextval('public.solicitud_credito_id_seq'::regclass);


--
-- Name: solicitud_creditos id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.solicitud_creditos ALTER COLUMN id SET DEFAULT nextval('public.solicitud_creditos_id_seq'::regclass);


--
-- Data for Name: asignacion; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.asignacion (id, solicitud_id, operador_id, estado, fecha_asignacion) FROM stdin;
\.


--
-- Data for Name: cliente; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cliente (id, nombre, apellido, correo, contrasena, fecha_registro) FROM stdin;
\.


--
-- Data for Name: clientes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.clientes (id, nombre, apellido, correo, contrasena, fecha_registro) FROM stdin;
\.


--
-- Data for Name: empresa; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.empresa (id, cliente_id, nombre, direccion, ciudad) FROM stdin;
\.


--
-- Data for Name: empresas; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.empresas (id, cliente_id, nombre, direccion, ciudad) FROM stdin;
\.


--
-- Data for Name: operadors; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.operadors (id, nombre, apellido, correo, total_asignados) FROM stdin;
\.


--
-- Data for Name: solicitud_credito; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.solicitud_credito (id, cliente_id, monto_solicitado, plazo_meses, destino_credito) FROM stdin;
\.


--
-- Data for Name: solicitud_creditos; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.solicitud_creditos (id, cliente_id, monto_solicitado, plazo_meses, destino_credito, estado, operador_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password, rol, activo) FROM stdin;
\.


--
-- Name: asignacion_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.asignacion_id_seq', 1, false);


--
-- Name: cliente_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cliente_id_seq', 1, false);


--
-- Name: clientes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.clientes_id_seq', 1, false);


--
-- Name: empresa_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.empresa_id_seq', 1, false);


--
-- Name: empresas_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.empresas_id_seq', 1, false);


--
-- Name: operadors_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.operadors_id_seq', 1, false);


--
-- Name: solicitud_credito_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.solicitud_credito_id_seq', 1, false);


--
-- Name: solicitud_creditos_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.solicitud_creditos_id_seq', 1, false);


--
-- Name: asignacion asignacion_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.asignacion
    ADD CONSTRAINT asignacion_pkey PRIMARY KEY (id);


--
-- Name: cliente cliente_correo_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT cliente_correo_key UNIQUE (correo);


--
-- Name: cliente cliente_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.cliente
    ADD CONSTRAINT cliente_pkey PRIMARY KEY (id);


--
-- Name: clientes clientes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clientes
    ADD CONSTRAINT clientes_pkey PRIMARY KEY (id);


--
-- Name: empresa empresa_cliente_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresa
    ADD CONSTRAINT empresa_cliente_id_key UNIQUE (cliente_id);


--
-- Name: empresa empresa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresa
    ADD CONSTRAINT empresa_pkey PRIMARY KEY (id);


--
-- Name: empresas empresas_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresas
    ADD CONSTRAINT empresas_pkey PRIMARY KEY (id);


--
-- Name: operadors operadors_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.operadors
    ADD CONSTRAINT operadors_pkey PRIMARY KEY (id);


--
-- Name: solicitud_credito solicitud_credito_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.solicitud_credito
    ADD CONSTRAINT solicitud_credito_pkey PRIMARY KEY (id);


--
-- Name: solicitud_creditos solicitud_creditos_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.solicitud_creditos
    ADD CONSTRAINT solicitud_creditos_pkey PRIMARY KEY (id);


--
-- Name: clientes uni_clientes_correo; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.clientes
    ADD CONSTRAINT uni_clientes_correo UNIQUE (correo);


--
-- Name: operadors uni_operadors_correo; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.operadors
    ADD CONSTRAINT uni_operadors_correo UNIQUE (correo);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_empresas_cliente_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_empresas_cliente_id ON public.empresas USING btree (cliente_id);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: empresa empresa_cliente_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.empresa
    ADD CONSTRAINT empresa_cliente_id_fkey FOREIGN KEY (cliente_id) REFERENCES public.cliente(id) ON DELETE CASCADE;


--
-- Name: solicitud_credito solicitud_credito_cliente_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.solicitud_credito
    ADD CONSTRAINT solicitud_credito_cliente_id_fkey FOREIGN KEY (cliente_id) REFERENCES public.cliente(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict 7CSwmnCTc4t5UHY9WPDgnt2PiHQHfUvVmPzdOt132gHWwa9EiJd7Ybf7jXCjgR1

