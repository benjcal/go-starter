CREATE OR REPLACE FUNCTION api.add_user(email TEXT, pwd TEXT, name TEXT) RETURNS VOID AS
$$
BEGIN
    INSERT INTO data.users(email, pwd_hash, name)
    VALUES (add_user.email, crypt(add_user.pwd, gen_salt('bf', 8)), add_user.name);
END;

$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION api.valid_user(email_in TEXT, pwd TEXT)
    RETURNS TABLE
            (
                id    INT,
                email TEXT,
                name  TEXT
            )
AS
$$
BEGIN
    RETURN QUERY SELECT u.id, u.email, u.name FROM data.users u WHERE u.email = valid_user.email_in AND u.pwd_hash = crypt(valid_user.pwd, u.pwd_hash);
END;
$$ LANGUAGE plpgsql;