-- create "cells" table
CREATE TABLE `cells` (`record_id` uuid NOT NULL, `data_field_id` uuid NOT NULL, PRIMARY KEY (), CONSTRAINT `cells_records_record` FOREIGN KEY (`record_id`) REFERENCES `records` (`id`) ON DELETE NO ACTION, CONSTRAINT `cells_data_fields_data_field` FOREIGN KEY (`data_field_id`) REFERENCES `data_fields` (`id`) ON DELETE NO ACTION);
-- create "data_fields" table
CREATE TABLE `data_fields` (`id` uuid NOT NULL, PRIMARY KEY (`id`));
-- create "records" table
CREATE TABLE `records` (`id` uuid NOT NULL, PRIMARY KEY (`id`));
