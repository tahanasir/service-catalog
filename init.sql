USE service_catalog;

-- Create tables if they don't exist
CREATE TABLE IF NOT EXISTS services (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    description TEXT,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS versions (
    service_id VARCHAR(36) NOT NULL,
    name VARCHAR(32) NOT NULL,
    changelog TEXT,
    PRIMARY KEY (service_id, name),
    FOREIGN KEY (service_id) REFERENCES services(id) ON DELETE CASCADE
);

-- Insert sample data
INSERT INTO services (id, name, description)
VALUES
    ('d97d82c3-b514-4f7b-943b-44686d2a3617', 'Locate Us', 'Description of Locate Us'),
    ('5b9a2437-0834-4f01-82b5-f92b7c79d43e', 'Collect Monday', ''),
    ('f5ed39e5-3726-4f2c-bce6-5c1f858f23c7', 'Contact Us', 'Description of Contact Us'),
    ('d07a8d4f-85dc-45cb-a68a-27e13f5f3f62', 'Contact Us', 'Description of Contact Us 2'),
    ('c1f84934-d26a-4022-bd34-d0e5b73a4a6b', 'FX Rates International', 'Description of FX Rates International'),
    ('eb3ffbab-0cc9-4c1e-8f2c-6b137f1fcdb3', 'FX Rates International', 'Description of FX Rates International Duplicate'),
    ('f0e5a822-d51d-42a5-9b70-cf711f0f1234', 'Notifications', ''),
    ('75be40b0-037f-4e74-a074-fdb8c037784e', 'Notifications', ''),
    ('f0ba092d-7e1e-4d4f-b1cf-bb883f7b071d', 'Priority Services', ''),
    ('c4b40517-674c-4a46-8e9a-11f280fd51d2', 'Reporting', 'Description of Reporting'),
    ('b3f5b1d7-ee10-45a7-b0f8-04e038f6f10b', 'Security', 'Description of Security'),
    ('5dd1257f-d594-47d6-812a-fec4bc776f8f', 'Security', 'Description of Security 2');

INSERT INTO versions (service_id, name, changelog)
VALUES
    ('d97d82c3-b514-4f7b-943b-44686d2a3617', 'v1 - Locate Us', 'Changelog for v1 of Locate Us'),
    ('d97d82c3-b514-4f7b-943b-44686d2a3617', 'v2 - Locate Us', 'Changelog for v2 of Locate Us'),
    ('d97d82c3-b514-4f7b-943b-44686d2a3617', 'v3 - Locate Us', 'Changelog for v3 of Locate Us'),
    ('5b9a2437-0834-4f01-82b5-f92b7c79d43e', 'v1 - Collect Monday', 'Changelog for v1 of Collect Monday'),
    ('5b9a2437-0834-4f01-82b5-f92b7c79d43e', 'v2 - Collect Monday', 'Changelog for v2 of Collect Monday'),
    ('f5ed39e5-3726-4f2c-bce6-5c1f858f23c7', 'v1 - Contact Us', 'Changelog for v1 of Contact Us');
