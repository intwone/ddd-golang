CREATE TYPE "user_role" AS ENUM ('student', 'instructor');

CREATE TABLE IF NOT EXISTS "users" (
  "user_id" UUID NOT NULL,
  "name" TEXT NOT NULL,
  "role" "user_role" DEFAULT 'student' NOT NULL,

  CONSTRAINT "users_pkey" PRIMARY KEY ("user_id")
);

CREATE TABLE IF NOT EXISTS questions (
  "question_id" UUID NOT NULL,
  "author_id" UUID NOT NULL,
  "best_answer_id" UUID,
  "slug" TEXT NOT NULL,
  "title" TEXT NOT NULL,
  "content" TEXT NOT NULL,
  "is_active" BOOLEAN NOT NULL DEFAULT true,
  "created_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,

  CONSTRAINT "questions_pkey" PRIMARY KEY ("question_id")
);

CREATE TABLE "answers" (
  "answer_id" UUID NOT NULL,
  "author_id" UUID NOT NULL,
  "question_id" UUID NOT NULL,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,

  CONSTRAINT "answers_pkey" PRIMARY KEY ("answer_id")
);

CREATE TABLE "comments" (
  "comment_id" UUID NOT NULL,
  "author_id" UUID NOT NULL,
  "question_id" UUID,
  "answer_id" UUID,
  "content" TEXT NOT NULL,
  "created_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,
  "updated_at" TIMESTAMPTZ(6) DEFAULT CURRENT_TIMESTAMP NOT NULL,

  CONSTRAINT "comments_pkey" PRIMARY KEY ("comment_id")
);

CREATE TABLE "attachments" (
  "attachment_id" UUID NOT NULL,
  "question_id" UUID,
  "answer_id" UUID,
  "title" TEXT NOT NULL,
  "link" TEXT NOT NULL,

  CONSTRAINT "attachments_pkey" PRIMARY KEY ("attachment_id")
);
