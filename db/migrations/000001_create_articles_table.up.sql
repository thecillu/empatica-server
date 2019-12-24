CREATE TABLE IF NOT EXISTS articles(
   id int(11) NOT NULL,
   title VARCHAR (100),
   description VARCHAR (100),
   content VARCHAR (300),
   PRIMARY KEY (id)
);


INSERT INTO articles (id, title, description, content)
VALUES(1, "article 1", "description of article 1", "long content of the article 1");

INSERT INTO articles (id, title, description, content)
VALUES(2, "article 2", "description of article 2", "long content of the article 2");

INSERT INTO articles (id, title, description, content)
VALUES(3, "article 3", "description of article 3", "long content of the article 3");