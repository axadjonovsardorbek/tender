DO $$ 
BEGIN 
    -- Drop ENUM types if they exist
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

-- Drop tables if they exist
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS bids CASCADE;
DROP TABLE IF EXISTS tenders CASCADE;
DROP TABLE IF EXISTS users CASCADE;

-- Drop trigger functions if they exist
DROP FUNCTION IF EXISTS validate_relation_id CASCADE;
DROP FUNCTION IF EXISTS check_tender_status CASCADE;
DROP FUNCTION IF EXISTS check_bid_limit CASCADE;
DROP FUNCTION IF EXISTS prevent_status_change_after_awarded CASCADE;

-- Drop triggers if they exist
DROP TRIGGER IF EXISTS check_relation_id ON notifications CASCADE;
DROP TRIGGER IF EXISTS enforce_tender_status ON bids CASCADE;
DROP TRIGGER IF EXISTS enforce_bid_limit ON bids CASCADE;
DROP TRIGGER IF EXISTS enforce_awarded_status ON tenders CASCADE;
