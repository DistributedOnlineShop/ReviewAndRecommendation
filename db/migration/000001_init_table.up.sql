CREATE TABLE "reviews" (
  "review_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "rating" NUMERIC(7,2),
  "comment" TEXT NOT NULL,
  "status" VARCHAR NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "wishlists" (
  "wl_id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "product_id" VARCHAR(12) NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);

CREATE TABLE "coupons" (
  "coupon_id" UUID PRIMARY KEY NOT NULL,
  "code" VARCHAR(50) UNIQUE,
  "discount" NUMERIC(7,2),
  "min_purchase" NUMERIC(7,2),
  "start_at" TIMESTAMP(0) NOT NULL,
  "expires_at" TIMESTAMP(0) NOT NULL,
  "created_at" TIMESTAMP(0) NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMP(0)
);
