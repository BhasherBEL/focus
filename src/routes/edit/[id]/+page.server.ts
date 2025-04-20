import { db } from '$lib/server/db';
import { projectTable } from '$lib/server/db/schema';
import { fail, redirect, error, type Actions, type ServerLoad } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';

export const load: ServerLoad = async ({ params }) => {
	const idStr = params.id;
	if (!idStr) throw error(400, 'Missing project id');

	const id = parseInt(idStr, 10);
	if (isNaN(id)) throw error(400, 'Invalid project id');

	console.log('Loading project with id:', id);

	const project = await db.query.projectTable.findFirst({ where: eq(projectTable.id, id) });
	if (!project) throw error(404, 'Project not found');

	return { project };
};

export const actions: Actions = {
	default: async ({ request, params }) => {
		const idStr = params.id;
		if (!idStr) return fail(400, { error: 'Missing project id.' });

		const id = parseInt(idStr, 10);
		if (isNaN(id)) return fail(400, { error: 'Invalid project id.' });

		const formData = await request.formData();
		const title = formData.get('title');

		if (!title || typeof title !== 'string' || !title.trim()) {
			return fail(400, { error: 'Project title is required.' });
		}

		const updated = await db
			.update(projectTable)
			.set({ title })
			.where(eq(projectTable.id, id))
			.returning();

		if (!updated || !updated[0]) {
			return fail(500, { error: 'Failed to update project.' });
		}

		return redirect(303, '/');
	}
};
