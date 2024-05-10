CREATE TABLE IF NOT EXISTS owldinal_nft_mint_box_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    box_id       BIGINT    NOT NULL,
    mint_type    BIGINT    NOT NULL,
    token_uri    TEXT,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user_address (user),
    INDEX idx_box_id (box_id),
    INDEX idx_block_number_mint (block_number)
);

CREATE TABLE IF NOT EXISTS owldinal_nft_transfer_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    from_user    CHAR(42)  NOT NULL,
    to_user      CHAR(42)  NOT NULL,
    box_id       BIGINT    NOT NULL,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_from_user (from_user),
    INDEX idx_to_user (to_user),
    INDEX idx_box_id (box_id),
    INDEX idx_block_number_mint (block_number)
);

CREATE TABLE IF NOT EXISTS gen_one_box_mint_box_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    token_id     BIGINT    NOT NULL,
    box_type     TINYINT   NOT NULL,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user),
    INDEX idx_token_id (token_id)
);

CREATE TABLE IF NOT EXISTS gen_one_box_transfer_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    from_user    CHAR(42)  NOT NULL,
    to_user      CHAR(42)  NOT NULL,
    token_id     BIGINT    NOT NULL,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_from_user (from_user),
    INDEX idx_to_user (to_user),
    INDEX idx_token_id (token_id)
);

CREATE TABLE IF NOT EXISTS owl_game_join_game_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    invite_code  INT       NOT NULL,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_bind_invitation_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    inviter      CHAR(42)  NOT NULL,
    invitee      CHAR(42)  NOT NULL,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_inviter (inviter),
    INDEX idx_invitee (invitee)
);

CREATE TABLE IF NOT EXISTS owl_game_prize_pool_increased_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    amount       DECIMAL(36, 18),
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);

CREATE TABLE IF NOT EXISTS owl_game_prize_pool_decreased_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    amount       DECIMAL(36, 18),
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);

CREATE TABLE IF NOT EXISTS owl_game_request_mint_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    user         CHAR(42)        NOT NULL,
    request_id   BIGINT UNSIGNED NOT NULL,
    count        BIGINT UNSIGNED NOT NULL,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);

CREATE TABLE IF NOT EXISTS owl_game_mint_mystery_box_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    user         CHAR(42)        NOT NULL,
    count        BIGINT UNSIGNED NOT NULL,
    request_id   BIGINT UNSIGNED NOT NULL,
    token_ids    TEXT,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);


CREATE TABLE IF NOT EXISTS owl_game_stake_owldinal_nft_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    token_ids    TEXT,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_unstake_owldinal_nft_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    token_ids    TEXT,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_stake_mystery_box_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    token_ids    TEXT,
    buff_level   TINYINT UNSIGNED,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_unstake_mystery_box_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    token_id     BIGINT    NOT NULL,
    box_type     TINYINT   NOT NULL,
    rewards      DECIMAL(36, 18),
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user),
    INDEX idx_token_id (token_id)
);

CREATE TABLE IF NOT EXISTS owl_game_rebate_rewards_increased_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    user         CHAR(42)        NOT NULL,
    invitee      CHAR(42)        NOT NULL,
    mint_count   BIGINT UNSIGNED NOT NULL,
    amount       DECIMAL(36, 18) NOT NULL,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user),
    INDEX idx_invitee (invitee)
);

CREATE TABLE IF NOT EXISTS owl_game_unlockable_rebate_increased_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    user         CHAR(42)        NOT NULL,
    mint_count   BIGINT UNSIGNED NOT NULL,
    amount       DECIMAL(36, 18) NOT NULL,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_owl_token_burned_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    user         CHAR(42)        NOT NULL,
    mint_count   BIGINT UNSIGNED NOT NULL,
    amount       DECIMAL(36, 18) NOT NULL,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

CREATE TABLE IF NOT EXISTS owl_game_fruit_reward_update_events
(
    id                BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash           CHAR(66)        NOT NULL,
    log_index         BIGINT          NOT NULL,
    block_number      BIGINT          NOT NULL,
    block_hash        CHAR(66)        NOT NULL,
    count             BIGINT UNSIGNED NOT NULL,
    total_fruit_count BIGINT UNSIGNED NOT NULL,
    total_elf_count   BIGINT UNSIGNED NOT NULL,
    amount            DECIMAL(36, 18) NOT NULL,
    created_at        TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);

