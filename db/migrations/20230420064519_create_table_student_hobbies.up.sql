CREATE TABLE student_hobbies(
                        id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        student_id int,
                        hobby_id int,
                        FOREIGN KEY (student_id) REFERENCES students(id),
                        FOREIGN KEY (hobby_id) REFERENCES hobbies(id)
)ENGINE = InnoDB;