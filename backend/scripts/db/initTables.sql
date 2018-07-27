CREATE TABLE spenderDB.User (
  ID int(11) NOT NULL AUTO_INCREMENT,
  Email varchar(50) NOT NULL,
  Password varchar(255) NOT NULL,
  Verified bool DEFAULT FALSE,
  CONSTRAINT fk_customer_pkey PRIMARY KEY (ID)
)

CREATE TABLE spenderDB.Spend (
  ID int(11) NOT NULL AUTO_INCREMENT,
  Amount double(8,2) NOT NULL,
  Currency varchar(20) NOT NULL,
  UserId int(11)  NOT NULL,
  Date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT fk_spend_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_customer_id FOREIGN KEY (UserId) REFERENCES spenderDB.User (ID) MATCH SIMPLE ON UPDATE NO ACTION ON DELETE NO ACTION
)