CREATE TABLE IF NOT EXISTS students (
  student_id UUID NOT NULL,
  name VARCHAR(255) NOT NULL,
  
  CONSTRAINT "student_pkey" PRIMARY KEY ("student_id")
);

CREATE TABLE IF NOT EXISTS questions (
  question_id UUID NOT NULL,
  author_id UUID NOT NULL REFERENCES students(student_id),
  best_answer_id UUID,
  slug VARCHAR(200) NOT NULL,
  title VARCHAR(100) NOT NULL,
  content TEXT NOT NULL,
  is_active BOOLEAN NOT NULL,
  created_at TIMESTAMPTZ(6) DEFAULT current_timestamp NOT NULL,
  updated_at TIMESTAMPTZ(6) DEFAULT current_timestamp NOT NULL,

  CONSTRAINT "question_pkey" PRIMARY KEY ("question_id")
);