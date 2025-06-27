import { pgTable, serial, text, integer } from 'drizzle-orm/pg-core';

export const user = pgTable('user', {
	id: serial('id').primaryKey(),
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
