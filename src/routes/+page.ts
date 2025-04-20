import Project from '$lib/types/Project';
import type { Load } from '@sveltejs/kit';

export const load: Load = async ({ data }) => {
	const projects = Project.parseAll(data?.projectSchemas);

	return {
		projects
	};
};
