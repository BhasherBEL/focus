import type { Card } from '../stores/interfaces';
import api, { processError } from '../utils/api';
import status from '../utils/status';

export async function newCardApi(projectId: number): Promise<Card> {
	const response = await api.post(`/v1/cards`, {
		project_id: projectId,
		title: 'Untitled',
		content: ''
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create card');
		return Promise.reject();
	}

	const id: number = response.data.id;

	return {
		id: id,
		projectId: projectId,
		title: 'Untitled',
		content: '',
		tags: []
	};
}

export async function deleteCardApi(cardID: number): Promise<void> {
	const response = await api.delete(`/v2/cards/${cardID}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete card');
		return Promise.reject();
	}
}
