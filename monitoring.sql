--
-- PostgreSQL database dump
--

-- Dumped from database version 12.1
-- Dumped by pg_dump version 12rc1

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
-- Name: category; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.category (
    id bigint NOT NULL,
    code_level character varying NOT NULL,
    name character varying NOT NULL,
    parent character varying,
    update_at timestamp without time zone,
    additional_input_1 character varying DEFAULT '-'::character varying,
    additional_input_2 character varying DEFAULT '-'::character varying,
    additional_input_3 character varying DEFAULT '-'::character varying
);


ALTER TABLE monitoring_maintenance.category OWNER TO postgres;

--
-- Name: kategori_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.kategori_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.kategori_id_seq OWNER TO postgres;

--
-- Name: kategori_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.kategori_id_seq OWNED BY monitoring_maintenance.category.id;


--
-- Name: role; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.role (
    id bigint NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE monitoring_maintenance.role OWNER TO postgres;

--
-- Name: role_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.role_id_seq OWNER TO postgres;

--
-- Name: role_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.role_id_seq OWNED BY monitoring_maintenance.role.id;


--
-- Name: ticket; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.ticket (
    id bigint NOT NULL,
    judul character varying NOT NULL,
    username_pembuat character varying NOT NULL,
    username_pembalas character varying,
    prioritas character varying NOT NULL,
    tgl_dibuat timestamp without time zone NOT NULL,
    tgl_diperbarui timestamp without time zone NOT NULL,
    total_waktu character varying,
    status character varying NOT NULL,
    ticket_code character varying,
    category character varying NOT NULL,
    lokasi character varying NOT NULL,
    terminal_id character varying NOT NULL,
    email character varying,
    assigned_to character varying,
    assigned_to_team character varying
);


ALTER TABLE monitoring_maintenance.ticket OWNER TO postgres;

--
-- Name: table1_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.table1_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.table1_id_seq OWNER TO postgres;

--
-- Name: table1_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.table1_id_seq OWNED BY monitoring_maintenance.ticket.id;


--
-- Name: task_list; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.task_list (
    id bigint NOT NULL,
    ticket_code character varying(255),
    attachment character varying(255),
    description character varying(255),
    task_name character varying(255),
    longitude character varying(255),
    latitude character varying(255),
    assigned_by character varying(255),
    status character varying(255),
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    index character varying(255)
);


ALTER TABLE monitoring_maintenance.task_list OWNER TO postgres;

--
-- Name: task_list_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.task_list_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.task_list_id_seq OWNER TO postgres;

--
-- Name: task_list_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.task_list_id_seq OWNED BY monitoring_maintenance.task_list.id;


--
-- Name: team; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.team (
    id bigint NOT NULL,
    name character varying(255),
    created_at timestamp without time zone DEFAULT now() NOT NULL
);


ALTER TABLE monitoring_maintenance.team OWNER TO postgres;

--
-- Name: team_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.team_id_seq OWNER TO postgres;

--
-- Name: team_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.team_id_seq OWNED BY monitoring_maintenance.team.id;


--
-- Name: ticket_isi; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.ticket_isi (
    id bigint NOT NULL,
    username_pengirim character varying NOT NULL,
    isi character varying NOT NULL,
    attachment1 character varying,
    tgl_dibuat timestamp without time zone NOT NULL,
    ticket_code character varying(255),
    attachment2 character varying
);


ALTER TABLE monitoring_maintenance.ticket_isi OWNER TO postgres;

--
-- Name: ticket_isi_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.ticket_isi_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.ticket_isi_id_seq OWNER TO postgres;

--
-- Name: ticket_isi_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.ticket_isi_id_seq OWNED BY monitoring_maintenance.ticket_isi.id;


--
-- Name: users; Type: TABLE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE TABLE monitoring_maintenance.users (
    id bigint NOT NULL,
    name character varying(50),
    username character varying(50) NOT NULL,
    email character varying(50),
    password character varying(255) NOT NULL,
    changepass integer,
    updated_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    role character varying,
    team character varying
);


ALTER TABLE monitoring_maintenance.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: monitoring_maintenance; Owner: postgres
--

CREATE SEQUENCE monitoring_maintenance.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE monitoring_maintenance.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: monitoring_maintenance; Owner: postgres
--

ALTER SEQUENCE monitoring_maintenance.users_id_seq OWNED BY monitoring_maintenance.users.id;


--
-- Name: category id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.category ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.kategori_id_seq'::regclass);


--
-- Name: role id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.role ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.role_id_seq'::regclass);


--
-- Name: task_list id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.task_list ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.task_list_id_seq'::regclass);


--
-- Name: team id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.team ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.team_id_seq'::regclass);


--
-- Name: ticket id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.ticket ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.table1_id_seq'::regclass);


--
-- Name: ticket_isi id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.ticket_isi ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.ticket_isi_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.users ALTER COLUMN id SET DEFAULT nextval('monitoring_maintenance.users_id_seq'::regclass);


--
-- Data for Name: category; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.category VALUES
	(5, '0.1.4', 'E-Ktp Reader', '0.1', '2021-06-23 11:02:42.709662', '-', '-', '-'),
	(7, '0.1.6', 'Signpad', '0.1', '2021-06-23 11:02:47.816475', '-', '-', '-'),
	(8, '0.1.7', 'Stylus Signpad', '0.1', '2021-06-23 11:02:56.129135', '-', '-', '-'),
	(15, '0.1.9', 'Printer', '0.1', '2021-06-23 11:03:23.712066', '-', '-', '-'),
	(23, '0.1.10', 'PC', '0.1', '2021-06-23 11:03:32.051304', '-', '-', '-'),
	(25, '0.1.12', 'QR code scanner', '0.1', '2021-06-23 11:03:39.280013', '-', '-', '-'),
	(27, '0.1.14', 'NFC reader', '0.1', '2021-06-23 11:03:46.178521', '-', '-', '-'),
	(17, '0.1.9.2', 'Laser', '0.1.9', '2021-06-23 11:04:23.208319', '-', '-', '-'),
	(18, '0.1.9.3', 'passbook', '0.1.9', '2021-06-23 11:04:26.342223', '-', '-', '-'),
	(19, '0.1.9.4', 'papper', '0.1.9', '2021-06-23 11:04:29.197017', '-', '-', '-'),
	(20, '0.1.9.5', 'ink', '0.1.9', '2021-06-23 11:04:32.700257', '-', '-', '-'),
	(21, '0.1.9.6', 'ribbon', '0.1.9', '2021-06-23 11:04:35.540297', '-', '-', '-'),
	(22, '0.1.9.7', 'etc related', '0.1.9', '2021-06-23 11:04:38.49505', '-', '-', '-'),
	(6, '0.1.5', 'Fingerprint', '0.1', '2021-06-23 11:02:42.709664', '-', '-', '-'),
	(28, '0.2', 'Software', '0', '2021-06-23 11:06:56.427154', '-', '-', '-'),
	(29, '0.2.1', 'login', '0.2', '2021-06-23 11:07:03.21756', '-', '-', '-'),
	(14, '0.1.3.1', 'etc related', '0.1.3', '2021-06-23 13:51:54.146', '-', '-', '-'),
	(13, '0.1.3.1.1', 'card transporter ', '0.1.3.1', '2021-06-23 13:51:54.175', '-', '-', '-'),
	(12, '0.1.3.1.2', 'card destroyer ', '0.1.3.1', '2021-06-23 13:51:54.176', '-', '-', '-'),
	(30, '0.2.2', 'reset password', '0.2', '2021-06-23 11:07:08.041488', '-', '-', '-'),
	(31, '0.2.3', 'update patch', '0.2', '2021-06-23 11:07:12.511112', '-', '-', '-'),
	(32, '0.2.4', 'Setting koneksi jaringan', '0.2', '2021-06-23 11:07:18.892253', '-', '-', '-'),
	(33, '0.2.5', 'UI', '0.2', '2021-06-23 11:07:28.061567', '-', '-', '-'),
	(37, '0.2.9', 'passbook print', '0.2', '2021-06-23 11:07:39.822228', '-', '-', '-'),
	(38, '0.2.10', 'Lain-lain', '0.2', '2021-06-23 11:07:42.559173', '-', '-', '-'),
	(11, '0.1.3.1.3', 'card reader ', '0.1.3.1', '2021-06-23 13:51:54.177', '-', '-', '-'),
	(10, '0.1.3.1.4', 'card hover ', '0.1.3.1', '2021-06-23 13:51:54.178', '-', '-', '-'),
	(9, '0.1.3.1.5', 'Card Dispenser ', '0.1.3.1', '2021-06-23 13:51:54.179', '-', '-', '-'),
	(55, '0.2.5.1', 'Button', '0.2.5', '2021-07-07 16:49:50.948', '-', '-', '-'),
	(16, '0.2.6.1', 'thermalll', '0.2.6', '2021-07-30 08:31:05.059', '-', '-', '-'),
	(26, '0.1.15', 'Bill Acceptor', '0.1', '2021-06-23 12:35:21.36', '-', '-', '-'),
	(52, '0.2.6.3', 'test', '0.2.6', '2021-08-07 16:16:15.228', 'Ini yang pertama', 'Ini yang kedua', 'Ini yang ketiga'),
	(57, '0.2.5.3', 'Keyboard', '0.2.5', '2021-08-07 17:42:45.161', 'Layout keyboard berantakan', '-', '-'),
	(34, '0.2.6', 'Ganti Kartu', '0.2', '2021-06-23 11:07:31.286907', '-', '-', '-'),
	(1, '0.1', 'Hardware', '0', '2021-06-23 11:02:01.231047', '-', '-', '-'),
	(2, '0.1.1', 'Casing', '0.1', '2021-06-23 11:02:19.236979', '-', '-', '-'),
	(3, '0.1.2', 'Monitor', '0.1', '2021-06-23 11:02:30.223515', '-', '-', '-'),
	(4, '0.1.3', 'EDC', '0.1', '2021-06-23 11:02:38.890407', '-', '-', '-');


--
-- Data for Name: role; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.role VALUES
	(2, 'USER'),
	(3, 'SUPERADMIN'),
	(1, 'TEKNISI'),
	(7, 'THIRD PARTY');


--
-- Data for Name: task_list; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.task_list VALUES
	(1, '9J1-F8A-A6G3', 'test', 'test', 'test', 'test', 'test', 'test', 'test', '2021-09-22 09:15:34.541872', '1'),
	(7, 'N9K-F8A-U6D8', 'ektpekspor.jpg', 'test', 'coba', '2', '2', '17', 'Open', '2021-09-29 10:31:51.03632', '1'),
	(9, 'N9K-F8A-U6D8', 'bri-gold-mastercard.jpg', 'test', 'coba', '2', '2', '17', 'Open', '2021-10-01 11:04:39.552876', '2'),
	(13, 'N9K-F8A-U6D8', 'index.png', 'test', 'coba', '2', '2', '17', 'Open', '2021-10-04 11:02:42.55872', '3'),
	(20, 'N9K-F8A-U6D8', 'BRIPLATINUM-FRONT.png', 'test', 'coba', '2', '2', '17', 'Open', '2021-10-06 10:14:11.421096', '0'),
	(21, 'N9K-F8A-U6D9', 'BRIPLATINUM-FRONT.png', 'test', 'coba', '2', '2', '17', 'Open', '2021-10-06 10:14:25.334833', '0'),
	(22, 'N9K-F8A-U6D9', 'barcode(1).png', 'test', 'coba', '2', '2', '17', 'Open', '2021-10-07 09:39:07.433107', '0'),
	(24, 'N9K-F8A-U6D9', 'barcode(1).png', 'hari ini', 'coba', '2', '2', '17', 'Open', '2021-10-07 09:42:37.993666', '0'),
	(25, 'N9K-F8A-U6D9', '-', 'hari ini', 'coba', '2', '2', '17', 'Open', '2021-10-11 15:51:49.844768', '0'),
	(26, '', '-', 'test', 'coba', '0', '0', '17', 'Open', '2021-10-11 16:05:49.941846', ''),
	(27, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:05:56.645723', ''),
	(28, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:06:10.924008', '0'),
	(29, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:08:45.020828', '0'),
	(30, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:08:50.832634', '0'),
	(31, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:08:54.905581', '0'),
	(32, '', '-', 'test', 'coba', '', '', '17', 'Open', '2021-10-11 16:29:30.105334', '0'),
	(33, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:31:36.038924', '1'),
	(34, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:32:07.288605', '1'),
	(35, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:33:38.543302', '1'),
	(36, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:36:20.787443', '1'),
	(37, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:38:59.583761', '1'),
	(38, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:39:02.60486', '1'),
	(39, '', '-', 'Membuat Ticket Baru', 'Create_ticket', '0', '0', '17', 'Unassigned', '2021-10-11 16:39:23.75622', '1');


--
-- Data for Name: team; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.team VALUES
	(1, 'Team A', '2021-09-23 20:24:14.832167'),
	(2, 'Team B', '2021-09-23 20:24:22.65634');


--
-- Data for Name: ticket; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.ticket VALUES
	(26, 'TEs321', '17', '17', 'Critical', '2021-10-11 14:12:53.079808', '2021-10-11 14:12:53.079808', '00:00:00', 'Unassigned', '002-002-002', '', 'Tebet', 'T9999', 'tesmail@mail.com', 'Unassigned', '2'),
	(27, 'test judul 123', 'Adia', 'Adia', 'Critical', '2021-10-11 14:22:23.051192', '2021-10-11 14:22:23.051192', '00:00:00', 'On Hold', '856-544-9092', '1', '12qewqew', 't003', 'test@mail.com', '2', '1'),
	(28, 'Tes3232', '17', '17', 'Medium', '2021-10-11 15:58:51.102921', '2021-10-11 15:58:51.102921', '00:00:00', 'Unassigned', '002-002-003', '55', 'Tebet', 'T999999', 'tes@mail.com', 'Unassigned', '1'),
	(14, 'pembayaran lewat edc', 'adminer', 'adminer', 'Medium', '2021-10-29 16:37:43.678', '2021-10-09 02:08:00.702', '23:26:57', 'Waiting Reply', 'A76-589-55C8', '3', 'Pimmm', 't707', 'shop@mail.com', 'adminer', '2'),
	(10, 'Tidak dapat melakukan pembayan Debit dan Kredit via EDC', 'Budi', 'Anto', 'Critical', '2021-10-29 14:35:06.156', '2021-10-09 02:08:33.021', '00:00:00', 'Resolved', 'AH5-509-YHC8', '36', 'Rawamangun', 't003', 'test@mail.com', 'Unassigned', '1'),
	(12, 'monitor tidak bergerak', 'adminer', 'adminer', 'Medium', '2021-10-29 16:19:57.476', '2021-10-29 16:19:57.476', '00:00:00', 'New', '9H3-509-KW2M', '1', 'alia', 't99', 'test@mail.com', 'adminer', '2'),
	(3, 'Akuu', 'Andi', 'adminer', 'Critical', '2021-10-04 10:21:09.719', '2021-10-09 13:09:23.512', '97d 2h 48m 13s', 'Resolved', 'N9K-F8A-U6D8', '5', 'GATSU', 't001', 'test2@mail.com', 'nmulyana123', '2'),
	(19, 'test judul', 'Anton', 'Anton', 'Critical', '2021-10-24 11:33:53.353232', '2021-10-24 11:33:53.353232', '00:00:00', 'Unassigned', '856-544-78US', '1', 'Tebet', 't003', 'test@mail.com', '17', '1'),
	(13, 'Ini judul', 'adminer', 'adminer', 'Medium', '2021-10-29 16:24:17.703', '2021-10-09 13:51:29.718', '40d 21h 27m 12s', 'Unassigned', '9J1-F8A-A6G3', '9', 'PIM', 't8080', 'shop@gmail.com', 'Unassigned', '3'),
	(17, 'Berikut adalah judulnya', 'adminer', 'adminer', 'Medium', '2021-10-09 03:43:17.618', '2021-10-09 12:51:37.109', '09:8:19', 'Uninvestigated', '454-544-44B3', '52', 'Kebayoran', 't0000999', 'dikaadia@gmail.com', 'Unassigned', '2'),
	(15, 'Ini adalah judul dari ticket ini', 'adminer', 'adminer', 'Medium', '2021-10-16 14:45:14.911', '2021-10-16 14:45:14.911', '00:00:00', 'Investigated', 'DD8-D87-874A', '2', 'Semarang ', 't0009', 'smg@gmail.com', 'finnet', '1'),
	(16, 'payment melalui EDC terjadi backout', 'adminer', 'adminer', 'Low', '2021-10-02 09:56:08.762', '2021-10-09 02:09:14.678', '00:00:00', 'Process', '8F8-F8A-8AE5', '12', 'assd', 'asd', 'mail@gmail.com', 'finnet', '1'),
	(5, 'Kendala MyG grapari indramayu sering restart otomatis', 'Kale', 'adminer', 'High', '2021-10-21 10:25:47.171', '2021-10-09 02:08:52.93', '00:42:23', 'Closed', 'LP1-544-8CO2', '12', 'Rawamangun', 't002', 'coba@mail.com', 'adminer', '1'),
	(6, 'CPU Konfirm Restart', 'Rawamangun', 'Rawamangun', 'Low', '2021-10-23 14:30:17.336', '2021-10-09 02:07:35.669', '00:00:00', 'Finish', '856-544-55B3', '38', 'Rawamangun', 't003', 'test@mail.com', 'Unassigned', '2'),
	(29, 'Test432', '17', '17', 'High', '2021-10-11 16:01:02.568896', '2021-10-11 16:01:02.568896', '00:00:00', 'Unassigned', '002-002-004', '13', 'Tebet', 'T999999', 'tes@mail.com', 'Unassigned', '1'),
	(30, 'test543', '17', '17', 'Low', '2021-10-11 16:02:52.864485', '2021-10-11 16:02:52.864485', '00:00:00', 'Unassigned', '002-002-005', '5', 'Tebet', 'T999999', 'email@mail.com', 'Unassigned', '1'),
	(31, 'teset123', '17', '17', 'Critical', '2021-10-11 16:04:26.27581', '2021-10-11 16:04:26.27581', '00:00:00', 'Unassigned', '002-002-006', '26', 'Tebet', 'T999999', 'mail@mail.com', 'Unassigned', '1'),
	(32, 'asd', '17', '17', 'Low', '2021-10-11 16:12:22.573756', '2021-10-11 16:12:22.573756', '00:00:00', 'Unassigned', '002-002-007', '2', 'Tebet', 'T999999', 'mail.com@mail.com', 'Unassigned', '1'),
	(24, 'test judul 123', 'Adia', 'Adia', 'Critical', '2021-10-11 14:09:29.89405', '2021-10-11 14:09:29.89405', '00:00:00', 'On Hold', '856-544-9091', '1', '12qewqew', 't003', 'test@mail.com', '2', '1'),
	(33, 'sdf', '17', '17', 'Low', '2021-10-11 16:14:49.729417', '2021-10-11 16:14:49.729417', '00:00:00', 'Unassigned', '002-002-008', '4', 'Tebet', 'T999999', 'mail@mail.com', 'Unassigned', '1'),
	(25, 'Tes123', '17', '17', 'Low', '2021-10-11 14:11:17.831806', '2021-10-11 14:11:17.831806', '00:00:00', 'Unassigned', '001-001-001', '', 'Tebet', 'T0009', 'tes@mail.com', 'Unassigned', '1');


--
-- Data for Name: ticket_isi; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.ticket_isi VALUES
	(27, 'adminer', 'Ini adalah balasan

Terimakasih', '', '2021-07-28 11:08:10.549', 'LP1-544-8CO2', ''),
	(1, 'Andi', 'mohon bantuan, roll penarikan dispenser kartu tidak keluar sehingga perso tidak dapat dilakukan.', '', '2021-05-04 10:21:09.719', 'N9K-F8A-U6D8', NULL),
	(2, 'Trilogi Helpdesk', 'Dear Rekan,

Untuk kendala diatas apakah masih error code 10?

apakah bisa di eject to front?

Thanks & Regards,
Support Trilogi', '', '2021-05-04 10:57:54', 'N9K-F8A-U6D8', NULL),
	(3, 'Andi', 'Dear pak yokka,

sudah dicoba dari kemarin untuk eject to front bin dsb, tetap sama tidak bergerak dispenser tersebut.
dengan keterangan terlampir.

terima kasih
Regards,

ANTON SUHARTO
082211010046
Team Leader GraPARI Shop Bandung 2 – Dago', '', '2021-05-04 10:58:48.779', 'N9K-F8A-U6D8', NULL),
	(14, 'adminer', 'Ini adalah balasan dari pesan diatas', 'C:\monitoring_ticket_myg\file\/20210504102109719/ini_file_yang_pertama.jpg', '2021-07-06 08:27:47.157', 'N9K-F8A-U6D8', 'C:\monitoring_ticket_myg\file\/20210504102109719/ini_file_yang_kedua.xlsx'),
	(26, 'adminer', 'Sedangkan ini adalah ini dari ticketnya

Dan ini adalah baris kedua', 'C:\monitoring_ticket_myg\file\/DD8-D87-874A/144455.jpg', '2021-07-16 14:45:14.911', 'DD8-D87-874A', 'C:\monitoring_ticket_myg\file\/DD8-D87-874A/144509.png'),
	(28, 'adminer', 'as', '', '2021-08-02 09:56:08.762', '8F8-F8A-8AE5', ''),
	(29, 'adminer', 'Ini yang pertama: Jawaban yang pertama
      Ini yang kedua: Jawaban yang kedua
      Ini yang ketiga: Jawaban yang ketiga
      Mohon bantuannya untuk segera di proses

Adia', '', '2021-08-09 03:43:17.618', '454-544-44B3', ''),
	(6, 'Rawamangun', 'test isi', 'test attachment', '2021-06-23 14:30:17.336', '856-544-55B3', NULL),
	(13, 'adminer', 'adsfthrtyjtj', 'C:\monitoring_ticket_myg\file\/20210629163743678/163728.pdf', '2021-06-29 16:37:43.678', 'A76-589-55C8', 'C:\monitoring_ticket_myg\file\/20210629163743678/163741.jpg'),
	(25, 'adminer', 'kuy', '', '2021-07-13 16:04:41.191', 'A76-589-55C8', ''),
	(10, 'Rawamangun', 'test isi', '', '2021-06-29 14:35:06.156', 'AH5-509-YHC8', 'C:\monitoring_ticket_myg\file\/2021-06-29/Adiaaaa.pdf'),
	(11, 'adminer', 'Isiiiiii', '', '2021-06-29 16:19:57.476', '9H3-509-KW2M', ''),
	(5, 'Kale', 'Shop : Grapari Mitra Boyolali
Tanggal trx : 22 maret 2021
MID EDC : 000100205043524
TID EDC : 05352499
SN EDC : 163237333221029001232378
nominal trx : 110.000 (628112658004 abdul )
no rekening : 900-002-398-0841
bank issuer : mandiri
no debit/cc : 4616993298394960

Lampiran : Struk, foto kartu debit / cc, bukti mutasi/bukti potong sal

atas support dan bantuanya kami ucapkan terima kasih', 'test attachment', '2021-06-21 10:25:47.171', 'LP1-544-8CO2', NULL),
	(12, 'adminer', 'Ini isis', 'C:\monitoring_ticket_myg\file\/2021-06-29/162414', '2021-06-29 16:24:17.703', '9J1-F8A-A6G3', 'C:\monitoring_ticket_myg\file\/2021-06-29/162401'),
	(30, 'adminer', 'Berikut adalah balasannya', '', '2021-08-09 11:51:03.886', '454-544-44B3', ''),
	(40, 'adminer', 'Tiket telah berhasil diselesaikan', '', '2021-08-09 13:09:23.512', 'N9K-F8A-U6D8', ''),
	(41, 'adminer', 'Berikut adalah balasan dari kami', '', '2021-08-09 13:11:20.041', '9J1-F8A-A6G3', ''),
	(46, 'Anton', 'test isi', '-', '2021-09-24 11:33:53.353232', '856-544-78US', '-'),
	(50, 'Adia', 'oooooooo', '-', '2021-10-11 14:09:29.89405', '856-544-9091', '-'),
	(51, '17', 'Tes 123', '-', '2021-10-11 14:11:17.831806', '001-001-001', '-'),
	(52, '17', 'Tes 321', '-', '2021-10-11 14:12:53.079808', '002-002-002', '-'),
	(53, 'Adia', 'oooooooo', '-', '2021-10-11 14:22:23.051192', '856-544-9092', '-'),
	(54, '17', 'Tes3232', '-', '2021-10-11 15:58:51.102921', '002-002-003', '-'),
	(55, '17', 'test432', '-', '2021-10-11 16:01:02.568896', '002-002-004', '-'),
	(56, '17', 'test543', '-', '2021-10-11 16:02:52.864485', '002-002-005', '-'),
	(57, '17', 'teset123', '-', '2021-10-11 16:04:26.27581', '002-002-006', '-'),
	(58, '17', 'asd', '-', '2021-10-11 16:12:22.573756', '002-002-007', '-'),
	(59, '17', 'sdf', '-', '2021-10-11 16:14:49.729417', '002-002-008', '-');


--
-- Data for Name: users; Type: TABLE DATA; Schema: monitoring_maintenance; Owner: postgres
--

INSERT INTO monitoring_maintenance.users VALUES
	(27, 'finnet', 'finnet', 'adia@mail.com', '$2a$10$FTR.9E8Ck5jOER25QWgvQOkMpZepdys5Q9g9FUriw97J2hR4fg1bS', 1, '2021-08-16 11:50:09.251', '2021-08-09 15:59:57.256', '7', '2'),
	(29, 'admin2', 'admin', 'adiamahardika@mail.com', '$2a$10$DyVmDf4wK5Tq0r34XHrUs.P7AZ/Pr1Pk9t.mMYim3NaCylYMfA.JW', NULL, '2021-10-01 16:18:04.140195', '2021-10-01 16:18:04.140195', '3', '1'),
	(17, 'Adia Mahardika', 'adminer', 'nmulyana1802@gmail.com', '$2a$10$.RmOY.SDATv/EMQBXYRG6eUxXvUADBIK2vujWqTa6qZ1W21V8bqCi', 1, '2021-10-06 11:07:18.615476', '2021-04-21 10:12:22.925', '3', '1');


--
-- Name: kategori_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.kategori_id_seq', 57, true);


--
-- Name: role_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.role_id_seq', 7, true);


--
-- Name: table1_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.table1_id_seq', 33, true);


--
-- Name: task_list_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.task_list_id_seq', 39, true);


--
-- Name: team_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.team_id_seq', 2, true);


--
-- Name: ticket_isi_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.ticket_isi_id_seq', 59, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: monitoring_maintenance; Owner: postgres
--

SELECT pg_catalog.setval('monitoring_maintenance.users_id_seq', 29, true);


--
-- Name: category kategori_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.category
    ADD CONSTRAINT kategori_pkey PRIMARY KEY (id);


--
-- Name: role role_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.role
    ADD CONSTRAINT role_pkey PRIMARY KEY (id);


--
-- Name: ticket table1_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.ticket
    ADD CONSTRAINT table1_pkey PRIMARY KEY (id);


--
-- Name: task_list task_list_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.task_list
    ADD CONSTRAINT task_list_pkey PRIMARY KEY (id);


--
-- Name: team team_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.team
    ADD CONSTRAINT team_pkey PRIMARY KEY (id);


--
-- Name: ticket_isi ticket_isi_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.ticket_isi
    ADD CONSTRAINT ticket_isi_pkey PRIMARY KEY (id);


--
-- Name: users uk6dotkott2kjsp8vw4d0m25fb7; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.users
    ADD CONSTRAINT uk6dotkott2kjsp8vw4d0m25fb7 UNIQUE (email);


--
-- Name: users ukr43af9ap4edm43mmtq01oddj6; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.users
    ADD CONSTRAINT ukr43af9ap4edm43mmtq01oddj6 UNIQUE (username);


--
-- Name: users username_uniqe; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.users
    ADD CONSTRAINT username_uniqe UNIQUE (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: monitoring_maintenance; Owner: postgres
--

ALTER TABLE ONLY monitoring_maintenance.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

