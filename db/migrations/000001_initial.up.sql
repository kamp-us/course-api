CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Istanbul";

-- Create courses table
CREATE TABLE courses (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_by UUID NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT NULL
);

-- Add indexes
CREATE INDEX active_courses ON courses(name) WHERE deleted_at = NULL;
CREATE INDEX idx_courses_slug ON courses(slug);

CREATE TABLE videos (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    uri VARCHAR(255) NOT NULL
);

-- Add indexes
CREATE INDEX active_videos ON videos(slug) WHERE deleted_at = NULL;
CREATE INDEX idx_videos_slug ON videos(slug);

-- Create lessons table
CREATE TABLE lessons (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT NULL,

    course_id UUID NULL,
    CONSTRAINT fk_course FOREIGN KEY(course_id) REFERENCES courses(id) ON DELETE CASCADE,

    video_id UUID NULL,
    CONSTRAINT fk_video FOREIGN KEY(video_id) REFERENCES videos(id) ON DELETE SET NULL
);

-- Add indexes
CREATE INDEX active_lessons ON lessons(slug) WHERE deleted_at = NULL;
CREATE INDEX idx_lessons_slug ON lessons(slug);

-- Create categories table
CREATE TABLE categories (
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP NULL,
    deleted_at TIMESTAMP NULL,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    description TEXT NULL
);

-- Add indexes
CREATE INDEX active_categories ON categories(slug) WHERE deleted_at = NULL;
CREATE INDEX idx_categories_slug ON categories(slug);

CREATE TABLE categories_courses (
  category_id UUID NOT NULL,
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id),

  course_id UUID NOT NULL,
  CONSTRAINT fk_course FOREIGN KEY(course_id) REFERENCES courses(id)
);

CREATE TABLE categories_lessons (
  category_id UUID NOT NULL,
  CONSTRAINT fk_category FOREIGN KEY(category_id) REFERENCES categories(id),

  lesson_id UUID NOT NULL,
  CONSTRAINT fk_lesson FOREIGN KEY(lesson_id) REFERENCES lessons(id)
);
