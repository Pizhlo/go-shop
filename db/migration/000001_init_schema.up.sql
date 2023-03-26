CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "favourites" int[]
);

CREATE TABLE "product" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "category" int NOT NULL,
  "photo" bytea NOT NULL,
  "price" int NOT NULL,
  "shop" int NOT NULL
);

CREATE TABLE "category" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "shop" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "products" int[]
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "products" int[] NOT NULL,
  "quantities" int[] NOT NULL,
  "user" int NOT NULL,
  "sum" int NOT NULL,
  "paid" bool NOT NULL,
  "status" varchar NOT NULL,
  "created" timestamptz DEFAULT (now()) NOT NULL
);


ALTER TABLE "product" ADD FOREIGN KEY ("category") REFERENCES "category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("shop") REFERENCES "shop" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");