import { db } from '$lib/server/db';
import { project, type ProjectSchema } from '$lib/server/db/schema';

export async function getProjects(): Promise<ProjectSchema[]> {
	return await db.select().from(project).all();
}