CREATE TABLE IF NOT EXISTS owl_game_elf_reward_update_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)        NOT NULL,
    log_index    BIGINT          NOT NULL,
    block_number BIGINT          NOT NULL,
    block_hash   CHAR(66)        NOT NULL,
    count        BIGINT UNSIGNED NOT NULL,
    amount       DECIMAL(36, 18) NOT NULL,
    created_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP       NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index)
);

CREATE TABLE IF NOT EXISTS owl_game_rebate_claimed_events
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    tx_hash      CHAR(66)  NOT NULL,
    log_index    BIGINT    NOT NULL,
    block_number BIGINT    NOT NULL,
    block_hash   CHAR(66)  NOT NULL,
    user         CHAR(42)  NOT NULL,
    amount       DECIMAL(36, 18),
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (tx_hash, log_index),
    INDEX idx_user (user)
);

--

CREATE TABLE owldinal_nft_tokens
(
    id           BIGINT AUTO_INCREMENT PRIMARY KEY,
    token_id     BIGINT    NOT NULL,
    owner        VARCHAR(42),
    token_uri    TEXT,
    is_staking   BOOLEAN,
    staking_time TIMESTAMP      DEFAULT NULL,
    buffing_ids  TEXT,
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (token_id),
    INDEX idx_token_id (token_id),
    INDEX idx_owner (owner)
);

CREATE TABLE mystery_box_tokens
(
    id              BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    token_id        BIGINT UNSIGNED  NOT NULL,
    owner           VARCHAR(42),
    box_type        TINYINT UNSIGNED NOT NULL,
    is_staking      BOOLEAN,
    staking_time    TIMESTAMP             DEFAULT NULL,
    buff_level      TINYINT UNSIGNED,
    current_rewards DECIMAL(36, 18),
    total_rewards   DECIMAL(36, 18),
    claimed_rewards DECIMAL(36, 18),
    created_at      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at      TIMESTAMP        NULL DEFAULT NULL,
    INDEX idx_token_id (token_id),
    INDEX idx_owner (owner)
);

CREATE TABLE user_infos
(
    id                  BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    address             VARCHAR(42),
    invite_code         VARCHAR(255),
    buff_level          INT,
    invite_count        INT,
    total_earned        DECIMAL(36, 18),
    unclaimed_referral  DECIMAL(36, 18),
    unlockable_referral DECIMAL(36, 18),
    claimed_referral    DECIMAL(36, 18),
    created_at          TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at          TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at          TIMESTAMP NULL DEFAULT NULL,
    UNIQUE (address)
);

CREATE TABLE invite_relations
(
    id               BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tx_hash          CHAR(66)  NOT NULL,
    log_index        BIGINT    NOT NULL,
    block_number     BIGINT    NOT NULL,
    block_hash       CHAR(66)  NOT NULL,
    inviter          VARCHAR(42),
    invitee          VARCHAR(42),
    referral_rewards DECIMAL(36, 18),
    created_at       TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_inviter (inviter)
);

CREATE TABLE daily_pool_snapshots
(
    id                BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    date              DATETIME,
    total_pool_amount DECIMAL(36, 18),
    allocated_rewards DECIMAL(36, 18),
    total_market_cap  DECIMAL(36, 18),
    total_burn        DECIMAL(36, 18),
    created_at        TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP NULL DEFAULT NULL
);

CREATE TABLE apr_snapshots
(
    id                BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    tx_hash           CHAR(66)        NULL,
    log_index         BIGINT          NULL,
    block_number      BIGINT          NULL,
    block_hash        CHAR(66)        NULL,
    reward_count      BIGINT UNSIGNED NOT NULL,
    total_fruit_count BIGINT UNSIGNED NOT NULL,
    total_elf_count   BIGINT UNSIGNED NOT NULL,
    fruit_rewards     DECIMAL(36, 18) NOT NULL,
    fruit_apr         DOUBLE          NOT NULL,
    fruit_apy         DOUBLE          NOT NULL,
    elf_apr           DOUBLE          NOT NULL,
    elf_apy           DOUBLE          NOT NULL,
    created_at        TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP       NULL DEFAULT NULL
);


CREATE TABLE reward_pool_transaction_records
(
    id           BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user         VARCHAR(42),
    operation    VARCHAR(255),
    description  VARCHAR(255),
    count        INT,
    amount       DECIMAL(36, 18),
    tx_hash      VARCHAR(66),
    log_index    BIGINT UNSIGNED,
    block_number BIGINT UNSIGNED,
    block_hash   VARCHAR(66),
    created_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP      DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP NULL DEFAULT NULL,
    INDEX idx_tx_hash (tx_hash),
    INDEX idx_log_index (log_index),
    INDEX idx_block_number (block_number)
);

