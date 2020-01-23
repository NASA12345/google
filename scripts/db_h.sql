/* Creating required Tables */

CREATE TABLE hospital
(
    id                 SERIAL PRIMARY KEY,
    name               VARCHAR(100),
    max_patient_amount INT
);

CREATE TABLE location
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100),
    hospital_id INT,
    FOREIGN KEY (hospital_id) REFERENCES hospital (id)
);

CREATE TABLE patient
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100),
    illness        VARCHAR(200),
    birth_date     DATE,
    location_id    INT NOT NULL,
    FOREIGN KEY (location_id) REFERENCES location (id)
);
/* Inserting data into Hospital Table */

INSERT INTO hospital (name, max_patient_amount)
VALUES ('Fortis Hospital', 3000),
       ('Noble Heart Hospital', 1000);

/* Inserting data into Location Table */

INSERT INTO location (name, hospital_id)
VALUES  ('Laboratory', 1),
        ('Pharmacy', 1),
        ('Isolation Ward', 2),
        ('Registration Desk', 2);

/* Inserting data into Patient Table */

INSERT INTO patient (name, illness, birth_date, location_id)
VALUES ('Anjali', 'cough', '1997-03-04', 1),
       ('Aarushi', 'flu', '1994-11-07', 2),
       ('Pearl', 'Diabetes', '2003-05-12', 3),
       ('Prathamesh', 'Asthama', '1981-12-23', 4),
       ('Ayush', 'Dengue', '2005-09-17', 1),
       ('Arindam', 'Malaria', '2003-12-19', 2),
       ('Mahir', 'Brain fever', '1999-04-06', 3),
       ('Nikhil', 'Cough', '1999-08-09', 4),
       ('Devansh', 'Malaria', '2004-09-12', 1);


