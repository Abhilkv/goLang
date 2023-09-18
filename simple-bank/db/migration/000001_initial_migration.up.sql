-- Create Enum Type (Currency):
CREATE TYPE Currency AS ENUM (
  'USD',
  'EUR',
  'INR'
);

-- Create Accounts Table (A):
CREATE TABLE accounts (
  id bigserial PRIMARY KEY,
  owner varchar NOT NULL,
  balance bigint NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  currency Currency NOT NULL
);

CREATE INDEX idx_owner ON accounts (owner);

-- Create Entries Table:
CREATE TABLE entries (
  id bigserial PRIMARY KEY,
  account_id bigint UNIQUE,
  amount bigint,
  created_at timestamptz DEFAULT now()
);

CREATE INDEX idx_account_id ON entries (account_id);

-- Create Transfers Table:
CREATE TABLE transfers (
  id bigserial PRIMARY KEY,
  from_account_id bigint,
  to_account_id bigint,
  amount bigint NOT NULL CHECK (amount > 0),
  created_at timestamptz DEFAULT now()
);

CREATE INDEX idx_from_account_id ON transfers (from_account_id);
CREATE INDEX idx_to_account_id ON transfers (to_account_id);
CREATE INDEX idx_transfer_accounts ON transfers (from_account_id, to_account_id);

-- Add a comment to the transfers table column:
COMMENT ON COLUMN transfers.amount IS 'Must be positive';

-- Add Foreign Key Constraints:
ALTER TABLE entries ADD FOREIGN KEY (account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD FOREIGN KEY (from_account_id) REFERENCES accounts (id);
ALTER TABLE transfers ADD FOREIGN KEY (to_account_id) REFERENCES accounts (id);
