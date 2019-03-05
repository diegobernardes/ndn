CREATE TYPE input_kind AS ENUM ('http.push');
CREATE TYPE output_kind AS ENUM ('http.pull');

CREATE TABLE input (
  id   BIGSERIAL
  kind input_kind NOT NULL
);

CREATE TABLE output (
  id   BIGSERIAL,
  kind output_kind NOT NULL
);

CREATE TABLE output_http_push_checkpoints (
  id    BIGINT NOT NULL,
  value VARCHAR NOT NULL,

  FOREIGN KEY (id) REFERENCES output (id) ON DELETE CASCADE
);

CREATE TABLE notification (
  id         BIGSERIAL,
  entity     VARCHAR NOT NULL,
  entity_id  VARCHAR NOT NULL,
  version    VARCHAR NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL,
);
