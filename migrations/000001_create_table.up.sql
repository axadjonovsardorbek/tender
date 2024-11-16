DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum') THEN
        CREATE TYPE role_enum AS ENUM(
            'client', 
            'contractor'
        );
    END IF;    
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_tender') THEN
        CREATE TYPE status_tender AS ENUM(
            'open', 
            'closed',
            'awarded'
        );
    END IF;    
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_bids') THEN
        CREATE TYPE status_bids AS ENUM(
            'pending', 
            'accepted',
            'rejected'
        );
    END IF;  
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'type_notification') THEN
        CREATE TYPE type_notification AS ENUM(
            'tender', 
            'bid'
        );
    END IF;  
END $$;

SET TIME ZONE 'Asia/Tashkent';

CREATE TABLE IF NOT EXISTS users(
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role role_enum DEFAULT 'client',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at BIGINT NOT NULL DEFAULT 0,
    UNIQUE(username, deleted_at),
    UNIQUE(email, deleted_at)
);

CREATE TABLE IF NOT EXISTS tenders (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    client_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    deadline TIMESTAMP NOT NULL,
    budget BIGINT NOT NULL CHECK (budget > 0),
    status status_tender DEFAULT 'open',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS bids (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    tender_id UUID NOT NULL REFERENCES tenders(id) ON DELETE CASCADE,
    contractor_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    price BIGINT NOT NULL CHECK (price > 0),
    delivery_time INTEGER NOT NULL CHECK (delivery_time > 0), -- kunlarda
    comments TEXT,
    status status_bids DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS notifications (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    message TEXT NOT NULL,
    relation_id UUID,  -- tender yoki bid bilan bog'liq bo'lishi mumkin
    type type_notification NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Trigger Function for Notification Validation
CREATE OR REPLACE FUNCTION validate_relation_id()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.type = 'tender' AND NOT EXISTS (SELECT 1 FROM tenders WHERE id = NEW.relation_id) THEN
        RAISE EXCEPTION 'Invalid relation_id for tender';
    ELSIF NEW.type = 'bid' AND NOT EXISTS (SELECT 1 FROM bids WHERE id = NEW.relation_id) THEN
        RAISE EXCEPTION 'Invalid relation_id for bid';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER check_relation_id
BEFORE INSERT OR UPDATE ON notifications
FOR EACH ROW
EXECUTE FUNCTION validate_relation_id();

