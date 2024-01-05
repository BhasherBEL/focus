import { parseCards, type Card, type Project } from '$lib/stores/interfaces';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

export async function getProjectAPI(projectId: number): Promise<Project> {
	const response = await api.get(`/v1/projects/${projectId}`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch project');
		return Promise.reject();
	}

	return response.data;
}

export async function getProjectCardsAPI(projectId: number): Promise<Card[]> {
	const response = await api.get(`/v1/projects/${projectId}/cards`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch cards');
		return Promise.reject();
	}

	return parseCards(response.data);
}
