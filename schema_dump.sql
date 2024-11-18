--
-- PostgreSQL database dump
--

-- Dumped from database version 12.19
-- Dumped by pg_dump version 14.14 (Homebrew)

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
-- Name: accounts; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.accounts (
    _id bigint NOT NULL,
    balance bigint DEFAULT 0 NOT NULL,
    owner character varying NOT NULL,
    currency character varying NOT NULL,
    created_at timestamp with time zone DEFAULT '2024-11-13 09:12:35.925309+00'::timestamp with time zone NOT NULL
);


ALTER TABLE public.accounts OWNER TO root;

--
-- Name: accounts__id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.accounts__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accounts__id_seq OWNER TO root;

--
-- Name: accounts__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.accounts__id_seq OWNED BY public.accounts._id;


--
-- Name: entries; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.entries (
    _id bigint NOT NULL,
    account_id bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT '2024-11-13 09:12:35.925309+00'::timestamp with time zone NOT NULL
);


ALTER TABLE public.entries OWNER TO root;

--
-- Name: entries__id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.entries__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.entries__id_seq OWNER TO root;

--
-- Name: entries__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.entries__id_seq OWNED BY public.entries._id;


--
-- Name: entries_account_id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.entries_account_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.entries_account_id_seq OWNER TO root;

--
-- Name: entries_account_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.entries_account_id_seq OWNED BY public.entries.account_id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: transfers; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.transfers (
    _id bigint NOT NULL,
    from_account bigint NOT NULL,
    to_account bigint NOT NULL,
    amount bigint NOT NULL,
    created_at timestamp with time zone DEFAULT '2024-11-13 09:12:35.925309+00'::timestamp with time zone NOT NULL
);


ALTER TABLE public.transfers OWNER TO root;

--
-- Name: COLUMN transfers.amount; Type: COMMENT; Schema: public; Owner: root
--

COMMENT ON COLUMN public.transfers.amount IS 'must be positive';


--
-- Name: transfers__id_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.transfers__id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transfers__id_seq OWNER TO root;

--
-- Name: transfers__id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.transfers__id_seq OWNED BY public.transfers._id;


--
-- Name: transfers_from_account_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.transfers_from_account_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transfers_from_account_seq OWNER TO root;

--
-- Name: transfers_from_account_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.transfers_from_account_seq OWNED BY public.transfers.from_account;


--
-- Name: transfers_to_account_seq; Type: SEQUENCE; Schema: public; Owner: root
--

CREATE SEQUENCE public.transfers_to_account_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transfers_to_account_seq OWNER TO root;

--
-- Name: transfers_to_account_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: root
--

ALTER SEQUENCE public.transfers_to_account_seq OWNED BY public.transfers.to_account;


--
-- Name: accounts _id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.accounts ALTER COLUMN _id SET DEFAULT nextval('public.accounts__id_seq'::regclass);


--
-- Name: entries _id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.entries ALTER COLUMN _id SET DEFAULT nextval('public.entries__id_seq'::regclass);


--
-- Name: entries account_id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.entries ALTER COLUMN account_id SET DEFAULT nextval('public.entries_account_id_seq'::regclass);


--
-- Name: transfers _id; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers ALTER COLUMN _id SET DEFAULT nextval('public.transfers__id_seq'::regclass);


--
-- Name: transfers from_account; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers ALTER COLUMN from_account SET DEFAULT nextval('public.transfers_from_account_seq'::regclass);


--
-- Name: transfers to_account; Type: DEFAULT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers ALTER COLUMN to_account SET DEFAULT nextval('public.transfers_to_account_seq'::regclass);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (_id);


--
-- Name: entries entries_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.entries
    ADD CONSTRAINT entries_pkey PRIMARY KEY (_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: transfers transfers_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_pkey PRIMARY KEY (_id);


--
-- Name: accounts_owner_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX accounts_owner_idx ON public.accounts USING btree (owner);


--
-- Name: entries_account_id_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX entries_account_id_idx ON public.entries USING btree (account_id);


--
-- Name: transfers_amount_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX transfers_amount_idx ON public.transfers USING btree (amount);


--
-- Name: transfers_from_account_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX transfers_from_account_idx ON public.transfers USING btree (from_account);


--
-- Name: transfers_from_account_to_account_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX transfers_from_account_to_account_idx ON public.transfers USING btree (from_account, to_account);


--
-- Name: transfers_to_account_idx; Type: INDEX; Schema: public; Owner: root
--

CREATE INDEX transfers_to_account_idx ON public.transfers USING btree (to_account);


--
-- Name: entries entries_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.entries
    ADD CONSTRAINT entries_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.accounts(_id);


--
-- Name: transfers transfers_from_account_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_from_account_fkey FOREIGN KEY (from_account) REFERENCES public.accounts(_id);


--
-- Name: transfers transfers_to_account_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transfers
    ADD CONSTRAINT transfers_to_account_fkey FOREIGN KEY (to_account) REFERENCES public.accounts(_id);


--
-- PostgreSQL database dump complete
--

