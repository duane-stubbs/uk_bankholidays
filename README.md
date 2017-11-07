# UK Bank Holidays
Get UK bank holidays from government gateway API

Download an parse from UK government API

URL: https://www.gov.uk/bank-holidays.json
Database:
CREATE TABLE `bank_holidays` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `description` varchar(50) DEFAULT NULL,
  `event_date` date DEFAULT NULL,
  `country` varchar(25) DEFAULT NULL,
  `entered` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `event_date` (`event_date`,`country`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
