import { db } from '$lib/server/db';
import { projectTable } from '$lib/server/db/schema';
import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { eq } from 'drizzle-orm';

export const load: PageServerLoad = async () => {
	const projectSchemas = await db.query.projectTable.findMany();

	return {
		projectSchemas
	};
};

export const actions: Actions = {
	delete: async ({ request }) => {
		const formData = await request.formData();
		const idStr = formData.get('id');

		if (!idStr || typeof idStr !== 'string') {
			return fail(400, { error: 'Project id is required.' });
		}

		const id = parseInt(idStr, 10);
		if (isNaN(id)) {
			return fail(400, { error: 'Invalid project id.' });
		}

		await db.delete(projectTable).where(eq(projectTable.id, id));

		return { success: true };
	}
};
