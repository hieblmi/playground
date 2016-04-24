INSERT INTO person (id, first_name, last_name, age)
	VALUES 
	(0, "Michael", "Hiebl", 33),
	(1, "Jessica", "McArdle", 30),
	(2, "Stefan", "Leipold", 29),
	(3, "Franklin", "Roosevelt", 101),
	(4, "Barack", "Obama", 47);

INSERT INTO pet (id, name, breed, age, dead)
	VALUES
	(0, "Xavi", "dog", 2, 0),
	(1, "Xena", "cat", 3, 0),
	(2, "Pinky", "parrot", 6, 0),
	(3, "Stinky", "skunk", 4, 0);

INSERT INTO car (id, brand, model, hp)
	VALUES
	(0, "BMW", "M4", 431),
	(1, "Fiat", "Zwickzwack", 2),
	(2, "Lexus", "i340", 300),
	(3, "VW", "Jetta", 180);

INSERT INTO person_pet (person_id, pet_id)
    VALUES 
    (0, 1),
    (0, 3),
    (1, 1),
    (1, 3),
    (3, 1),
    (4, 2);

INSERT INTO person_car (person_id, car_id)
    VALUES 
    (0, 0),
    (1, 1),
    (2, 2),
    (3, 3),
    (4, 3);

