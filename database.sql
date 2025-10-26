CREATE TABLE config (
    key   INTEGER PRIMARY KEY,
    value JSONB NOT NULL
);

CREATE TABLE application (
    id         SERIAL PRIMARY KEY,
    secret_id  TEXT   NOT NULL UNIQUE,
    secret_key TEXT   NOT NULL UNIQUE,
    name       TEXT   NOT NULL
);

CREATE TABLE permission (
    id            SERIAL  PRIMARY KEY ,
    application   INTEGER NOT NULL REFERENCES application(id) ON DELETE CASCADE,
    interface     TEXT    NOT NULL,
    disable_key   BOOLEAN NOT NULL,
    disable_token BOOLEAN NOT NULL,
    blacklist     JSONB,
    whitelist     JSONB,
    qps           BIGINT,
    qpm           BIGINT
);

CREATE TABLE package (
    id             SERIAL  PRIMARY KEY,
    application    INTEGER NOT NULL REFERENCES application(id) ON DELETE CASCADE,
    interface      TEXT    NOT NULL,
    total          BIGINT  NOT NULL CHECK (total >= 0),
    used           BIGINT  NOT NULL CHECK (used >= 0) DEFAULT 0,
    unlimit        BOOLEAN NOT NULL DEFAULT false,
    priority       INTEGER NOT NULL,
    available_from BIGINT  NOT NULL,
    available_to   BIGINT  NOT NULL CHECK (available_from < available_to),
    CONSTRAINT chk_used_le_total CHECK (used <= total)
);

CREATE TABLE links (
    id       BIGSERIAL PRIMARY KEY,
    link     TEXT NOT NULL,
    validity BIGINT DEFAULT NULL
);

CREATE INDEX CONCURRENTLY idx_links_link_hash ON links USING hash(link);
CREATE INDEX CONCURRENTLY idx_links_validity ON links (validity) WHERE validity IS NOT NULL;
CREATE INDEX CONCURRENTLY idx_links_link_validity ON links (link, validity);
CREATE INDEX CONCURRENTLY idx_links_valid_null ON links (id) WHERE validity IS NULL;

CREATE INDEX CONCURRENTLY idx_permission_app_interface ON permission (application, interface);
CREATE INDEX CONCURRENTLY idx_permission_interface ON permission (interface);
CREATE INDEX CONCURRENTLY idx_permission_rate_limit ON permission (qps, qpm) WHERE qps IS NOT NULL OR qpm IS NOT NULL;

CREATE INDEX CONCURRENTLY idx_package_app_time ON package (application, available_from, available_to);
CREATE INDEX CONCURRENTLY idx_package_priority_available ON package (priority, available_from, available_to) WHERE unlimit = false;
CREATE INDEX CONCURRENTLY idx_package_remaining ON package ((total - used), priority) WHERE unlimit = false AND total > used;

DELETE FROM links WHERE validity < EXTRACT(EPOCH FROM NOW());