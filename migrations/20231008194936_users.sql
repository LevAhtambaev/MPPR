-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users (
   "vk_id" integer not null primary key,
   "state" varchar(255) not null,
    city_driving BOOLEAN default false,
    highway_driving BOOLEAN default false,
    fuel_efficient BOOLEAN default false,
    spacious BOOLEAN default false,
    powerful BOOLEAN default false,
    easy_to_drive BOOLEAN default false,
    comfortable BOOLEAN default false,
    sporty BOOLEAN default false,
    modern_design BOOLEAN default false,
    high_end_luxury BOOLEAN default false,
    environmentally_friendly BOOLEAN default false,
    fwd BOOLEAN default false,
    rwd BOOLEAN default false,
    awd BOOLEAN default false,
    rear_view_camera BOOLEAN default false,
    tire_pressure_monitoring BOOLEAN default false,
    automatic_parking BOOLEAN default false,
    navigation_system BOOLEAN default false,
    heated_seats BOOLEAN default false,
    electric_motor BOOLEAN default false
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
