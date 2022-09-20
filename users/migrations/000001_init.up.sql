CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar(100) NOT NULL,
  "username" varchar(45) NOT NULL,
  "hashed_password" varchar(255) NOT NULL,
  "address" text DEFAULT null,
  "phone_number" varchar(14) NOT NULL,
  "level" int DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar(255) NOT NULL,
  "rating" int NOT NULL,
  "author" varchar(255) NOT NULL,
  "stock" int NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "created_by" int
);

CREATE TABLE "user_book" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "book_id" int,
  "due_date" date NOT NULL,
  "charge" int DEFAULT 0,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "user_book" ("user_id", "book_id");

ALTER TABLE "books" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_book" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_book" ADD FOREIGN KEY ("book_id") REFERENCES "books" ("id") ON DELETE CASCADE;