CREATE TABLE request_mint_jobs
(
    id                   BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user                 VARCHAR(255)    NOT NULL,
    request_id           BIGINT UNSIGNED NOT NULL,
    count                BIGINT UNSIGNED NOT NULL,
    status               INT UNSIGNED    NOT NULL,
    has_confirmed        BOOLEAN         NOT NULL,
    retry_count          INT UNSIGNED    NOT NULL,
    result               TEXT,
    request_tx_hash      CHAR(66),
    request_log_index    INT UNSIGNED,
    request_block_number BIGINT UNSIGNED,
    request_block_hash   CHAR(66),
    job_tx_hash          CHAR(66),
    job_log_index        INT UNSIGNED,
    job_block_number     BIGINT UNSIGNED,
    job_block_hash       CHAR(66),
    created_at           TIMESTAMP            DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP            DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at           TIMESTAMP       NULL DEFAULT NULL
);

CREATE TABLE `transfer_rewards`
(
    `id`                    BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `user`                  VARCHAR(255)     NOT NULL,
    `token_id`              BIGINT UNSIGNED  NOT NULL,
    `box_type`              TINYINT UNSIGNED NOT NULL,
    `rewards`               DECIMAL(36, 18)  NOT NULL,
    `burned_rewards`        DECIMAL(36, 18)  NOT NULL,
    `buff_level`            TINYINT UNSIGNED NOT NULL,
    `moon_boost`            BOOLEAN          NOT NULL,
    `is_confirmed`          BOOLEAN          NOT NULL,
    `status`                INT UNSIGNED     NOT NULL,
    `result`                TEXT,
    `claim_tx_hash`         VARCHAR(66),
    `claim_log_index`       INT UNSIGNED,
    `claim_block_number`    BIGINT UNSIGNED,
    `claim_block_hash`      VARCHAR(66),
    `transfer_tx_hash`      VARCHAR(66),
    `transfer_log_index`    INT UNSIGNED,
    `transfer_block_number` BIGINT UNSIGNED,
    `transfer_block_hash`   VARCHAR(66),
    `burn_tx_hash`          VARCHAR(66),
    `created_at`            TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    `updated_at`            TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`            TIMESTAMP        NULL DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS transfer_multiple_rewards
(
    id                    BIGINT AUTO_INCREMENT PRIMARY KEY,
    user                  CHAR(42)         NOT NULL,
    token_id_list         TEXT,
    claimed_rewards       DECIMAL(36, 18)  NOT NULL,
    elf_rewards           DECIMAL(36, 18)  NOT NULL,
    burned_rewards        DECIMAL(36, 18)  NOT NULL,
    buff_level            TINYINT UNSIGNED NOT NULL,
    moon_boost            BOOLEAN          NOT NULL,
    has_confirmed         BOOLEAN          NOT NULL,
    status                INT UNSIGNED     NOT NULL,
    result                TEXT,
    transfer_tx_hash      CHAR(66),
    transfer_log_index    INT UNSIGNED     NOT NULL,
    transfer_block_number BIGINT UNSIGNED  NOT NULL,
    transfer_block_hash   CHAR(66),
    burn_tx_hash          CHAR(66),
    created_at            TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at            TIMESTAMP        NULL DEFAULT NULL
);

-- initial date
INSERT INTO daily_pool_snapshots (date, total_pool_amount, allocated_rewards, total_market_cap, total_burn)
VALUES ('1970-01-01', 0.0, 0.0, 0.0, 0.0);


INSERT INTO apr_snapshots (reward_count, total_fruit_count, total_elf_count, fruit_rewards, fruit_apr, fruit_apy,
                           elf_apr, elf_apy)
VALUES (0, 0, 0, 0.0, 0.0, 0.0, 0.0, 0.0);


CREATE TABLE block_records
(
    `id`           BIGINT AUTO_INCREMENT PRIMARY KEY,
    `record_type`  TINYINT UNSIGNED NOT NULL,
    `block_height` BIGINT UNSIGNED  NOT NULL,
    `created_at`   TIMESTAMP             DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP             DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`   TIMESTAMP        NULL DEFAULT NULL
);

INSERT INTO block_records (record_type, block_height)
VALUES (0, 8000000),
       (1, 8000000);