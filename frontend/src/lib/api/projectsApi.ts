import type Card from '$lib/types/Card';
import type Project from '$lib/types/Project';
import ProjectTag from '$lib/types/ProjectTag';
import api, { processError } from '$lib/utils/api';
import { parseCards } from '$lib/utils/parser';
import status from '$lib/utils/status';

async function get(projectId: number): Promise<number | null> {
	const response = await api.get(`/v1/projects/${projectId}`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch project');
		return null;
	}

	return response.data;
}

async function getCards(projectId: number): Promise<Card[]> {
	const response = await api.get(`/v1/projects/${projectId}/cards`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch cards');
		return Promise.reject();
	}

	return parseCards(response.data);
}

async function getTags(project: Project): Promise<ProjectTag[]> {
	const response = await api.get(`/v1/projects/${project.id}/tags`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch tags');
		return [];
	}

	const projectTags: ProjectTag[] = ProjectTag.parseAll(response.data, project);

	return projectTags;
}
