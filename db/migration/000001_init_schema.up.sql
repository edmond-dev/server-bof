CREATE TABLE "users" (
 "id" bigserial PRIMARY KEY,
 "user_id" varchar UNIQUE NOT NULL,
 "first_name" varchar NOT NULL,
 "last_name" varchar NOT NULL,
 "email" varchar UNIQUE NOT NULL,
 "password" varchar NOT NULL,
 "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" bigserial PRIMARY KEY,
  "category_id" varchar UNIQUE NOT NULL,
  "category_name" varchar UNIQUE NOT NULL
);

CREATE TABLE "products" (
    "id" bigserial PRIMARY KEY,
    "product_id" varchar UNIQUE NOT NULL,
    "image_url_public_id" varchar NOT NULL,
    "image_url_secure_id" varchar NOT NULL,
    "product_name" varchar NOT NULL,
    "product_description" varchar NOT NULL,
    "price" bigint NOT NULL,
    "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "reviews" (
   "id" bigserial PRIMARY KEY,
   "review_id" varchar UNIQUE NOT NULL,
   "user_review_id" varchar,
   "review" text NOT NULL,
   "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "orders" (
  "id" bigserial PRIMARY KEY,
  "order_id" varchar UNIQUE NOT NULL,
  "quantity" int DEFAULT 1,
  "total_amount" bigint,
  "created_at" varchar
);

CREATE TABLE "address" (
   "id" bigserial PRIMARY KEY,
   "address_user_id" varchar UNIQUE NOT NULL,
   "order_id" varchar,
   "email" varchar NOT NULL,
   "street" varchar NOT NULL,
   "address_line_1" varchar NOT NULL,
   "address_line_2" varchar NOT NULL,
   "city" varchar NOT NULL,
   "state" varchar NOT NULL,
   "zip_code" varchar NOT NULL
);

ALTER TABLE "products" ADD FOREIGN KEY ("product_id") REFERENCES "categories" ("category_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("review_id") REFERENCES "products" ("product_id");

ALTER TABLE "reviews" ADD FOREIGN KEY ("user_review_id") REFERENCES "users" ("user_id");

ALTER TABLE "orders" ADD FOREIGN KEY ("order_id") REFERENCES "products" ("product_id");

ALTER TABLE "address" ADD FOREIGN KEY ("address_user_id") REFERENCES "users" ("user_id");

ALTER TABLE "address" ADD FOREIGN KEY ("order_id") REFERENCES "orders" ("order_id");

CREATE INDEX ON "users" ("user_id");

CREATE INDEX ON "products" ("product_id");

CREATE INDEX ON "reviews" ("review_id");

CREATE INDEX ON "orders" ("order_id");

CREATE INDEX ON "address" ("address_user_id");

COMMENT ON COLUMN "orders"."created_at" IS 'When order created';
