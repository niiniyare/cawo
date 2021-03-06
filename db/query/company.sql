-- name: GetAirlineCompany :one
SELECT * FROM airline_company
WHERE company_id = $1 
LIMIT 1;

-- name: ListAirlineCompany :many
SELECT * FROM airline_company
ORDER BY name;

-- name: CreateAirlineCompany :one
INSERT INTO airline_company (
    company_name ,
    iata_code  ,
    icao_Code ,
    hub_airport,
    website
) VALUES (
  $1, $2, $3 , $4 ,$5
)
RETURNING *;

-- name: DeleteAirlineCompany :exec
DELETE FROM airline_company
WHERE company_id = $1;
/*

INSERT INTO airline_company (
  company_name,
  iata_code,
  main_airport
) VALUES (
 'Turkish Airlines', 'TK', 'IST'
)
RETURNING *;

\copy airports_data(airport_code,airport_name,country,city,coordinates,ACOA,timezone) FROM '/data/data/com.termux/files/home/downloads/airports_data.csv' DELIMITER ','  CSV HEADER;
INSERT INTO airports_data (airport_code, airport_name ,city ,coordinates,timezone)
VALUES ('IST', '{'name' : 'Istanbul Internationa Airport'}','{'name' : 'Istanbul'}', POINT(40.982555, 28.820829), 'Europe/Istanbul')
RETURNING *;

INSERT INTO airports_data (airport_code, airport_name ,city ,coordinates,timezone)
VALUES ('IST', 'Istanbul Internationa Airport', 'Istanbul', POINT(40.982555, 28.820829), 'Europe/Istanbul')
RETURNING *;

*/