-- public.portfolio definition

-- Drop table

-- DROP TABLE public.portfolio;

CREATE TABLE public.portfolio (
	id VARCHAR(20) NOT NULL DEFAULT unique_rowid(),
	name VARCHAR(50) NOT NULL,
	created_by VARCHAR(20) NOT NULL,
	image VARCHAR(100) NOT NULL,
	created_at TIMESTAMP NOT NULL,
	CONSTRAINT portfolios_pkey PRIMARY KEY (id ASC),
	UNIQUE INDEX portfolios_un (name ASC)
);

INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948548076061491201','portfolio1','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:00'),
	 ('948548082632163329','portfolio2','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:02'),
	 ('948548089273057281','portfolio3','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:04'),
	 ('948548095904186369','portfolio4','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:06'),
	 ('948548102525943809','portfolio5','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:08'),
	 ('948548109106544641','portfolio6','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:10'),
	 ('948548115671941121','portfolio7','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:12'),
	 ('948548122251198465','portfolio8','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:14'),
	 ('948548128825966593','portfolio9','948488886419128321','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:25:16'),
	 ('948548468074872833','portfolio10','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:26:59');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948548474686832641','portfolio11','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:01'),
	 ('948548481329561601','portfolio12','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:03'),
	 ('948548487919042561','portfolio13','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:05'),
	 ('948548494558658561','portfolio14','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:07'),
	 ('948548501133393921','portfolio15','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:09'),
	 ('948548507717435393','portfolio16','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:11'),
	 ('948548514298855425','portfolio17','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:13'),
	 ('948548520865333249','portfolio18','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:15'),
	 ('948548527463792641','portfolio19','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:17'),
	 ('948548534038331393','portfolio20','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:19');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948548540628205569','portfolio21','948489045346189313','https://www.c-sharpcorner.com/blogs/rownumber-function-with-partition-by-clause-in-sql-server1','2024-03-04 09:27:21'),
	 ('948548810090643457','portfolio22','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:44'),
	 ('948548816702996481','portfolio23','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:46'),
	 ('948548823332454401','portfolio24','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:48'),
	 ('948548829938843649','portfolio25','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:50'),
	 ('948548836552015873','portfolio26','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:52'),
	 ('948548843125768193','portfolio27','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:54'),
	 ('948548849691852801','portfolio28','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:56'),
	 ('948548856275894273','portfolio29','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:28:58'),
	 ('948548862849187841','portfolio30','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:00');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948548869437718529','portfolio31','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:02'),
	 ('948548876002754561','portfolio32','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:04'),
	 ('948548882583846913','portfolio33','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:06'),
	 ('948548889202720769','portfolio34','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:08'),
	 ('948548895811305473','portfolio35','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:10'),
	 ('948548902422413313','portfolio36','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:12'),
	 ('948548909057146881','portfolio37','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:14'),
	 ('948548915646038017','portfolio38','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:16'),
	 ('948548922231586817','portfolio39','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:18'),
	 ('948548928834666497','portfolio40','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:20');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948548935404584961','portfolio41','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:22'),
	 ('948548941972799489','portfolio42','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:24'),
	 ('948548948543045633','portfolio43','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:26'),
	 ('948548955111817217','portfolio44','948489047101276161','https://jwt.io/img/pic_logo.svg','2024-03-04 09:29:28'),
	 ('948549343628394497','portfolio45','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:26'),
	 ('948549346929573889','portfolio46','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:28'),
	 ('948549350238846977','portfolio47','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:29'),
	 ('948549353556410369','portfolio48','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:30'),
	 ('948549356866306049','portfolio49','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:31'),
	 ('948549360173318145','portfolio50','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:32');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948549363529121793','portfolio51','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:33'),
	 ('948549366856417281','portfolio52','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:34'),
	 ('948549370216480769','portfolio53','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:35'),
	 ('948549373573103617','portfolio54','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:36'),
	 ('948549376947355649','portfolio55','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:37'),
	 ('948549380288020481','portfolio56','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:38'),
	 ('948549383605977089','portfolio57','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:39'),
	 ('948549386908073985','portfolio58','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:40'),
	 ('948549390260895745','portfolio59','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:41'),
	 ('948549393562927105','portfolio60','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:42');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948549396882227201','portfolio61','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:43'),
	 ('948549400240324609','portfolio62','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:44'),
	 ('948549403600650241','portfolio63','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:45'),
	 ('948549406901698561','portfolio64','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:46'),
	 ('948549410249637889','portfolio65','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:47'),
	 ('948549413538988033','portfolio66','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:48'),
	 ('948549416830959617','portfolio67','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:49'),
	 ('948549420136988673','portfolio68','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:50'),
	 ('948549423448227841','portfolio69','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:51'),
	 ('948549426748784641','portfolio70','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:52');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948549430054518785','portfolio71','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:53'),
	 ('948549433408815105','portfolio72','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:54'),
	 ('948549436759408641','portfolio73','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:55'),
	 ('948549440069992449','portfolio74','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:56'),
	 ('948549443390832641','portfolio75','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:57'),
	 ('948549446726549505','portfolio76','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:58'),
	 ('948549450078486529','portfolio77','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:31:59'),
	 ('948549453445758977','portfolio78','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:00'),
	 ('948549456757030913','portfolio79','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:01'),
	 ('948549460111556609','portfolio80','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:02');
INSERT INTO public.portfolio (id,"name",created_by,image,created_at) VALUES
	 ('948549463428169729','portfolio81','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:03'),
	 ('948549466753236993','portfolio82','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:04'),
	 ('948549470063853569','portfolio83','948489062380503041','https://jwt.io/img/pic_logo.svg','2024-03-04 09:32:05'),
	 ('948549719256399873','portfolio84','948489062399279105','https://jwt.io/img/pic_logo.svg','2024-03-04 09:33:21'),
	 ('948549722544766977','portfolio85','948489062399279105','https://jwt.io/img/pic_logo.svg','2024-03-04 09:33:22'),
	 ('948549850653261825','portfolio86','948489062412845057','https://jwt.io/img/pic_logo.svg','2024-03-04 09:34:01'),
	 ('948549853939531777','portfolio87','948489062412845057','https://jwt.io/img/pic_logo.svg','2024-03-04 09:34:02');
