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
    deadline TIMESTAMP NOT NULL CHECK (deadline > NOW()),
    budget BIGINT NOT NULL CHECK (budget > 0),
    status status_tender DEFAULT 'open',
    file_url VARCHAR(255),
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

-- Tenderning statusi "open" ekanligini tekshiruvchi funksiya
CREATE OR REPLACE FUNCTION check_tender_status()
RETURNS TRIGGER AS $$
BEGIN
    -- Tenderning statusini tekshirish
    IF NOT EXISTS (
        SELECT 1
        FROM tenders
        WHERE id = NEW.tender_id AND status = 'open'
    ) THEN
        RAISE EXCEPTION 'Bid cannot be created because the tender status is not open';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


-- Triggerni yaratish
CREATE TRIGGER enforce_tender_status
BEFORE INSERT ON bids
FOR EACH ROW
EXECUTE FUNCTION check_tender_status();


-- Trigger Function to Check Bid Limit
CREATE OR REPLACE FUNCTION check_bid_limit()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the contractor has already created 5 bids in the last minute
    IF (
        SELECT COUNT(*) 
        FROM bids 
        WHERE contractor_id = NEW.contractor_id 
          AND created_at > NOW() - INTERVAL '1 minute'
    ) >= 5 THEN
        RAISE EXCEPTION 'A contractor can only create up to 5 bids in one minute.';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the Trigger
CREATE TRIGGER enforce_bid_limit
BEFORE INSERT ON bids
FOR EACH ROW
EXECUTE FUNCTION check_bid_limit();


-- Trigger funksiyasini yaratish
CREATE OR REPLACE FUNCTION prevent_status_change_after_awarded()
RETURNS TRIGGER AS $$
BEGIN
    -- Agar eski status 'awarded' bo'lsa va yangi status boshqa holat bo'lsa, xato qaytar
    IF OLD.status = 'awarded' AND NEW.status != 'awarded' THEN
        RAISE EXCEPTION 'Status cannot be changed once it is set to "awarded"';
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Triggerni yaratish
CREATE TRIGGER enforce_awarded_status
BEFORE UPDATE ON tenders
FOR EACH ROW
EXECUTE FUNCTION prevent_status_change_after_awarded();
