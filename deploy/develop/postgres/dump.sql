--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Debian 16.3-1.pgdg120+1)
-- Dumped by pg_dump version 16.3 (Debian 16.3-1.pgdg120+1)

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

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: activestatus; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.activestatus AS ENUM (
    'active',
    'pause'
);


ALTER TYPE public.activestatus OWNER TO dbuser;

--
-- Name: adassettype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.adassettype AS ENUM (
    'undefined',
    'image',
    'video',
    'html5'
);


ALTER TYPE public.adassettype OWNER TO dbuser;

--
-- Name: applicationtype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.applicationtype AS ENUM (
    'site',
    'application',
    'game'
);


ALTER TYPE public.applicationtype OWNER TO dbuser;

--
-- Name: approvestatus; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.approvestatus AS ENUM (
    'pending',
    'approved',
    'rejected'
);


ALTER TYPE public.approvestatus OWNER TO dbuser;

--
-- Name: auctiontype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.auctiontype AS ENUM (
    'undefined',
    'first_price',
    'second_price'
);


ALTER TYPE public.auctiontype OWNER TO dbuser;

--
-- Name: formattype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.formattype AS ENUM (
    'undefined',
    'direct',
    'proxy',
    'video',
    'banner',
    'html5',
    'native',
    'custom'
);


ALTER TYPE public.formattype OWNER TO dbuser;

--
-- Name: platformtype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.platformtype AS ENUM (
    'web',
    'desktop',
    'mobile',
    'smart.phone',
    'tablet',
    'smart.tv',
    'gamestation',
    'smart.watch',
    'vr',
    'smart.glasses',
    'smart.billboard'
);


ALTER TYPE public.platformtype OWNER TO dbuser;

--
-- Name: pricingmodel; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.pricingmodel AS ENUM (
    'undefined',
    'CPM',
    'CPC',
    'CPA'
);


ALTER TYPE public.pricingmodel OWNER TO dbuser;

--
-- Name: privatestatus; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.privatestatus AS ENUM (
    'public',
    'private'
);


ALTER TYPE public.privatestatus OWNER TO dbuser;

--
-- Name: processingstatus; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.processingstatus AS ENUM (
    'undefined',
    'progress',
    'processed',
    'error',
    'deleted'
);


ALTER TYPE public.processingstatus OWNER TO dbuser;

--
-- Name: rtbrequesttype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.rtbrequesttype AS ENUM (
    'undefined',
    'json',
    'xml',
    'protobuff',
    'postformencoded',
    'plain'
);


ALTER TYPE public.rtbrequesttype OWNER TO dbuser;

--
-- Name: transaction_status; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.transaction_status AS ENUM (
    'created',
    'executed',
    'execution_error',
    'cancelled'
);


ALTER TYPE public.transaction_status OWNER TO dbuser;

--
-- Name: transaction_type; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.transaction_type AS ENUM (
    'top_up',
    'withdrawal',
    'deduction',
    'credit',
    'pay_off_credit'
);


ALTER TYPE public.transaction_type OWNER TO dbuser;

--
-- Name: zonetype; Type: TYPE; Schema: public; Owner: dbuser
--

CREATE TYPE public.zonetype AS ENUM (
    'zone',
    'smartlink'
);


ALTER TYPE public.zonetype OWNER TO dbuser;

