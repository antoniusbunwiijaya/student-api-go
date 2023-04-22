CREATE TABLE students(
                         id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         age INT NOT NULL,
                         gender TINYINT NOT NULL,
                         created_at DATETIME NOT NULL,
                         major_id INT,
                         FOREIGN KEY (major_id) REFERENCES majors(id)
)ENGINE = InnoDB;