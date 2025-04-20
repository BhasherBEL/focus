import type Card from '$lib/types/Card';
import api, { processError } from '../utils/api';
import status from '../utils/status';

async function create(projectId: number): Promise<number | null> {
	const response = await api.post(`/v1/cards`, {
		project_id: projectId,
		title: 'Untitled',
		content: ''
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create card');
		return null;
	}

	return response.data.id;
}

async function update(
	cardId: number,
	projectId: number,
	title: string,
	content: string
): Promise<boolean> {
	const response = await api.put(`/v1/cards/${cardId}`, {
		project_id: projectId,
		title: title,
		content: content
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update card');
		return false;
	}

	return true;
}

async function delete_(cardID: number): Promise<boolean> {
	const response = await api.delete(`/v1/cards/${cardID}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete card');
		return false;
	}
	return true;
}

export default {
	create,
	update,
	delete: delete_
};