--
-- Name: keep_for(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.keep_for() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    -- param[0] = column name
    -- param[1] = interval
    BEGIN
        IF TG_WHEN <> 'BEFORE' THEN
            RAISE EXCEPTION 'keep_for() may only run as a BEFORE trigger';
        END IF ;

        IF TG_ARGV[0]::TEXT IS NULL THEN
            RAISE EXCEPTION 'keep_for() must be installed with a column name to check';
        END IF;

        IF TG_ARGV[1]::INTERVAL IS NULL THEN
            RAISE EXCEPTION 'keep_for() must be installed with an INTERVAL to keep data for';
        END IF;

        IF TG_OP = 'INSERT' THEN

            EXECUTE 'DELETE FROM ' || quote_ident(TG_TABLE_NAME) ||
                ' WHERE ' || TG_ARGV[0]::TEXT ||
                    ' < now() - INTERVAL ' || quote_literal(TG_ARGV[1]::TEXT) || ';';

        END IF;

        RETURN NEW;
    END;
$$;


ALTER FUNCTION public.keep_for() OWNER TO dbuser;

--
-- Name: notify_update_event(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.notify_update_event() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    DECLARE
        data json;
        notification json;
    BEGIN

        -- Convert the old or new row to JSON, based on the kind of action.
        -- Action = DELETE?             -> OLD row
        -- Action = INSERT or UPDATE?   -> NEW row
        IF (TG_OP = 'DELETE') THEN
            data = row_to_json(OLD);
        ELSE
            data = row_to_json(NEW);
        END IF;

        -- Contruct the notification as a JSON string.
        notification = json_build_object(
                          'schema', TG_TABLE_SCHEMA,
                          'table',  TG_TABLE_NAME,
                          'action', TG_OP,
                          'data',   data);


        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('update_events', notification::text);

        -- Result is ignored since this is an AFTER trigger
        RETURN NULL;
    END;
$$;


ALTER FUNCTION public.notify_update_event() OWNER TO dbuser;

--
-- Name: updated_at_column(); Type: FUNCTION; Schema: public; Owner: dbuser
--

CREATE FUNCTION public.updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
   NEW.updated_at = now(); 
   RETURN NEW;
END;
$$;


ALTER FUNCTION public.updated_at_column() OWNER TO dbuser;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: account_base; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_base (
    id bigint NOT NULL,
    approve_status integer DEFAULT 0 NOT NULL,
    title character varying(128) NOT NULL,
    description text,
    logo_uri character varying(1024) DEFAULT ''::character varying NOT NULL,
    policy_uri character varying(1024) DEFAULT ''::character varying NOT NULL,
    tos_uri character varying(1024) DEFAULT ''::character varying NOT NULL,
    client_uri character varying(1024) DEFAULT ''::character varying NOT NULL,
    contacts text[],
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT account_base_title_check CHECK (((title)::text ~* '^[^\s]+'::text))
);


ALTER TABLE public.account_base OWNER TO dbuser;

--
-- Name: account_base_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.account_base_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.account_base_id_seq OWNER TO dbuser;

--
-- Name: account_base_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.account_base_id_seq OWNED BY public.account_base.id;


--
-- Name: account_member; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_member (
    id bigint NOT NULL,
    approve_status integer DEFAULT 0 NOT NULL,
    account_id bigint NOT NULL,
    user_id bigint NOT NULL,
    is_admin boolean DEFAULT false NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone
);


ALTER TABLE public.account_member OWNER TO dbuser;

--
-- Name: account_member_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.account_member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.account_member_id_seq OWNER TO dbuser;

--
-- Name: account_member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.account_member_id_seq OWNED BY public.account_member.id;


--
-- Name: account_social; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_social (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    social_id character varying(128) NOT NULL,
    provider character varying(128) NOT NULL,
    email character varying(128) NOT NULL,
    first_name character varying(128) NOT NULL,
    last_name character varying(128) NOT NULL,
    username character varying(128) NOT NULL,
    avatar character varying(512) NOT NULL,
    link character varying(1024) NOT NULL,
    data jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.account_social OWNER TO dbuser;

--
-- Name: account_social_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.account_social_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.account_social_id_seq OWNER TO dbuser;

--
-- Name: account_social_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.account_social_id_seq OWNED BY public.account_social.id;


--
-- Name: account_social_session; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_social_session (
    account_social_id bigint NOT NULL,
    name character varying(128) DEFAULT 'default'::character varying NOT NULL,
    token_type character varying(128) NOT NULL,
    access_token character varying(512) NOT NULL,
    refresh_token character varying(512) NOT NULL,
    scopes text[] NOT NULL,
    expires_at timestamp without time zone,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.account_social_session OWNER TO dbuser;

--
-- Name: account_user; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_user (
    id bigint NOT NULL,
    approve_status integer DEFAULT 0 NOT NULL,
    email character varying(128) NOT NULL,
    password character varying(128) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    CONSTRAINT account_user_email_check CHECK (((email)::text ~* '^[^\s]+$'::text)),
    CONSTRAINT account_user_password_check CHECK (((length((password)::text) = 0) OR (length((password)::text) > 5)))
);


ALTER TABLE public.account_user OWNER TO dbuser;

--
-- Name: account_user_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.account_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.account_user_id_seq OWNER TO dbuser;

--
-- Name: account_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.account_user_id_seq OWNED BY public.account_user.id;


--
-- Name: account_user_password_reset; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.account_user_password_reset (
    user_id bigint NOT NULL,
    token character varying(128) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    expires_at timestamp without time zone DEFAULT (now() + '00:30:00'::interval) NOT NULL,
    CONSTRAINT account_user_password_reset_token_check CHECK (((token)::text ~* '^[^\s]+$'::text))
);


ALTER TABLE public.account_user_password_reset OWNER TO dbuser;

--
-- Name: adv_campaign; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_campaign (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    company_id bigint NOT NULL,
    creator_id bigint NOT NULL,
    moderator_id bigint NOT NULL,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    private public.privatestatus DEFAULT 'public'::public.privatestatus NOT NULL,
    daily_budget double precision DEFAULT 0 NOT NULL,
    budget double precision DEFAULT 0 NOT NULL,
    daily_test_budget double precision DEFAULT 0 NOT NULL,
    test_budget double precision DEFAULT 0 NOT NULL,
    context jsonb,
    zones integer[],
    domains text[],
    categories integer[],
    geos text[],
    languages text[],
    browsers integer[],
    os integer[],
    device_types integer[],
    devices integer[],
    sex integer[],
    age integer[],
    date_start timestamp without time zone,
    date_end timestamp without time zone,
    hours character varying(168),
    trace text[],
    trace_percent integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.adv_campaign OWNER TO dbuser;

--
-- Name: ad_campaign_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.ad_campaign_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ad_campaign_id_seq OWNER TO dbuser;

--
-- Name: ad_campaign_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.ad_campaign_id_seq OWNED BY public.adv_campaign.id;


--
-- Name: ad_format; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.ad_format (
    id bigint NOT NULL,
    codename character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    active public.activestatus NOT NULL,
    width integer,
    height integer,
    min_width integer,
    min_height integer,
    config jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.ad_format OWNER TO dbuser;

--
-- Name: ad_format_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.ad_format_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ad_format_id_seq OWNER TO dbuser;

--
-- Name: ad_format_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.ad_format_id_seq OWNED BY public.ad_format.id;


--
-- Name: adv_item; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_item (
    id bigint NOT NULL,
    campaign_id bigint NOT NULL,
    bids jsonb,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    format_id bigint NOT NULL,
    min_width integer DEFAULT 0 NOT NULL,
    min_height integer DEFAULT 0 NOT NULL,
    max_width integer DEFAULT 0 NOT NULL,
    max_height integer DEFAULT 0 NOT NULL,
    pricing_model public.pricingmodel DEFAULT 'undefined'::public.pricingmodel NOT NULL,
    bid_price double precision DEFAULT 0 NOT NULL,
    price double precision DEFAULT 0 NOT NULL,
    lead_price double precision DEFAULT 0 NOT NULL,
    daily_budget double precision DEFAULT 0 NOT NULL,
    budget double precision DEFAULT 0 NOT NULL,
    daily_test_budget double precision DEFAULT 0 NOT NULL,
    test_budget double precision DEFAULT 0 NOT NULL,
    context jsonb,
    weight integer DEFAULT 0 NOT NULL,
    frequency_capping integer DEFAULT 0 NOT NULL,
    hours character varying(168),
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.adv_item OWNER TO dbuser;

--
-- Name: ad_item_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.ad_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ad_item_id_seq OWNER TO dbuser;

--
-- Name: ad_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.ad_item_id_seq OWNED BY public.adv_item.id;


--
-- Name: adv_link; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_link (
    id bigint NOT NULL,
    link text NOT NULL,
    campaign_id bigint NOT NULL,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.adv_link OWNER TO dbuser;

--
-- Name: ad_link_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.ad_link_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.ad_link_id_seq OWNER TO dbuser;

--
-- Name: ad_link_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.ad_link_id_seq OWNED BY public.adv_link.id;


--
-- Name: adv_application; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_application (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    creator_id bigint NOT NULL,
    uri character varying(255) NOT NULL,
    type public.applicationtype NOT NULL,
    platform public.platformtype NOT NULL,
    premium boolean DEFAULT false NOT NULL,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    private public.privatestatus DEFAULT 'public'::public.privatestatus NOT NULL,
    categories bigint[],
    revenue_share numeric(3,10) DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    title text NOT NULL,
    description text,
    allowed_sources bigint[],
    disallowed_sources bigint[]
);


ALTER TABLE public.adv_application OWNER TO dbuser;

--
-- Name: adv_application_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.adv_application_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.adv_application_id_seq OWNER TO dbuser;

--
-- Name: adv_application_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.adv_application_id_seq OWNED BY public.adv_application.id;


--
-- Name: adv_category; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_category (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    iab_code character varying(255) NOT NULL,
    parent_id bigint,
    "position" bigint NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.adv_category OWNER TO dbuser;

--
-- Name: adv_category_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.adv_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.adv_category_id_seq OWNER TO dbuser;

--
-- Name: adv_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.adv_category_id_seq OWNED BY public.adv_category.id;


--
-- Name: adv_format; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_format (
    id bigint NOT NULL,
    codename character varying(255) NOT NULL,
    type character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    active public.activestatus NOT NULL,
    width integer,
    height integer,
    min_width integer,
    min_height integer,
    config jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.adv_format OWNER TO dbuser;

--
-- Name: adv_format_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.adv_format_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.adv_format_id_seq OWNER TO dbuser;

--
-- Name: adv_format_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.adv_format_id_seq OWNED BY public.adv_format.id;


--
-- Name: adv_zone; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.adv_zone (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    account_id bigint NOT NULL,
    type public.zonetype NOT NULL,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    default_code jsonb,
    context jsonb,
    min_ecpm numeric(10,5) DEFAULT 0 NOT NULL,
    min_ecpm_by_geo jsonb,
    fixed_purchase_price numeric(10,5) DEFAULT 0 NOT NULL,
    allowed_formats bigint[],
    allowed_types bigint[],
    allowed_sources bigint[],
    disallowed_sources bigint[],
    campaigns bigint[],
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    codename character varying(64) NOT NULL,
    description text
);


ALTER TABLE public.adv_zone OWNER TO dbuser;

--
-- Name: adv_zone_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.adv_zone_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.adv_zone_id_seq OWNER TO dbuser;

--
-- Name: adv_zone_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.adv_zone_id_seq OWNED BY public.adv_zone.id;


--
-- Name: auth_client; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.auth_client (
    id character varying(128) NOT NULL,
    account_id bigint NOT NULL,
    user_id bigint NOT NULL,
    title text NOT NULL,
    secret character varying(256) NOT NULL,
    redirect_uris text[],
    grant_types text[],
    response_types text[],
    scope text,
    audience text[],
    subject_type character varying(128),
    allowed_cors_origins text[],
    public boolean DEFAULT false NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.auth_client OWNER TO dbuser;

--
-- Name: auth_session; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.auth_session (
    id bigint NOT NULL,
    active boolean DEFAULT true NOT NULL,
    client_id character varying(128) NOT NULL,
    username character varying(128) DEFAULT ''::character varying NOT NULL,
    subject character varying(128) DEFAULT ''::character varying NOT NULL,
    request_id character varying(256) NOT NULL,
    access_token character varying(256) NOT NULL,
    access_token_expires_at timestamp without time zone NOT NULL,
    refresh_token character varying(256),
    refresh_token_expires_at timestamp without time zone NOT NULL,
    form text DEFAULT ''::text NOT NULL,
    requested_scope text[],
    granted_scope text[],
    requested_audience text[],
    granted_audience text[],
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.auth_session OWNER TO dbuser;

--
-- Name: auth_session_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.auth_session_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.auth_session_id_seq OWNER TO dbuser;

--
-- Name: auth_session_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.auth_session_id_seq OWNED BY public.auth_session.id;


--
-- Name: direct_access_tokens; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.direct_access_tokens (
    id bigint NOT NULL,
    token character varying(256) NOT NULL,
    description text NOT NULL,
    account_id bigint NOT NULL,
    user_id bigint,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    expires_at timestamp without time zone
);


ALTER TABLE public.direct_access_tokens OWNER TO dbuser;

--
-- Name: direct_access_tokens_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.direct_access_tokens_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.direct_access_tokens_id_seq OWNER TO dbuser;

--
-- Name: direct_access_tokens_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.direct_access_tokens_id_seq OWNED BY public.direct_access_tokens.id;


--
-- Name: history_actions; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.history_actions (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    request_id character varying(255) NOT NULL,
    user_id bigint NOT NULL,
    account_id bigint NOT NULL,
    name character varying(255) NOT NULL,
    message text NOT NULL,
    object_type character varying(255) NOT NULL,
    object_id bigint NOT NULL,
    object_ids character varying(255) NOT NULL,
    data jsonb NOT NULL,
    action_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.history_actions OWNER TO dbuser;

--
-- Name: m2m_account_member_role; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.m2m_account_member_role (
    member_id bigint NOT NULL,
    role_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.m2m_account_member_role OWNER TO dbuser;

--
-- Name: m2m_rbac_role; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.m2m_rbac_role (
    parent_role_id bigint NOT NULL,
    child_role_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE public.m2m_rbac_role OWNER TO dbuser;

--
-- Name: option; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.option (
    name character varying(256) NOT NULL,
    type character varying(32) NOT NULL,
    target_id bigint NOT NULL,
    value jsonb NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.option OWNER TO dbuser;

--
-- Name: rbac_role; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.rbac_role (
    id bigint NOT NULL,
    name character varying(256) NOT NULL,
    title character varying(256) NOT NULL,
    description text DEFAULT ''::text NOT NULL,
    context jsonb,
    permissions text[],
    access_level integer DEFAULT 0 NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    CONSTRAINT rbac_role_name_check CHECK (((name)::text ~* '^[\w\d@_:\.-]+$'::text))
);


ALTER TABLE public.rbac_role OWNER TO dbuser;

--
-- Name: rbac_role_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.rbac_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rbac_role_id_seq OWNER TO dbuser;

--
-- Name: rbac_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.rbac_role_id_seq OWNED BY public.rbac_role.id;


--
-- Name: rtb_access_point; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.rtb_access_point (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    codename character varying(255) NOT NULL,
    title character varying(255) NOT NULL,
    description text NOT NULL,
    revenue_share_reduce numeric(3,10) DEFAULT 0 NOT NULL,
    auction_type public.auctiontype DEFAULT 'undefined'::public.auctiontype NOT NULL,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    flags jsonb,
    protocol character varying(255) NOT NULL,
    timeout integer DEFAULT 0 NOT NULL,
    rps integer DEFAULT 0 NOT NULL,
    domain_default text,
    request_type public.rtbrequesttype DEFAULT 'undefined'::public.rtbrequesttype NOT NULL,
    headers jsonb,
    max_bid numeric(10,5) DEFAULT 0 NOT NULL,
    fixed_purchase_price numeric(10,5) DEFAULT 0 NOT NULL,
    formats text[],
    device_types bigint[],
    devices bigint[],
    os bigint[],
    browsers bigint[],
    categories bigint[],
    carriers bigint[],
    countries text[],
    languages text[],
    apps bigint[],
    zones bigint[],
    domains text[],
    rtb_sources bigint[],
    secure integer DEFAULT 0 NOT NULL,
    adblock integer DEFAULT 0 NOT NULL,
    private_browsing integer DEFAULT 0 NOT NULL,
    ip integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.rtb_access_point OWNER TO dbuser;

--
-- Name: TABLE rtb_access_point; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON TABLE public.rtb_access_point IS 'RTB Access Point';


--
-- Name: COLUMN rtb_access_point.secure; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_access_point.secure IS 'Secure flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_access_point.adblock; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_access_point.adblock IS 'Adblock flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_access_point.private_browsing; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_access_point.private_browsing IS 'Private browsing flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_access_point.ip; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_access_point.ip IS 'IP flag (0 - any, 1 - IPv4, 2 - IPv6)';


--
-- Name: rtb_access_point_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.rtb_access_point_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rtb_access_point_id_seq OWNER TO dbuser;

--
-- Name: rtb_access_point_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.rtb_access_point_id_seq OWNED BY public.rtb_access_point.id;


--
-- Name: rtb_source; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.rtb_source (
    id bigint NOT NULL,
    account_id bigint NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    status public.approvestatus DEFAULT 'pending'::public.approvestatus NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    flags jsonb,
    protocol character varying(255) NOT NULL,
    minimal_weight numeric NOT NULL,
    url text NOT NULL,
    method character varying(10) DEFAULT 'POST'::character varying NOT NULL,
    request_type public.rtbrequesttype DEFAULT 'undefined'::public.rtbrequesttype NOT NULL,
    headers jsonb,
    rps integer DEFAULT 0 NOT NULL,
    timeout integer NOT NULL,
    accuracy numeric,
    price_correction_reduce numeric,
    auction_type public.auctiontype DEFAULT 'undefined'::public.auctiontype NOT NULL,
    min_bid numeric,
    max_bid numeric,
    formats text[],
    device_types bigint[],
    devices bigint[],
    os bigint[],
    browsers bigint[],
    carriers bigint[],
    categories bigint[],
    countries text[],
    languages text[],
    apps bigint[],
    domains text[],
    zones bigint[],
    external_zones bigint[],
    secure integer DEFAULT 0 NOT NULL,
    adblock integer DEFAULT 0 NOT NULL,
    private_browsing integer DEFAULT 0 NOT NULL,
    ip integer DEFAULT 0 NOT NULL,
    config jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.rtb_source OWNER TO dbuser;

--
-- Name: TABLE rtb_source; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON TABLE public.rtb_source IS 'RTB Source';


--
-- Name: COLUMN rtb_source.secure; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_source.secure IS 'Secure flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_source.adblock; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_source.adblock IS 'Adblock flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_source.private_browsing; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_source.private_browsing IS 'Private browsing flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: COLUMN rtb_source.ip; Type: COMMENT; Schema: public; Owner: dbuser
--

COMMENT ON COLUMN public.rtb_source.ip IS 'IP flag (0 - any, 1 - only, 2 - exclude)';


--
-- Name: rtb_source_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.rtb_source_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.rtb_source_id_seq OWNER TO dbuser;

--
-- Name: rtb_source_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.rtb_source_id_seq OWNED BY public.rtb_source.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO dbuser;

--
-- Name: schema_migrations_dev; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.schema_migrations_dev (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations_dev OWNER TO dbuser;

--
-- Name: schema_migrations_initial; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.schema_migrations_initial (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations_initial OWNER TO dbuser;

--
-- Name: schema_migrations_initial2; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.schema_migrations_initial2 (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations_initial2 OWNER TO dbuser;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    key character varying(128) NOT NULL,
    type public.transaction_type NOT NULL,
    amount numeric NOT NULL,
    status public.transaction_status NOT NULL,
    gateway_id text NOT NULL,
    gateway_payment_id text NOT NULL,
    gateway_info text NOT NULL,
    message text NOT NULL,
    invoice_id integer NOT NULL,
    invoice_key text NOT NULL,
    invoice_number text NOT NULL
);


ALTER TABLE public.transactions OWNER TO dbuser;

--
-- Name: type_browser; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.type_browser (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    versions jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    description text NOT NULL
);


ALTER TABLE public.type_browser OWNER TO dbuser;

--
-- Name: type_browser_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.type_browser_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.type_browser_id_seq OWNER TO dbuser;

--
-- Name: type_browser_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.type_browser_id_seq OWNED BY public.type_browser.id;


--
-- Name: type_device_maker; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.type_device_maker (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    match_exp text DEFAULT ''::text NOT NULL
);


ALTER TABLE public.type_device_maker OWNER TO dbuser;

--
-- Name: type_device_maker_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.type_device_maker_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.type_device_maker_id_seq OWNER TO dbuser;

--
-- Name: type_device_maker_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.type_device_maker_id_seq OWNED BY public.type_device_maker.id;


--
-- Name: type_device_model; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.type_device_model (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    maker_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone,
    type_id bigint DEFAULT 0 NOT NULL,
    versions jsonb
);


ALTER TABLE public.type_device_model OWNER TO dbuser;

--
-- Name: type_device_model_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.type_device_model_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.type_device_model_id_seq OWNER TO dbuser;

--
-- Name: type_device_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.type_device_model_id_seq OWNED BY public.type_device_model.id;


--
-- Name: type_os; Type: TABLE; Schema: public; Owner: dbuser
--

CREATE TABLE public.type_os (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    description text NOT NULL,
    match_exp text NOT NULL,
    active public.activestatus DEFAULT 'pause'::public.activestatus NOT NULL,
    versions jsonb,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    deleted_at timestamp without time zone
);


ALTER TABLE public.type_os OWNER TO dbuser;

--
-- Name: type_os_id_seq; Type: SEQUENCE; Schema: public; Owner: dbuser
--

CREATE SEQUENCE public.type_os_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.type_os_id_seq OWNER TO dbuser;

--
-- Name: type_os_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: dbuser
--

ALTER SEQUENCE public.type_os_id_seq OWNED BY public.type_os.id;


--
-- Name: account_base id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_base ALTER COLUMN id SET DEFAULT nextval('public.account_base_id_seq'::regclass);


--
-- Name: account_member id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_member ALTER COLUMN id SET DEFAULT nextval('public.account_member_id_seq'::regclass);


--
-- Name: account_social id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_social ALTER COLUMN id SET DEFAULT nextval('public.account_social_id_seq'::regclass);


--
-- Name: account_user id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_user ALTER COLUMN id SET DEFAULT nextval('public.account_user_id_seq'::regclass);


--
-- Name: ad_format id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.ad_format ALTER COLUMN id SET DEFAULT nextval('public.ad_format_id_seq'::regclass);


--
-- Name: adv_application id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_application ALTER COLUMN id SET DEFAULT nextval('public.adv_application_id_seq'::regclass);


--
-- Name: adv_campaign id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_campaign ALTER COLUMN id SET DEFAULT nextval('public.ad_campaign_id_seq'::regclass);


--
-- Name: adv_category id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_category ALTER COLUMN id SET DEFAULT nextval('public.adv_category_id_seq'::regclass);


--
-- Name: adv_format id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_format ALTER COLUMN id SET DEFAULT nextval('public.adv_format_id_seq'::regclass);


--
-- Name: adv_item id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_item ALTER COLUMN id SET DEFAULT nextval('public.ad_item_id_seq'::regclass);


--
-- Name: adv_link id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_link ALTER COLUMN id SET DEFAULT nextval('public.ad_link_id_seq'::regclass);


--
-- Name: adv_zone id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_zone ALTER COLUMN id SET DEFAULT nextval('public.adv_zone_id_seq'::regclass);


--
-- Name: auth_session id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_session ALTER COLUMN id SET DEFAULT nextval('public.auth_session_id_seq'::regclass);


--
-- Name: direct_access_tokens id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.direct_access_tokens ALTER COLUMN id SET DEFAULT nextval('public.direct_access_tokens_id_seq'::regclass);


--
-- Name: rbac_role id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rbac_role ALTER COLUMN id SET DEFAULT nextval('public.rbac_role_id_seq'::regclass);


--
-- Name: rtb_access_point id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_access_point ALTER COLUMN id SET DEFAULT nextval('public.rtb_access_point_id_seq'::regclass);


--
-- Name: rtb_source id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_source ALTER COLUMN id SET DEFAULT nextval('public.rtb_source_id_seq'::regclass);


--
-- Name: type_browser id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_browser ALTER COLUMN id SET DEFAULT nextval('public.type_browser_id_seq'::regclass);


--
-- Name: type_device_maker id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_device_maker ALTER COLUMN id SET DEFAULT nextval('public.type_device_maker_id_seq'::regclass);


--
-- Name: type_device_model id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_device_model ALTER COLUMN id SET DEFAULT nextval('public.type_device_model_id_seq'::regclass);


--
-- Name: type_os id; Type: DEFAULT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_os ALTER COLUMN id SET DEFAULT nextval('public.type_os_id_seq'::regclass);


--
-- Data for Name: account_base; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_base (id, approve_status, title, description, logo_uri, policy_uri, tos_uri, client_uri, contacts, created_at, updated_at, deleted_at) FROM stdin;
2	1	Test						{}	2024-08-06 20:10:48.656557	2024-08-06 20:11:21.430865	\N
1	1	system	System account	https://www.pcgameshardware.de/screenshots/1280x1024/2023/07/Mini-Operating-System-9-mit-Mini-Navigation-inklusive-Ladeplanung-pcgh.jpg				{}	2024-07-20 17:56:42.281845	2024-08-24 15:54:17.920972	\N
\.


--
-- Data for Name: account_member; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_member (id, approve_status, account_id, user_id, is_admin, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	1	t	2024-07-20 17:56:42.281845+00	2024-08-06 20:10:12.447916+00	\N
2	1	2	1	t	2024-08-06 20:10:48.657432+00	2024-09-07 13:16:53.604524+00	\N
\.


--
-- Data for Name: account_social; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_social (id, user_id, social_id, provider, email, first_name, last_name, username, avatar, link, data, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: account_social_session; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_social_session (account_social_id, name, token_type, access_token, refresh_token, scopes, expires_at, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: account_user; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_user (id, approve_status, email, password, created_at, updated_at, deleted_at) FROM stdin;
1	1	super@project.com	$2a$10$1Eh45WHxlcO4Tb90mVc62Og.T8Q81QJTpGyF9pT1EwSkQUPA8XyAS	2024-07-20 17:56:42.281845	2024-08-06 20:10:16.338053	\N
\.


--
-- Data for Name: account_user_password_reset; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.account_user_password_reset (user_id, token, created_at, expires_at) FROM stdin;
\.


--
-- Data for Name: ad_format; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.ad_format (id, codename, type, title, active, width, height, min_width, min_height, config, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: adv_application; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_application (id, account_id, creator_id, uri, type, platform, premium, status, active, private, categories, revenue_share, created_at, updated_at, deleted_at, title, description, allowed_sources, disallowed_sources) FROM stdin;
4	2	1	test.com	site	web	f	approved	active	public	{5}	0.0000000000	2024-10-15 16:51:18.854	2024-10-20 13:20:54.696506	\N	test1		\N	\N
3	2	1	test.com	site	web	f	pending	pause	public	{}	0.0000000000	2024-10-14 16:23:07.93	2024-10-20 13:50:10.506574	\N	Test		\N	\N
\.


--
-- Data for Name: adv_campaign; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_campaign (id, title, company_id, creator_id, moderator_id, status, active, private, daily_budget, budget, daily_test_budget, test_budget, context, zones, domains, categories, geos, languages, browsers, os, device_types, devices, sex, age, date_start, date_end, hours, trace, trace_percent, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: adv_category; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_category (id, name, description, iab_code, parent_id, "position", active, created_at, updated_at, deleted_at) FROM stdin;
4	Cars			\N	0	active	2024-08-06 15:29:11.569444	2024-08-06 19:43:04.601636	\N
5	Toyota	New description	AIB-12	4	0	active	2024-08-06 15:49:29.174259	2024-08-06 19:52:17.101354	\N
6	New		AIB-100!	5	0	active	2024-08-06 20:08:45.962658	2024-08-06 20:09:05.058746	\N
\.


--
-- Data for Name: adv_format; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_format (id, codename, type, title, description, active, width, height, min_width, min_height, config, created_at, updated_at, deleted_at) FROM stdin;
1	direct	direct	Direct (Popunders/Popups)	New description h	active	0	0	0	0	{"fields": [{"name": "fff", "title": "d"}]}	2024-08-21 16:24:54.365901	2024-08-22 20:59:14.618591	\N
2	native	native	Native	Native ad format	active	0	0	0	0	{"assets": [{"id": 1, "name": "main", "width": 1500, "height": 1500, "thumbs": ["250x", "350x", "500x"], "required": true, "min_width": 50, "min_height": 50, "adjust_size": true, "allowed_types": ["image", "video"]}, {"id": 2, "name": "logo", "width": 100, "height": 100, "min_width": 50, "min_height": 50, "allowed_types": ["image"]}], "fields": [{"id": 101, "max": 40, "min": 5, "name": "title", "type": "string", "title": "Title", "required": true}, {"id": 102, "max": 80, "min": 5, "name": "description", "type": "string", "title": "Description", "required": true}, {"id": 103, "max": 30, "name": "brandname", "type": "string", "title": "Brandname"}, {"id": 104, "name": "phone", "type": "phone", "title": "Phone"}, {"id": 105, "name": "url", "type": "url", "title": "Promotion URL"}]}	2024-08-24 11:22:23.650644	2024-08-24 11:26:55.736068	\N
\.


--
-- Data for Name: adv_item; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_item (id, campaign_id, bids, status, active, format_id, min_width, min_height, max_width, max_height, pricing_model, bid_price, price, lead_price, daily_budget, budget, daily_test_budget, test_budget, context, weight, frequency_capping, hours, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: adv_link; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_link (id, link, campaign_id, status, active, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: adv_zone; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.adv_zone (id, title, account_id, type, status, active, default_code, context, min_ecpm, min_ecpm_by_geo, fixed_purchase_price, allowed_formats, allowed_types, allowed_sources, disallowed_sources, campaigns, created_at, updated_at, deleted_at, codename, description) FROM stdin;
2	top	2	zone	approved	active	null	null	0.00000	null	0.00000	\N	\N	\N	\N	\N	2024-10-21 10:40:31.286395	2024-10-21 11:02:30.484307	\N	2nkBXMzDH71gxXKlBxH9GiG8eGc	
\.


--
-- Data for Name: auth_client; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.auth_client (id, account_id, user_id, title, secret, redirect_uris, grant_types, response_types, scope, audience, subject_type, allowed_cors_origins, public, expires_at, created_at, updated_at, deleted_at) FROM stdin;
system	1	1	system gate	$2a$10$IxMdI6d.LIRZPpSfEwNoeu4rY3FhDREsxFJXikcgdRRAStxUlsuEO	{http://localhost:3846/callback}	{password,client_credentials,authorization_code,implicit,refresh_token}	{id_token,code,token}	superuser menu openid offline	{}	public	{http://localhost:3846/users}	f	2034-07-20 17:56:42.281845	2024-07-20 17:56:42.281845	2024-07-20 17:56:42.281845	\N
\.


--
-- Data for Name: auth_session; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.auth_session (id, active, client_id, username, subject, request_id, access_token, access_token_expires_at, refresh_token, refresh_token_expires_at, form, requested_scope, granted_scope, requested_audience, granted_audience, created_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: direct_access_tokens; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.direct_access_tokens (id, token, description, account_id, user_id, created_at, expires_at) FROM stdin;
2	2u0N0^1{0x2w2&0D1J0S050|2f0!1>2v2R1z0W2r05061v291m0%1l1T2U0k0W1N		1	1	2024-07-25 09:50:45.771072	2024-07-25 08:50:51.475909
1	2X2J2M110p0`2|1M132o0s1b1v0|0s0C1!2n1g1u0(2^2U01292b2a1j2i2<1Q01		1	1	2024-07-25 09:48:53.50894	2024-07-25 08:50:52.99145
3	0d1>0<2#0Q190R242Q2T0%2`1t2)2t2z0c1V2f2L1*162e1D0G212*1w0D1s2I2!		1	1	2024-08-06 20:10:00.408845	2024-08-07 10:22:22.100827
\.


--
-- Data for Name: history_actions; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.history_actions (id, request_id, user_id, account_id, name, message, object_type, object_id, object_ids, data, action_at) FROM stdin;
ca1dd1d9-6735-474e-b0ec-bee3d2bbc8e4	e4cfa534-cf28-4bb7-bbe6-78da3ac5536b	1	1	update		Format	1	1	{"id": 1, "type": "direct", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:09:40.100432Z", "description": "New description h"}	2024-08-22 20:10:17.986681
d05d6630-a22d-42a0-a83b-67f16cd185a5	379478cb-0902-4943-a7d0-4d31afe1c9ca	1	1	update		Format	1	1	{"id": 1, "type": "direct", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:10:17.989401Z", "description": "New description h"}	2024-08-22 20:13:45.016864
370df6f5-cdac-41f8-8f8d-f95059142139	577c9270-5883-4055-b7a1-d6b3144d8756	1	1	update		Format	1	1	{"id": 1, "type": "direct", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:13:45.018964Z", "description": "New description h"}	2024-08-22 20:36:24.062054
18efe580-7c69-4eca-b5a9-1752a01868e9	a0268711-87ce-43ba-9cb6-879af93e404f	1	1	update		Format	1	1	{"id": 1, "type": "banner", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:36:24.063637Z", "description": "New description h"}	2024-08-22 20:59:00.86778
3c692680-af4b-4af3-b84d-e1cc8dfd893e	5bd83428-6866-4ade-8ff5-456977f65da3	1	1	update		Format	1	1	{"id": 1, "type": "direct", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:59:00.868592Z", "description": "New description h"}	2024-08-22 20:59:07.982252
645433be-bc27-4be7-902f-6133151ac82a	85553693-fb46-4f65-97cd-619f2997c271	1	1	update		Format	1	1	{"id": 1, "type": "direct", "title": "Direct (Popunders/Popups)", "width": 0, "active": "active", "config": {"fields": [{"name": "fff", "title": "d"}]}, "height": 0, "codename": "direct", "min_width": 0, "created_at": "2024-08-21T16:24:54.365901Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-22T20:59:07.983695Z", "description": "New description h"}	2024-08-22 20:59:14.617589
b922e52c-81b8-44fd-988c-b9db4dc2f3ed	4ff451c9-90b9-40c3-ad49-f3dbeb2727bb	1	1	create		Format	0	0	{"id": 0, "type": "native", "title": "Native", "width": 0, "active": "pause", "config": {}, "height": 0, "codename": "native", "min_width": 0, "created_at": "2024-08-24T11:22:23.650644458Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-24T11:22:23.650644458Z", "description": ""}	2024-08-24 11:22:23.650728
3fce8213-6070-4d5f-834e-e8ed2c2a6dd2	287f02d8-735b-4cab-80cb-6f2f2269a419	1	1	update		Format	2	2	{"id": 2, "type": "native", "title": "Native", "width": 0, "active": "pause", "config": {}, "height": 0, "codename": "native", "min_width": 0, "created_at": "2024-08-24T11:22:23.650644Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-24T11:22:23.650644Z", "description": "Native ad format"}	2024-08-24 11:22:53.872767
c1a024af-dc35-4967-aea1-c9f34c259cc3	89d96eb2-b878-4b86-84bf-531d72cda5ee	1	1	update		Format	2	2	{"id": 2, "type": "native", "title": "Native", "width": 0, "active": "active", "config": {}, "height": 0, "codename": "native", "min_width": 0, "created_at": "2024-08-24T11:22:23.650644Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-24T11:22:53.873772Z", "description": "Native ad format"}	2024-08-24 11:26:12.824807
935e33d5-5a1e-451a-9710-ac7899add8a3	156871f8-cb30-4778-97c1-2b9bc20bfe3b	1	1	update		Format	2	2	{"id": 2, "type": "native", "title": "Native", "width": 0, "active": "active", "config": {"assets": [{"id": 1, "name": "main", "width": 1500, "height": 1500, "thumbs": ["250x", "350x", "500x"], "required": true, "min_width": 50, "min_height": 50, "adjust_size": true, "allowed_types": ["image", "video"]}, {"id": 2, "name": "logo", "width": 100, "height": 100, "min_width": 50, "min_height": 50, "allowed_types": ["image"]}], "fields": [{"id": 101, "max": 40, "min": 5, "name": "title", "type": "string", "title": "Title", "required": true}, {"id": 102, "max": 80, "min": 5, "name": "description", "type": "string", "title": "Description", "required": true}, {"id": 103, "max": 30, "name": "brandname", "type": "string", "title": "Brandname"}, {"id": 104, "name": "phone", "type": "phone", "title": "Phone"}, {"id": 105, "name": "url", "type": "url", "title": "Promotion URL"}]}, "height": 0, "codename": "native", "min_width": 0, "created_at": "2024-08-24T11:22:23.650644Z", "deleted_at": null, "min_height": 0, "updated_at": "2024-08-24T11:26:12.826019Z", "description": "Native ad format"}	2024-08-24 11:26:55.734957
35538561-7ce8-45e0-92ab-9c4c305a3ea6	8c7979ec-9c57-43bb-86e5-473e5ea627ff	1	1	create		Browser	0	0	{"id": 0, "name": "Firefox", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:39:40.726346091Z", "deleted_at": null, "updated_at": "2024-08-24T14:39:40.726346091Z", "description": ""}	2024-08-24 14:39:40.726425
4ed1d9be-65df-4ce5-920a-7196dffdf25f	f9f4b5c8-37fd-4e8b-ad32-3552244c79da	1	1	update		Browser	1	1	{"id": 1, "name": "Chrome", "active": "active", "versions": [{"max": "1.999", "min": "0.0.1", "name": "1"}], "match_exp": "", "created_at": "2024-08-04T10:30:59.308409Z", "deleted_at": null, "updated_at": "2024-08-06T19:59:54.894086Z", "description": "Chrome is good"}	2024-08-24 14:40:23.918736
8b606ec1-638d-414f-b657-07197d65eba9	ee9be47e-a9f7-476d-803b-d60aa5bbce52	1	1	update		Browser	1	1	{"id": 1, "name": "Chrome", "active": "active", "versions": [{"max": "1.999", "min": "0.0.1", "name": "1"}], "match_exp": "", "created_at": "2024-08-04T10:30:59.308409Z", "deleted_at": null, "updated_at": "2024-08-24T14:40:23.92008Z", "description": "Chrome is good"}	2024-08-24 14:40:26.699724
754f1a81-c7a9-4a7c-9f84-d7d2af9645e4	3e11e44f-6c30-4609-9db9-0b331d282d75	1	1	create		OS	0	0	{"id": 0, "name": "Windows", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:52:35.273928338Z", "deleted_at": null, "updated_at": "2024-08-24T14:52:35.273928338Z", "description": ""}	2024-08-24 14:52:35.274006
10eabdbc-cb1b-4ef5-986a-3db808eda279	3199469a-fb26-4372-9e19-a5f01bb08446	1	1	create		OS	0	0	{"id": 0, "name": "Windows", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:52:38.008479215Z", "deleted_at": null, "updated_at": "2024-08-24T14:52:38.008479215Z", "description": ""}	2024-08-24 14:52:38.008492
106f1dac-c10b-4e58-934a-884d0ba55c6c	c7b8287e-02f8-4831-a249-3e0fe394a75c	1	1	create		OS	0	0	{"id": 0, "name": "Windows", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:52:49.202593262Z", "deleted_at": null, "updated_at": "2024-08-24T14:52:49.202593262Z", "description": ""}	2024-08-24 14:52:49.202606
60f530fd-cf05-4aac-aa67-23dba7d4182a	e2c1e24a-c73b-4c62-b68d-8b78f872cc8b	1	1	create		OS	0	0	{"id": 0, "name": "", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:53:28.617252377Z", "deleted_at": null, "updated_at": "2024-08-24T14:53:28.617252377Z", "description": ""}	2024-08-24 14:53:28.617266
a4e84b1b-5692-4dec-90bc-bdd60fd52358	1c27a22b-a3f1-40ec-917b-23104f9c9b30	1	1	create		OS	0	0	{"id": 0, "name": "Windows", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:53:34.547354255Z", "deleted_at": null, "updated_at": "2024-08-24T14:53:34.547354255Z", "description": ""}	2024-08-24 14:53:34.547391
d0a820eb-529c-4095-a507-8957dd52327f	ac34f623-7440-4181-ba02-eae3afd121d6	1	1	create		OS	0	0	{"id": 0, "name": "Windows", "active": "pause", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:58:35.108026088Z", "deleted_at": null, "updated_at": "2024-08-24T14:58:35.108026088Z", "description": ""}	2024-08-24 14:58:35.108043
6aa32fe6-900d-4874-a8db-13878b31a2e9	d0f6aa58-a2d4-4db2-bcc8-4a23d99c9cb7	1	1	update		OS	1	1	{"id": 1, "name": "Windows", "active": "active", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:58:35.108026Z", "deleted_at": null, "updated_at": "2024-08-24T14:58:35.108026Z", "description": ""}	2024-08-24 14:59:13.786195
dd6637b9-5f66-47f1-83e3-800337dc70c4	e4b5097d-3fef-45d3-8df8-693a391ef4e5	1	1	update		OS	1	1	{"id": 1, "name": "Windows", "active": "active", "versions": [], "match_exp": "", "created_at": "2024-08-24T14:58:35.108026Z", "deleted_at": null, "updated_at": "2024-08-24T14:59:13.787413Z", "description": "Windows is irritan"}	2024-08-24 15:00:56.204683
15c2a232-c7cb-4d6d-8fc3-135e9c926515	f00c8701-ee6f-4561-90df-407d8fe73cb2	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": null, "rps": 0, "url": "", "apps": null, "flags": null, "title": "title!", "zones": null, "active": "pause", "config": null, "method": "", "secure": 0, "status": "pending", "adblock": 0, "devices": null, "domains": null, "formats": null, "headers": null, "max_bid": 0, "min_bid": 0, "timeout": 0, "accuracy": 0, "browsers": null, "carriers": null, "protocol": "", "countries": null, "languages": null, "account_id": 0, "categories": null, "created_at": "2024-08-24T15:10:10.990512049Z", "deleted_at": null, "updated_at": "2024-08-24T15:10:10.990512049Z", "description": "", "auction_type": "undefined", "device_types": null, "request_type": "undefined", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0}	2024-08-24 15:10:10.990665
4c899e95-d330-4798-acb1-981451bbdfd8	d5592e83-7e7c-4c44-9d68-46e6b43072a9	1	1	update		Account	1	1	{"": null, "id": 1, "title": "system", "tos_uri": "", "contacts": [], "logo_uri": "https://www.pcgameshardware.de/screenshots/1280x1024/2023/07/Mini-Operating-System-9-mit-Mini-Navigation-inklusive-Ladeplanung-pcgh.jpg", "client_uri": "", "created_at": "0001-01-01T00:00:00Z", "deleted_at": null, "policy_uri": "", "updated_at": "0001-01-01T00:00:00Z", "description": "System account", "approve_status": 0}	2024-08-24 15:54:17.919174
f845b828-8b98-48be-9630-dac80b00d65d	b986dce6-03f5-4833-bace-d3af32261f35	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:03:56.335174959Z", "deleted_at": null, "updated_at": "2024-08-24T17:03:56.335174959Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-24 17:03:56.335335
58bb7236-4cbd-4f78-aa06-7b58bb34e326	364e6401-f23a-4185-90db-bfad86dd384c	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.0001, "min_bid": 0.00001, "timeout": 1000, "accuracy": 0.1, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:06:24.711837916Z", "deleted_at": null, "updated_at": "2024-08-24T17:06:24.711837916Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-24 17:06:24.711854
75b0e279-842a-40a3-9fe1-3ac62e850527	2bf30b95-210f-4c67-bfe7-e6291cb8a2fa	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.0000001, "min_bid": 0.00000001, "timeout": 1000, "accuracy": 0.1, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:06:36.58372613Z", "deleted_at": null, "updated_at": "2024-08-24T17:06:36.58372613Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-24 17:06:36.583743
ffd8793e-9a76-4aee-9728-e9886beebbea	35f85790-31ee-462e-ad27-ccd571ef4a0b	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.00000001, "min_bid": 0.00000001, "timeout": 1000, "accuracy": 0.1, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:07:03.622365962Z", "deleted_at": null, "updated_at": "2024-08-24T17:07:03.622365962Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-24 17:07:03.622385
44da96c1-8bed-43af-aa96-16ece7d00c6c	331e9b89-bcbc-43ce-85ce-baf73eb7674f	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.00000001, "min_bid": 0.00000001, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:07:23.790585222Z", "deleted_at": null, "updated_at": "2024-08-24T17:07:23.790585222Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:07:23.7906
90c6d32f-366c-4d41-ac8d-eed386b13447	40f624b8-450f-4f8c-ac2d-b35198c56b33	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.00000001, "min_bid": 0.00000001, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:09:47.30613576Z", "deleted_at": null, "updated_at": "2024-08-24T17:09:47.30613576Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:09:47.306153
9e485684-ceb9-4ffe-8c54-5785a7c1dec9	c4120861-1b37-451b-be6b-77d3e30c50ea	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 0.00000001, "min_bid": 0.00000001, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:14:43.835435425Z", "deleted_at": null, "updated_at": "2024-08-24T17:14:43.835435425Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:14:43.835646
4d37ebe7-5e1b-4769-9bb8-b95e04b4f673	89aa04a6-9913-4536-9b6e-cbfdbfe7e154	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:14:58.54079821Z", "deleted_at": null, "updated_at": "2024-08-24T17:14:58.54079821Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:14:58.540815
08f67c9d-5de9-4d3d-a733-ae909006208e	c14e68a0-83eb-4871-b0ca-ce6d1a2061b0	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:17:18.918727386Z", "deleted_at": null, "updated_at": "2024-08-24T17:17:18.918727386Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:17:18.918747
bfda5c26-0cb2-468e-aed6-f5c26bcd6c13	6ad7ccb5-7bf0-46a1-8f53-5098ad2b19e2	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T14:35:45.486602469Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T14:35:45.486602469Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 14:35:45.486643
70f6d0b7-8324-4a22-ab86-d9a265fc242f	3d812978-5327-470a-9f80-f3456c118dc5	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 0, "categories": [1], "created_at": "2024-08-24T17:18:46.25816101Z", "deleted_at": null, "updated_at": "2024-08-24T17:18:46.25816101Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:18:46.25818
a8efce07-95c9-42aa-aeb5-ad85a6562936	6e4d28f0-ba4f-43ee-b71a-fcc8ecb50206	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": null, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555547Z", "deleted_at": null, "updated_at": "2024-08-24T17:19:05.814555547Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:19:05.814809
32934b97-306c-479e-9b16-6ca26f5a9831	e51dabe7-7f72-4549-81ef-3c9e8461f039	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "qwewqedsa", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:19:05.814555Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:28:09.818883
7bb9a72d-ebab-4336-912a-227f34a3f7e2	c60cdcf9-3265-41c0-956c-4ea52e64f766	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:28:09.821463Z", "description": "", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:28:16.497679
42d86c6a-128b-4284-9c57-2fe9e3fe0b02	c4779f69-9833-483a-ac21-13ea02aaf084	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:28:16.499058Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:30:46.43286
a978795a-45bf-453b-b3f7-38af70f62f15	dbada02c-6781-4adc-bfc6-863de8df817f	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:30:46.434324Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:31:19.399746
be721866-56a7-4365-89a9-20f49d6f32ea	578ad1cc-4322-491c-aba7-076b949f8627	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T15:07:28.959146878Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T15:07:28.959146878Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 15:07:28.959239
8bcb36c5-1d50-4231-8f43-bf045f8bd88a	8b580324-0f4c-433f-a8b2-2feceeaea0be	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["US", "JP", "RU", "O1"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:31:19.401048Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:33:21.359713
3ef8e3bd-9057-4a3a-b6c6-28b2854f7b98	c75c3423-1a2f-42f5-8d81-a5dd0b6ecd7f	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["banner", "native", "video"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["O1", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:33:21.361633Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:35:55.148694
248fd8bd-6c7f-41fb-bf73-a2a4e9e628d1	483e39fc-16a8-4d24-8909-729389300b0d	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "O1", "JP", "RU", "US", "AF", "**", "AL"], "languages": ["en"], "account_id": 1, "categories": [1], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:35:55.150096Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:36:11.266189
777a5dd7-8bbb-448e-8e8f-20e30aae313d	76fdfda6-9178-48c1-af02-e2bf0f02de96	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["**", "A2", "O1", "AF", "AL", "JP", "RU", "US", "AG"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:36:11.267425Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:48:44.805027
fd932cf8-57ad-46cc-882a-d9f52e73f8d7	6fd48e2f-ddda-4f22-857b-ef820448c98b	1	1	update		RTBSource	3	3	{"id": 3, "ip": 1, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 2, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["**", "A2", "O1", "AF", "AL", "JP", "RU", "US", "AG"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:48:44.806471Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 2, "price_correction_reduce": 0.01}	2024-08-24 17:48:55.160769
c1d275ab-e988-4e7a-9552-a160dfaefd0d	4923b5b3-90ce-44bf-a62e-9caf14a3f0d1	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["**", "A2", "O1", "AF", "AL", "JP", "RU", "US", "AG"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:48:55.162165Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:49:05.37583
98bddc73-3317-4177-b212-e1a59419ce2a	743e2564-c4f5-4a3a-8751-c381094e0a74	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T15:34:38.551053466Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T15:34:38.551053466Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 15:34:38.551066
a486de79-90a1-4a10-86a7-15c07dc6efa5	4184710e-860a-466a-976c-221ac3512508	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["**", "A2", "O1", "AF", "AL", "JP", "RU", "US", "AG"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:49:05.376914Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:49:22.403907
438422bf-6701-4e58-a44b-06a26d353332	db92af88-b6a3-4c7f-9da6-f6e042316862	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["O1", "AF", "AG", "AL", "JP", "RU", "US"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:49:22.406004Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:51:16.93905
ca8eea7d-24b6-4dbe-b58e-fcac53ad21bf	986622ee-1824-433b-800d-26736fdf95c3	1	1	update		RTBSource	3	3	{"id": 3, "ip": 1, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 2, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:51:16.940213Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 2, "price_correction_reduce": 0.01}	2024-08-24 17:52:33.646134
9c72558b-f596-4459-bfc2-5bf67c20e961	35b4e337-a69b-4d35-beaf-5046e693bd4a	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1], "domains": [], "formats": ["native", "direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US"], "languages": ["en"], "account_id": 1, "categories": [4, 6], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:52:33.64748Z", "description": "New desc", "auction_type": "first_price", "device_types": [1], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:53:04.053939
69d887c3-6ba3-4801-97f7-205ecb407aec	c0ea3eac-6d01-47ee-ba99-282f9fb1927d	1	1	update		RTBSource	3	3	{"id": 3, "ip": 1, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 2, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:53:04.054963Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 2, "price_correction_reduce": 0.01}	2024-08-24 17:59:24.656037
010ee126-bfa7-449c-b05c-ca0b3d9a14f8	fa762bd9-711e-4d30-a4e1-eabea6f5ca4b	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:59:24.658717Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:59:29.512168
e7e353c1-967a-4ddd-bdb8-3a0321c2ecd3	6ddf0920-df69-4fec-930c-a78fdb1321b6	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:59:29.513529Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 17:59:43.25977
2c716c2a-d717-4245-9b53-3bcd4fbc6eaf	7ecf9f59-8be7-4529-9b4f-f29fad8247de	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T17:59:43.260974Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 18:04:57.514419
85f981af-c906-4c8b-ad9e-31d06cabb66c	44244544-719c-4ac2-8f56-df23aae5efcf	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["AF", "AG", "AL", "JP", "RU", "US", "A2"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T18:04:57.515299Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 18:06:05.318045
c3d93e59-6e53-41bd-b56c-23f8a460a524	4c5183f9-8816-4a1f-808b-f2a050813286	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T18:06:05.319738Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-24 18:07:47.35321
b4ab56ff-92e6-4d09-ad24-93a51dcdfd05	51258bea-ec7d-47c2-9868-c888921e1e61	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-24T18:07:47.355257Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 09:58:20.711147
e264f4ae-598c-4071-a618-d9a032aee7b6	0cfbc611-3701-4bad-a3cd-6013692c819a	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T09:58:20.713411Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 10:01:16.363218
3762d7d1-19f2-4891-ac48-463244393696	b41bbbe9-1eb2-4e18-bb7f-de0813cbb049	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T10:01:16.36465Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 10:02:28.140312
8a37b29d-d441-4ef2-aba2-49b6acbc463c	2998ab09-2205-4333-925c-4e6e6d0946cf	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T10:02:28.141617Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 10:02:35.462442
bd71f574-b982-46b3-a991-d915bc6f2e69	60290266-eea4-400e-b19d-2b63c0e9317d	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 1, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T10:02:35.463523Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 11:01:28.189456
1a45e90a-561b-4865-a602-d64c5f7cc8ab	e1fe649d-f0f6-4a84-b9b2-b52c5f71a46b	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:01:28.190796Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 11:01:37.20956
849189a9-7028-4a51-8482-93762949afb8	131c3ec3-d421-43fc-896d-9dbe3b1edb8b	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI"], "languages": ["en"], "account_id": 1, "categories": [4], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:01:37.210912Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 11:02:13.2514
96335a7e-0e09-4e11-a923-d1b4ceff2888	d1aabb32-c817-459b-b9ac-6ecc3b108e6c	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN"], "languages": ["en"], "account_id": 1, "categories": [4], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:02:13.254717Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 11:03:18.172259
0f0fadee-7dda-4624-89e4-f4d31ca7da15	e3793b26-5b92-45f5-a6d8-bd8bea30f3ed	1	1	update		RTBSource	3	3	{"id": 3, "ip": 2, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 2, "status": "pending", "adblock": 2, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:03:18.173425Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 1, "price_correction_reduce": 0.01}	2024-08-25 11:08:56.18857
8254499b-2933-472e-923b-3f9a5783eb9f	6c3379aa-c79e-4e4b-bff5-2f1d6c4ae490	1	1	update		RTBSource	3	3	{"id": 3, "ip": 2, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 2, "status": "pending", "adblock": 2, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:08:56.190273Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 1, "price_correction_reduce": 0.01}	2024-08-25 11:09:03.557847
0217cce3-4d69-4a70-9efe-4af9ac2db022	81cb8641-947f-45a2-8dde-a12add89e46b	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T11:09:03.559157Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 11:09:17.531087
39172dc5-7934-4570-9cf5-b4f52675c08f	94a80d04-fc6e-4651-bd61-667056920afc	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://asds.sd", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T12:01:24.961058Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 13:18:59.478369
59dda84b-095e-4e75-acae-343cc0d5c858	0b06602c-995e-48a2-b1dd-9b07c80ed86f	1	1	create		RTBAccessPoint	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "sources": null, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 0, "categories": [], "created_at": "2024-08-25T14:30:35.99247918Z", "deleted_at": null, "updated_at": "2024-08-25T14:30:35.99247918Z", "description": "", "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:30:35.992682
e8be3e1b-cfa1-4e5a-b79c-01f9f4621c4d	5735d5bf-01d7-4a8d-a1eb-b77d5c7a3f5d	1	1	create		RTBAccessPoint	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "sources": null, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 0, "categories": [], "created_at": "2024-08-25T14:31:12.839808169Z", "deleted_at": null, "updated_at": "2024-08-25T14:31:12.839808169Z", "description": "", "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:31:12.839824
665e5464-5ff5-4f58-b8fe-59a546ef9ad2	617087a8-641e-47ca-9fb3-3801aacb5244	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T16:13:37.689301049Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T16:13:37.689301049Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:13:37.689316
683e5aef-be79-412d-90e6-7053df9b2f75	5ee33f16-78c1-4686-8c04-92a1d8967964	1	1	create		RTBAccessPoint	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 0, "categories": [], "created_at": "2024-08-25T14:33:35.917091055Z", "deleted_at": null, "updated_at": "2024-08-25T14:33:35.917091055Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:33:35.917314
914a789c-9fd9-4bc4-bff3-d73d30424d37	45a541e0-08b6-42e7-9841-e8f0574b93ae	1	1	create		RTBAccessPoint	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 0, "categories": [], "created_at": "2024-08-25T14:38:26.396386384Z", "deleted_at": null, "updated_at": "2024-08-25T14:38:26.396386384Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:38:26.396611
bc6bc6c8-295d-4828-8a67-bb8755ce60da	e3f5f920-a523-415c-a21a-d2fd13158a4a	1	1	create		RTBAccessPoint	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746004Z", "deleted_at": null, "updated_at": "2024-08-25T14:39:47.012746004Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:39:47.012955
aa6f9911-ae88-4b9e-8e11-caca6e2d15cf	ff170f3b-3184-45e4-9bbf-498b6d48f626	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T14:39:47.012746Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:44:01.023026
580b7bff-02ba-4c8e-9ca2-7adf296ceff2	5325db34-07d8-443d-9dca-594836daefe7	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "test", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T14:44:01.024771Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:47:11.536729
9a6d7783-a6c1-44f3-8fa8-1d686c926e4a	14c1293d-0391-4cdc-a7b0-6b1c566d6cc3	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": ["direct"], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "test", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T14:47:11.539378Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:50:21.3681
45764b4e-41cb-45a9-9842-e42ba438d4e5	cb9d87cf-9e7d-434d-8b11-7a4759952993	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test", "zones": [], "active": "pause", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "test", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T14:50:21.369369Z", "description": "", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 14:50:29.771846
f0e0233c-efdd-4720-8849-1ef67e57aff0	948bdbd1-3844-4072-b252-b3c1b331ce5e	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:18:04.206056505Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T16:18:04.206056505Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:18:04.206162
c4c048be-4448-4310-8b19-ada4043d0f21	aeeab412-f26e-423b-9b03-38ffd512442e	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test", "zones": [], "active": "active", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "test", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T14:54:19.464375Z", "description": "Desc", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 15:17:52.535062
e5866461-e6d0-49a5-afda-8ff866663e7b	5a6cfd45-0901-4ecb-9df0-6b81c76d6fec	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T15:06:15.197379Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 18:57:09.894382
c078796d-f65d-4c6f-a96f-dd3a823638c1	1a9fd935-18b6-4dae-a9d3-d98970b14406	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.001, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T18:57:09.897052Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 18:57:14.03359
9b271bd8-58e9-4b49-beaa-f1985918eba7	3a32202c-612e-4d5f-921f-1bfb70686f9b	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [1], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [1, 3], "domains": [], "formats": ["direct", "native"], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [1], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": ["en"], "account_id": 1, "categories": [4, 6, 5], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T18:57:14.03484Z", "description": "New desc", "auction_type": "first_price", "device_types": [1, 4], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 18:57:25.432652
b5836b54-84bc-4b9b-ac05-916280fba030	971c04bc-c04f-4ae8-9820-48937f614bcc	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.01, "browsers": [], "carriers": [1], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": ["en"], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T18:57:25.43353Z", "description": "New desc", "auction_type": "first_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-08-25 18:57:50.351942
beca94df-711f-40cc-ae89-656a6fd5999d	26ec4aa6-97ce-4e7e-9640-2648675cd626	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T19:22:35.391456Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-25 19:23:45.652837
7972f613-9e5c-4149-8906-4b8c7a04cc32	14ee2977-320a-4cd2-adf2-5af0402a9c02	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T19:23:45.654633Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-25 19:23:50.554981
eb0cc869-a16d-4114-ba28-a624bee3f068	73fa1210-6f2c-448d-a919-3bda100cb09c	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T19:23:50.556315Z", "description": "New desc", "auction_type": "first_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-25 19:42:18.051389
64eadbed-30bc-4209-95e5-596da5f5c4ac	e1f935ba-bfff-43fe-aa15-95021bb67dbb	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T19:42:18.052688Z", "description": "New desc", "auction_type": "first_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-25 19:51:03.992181
ff3d43e3-2d41-47ad-9766-bb37e42b595a	c35d80ec-e5f4-402e-ba76-8bb865c060ae	1	1	update		RTBAccessPoint	2	2	{"id": 2, "ip": 0, "os": [], "rps": 100, "apps": [], "flags": {}, "title": "test AP", "zones": [], "active": "active", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "timeout": 1000, "browsers": [], "carriers": [], "codename": "test", "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T14:39:47.012746Z", "deleted_at": null, "updated_at": "2024-08-25T15:18:56.774392Z", "description": "Desc", "rtb_sources": null, "auction_type": "first_price", "device_types": [], "request_type": "json", "domain_default": "", "private_browsing": 0, "fixed_purchase_price": 0, "revenue_share_reduce": 0}	2024-08-25 19:51:12.810096
5e3c06a5-6654-424c-b204-24440173cb9a	d7eb7e96-3414-47a8-a6d9-5dc3d5768856	1	1	create		RTBSource	0	0	{"id": 0, "ip": 0, "os": [], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test2", "zones": [], "active": "pause", "config": null, "method": "POST", "secure": 0, "status": "pending", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T20:24:00.752410343Z", "deleted_at": null, "updated_at": "2024-08-25T20:24:00.752410343Z", "description": "", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-08-25 20:24:00.752428
6e6ebf2d-0690-40da-af7c-9f8c12d96871	fadc9cba-a02b-4e5a-9fa7-747ba79eb140	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-08-25T19:51:03.993376Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.1}	2024-09-05 16:17:45.616575
469f3977-5f95-440f-a4c2-d5e3c32078ac	ff7895be-22fc-49b8-96c0-53a0c4455ad8	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:23:07.930633257Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:23:07.930633257Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:23:07.930719
c91dccc9-f09e-4417-806c-e73aee5677eb	367b2243-bd3f-4cb3-8cbb-50fc0cf3852e	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.1, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-09-05T16:17:45.622079Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:12:12.392604
774c09d4-a6d1-4108-a5b9-0e1c53769cca	ff26afb1-7a16-4b98-8db0-0ea3e2cf3a7e	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 0.1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-09-05T18:12:12.395715Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:12:24.825657
753492a0-df62-4a40-be26-b6ab9ee85796	1a016fe7-bc91-47c6-b535-e9e7e75d46d5	1	2	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-09-05T18:12:24.827177Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:12:35.707122
d52badcd-7e38-4536-baef-85b703012d1d	8a0f0dd2-bc43-4798-9691-0a3a736d7ff4	1	1	update		RTBSource	3	3	{"id": 3, "ip": 0, "os": [], "rps": 500, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test source", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": ["A2", "AF", "AG", "AL", "JP", "RU", "US", "**", "AD", "AI", "CN", "A1"], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-24T17:19:05.814555Z", "deleted_at": null, "updated_at": "2024-09-05T18:12:35.708214Z", "description": "New desc", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:19:42.213251
9d88abfa-adda-4872-9851-02af8f42397f	0d55bd76-62c5-431e-aff6-da5f7eefe7e3	1	1	update		RTBSource	4	4	{"id": 4, "ip": 0, "os": [], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test2", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 1, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T20:24:00.75241Z", "deleted_at": null, "updated_at": "2024-08-25T20:24:06.90334Z", "description": "", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:20:52.211259
40bacc34-7189-4124-a14f-c70b9d8959ca	29aef78c-93b3-4548-b96c-9e25cda18af1	1	1	update		RTBSource	4	4	{"id": 4, "ip": 0, "os": [], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test2", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T20:24:00.75241Z", "deleted_at": null, "updated_at": "2024-09-05T18:20:52.212457Z", "description": "", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:21:10.472582
9357d7a2-0a2e-4441-854b-a08926fe5bc1	de822f68-1a82-47eb-8064-948c69de3ad1	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:31:09.852181841Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:31:09.852181841Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:31:09.852707
c8622ca4-6368-4f9f-85a0-5af809ab7e90	36df2d5d-ffff-412f-8c2d-2b58a077f525	1	1	update		RTBSource	4	4	{"id": 4, "ip": 0, "os": [], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test2", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T20:24:00.75241Z", "deleted_at": null, "updated_at": "2024-09-05T18:21:10.474132Z", "description": "", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:23:04.595768
c34a94b6-177b-49de-b153-89a11532797d	b94a8ce1-f18e-4d91-a241-b6fd05fd57fa	1	1	update		RTBSource	4	4	{"id": 4, "ip": 0, "os": [], "rps": 100, "url": "https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars", "apps": [], "flags": {}, "title": "Test2", "zones": [], "active": "active", "config": null, "method": "POST", "secure": 0, "status": "approved", "adblock": 0, "devices": [], "domains": [], "formats": [], "headers": null, "max_bid": 2, "min_bid": 0.01, "timeout": 1000, "accuracy": 1, "browsers": [], "carriers": [], "protocol": "openrtb", "countries": [], "languages": [], "account_id": 1, "categories": [], "created_at": "2024-08-25T20:24:00.75241Z", "deleted_at": null, "updated_at": "2024-09-05T18:23:04.596913Z", "description": "", "auction_type": "second_price", "device_types": [], "request_type": "json", "external_zones": null, "minimal_weight": 0, "private_browsing": 0, "price_correction_reduce": 0.01}	2024-09-05 18:23:25.41857
dc09e0b0-9838-486c-822b-c96c3987ce0a	0	0	0	update		AccountMember	2	2	{"": [{"ID": 1, "Name": "system:admin", "Title": "System admins", "Context": null, "CreatedAt": "2024-07-20T17:56:42.286614Z", "DeletedAt": null, "UpdatedAt": "2024-07-20T17:56:42.286614Z", "ChildRoles": null, "AccessLevel": 100, "Description": "System administrators have full access to all system resources", "PermissionPatterns": ["*"]}, {"ID": 6, "Name": "account:admin", "Title": "Account admins", "Context": null, "CreatedAt": "2024-07-20T17:56:42.286614Z", "DeletedAt": null, "UpdatedAt": "2024-07-20T17:56:42.286614Z", "ChildRoles": null, "AccessLevel": 2, "Description": "Account administrators have full access to all account resources", "PermissionPatterns": ["*.*.{account|owner}", "*.*.*.{account|owner}", "role.check", "user.password.reset.{account|owner}", "permission.list", "account.member.roles.set.account", "certificate.view.statistic.{account|owner}"]}], "id": 2, "user_id": 1, "is_admin": true, "account_id": 2, "created_at": "2024-08-06T20:10:48.657432Z", "deleted_at": null, "updated_at": "2024-08-06T20:10:48.657432Z", "approve_status": 1}	2024-09-07 13:16:53.604957
0b10babf-d8cd-4f43-ba40-004b8cf0f4b8	e64c86d3-4c9f-4516-b328-fef97b9ec43f	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T11:51:59.125037087Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T11:51:59.125037087Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 11:51:59.125155
0dce058c-ea9a-4241-9e79-ce093439aca0	c95f6f6c-27b3-40eb-aeea-a84caf23f4bd	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T12:01:49.731241055Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T12:01:49.731241055Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 12:01:49.731339
bfbc8c63-2769-46e9-992c-a3c37eee71d7	17bcd673-f99a-4528-b419-1da60737303e	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T12:53:58.94345442Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T12:53:58.94345442Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 12:53:58.943469
aa54b8a5-3abe-4ea5-be01-4fc2e6bf392b	1e88f532-930b-44ea-819b-6d8000967042	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T12:54:50.870028208Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T12:54:50.870028208Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 12:54:50.870045
e7344c3a-1ec4-43e1-a270-66c2d810808b	71f9036d-573b-4108-a84a-b4aa4a3546b5	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T12:56:16.631950845Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T12:56:16.631950845Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 12:56:16.631984
8b5e4db8-15c6-4e0b-81d6-f8dc58ec6e47	35279efc-78bb-4e10-8ebb-adbd2eca7332	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T13:11:02.718116839Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T13:11:02.718116839Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 13:11:02.71813
dc7fac61-2dd2-4c30-a209-a92c7e953c17	2f4e8f47-3966-470a-927b-416ffadd2ed2	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 0, "categories": [], "created_at": "2024-10-14T13:11:17.571163429Z", "creator_id": 0, "deleted_at": null, "updated_at": "2024-10-14T13:11:17.571163429Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 13:11:17.571178
7646baf3-6031-4bf8-a7d5-97b6a095d6a3	78d47aad-289c-400d-9b9d-03659b69c235	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:32:57.093244835Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:32:57.093244835Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:32:57.093843
68a0bceb-d3b0-47e8-86e9-58120b10bfdd	d8dfdfca-a413-419a-af17-03ee49a2690c	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:33:14.884987635Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:33:14.884987635Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:33:14.885367
a00c8524-5030-42dc-bede-4384e6eff945	24a97cd2-9f0f-4cb8-b30e-5b132612c04d	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:33:17.191498178Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:33:17.191498178Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:33:17.191941
9a7de3ef-d5e1-4f6d-ac26-1071ad7209aa	e6a741c4-6474-4139-8353-7ccbbe84b61b	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:37:31.548553754Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:37:31.548553754Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:37:31.549158
333bf192-14f4-47b6-99f6-f1f6362e1dec	f0abc3d4-6fb2-4af7-a9ef-1bf8374eaec0	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:38:16.065085511Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:38:16.065085511Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:38:16.065937
b8ac490c-b926-4a42-9e46-9d21e6ec2ae4	1a3d90a6-7be9-417f-8b7d-3b3fb054dead	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:52:31.00944417Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:52:31.00944417Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-14 16:52:31.009836
b06a14ab-629d-4dd7-b118-859ab3bbde1c	d73f8703-810e-4afb-815d-003c93ae30e5	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T15:54:46.688730172Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T15:54:46.688730172Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 15:54:46.689408
e794e177-18af-4b3d-9a20-6ebd60f5a2a2	049c4e44-608a-4908-b7fa-32110c17f070	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:26:29.048992136Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:26:29.048992136Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:26:29.049372
c1f07ed2-a704-4c77-a8c5-2f700c5d971f	3f146be1-ca26-4522-8c82-5a5db1e8bfe6	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:27:34.458681222Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:27:34.458681222Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:27:34.459315
ece74f38-c9a8-41d6-9fe4-0a7bd70b1a30	c35d7e5b-5e3a-4f21-b945-861d30cd862c	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:31:12.840687128Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:31:12.840687128Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:31:12.8411
42ec9eb2-2059-4ca8-b211-c891005743a9	7e106060-3d43-4e0f-9580-bb09c17cecc9	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:34:02.855848596Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:34:02.855848596Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:34:02.856552
7992f03f-552c-44ac-9bd3-6cc944b573fc	3411e477-d090-4ffc-bbe9-c456190515c9	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:42:51.942644758Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:42:51.942644758Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:42:51.943322
ee0ce8bf-1cf2-4d07-8f11-3244e2b71482	ae2ed687-6075-46dd-9658-a4cb61dba62f	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:45:44.539869004Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:45:44.539869004Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:45:44.540051
37297ab6-6210-41e5-9823-c66608c7bf3a	ae2ed687-6075-46dd-9658-a4cb61dba62f	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:45:44.539869004Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:45:44.539869004Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:45:44.54189
efa17672-122a-4704-a534-3b8e74c91903	138e3238-bc90-4db7-a91a-80481972d018	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:49:20.733660799Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:49:20.733660799Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:49:20.733828
dc065de5-ce62-479a-aecf-05d8f86dff89	138e3238-bc90-4db7-a91a-80481972d018	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:49:20.733660799Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:49:20.733660799Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:49:20.735931
b758fcbc-8590-4394-9922-f740c28e71a4	138e3238-bc90-4db7-a91a-80481972d018	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:49:20.733660799Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:49:20.733660799Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:49:20.737156
797cf789-5dec-4e4a-990e-833b8aebae56	daf665fd-cf68-4383-857e-f626a1c8cc9e	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:51:18.854909381Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854909381Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:51:18.855077
c6adfb69-f177-4dd2-b140-6ee074e7efa8	daf665fd-cf68-4383-857e-f626a1c8cc9e	1	2	create		Application	0	0	{"id": 0, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-15T16:51:18.854909381Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854909381Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 16:51:18.857191
39a96de4-bb31-4517-9510-58dfa86dc815	7ff4d761-20f5-4ade-9934-42acf435bd53	1	2	update		Application	3	3	{"id": 3, "uri": "test.com", "type": "site", "title": "Test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:23:07.93Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:23:07.93Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 17:02:38.823413
4e4acf1c-6416-4b56-a2a2-4524e07093f2	7b4e401f-2fa6-451f-9f85-5ab3ef4bf48e	1	2	update		Application	3	3	{"id": 3, "uri": "test.com", "type": "site", "title": "Test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:23:07.93Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-14T16:23:07.93Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 17:03:32.562022
9e56b267-cdc5-4024-98ef-acf393a314a8	ee463e18-8780-4a0f-a3c8-b5f45dfcfa55	1	2	update		Application	3	3	{"id": 3, "uri": "test.com", "type": "site", "title": "Test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [], "created_at": "2024-10-14T16:23:07.93Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T17:03:32.563687Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-15 17:06:48.572405
98cdd97c-aa94-47d7-99f1-6e7f3314d4e5	df3d319d-6171-45eb-8527-2d7a8d50bd82	1	2	update		Application	3	3	{"id": 3, "uri": "test.com", "type": "site", "title": "Test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [4], "created_at": "2024-10-14T16:23:07.93Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T17:03:32.563687Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-16 10:08:46.192768
337da3fc-a43a-43b3-8b8c-452c79c8e67a	c7de2113-054e-49ca-8d7f-52fab3b9f86d	1	2	update		Application	4	4	{"id": 4, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [6], "created_at": "2024-10-15T16:51:18.854Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-16 10:08:54.201331
cc6631c8-4d11-4941-80c9-aa90524173b0	1dcfad4a-d00e-47cd-b88b-be5f17a98842	1	2	update		Application	4	4	{"id": 4, "uri": "test.com", "type": "site", "title": "test", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [5], "created_at": "2024-10-15T16:51:18.854Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-16 10:14:32.498506
42c80b97-615d-461e-bd64-2d33e8b7ead8	91a6a7e4-c8c2-49ec-a3e3-eaa5a63fcd84	1	2	update		Application	4	4	{"id": 4, "uri": "test.com", "type": "site", "title": "test1", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [5], "created_at": "2024-10-15T16:51:18.854Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-16 10:16:30.0273
7b251ad4-b38c-49eb-a13f-740d093505f0	62946eae-3946-442b-8db4-07b42a9bf20c	1	2	update		Application	4	4	{"id": 4, "uri": "test.com", "type": "site", "title": "test1", "active": "pause", "status": "pending", "premium": false, "private": "public", "platform": "web", "account_id": 2, "categories": [5], "created_at": "2024-10-15T16:51:18.854Z", "creator_id": 1, "deleted_at": null, "updated_at": "2024-10-15T16:51:18.854Z", "description": "", "revenue_share": 0, "allowed_sources": null, "disallowed_sources": null}	2024-10-16 10:17:45.247492
84d9eb20-b4f1-4a22-89c3-d519659c2de4	04397c03-581f-4fc9-a7b2-0efdb6b52fd9	1	2	create		Zone	0	0	{"id": 0, "type": "zone", "title": "top", "active": "pause", "status": "pending", "context": null, "codename": "2nkB6oMIwRACemshvsI31AX94jQ", "min_ecpm": 0, "campaigns": null, "account_id": 0, "created_at": "2024-10-21T10:37:00.515803887Z", "deleted_at": null, "updated_at": "2024-10-21T10:37:00.515803887Z", "description": "", "default_code": null, "allowed_types": null, "allowed_formats": null, "allowed_sources": null, "min_ecpm_by_geo": null, "disallowed_sources": null, "fixed_purchase_price": 0}	2024-10-21 10:37:00.516056
338c3330-72b1-4189-b311-5c4d83d277a1	e43bff5d-f8f7-4ec9-b64b-5abce3653f0d	1	2	create		Zone	0	0	{"id": 0, "type": "zone", "title": "top", "active": "pause", "status": "pending", "context": null, "codename": "2nkBDVRFJOzNDVnyIq86rvhMuaK", "min_ecpm": 0, "campaigns": null, "account_id": 0, "created_at": "2024-10-21T10:37:53.427458217Z", "deleted_at": null, "updated_at": "2024-10-21T10:37:53.427458217Z", "description": "", "default_code": null, "allowed_types": null, "allowed_formats": null, "allowed_sources": null, "min_ecpm_by_geo": null, "disallowed_sources": null, "fixed_purchase_price": 0}	2024-10-21 10:37:53.427962
5a150fea-b5d4-449d-9b96-55f41cb2163a	f4e4e075-258d-4b54-bfc6-153080c3e5f2	1	2	create		Zone	0	0	{"id": 0, "type": "zone", "title": "top", "active": "pause", "status": "pending", "context": null, "codename": "2nkBFfyBc8LIzByJ4ryAHEauGVl", "min_ecpm": 0, "campaigns": null, "account_id": 0, "created_at": "2024-10-21T10:38:11.753327378Z", "deleted_at": null, "updated_at": "2024-10-21T10:38:11.753327378Z", "description": "", "default_code": null, "allowed_types": null, "allowed_formats": null, "allowed_sources": null, "min_ecpm_by_geo": null, "disallowed_sources": null, "fixed_purchase_price": 0}	2024-10-21 10:38:11.75488
8e0a1d9f-a0ce-4e54-93dc-a75ce9a153aa	f69572b5-7d67-4267-8f05-031c5325d635	1	2	create		Zone	2	2	{"id": 2, "type": "zone", "title": "top", "active": "pause", "status": "pending", "context": null, "codename": "2nkBXMzDH71gxXKlBxH9GiG8eGc", "min_ecpm": 0, "campaigns": null, "account_id": 2, "created_at": "2024-10-21T10:40:31.286395679Z", "deleted_at": null, "updated_at": "2024-10-21T10:40:31.286395679Z", "description": "", "default_code": null, "allowed_types": null, "allowed_formats": null, "allowed_sources": null, "min_ecpm_by_geo": null, "disallowed_sources": null, "fixed_purchase_price": 0}	2024-10-21 10:40:31.288087
\.


--
-- Data for Name: m2m_account_member_role; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.m2m_account_member_role (member_id, role_id, created_at) FROM stdin;
1	1	2024-07-20 17:56:42.286614+00
1	6	2024-08-06 20:10:11.027616+00
1	7	2024-08-06 20:10:11.824773+00
1	9	2024-08-06 20:10:12.44865+00
2	1	2024-09-07 13:16:53.612221+00
2	6	2024-09-07 13:16:53.612221+00
\.


--
-- Data for Name: m2m_rbac_role; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.m2m_rbac_role (parent_role_id, child_role_id, created_at) FROM stdin;
\.


--
-- Data for Name: option; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.option (name, type, target_id, value, created_at, updated_at, deleted_at) FROM stdin;
\.


--
-- Data for Name: rbac_role; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.rbac_role (id, name, title, description, context, permissions, access_level, created_at, updated_at, deleted_at) FROM stdin;
1	system:admin	System admins	System administrators have full access to all system resources	\N	{*}	100	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
2	system:manager	System manager	System managers have full access to all system resources except for some sensitive operations	\N	{"*.{view|list|count|create|update|delete|restore|approve|reject|reset}.*",role.**,user.password.reset,account.member.**,permission.**}	90	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
3	system:analyst	System analyst	System analysts have read-only access to all system resources	\N	{"*.{view|list|count}.*","*.*.{view|list|count}.*",role.check,user.password.reset,permission.list}	80	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
4	system:compliance	System compliance	System compliance have access to all system resources with approve/reject permissions	\N	{"*.{view|list|count|approve|reject}.*","*.*.{view|list|count|approve|reject}.*",role.check,user.password.reset,permission.list}	70	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
5	system:viewer	System viewer	System viewers have read-only access to all system resources	\N	{"*.{view|list|count}.*",role.check,user.password.reset,permission.list}	60	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
6	account:admin	Account admins	Account administrators have full access to all account resources	\N	{"*.*.{account|owner}","*.*.*.{account|owner}",role.check,"user.password.reset.{account|owner}",permission.list,account.member.roles.set.account,"certificate.view.statistic.{account|owner}"}	2	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
7	account:writer	Account writer	Account writers have full access to all account resources except for some sensitive operations	\N	{"*.{view|list|restore}.{account|owner}","*.*.{view|list|restore}.{account|owner}",role.check,user.password.reset,permission.list,"certificate.view.statistic.{account|owner}"}	2	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
8	account:analyst	Account analyst	Account analysts have read-only access to all account resources and analytics	\N	{"*.{view|list}.{account|owner}","*.*.{view|list}.{account|owner}",role.check,user.password.reset,permission.list,"certificate.view.statistic.{account|owner}"}	2	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
9	account:viewer	Account viewer	Account viewers have read-only access to all account resources except analytics	\N	{"*.{view|list}.{account|owner}","*.*.{view|list}.{account|owner}",role.check,user.password.reset,permission.list,"certificate.view.statistic.{account|owner}"}	2	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
10	account:compliance	Account compliance	Account compliance have access to all account resources with approve/reject permissions	\N	{"*.{view|list|approve|reject}.{account|owner}","*.*.{view|list|approve|reject}.{account|owner}",role.check,user.password.reset,permission.list}	2	2024-07-20 17:56:42.286614+00	2024-07-20 17:56:42.286614+00	\N
\.


--
-- Data for Name: rtb_access_point; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.rtb_access_point (id, account_id, codename, title, description, revenue_share_reduce, auction_type, status, active, flags, protocol, timeout, rps, domain_default, request_type, headers, max_bid, fixed_purchase_price, formats, device_types, devices, os, browsers, categories, carriers, countries, languages, apps, zones, domains, rtb_sources, secure, adblock, private_browsing, ip, created_at, updated_at, deleted_at) FROM stdin;
2	1	test	test AP	Desc	0.0000000000	first_price	approved	active	{}	openrtb	1000	100		json	null	1.00000	0.00000	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	\N	0	0	0	0	2024-08-25 14:39:47.012746	2024-08-25 19:51:12.81155	\N
\.


--
-- Data for Name: rtb_source; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.rtb_source (id, account_id, title, description, status, active, flags, protocol, minimal_weight, url, method, request_type, headers, rps, timeout, accuracy, price_correction_reduce, auction_type, min_bid, max_bid, formats, device_types, devices, os, browsers, carriers, categories, countries, languages, apps, domains, zones, external_zones, secure, adblock, private_browsing, ip, config, created_at, updated_at, deleted_at) FROM stdin;
4	1	Test2		approved	active	{"trace": 1}	openrtb	0	https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars	POST	json	null	100	1000	1	0.01	second_price	0.01	2	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	{}	\N	0	0	0	0	null	2024-08-25 20:24:00.75241	2024-09-07 10:57:30.113441	\N
3	1	Test source	New desc	approved	active	{"trace": 1}	openrtb	0	https://revolvesyndicate.net/dsp/bidrequest/openrtb/trafficstars	POST	json	null	500	1000	1	0.01	second_price	0.01	2	{}	{}	{}	{}	{}	{}	{}	{A2,AF,AG,AL,JP,RU,US,**,AD,AI,CN,A1}	{}	{}	{}	{}	\N	0	0	0	0	null	2024-08-24 17:19:05.814555	2024-09-07 10:57:30.113441	\N
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.schema_migrations (version, dirty) FROM stdin;
10	f
\.


--
-- Data for Name: schema_migrations_dev; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.schema_migrations_dev (version, dirty) FROM stdin;
\.


--
-- Data for Name: schema_migrations_initial; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.schema_migrations_initial (version, dirty) FROM stdin;
9	f
\.


--
-- Data for Name: schema_migrations_initial2; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.schema_migrations_initial2 (version, dirty) FROM stdin;
1	f
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.transactions (id, created_at, key, type, amount, status, gateway_id, gateway_payment_id, gateway_info, message, invoice_id, invoice_key, invoice_number) FROM stdin;
\.


--
-- Data for Name: type_browser; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.type_browser (id, name, active, versions, created_at, updated_at, deleted_at, description) FROM stdin;
1	Chrome	active	[{"max": "1.999", "min": "0.0.1", "name": "1"}]	2024-08-04 10:30:59.308409	2024-08-24 14:40:26.700783	\N	Chrome is good
\.


--
-- Data for Name: type_device_maker; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.type_device_maker (id, name, description, active, created_at, updated_at, deleted_at, match_exp) FROM stdin;
1	Apple	Nice description 1	active	2024-08-04 12:43:57.755586	2024-08-06 11:25:24.429333	\N	Apple*
\.


--
-- Data for Name: type_device_model; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.type_device_model (id, name, description, active, maker_id, created_at, updated_at, deleted_at, type_id, versions) FROM stdin;
1	Test	Test description	active	1	2024-08-04 12:45:41.751358	2024-08-04 12:45:41.751358	\N	0	\N
2	Test 2	Test description	active	1	2024-08-04 12:52:09.934321	2024-08-04 12:52:09.934321	\N	1	\N
3	Test 3	Test description	active	1	2024-08-04 12:52:53.294122	2024-08-04 12:52:53.294122	\N	1	\N
4	Test 4	Test description	active	1	2024-08-04 13:00:39.004626	2024-08-04 13:00:39.004626	\N	1	\N
5	Test 5	Test description	active	1	2024-08-04 13:05:55.426891	2024-08-04 13:05:55.426891	\N	1	\N
6	Test 6	Test description 2	active	1	2024-08-04 14:54:09.742794	2024-08-04 14:54:09.742794	\N	1	[{"max": "0", "min": "1", "name": "1"}, {"max": "3", "min": "2", "name": "2"}]
7	Test 7	Test description 2	active	1	2024-08-04 14:54:38.887077	2024-08-04 14:54:38.887077	\N	1	[{"max": "0", "min": "1", "name": "1"}, {"max": "3", "min": "2", "name": "2"}]
\.


--
-- Data for Name: type_os; Type: TABLE DATA; Schema: public; Owner: dbuser
--

COPY public.type_os (id, name, description, match_exp, active, versions, created_at, updated_at, deleted_at) FROM stdin;
1	Windows	Windows is irritan		active	\N	2024-08-24 14:58:35.108026	2024-08-24 15:00:56.20586	\N
\.


--
-- Name: account_base_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.account_base_id_seq', 2, true);


--
-- Name: account_member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.account_member_id_seq', 2, true);


--
-- Name: account_social_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.account_social_id_seq', 1, false);


--
-- Name: account_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.account_user_id_seq', 1, true);


--
-- Name: ad_campaign_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.ad_campaign_id_seq', 1, false);


--
-- Name: ad_format_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.ad_format_id_seq', 1, false);


--
-- Name: ad_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.ad_item_id_seq', 1, false);


--
-- Name: ad_link_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.ad_link_id_seq', 1, false);


--
-- Name: adv_application_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.adv_application_id_seq', 4, true);


--
-- Name: adv_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.adv_category_id_seq', 6, true);


--
-- Name: adv_format_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.adv_format_id_seq', 2, true);


--
-- Name: adv_zone_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.adv_zone_id_seq', 2, true);


--
-- Name: auth_session_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.auth_session_id_seq', 1, false);


--
-- Name: direct_access_tokens_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.direct_access_tokens_id_seq', 3, true);


--
-- Name: rbac_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.rbac_role_id_seq', 10, true);


--
-- Name: rtb_access_point_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.rtb_access_point_id_seq', 2, true);


--
-- Name: rtb_source_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.rtb_source_id_seq', 4, true);


--
-- Name: type_browser_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.type_browser_id_seq', 1, true);


--
-- Name: type_device_maker_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.type_device_maker_id_seq', 1, true);


--
-- Name: type_device_model_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.type_device_model_id_seq', 7, true);


--
-- Name: type_os_id_seq; Type: SEQUENCE SET; Schema: public; Owner: dbuser
--

SELECT pg_catalog.setval('public.type_os_id_seq', 1, true);


--
-- Name: account_base account_base_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_base
    ADD CONSTRAINT account_base_pkey PRIMARY KEY (id);


--
-- Name: account_member account_member_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_member
    ADD CONSTRAINT account_member_pkey PRIMARY KEY (id);


--
-- Name: account_social account_social_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_social
    ADD CONSTRAINT account_social_pkey PRIMARY KEY (id);


--
-- Name: account_social_session account_social_session_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_social_session
    ADD CONSTRAINT account_social_session_pkey PRIMARY KEY (account_social_id, name);


--
-- Name: account_user account_user_email_key; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_user
    ADD CONSTRAINT account_user_email_key UNIQUE (email);


--
-- Name: account_user_password_reset account_user_password_reset_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_user_password_reset
    ADD CONSTRAINT account_user_password_reset_pkey PRIMARY KEY (token);


--
-- Name: account_user account_user_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_user
    ADD CONSTRAINT account_user_pkey PRIMARY KEY (id);


--
-- Name: adv_campaign ad_campaign_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_campaign
    ADD CONSTRAINT ad_campaign_pkey PRIMARY KEY (id);


--
-- Name: ad_format ad_format_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.ad_format
    ADD CONSTRAINT ad_format_pkey PRIMARY KEY (id);


--
-- Name: adv_item ad_item_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_item
    ADD CONSTRAINT ad_item_pkey PRIMARY KEY (id);


--
-- Name: adv_link ad_link_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_link
    ADD CONSTRAINT ad_link_pkey PRIMARY KEY (id);


--
-- Name: adv_application adv_application_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_application
    ADD CONSTRAINT adv_application_pkey PRIMARY KEY (id);


--
-- Name: adv_category adv_category_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_category
    ADD CONSTRAINT adv_category_pkey PRIMARY KEY (id);


--
-- Name: adv_format adv_format_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_format
    ADD CONSTRAINT adv_format_pkey PRIMARY KEY (id);


--
-- Name: adv_zone adv_zone_codename_key; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_zone
    ADD CONSTRAINT adv_zone_codename_key UNIQUE (codename);


--
-- Name: adv_zone adv_zone_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_zone
    ADD CONSTRAINT adv_zone_pkey PRIMARY KEY (id);


--
-- Name: auth_client auth_client_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_client
    ADD CONSTRAINT auth_client_pkey PRIMARY KEY (id);


--
-- Name: auth_session auth_session_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_session
    ADD CONSTRAINT auth_session_pkey PRIMARY KEY (id);


--
-- Name: direct_access_tokens direct_access_tokens_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.direct_access_tokens
    ADD CONSTRAINT direct_access_tokens_pkey PRIMARY KEY (id);


--
-- Name: direct_access_tokens direct_access_tokens_token_key; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.direct_access_tokens
    ADD CONSTRAINT direct_access_tokens_token_key UNIQUE (token);


--
-- Name: history_actions history_actions_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.history_actions
    ADD CONSTRAINT history_actions_pkey PRIMARY KEY (id);


--
-- Name: m2m_account_member_role m2m_account_member_role_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_account_member_role
    ADD CONSTRAINT m2m_account_member_role_pkey PRIMARY KEY (member_id, role_id);


--
-- Name: m2m_rbac_role m2m_rbac_role_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_rbac_role
    ADD CONSTRAINT m2m_rbac_role_pkey PRIMARY KEY (parent_role_id, child_role_id);


--
-- Name: option option_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.option
    ADD CONSTRAINT option_pkey PRIMARY KEY (name, type, target_id);


--
-- Name: rbac_role rbac_role_name_key; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rbac_role
    ADD CONSTRAINT rbac_role_name_key UNIQUE (name);


--
-- Name: rbac_role rbac_role_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rbac_role
    ADD CONSTRAINT rbac_role_pkey PRIMARY KEY (id);


--
-- Name: rtb_access_point rtb_access_point_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_access_point
    ADD CONSTRAINT rtb_access_point_pkey PRIMARY KEY (id);


--
-- Name: rtb_source rtb_source_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_source
    ADD CONSTRAINT rtb_source_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations_dev schema_migrations_dev_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.schema_migrations_dev
    ADD CONSTRAINT schema_migrations_dev_pkey PRIMARY KEY (version);


--
-- Name: schema_migrations_initial2 schema_migrations_initial2_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.schema_migrations_initial2
    ADD CONSTRAINT schema_migrations_initial2_pkey PRIMARY KEY (version);


--
-- Name: schema_migrations_initial schema_migrations_initial_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.schema_migrations_initial
    ADD CONSTRAINT schema_migrations_initial_pkey PRIMARY KEY (version);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: type_browser type_browser_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_browser
    ADD CONSTRAINT type_browser_pkey PRIMARY KEY (id);


--
-- Name: type_device_maker type_device_maker_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_device_maker
    ADD CONSTRAINT type_device_maker_pkey PRIMARY KEY (id);


--
-- Name: type_device_model type_device_model_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_device_model
    ADD CONSTRAINT type_device_model_pkey PRIMARY KEY (id);


--
-- Name: type_os type_os_pkey; Type: CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_os
    ADD CONSTRAINT type_os_pkey PRIMARY KEY (id);


--
-- Name: idx_account_member_unique_account_user; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE UNIQUE INDEX idx_account_member_unique_account_user ON public.account_member USING btree (account_id, user_id);


--
-- Name: idx_account_social_uniq_social_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE UNIQUE INDEX idx_account_social_uniq_social_id ON public.account_social USING btree (social_id, provider) WHERE (deleted_at IS NULL);


--
-- Name: idx_account_user_password_reset_expires_at; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_account_user_password_reset_expires_at ON public.account_user_password_reset USING btree (expires_at);


--
-- Name: idx_auth_session_uniq_access_token_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE UNIQUE INDEX idx_auth_session_uniq_access_token_id ON public.auth_session USING btree (access_token) WHERE (deleted_at IS NULL);


--
-- Name: idx_auth_session_uniq_refresh_token_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE UNIQUE INDEX idx_auth_session_uniq_refresh_token_id ON public.auth_session USING btree (refresh_token) WHERE ((length((COALESCE(refresh_token, ''::character varying))::text) > 0) AND (deleted_at IS NULL));


--
-- Name: idx_auth_session_uniq_request_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE UNIQUE INDEX idx_auth_session_uniq_request_id ON public.auth_session USING btree (request_id) WHERE (deleted_at IS NULL);


--
-- Name: idx_history_actions_account_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_account_id ON public.history_actions USING btree (account_id);


--
-- Name: idx_history_actions_at; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_at ON public.history_actions USING btree (action_at);


--
-- Name: idx_history_actions_name; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_name ON public.history_actions USING btree (name);


--
-- Name: idx_history_actions_object_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_object_id ON public.history_actions USING btree (object_id);


--
-- Name: idx_history_actions_object_ids; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_object_ids ON public.history_actions USING btree (object_ids);


--
-- Name: idx_history_actions_object_type; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_object_type ON public.history_actions USING btree (object_type);


--
-- Name: idx_history_actions_request_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_request_id ON public.history_actions USING btree (request_id);


--
-- Name: idx_history_actions_user_id; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX idx_history_actions_user_id ON public.history_actions USING btree (user_id);


--
-- Name: transactions_key_idx; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX transactions_key_idx ON public.transactions USING btree (key);


--
-- Name: transactions_key_type_status_invoice_key_idx; Type: INDEX; Schema: public; Owner: dbuser
--

CREATE INDEX transactions_key_type_status_invoice_key_idx ON public.transactions USING btree (key, type, status, invoice_key);


--
-- Name: account_user_password_reset keep_for_one_week; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER keep_for_one_week BEFORE INSERT OR UPDATE ON public.account_user_password_reset FOR EACH ROW EXECUTE FUNCTION public.keep_for('expires_at', '1 week');


--
-- Name: account_base notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.account_base FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: account_social notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.account_social FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: account_social_session notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.account_social_session FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: account_user notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.account_user FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: account_user_password_reset notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.account_user_password_reset FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: ad_format notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.ad_format FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_application notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_application FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_campaign notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_campaign FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_category notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_category FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_format notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_format FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_item notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_item FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_link notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_link FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: adv_zone notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.adv_zone FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: auth_client notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.auth_client FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: auth_session notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.auth_session FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: direct_access_tokens notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.direct_access_tokens FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: option notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.option FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: rtb_access_point notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.rtb_access_point FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: rtb_source notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.rtb_source FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: type_browser notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.type_browser FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: type_device_maker notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.type_device_maker FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: type_device_model notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.type_device_model FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: type_os notify_update_event_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER notify_update_event_trigger AFTER INSERT OR DELETE OR UPDATE ON public.type_os FOR EACH ROW EXECUTE FUNCTION public.notify_update_event();


--
-- Name: account_base updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.account_base FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: account_member updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.account_member FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: account_social updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.account_social FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: account_social_session updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.account_social_session FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: account_user updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.account_user FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: ad_format updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.ad_format FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_format updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.adv_format FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_item updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.adv_item FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: auth_client updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.auth_client FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: auth_session updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.auth_session FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: m2m_rbac_role updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.m2m_rbac_role FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: option updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.option FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: rbac_role updated_at_triger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_triger BEFORE UPDATE ON public.rbac_role FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_application updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.adv_application FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_campaign updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.adv_campaign FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_category updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.adv_category FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_link updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.adv_link FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: adv_zone updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.adv_zone FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: rtb_access_point updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.rtb_access_point FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: rtb_source updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.rtb_source FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: type_browser updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.type_browser FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: type_device_maker updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.type_device_maker FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: type_device_model updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.type_device_model FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: type_os updated_at_trigger; Type: TRIGGER; Schema: public; Owner: dbuser
--

CREATE TRIGGER updated_at_trigger BEFORE UPDATE ON public.type_os FOR EACH ROW EXECUTE FUNCTION public.updated_at_column();


--
-- Name: account_member account_member_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_member
    ADD CONSTRAINT account_member_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: account_member account_member_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_member
    ADD CONSTRAINT account_member_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: account_social_session account_social_session_account_social_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_social_session
    ADD CONSTRAINT account_social_session_account_social_id_fkey FOREIGN KEY (account_social_id) REFERENCES public.account_social(id) ON DELETE RESTRICT;


--
-- Name: account_social account_social_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_social
    ADD CONSTRAINT account_social_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: account_user_password_reset account_user_password_reset_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.account_user_password_reset
    ADD CONSTRAINT account_user_password_reset_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: adv_campaign ad_campaign_company_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_campaign
    ADD CONSTRAINT ad_campaign_company_id_fkey FOREIGN KEY (company_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: adv_campaign ad_campaign_creator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_campaign
    ADD CONSTRAINT ad_campaign_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: adv_campaign ad_campaign_moderator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_campaign
    ADD CONSTRAINT ad_campaign_moderator_id_fkey FOREIGN KEY (moderator_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: adv_item ad_item_campaign_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_item
    ADD CONSTRAINT ad_item_campaign_id_fkey FOREIGN KEY (campaign_id) REFERENCES public.adv_campaign(id) ON DELETE RESTRICT;


--
-- Name: adv_item ad_item_format_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_item
    ADD CONSTRAINT ad_item_format_id_fkey FOREIGN KEY (format_id) REFERENCES public.ad_format(id) ON DELETE RESTRICT;


--
-- Name: adv_link ad_link_campaign_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_link
    ADD CONSTRAINT ad_link_campaign_id_fkey FOREIGN KEY (campaign_id) REFERENCES public.adv_campaign(id) ON DELETE RESTRICT;


--
-- Name: adv_application adv_application_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_application
    ADD CONSTRAINT adv_application_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: adv_application adv_application_creator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_application
    ADD CONSTRAINT adv_application_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: adv_category adv_category_parent_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_category
    ADD CONSTRAINT adv_category_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES public.adv_category(id) ON DELETE SET NULL;


--
-- Name: adv_zone adv_zone_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.adv_zone
    ADD CONSTRAINT adv_zone_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: auth_client auth_client_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_client
    ADD CONSTRAINT auth_client_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: auth_client auth_client_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_client
    ADD CONSTRAINT auth_client_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: auth_session auth_session_client_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.auth_session
    ADD CONSTRAINT auth_session_client_id_fkey FOREIGN KEY (client_id) REFERENCES public.auth_client(id) ON DELETE RESTRICT;


--
-- Name: direct_access_tokens direct_access_tokens_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.direct_access_tokens
    ADD CONSTRAINT direct_access_tokens_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: direct_access_tokens direct_access_tokens_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.direct_access_tokens
    ADD CONSTRAINT direct_access_tokens_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.account_user(id) ON DELETE RESTRICT;


--
-- Name: m2m_account_member_role m2m_account_member_role_member_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_account_member_role
    ADD CONSTRAINT m2m_account_member_role_member_id_fkey FOREIGN KEY (member_id) REFERENCES public.account_member(id) ON DELETE RESTRICT;


--
-- Name: m2m_account_member_role m2m_account_member_role_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_account_member_role
    ADD CONSTRAINT m2m_account_member_role_role_id_fkey FOREIGN KEY (role_id) REFERENCES public.rbac_role(id) ON DELETE RESTRICT;


--
-- Name: m2m_rbac_role m2m_rbac_role_child_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_rbac_role
    ADD CONSTRAINT m2m_rbac_role_child_role_id_fkey FOREIGN KEY (child_role_id) REFERENCES public.rbac_role(id) ON DELETE RESTRICT;


--
-- Name: m2m_rbac_role m2m_rbac_role_parent_role_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.m2m_rbac_role
    ADD CONSTRAINT m2m_rbac_role_parent_role_id_fkey FOREIGN KEY (parent_role_id) REFERENCES public.rbac_role(id) ON DELETE RESTRICT;


--
-- Name: rtb_access_point rtb_access_point_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_access_point
    ADD CONSTRAINT rtb_access_point_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: rtb_source rtb_source_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.rtb_source
    ADD CONSTRAINT rtb_source_account_id_fkey FOREIGN KEY (account_id) REFERENCES public.account_base(id) ON DELETE RESTRICT;


--
-- Name: type_device_model type_device_model_maker_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: dbuser
--

ALTER TABLE ONLY public.type_device_model
    ADD CONSTRAINT type_device_model_maker_id_fkey FOREIGN KEY (maker_id) REFERENCES public.type_device_maker(id) ON DELETE RESTRICT;


--
-- PostgreSQL database dump complete
--

