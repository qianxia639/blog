DROP TABLE IF EXISTS comments;

ALTER TABLE IF EXISTS "comments" DROP CONSTRAINT IF EXISTS "comments_blog_id_fkey";

ALTER TABLE IF EXISTS "comments" DROP CONSTRAINT IF EXISTS "comments_comment_id_fkey";