CREATE TABLE foo (bar bool not null, "interval" interval not null);

INSERT INTO foo (bar, "interval")
VALUES (true, interval '5' minute),
       (true, interval '10' minute),
       (true, interval '15' minute),
       (true, interval '1' hour),
       (true, interval '1' day)