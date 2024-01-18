--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5
-- Dumped by pg_dump version 15.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- Name: brand; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brand (
    id integer NOT NULL,
    name character varying(50),
    type character varying(50),
    seat_capacity character varying(50),
    tank_capacity character varying(50),
    weight_capacity character varying(50)
);


ALTER TABLE public.brand OWNER TO postgres;

--
-- Name: brand_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.brand_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.brand_id_seq OWNER TO postgres;

--
-- Name: brand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.brand_id_seq OWNED BY public.brand.id;


--
-- Name: vehicle; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicle (
    id integer NOT NULL,
    cost character varying(50) NOT NULL,
    yom character varying(50) NOT NULL,
    mileage character varying(50) NOT NULL
);


ALTER TABLE public.vehicle OWNER TO postgres;

--
-- Name: vehicle_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicle_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vehicle_id_seq OWNER TO postgres;

--
-- Name: vehicle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicle_id_seq OWNED BY public.vehicle.id;


--
-- Name: brand id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand ALTER COLUMN id SET DEFAULT nextval('public.brand_id_seq'::regclass);


--
-- Name: vehicle id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle ALTER COLUMN id SET DEFAULT nextval('public.vehicle_id_seq'::regclass);


--
-- Data for Name: brand; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.brand (id, name, type, seat_capacity, tank_capacity, weight_capacity) FROM stdin;
1	BMW	bus	14	40.5	2400
2	Mercedes	car	4	18	1000
\.


--
-- Data for Name: vehicle; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vehicle (id, cost, yom, mileage) FROM stdin;
1	20000	2002	1000
2	10000	1983	5000
\.


--
-- Name: brand_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.brand_id_seq', 2, true);


--
-- Name: vehicle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicle_id_seq', 2, true);


--
-- Name: vehicle vehicle_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicle
    ADD CONSTRAINT vehicle_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

