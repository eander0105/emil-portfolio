import { pgTable, serial, text, integer, timestamp } from 'drizzle-orm/pg-core';

export const user = pgTable('user', {
	id: text('id').primaryKey(),
	age: integer('age'),
});

export const translations = pgTable('translations', {
	id: serial('id').primaryKey(),
	key: text('key'),
	language: text('language'),
	translation: text('translation'),
});

export const project = pgTable('project', {
	title: text('title_key').references(() => translations.id),
	description: integer('description_key').references(() => translations.id),
	githubLink: text('github_link'),
});

export const session = pgTable('session', {
	id: text('id').primaryKey(),
	userId: text('user_id')
		.notNull()
		.references(() => user.id),
	expiresAt: timestamp('expires_at', { withTimezone: true, mode: 'date' }).notNull(),
});

export type Session = typeof session.$inferSelect;

export type User = typeof user.$inferSelect;
