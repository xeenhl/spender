CREATE TABLE spenderDB.User (
  ID int(11) NOT NULL AUTO_INCREMENT,
  Email varchar(50) DEFAULT NULL,
  Password varchar(255) DEFAULT NULL,
  Verified bool DEFAULT NULL,
  PRIMARY KEY (ID)
)

CREATE TABLE spenderDB.Spend (
  ID int(11) NOT NULL AUTO_INCREMENT,
  Amount double(8,2) NOT NULL,
  Currency varchar(20) NOT NULL,
  User int(11)  NOT NULL,
  PRIMARY KEY (ID)
)