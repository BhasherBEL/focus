import { sqliteTable, text, integer, primaryKey } from 'drizzle-orm/sqlite-core';

export const projectTable = sqliteTable('projects', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	title: text('title').notNull()
});

export type ProjectSchema = typeof projectTable.$inferSelect;

export const projectTagTable = sqliteTable('tags', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	project_id: integer('project_id')
		.notNull()
		.references(() => projectTable.id),
	title: text('title').notNull(),
	type: integer('type').notNull()
});

export const tagOptionTable = sqliteTable('tagoptions', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	tag_id: integer('tag_id')
		.notNull()
		.references(() => projectTagTable.id),
	value: text('value').notNull()
});

export const cardTable = sqliteTable('cards', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	project_id: integer('project_id')
		.notNull()
		.references(() => projectTable.id),
	title: text('title').notNull(),
	content: text('content').notNull()
});

export const cardTagTable = sqliteTable(
	'cardtags',
	{
		card_id: integer('card_id')
			.notNull()
			.references(() => cardTable.id),
		tag_id: integer('tag_id')
			.notNull()
			.references(() => projectTagTable.id),
		option_id: integer('option_id').references(() => tagOptionTable.id),
		value: text('value')
	},
	(table) => [primaryKey({ columns: [table.card_id, table.tag_id] })]
);

export const viewTable = sqliteTable('views', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	project_id: integer('project_id')
		.notNull()
		.references(() => projectTable.id),
	primary_tag_id: integer('primary_tag_id').references(() => projectTagTable.id),
	secondary_tag_id: integer('secondary_tag_id').references(() => projectTagTable.id),
	title: text('title').notNull(),
	sort_tag_id: integer('sort_tag_id').references(() => projectTagTable.id),
	sort_direction: integer('sort_direction')
});

export const filterTable = sqliteTable('filters', {
	id: integer('id').primaryKey({ autoIncrement: true }).notNull(),
	view_id: integer('view_id')
		.notNull()
		.references(() => viewTable.id),
	tag_id: integer('tag_id')
		.notNull()
		.references(() => projectTagTable.id),
	filter_type: integer('filter_type').notNull(),
	option_id: integer('option_id').references(() => tagOptionTable.id)
});
