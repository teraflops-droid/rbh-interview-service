CREATE TABLE tb_card (
                         id INT AUTO_INCREMENT PRIMARY KEY,                        -- Auto-incrementing primary key
                         title VARCHAR(100),                                            -- Title field
                         description TEXT,                                              -- Description field
                         status VARCHAR(20),       -- Status as ENUM with default value
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                 -- CreatedAt with default current timestamp
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- UpdatedAt with auto-update on modification
                         created_by VARCHAR(100),                                       -- CreatedBy field
                         updated_by VARCHAR(100)                                       -- UpdatedBy field
);
CREATE TABLE tb_comment (
                            id INT AUTO_INCREMENT PRIMARY KEY,  -- Auto-incrementing primary key
                            description TEXT,                           -- Description field
                            create_by VARCHAR(100),                     -- CreateBy as varchar with a length of 100
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- CreatedAt with default current timestamp
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- UpdatedAt with auto-update on modification
                            card_id INT,                                -- Foreign key to Card (card_id)
                            FOREIGN KEY (card_id) REFERENCES tb_card(id) ON UPDATE CASCADE ON DELETE SET NULL  -- Foreign key constraint
);
CREATE TABLE tb_user (
                         username VARCHAR(100) PRIMARY KEY,         -- username is the primary key
                         password VARCHAR(200) NOT NULL,            -- password field with varchar(200)
                         is_active BOOLEAN NOT NULL                -- is_active field as boolean
);

CREATE TABLE tb_user_role (
                              id VARCHAR(36) PRIMARY KEY,         -- user_role_id is the primary key with UUID format
                              role VARCHAR(10) NOT NULL,                 -- role field with varchar(10)
                              username VARCHAR(100) NOT NULL,            -- username linked to the user entity
                              CONSTRAINT fk_user_role_username FOREIGN KEY (username)
                                  REFERENCES tb_user(username)           -- Foreign key reference to tb_user
                                  ON UPDATE CASCADE
                                  ON DELETE CASCADE
);
