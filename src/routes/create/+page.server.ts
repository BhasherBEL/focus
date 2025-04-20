import { db } from '$lib/server/db';
import { projectTable } from '$lib/server/db/schema';

import { fail, redirect, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	create: async ({ request }) => {
		const formData = await request.formData();
		const title = formData.get('title');

		if (!title || typeof title !== 'string' || !title.trim()) {
			return fail(400, { error: 'Project title is required.' });
		}

		const [project] = await db.insert(projectTable).values({ title }).returning();
		if (!project || !project.id) {
			return fail(500, { error: 'Failed to create project.' });
		}

		return redirect(303, '/');
	}
};
