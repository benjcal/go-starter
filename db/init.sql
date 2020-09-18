\i data/data.sql
\i func/func.sql

SELECT api.add_user('aa@bb.cc', 'qwerty', 'Foo Bar');
SELECT * FROM api.valid_user('aa@bb.cc', 'qwerty');
