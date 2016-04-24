SELECT pet.name, pet.breed from pet, person_pet, person where
pet.id = person_pet.pet_id and
person_pet.person_id = person.id and
person.first_name = "Michael";