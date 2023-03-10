CREATE TABLE Person(Id SERIAL NOT NULL PRIMARY KEY ,First_Name varchar(250) Not NULL,Last_Name varchar(250) NOT NULL,Age int, Address varchar(250));


INSERT INTO Person(Id,First_Name, Last_Name ,Age,Address) VALUES (Default,'Arbind','Chauhan',25,'Sant Kabir Nagar');
INSERT INTO Person(Id,First_Name, Last_Name ,Age,Address) VALUES (Default,'Meer',' kINoor ',22,'Mh');
INSERT INTO Person(Id,First_Name, Last_Name ,Age,Address) VALUES (Default,'Khupte','Aniket ',28,'Mh');
INSERT INTO Person(Id,First_Name, Last_Name ,Age,Address) VALUES (Default,'Nayan','DB',22,'Goa');
INSERT INTO Person(Id,First_Name, Last_Name ,Age,Address) VALUES (Default,'xyz','NR ',21,'NT');
select * from Person;