CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL UNIQUE,
  "email" varchar NOT NULL UNIQUE,
  "hashed_password" varchar NOT NULL,
  "favourites" int[],
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL
);

CREATE TABLE "product" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "category_id" int NOT NULL,
  "photo" bytea NOT NULL,
  "price" int NOT NULL,
  "shop_id" int NOT NULL,
  "raiting" int
);

CREATE TABLE "category" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "shop" (
  "id" serial PRIMARY KEY,
  "name" varchar NOT NULL,
  "products" int[],
  "owner" int NOT NULL
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


ALTER TABLE "product" ADD FOREIGN KEY ("category_id") REFERENCES "category" ("id");

ALTER TABLE "product" ADD FOREIGN KEY ("shop_id") REFERENCES "shop" ("id");

ALTER TABLE "orders" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

ALTER TABLE "shop" ADD FOREIGN KEY ("owner") REFERENCES "users" ("id");