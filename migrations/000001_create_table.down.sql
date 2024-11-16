DROP TRIGGER IF EXISTS check_relation_id ON notifications CASCADE;
DROP FUNCTION IF EXISTS validate_relation_id() CASCADE;

DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS bids CASCADE;
DROP TABLE IF EXISTS tenders CASCADE;
DROP TABLE IF EXISTS users CASCADE;

DO $$ 
BEGIN 
    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        DROP TYPE role_enum CASCADE;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_tender') THEN
        DROP TYPE status_tender CASCADE;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_bids') THEN
        DROP TYPE status_bids CASCADE;
    END IF;

    IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_notification') THEN
        DROP TYPE type_notification CASCADE;
    END IF;
END $$;
