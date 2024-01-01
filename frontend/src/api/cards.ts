import type { Card, TagValue } from '../stores/interfaces';
import api, { processError } from '../utils/api';
import status from '../utils/status';

export async function newCardApi(projectId: number, tags: TagValue[]): Promise<Card> {
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

	tags.forEach((tag) => (tag.card_id = id));

	return {
		id: id,
		projectId: projectId,
		title: 'Untitled',
		content: '',
		tags: tags
	};
}

export async function deleteCardApi(cardID: number): Promise<void> {
	const response = await api.delete(`/v1/cards/${cardID}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete card');
		return Promise.reject();
	}
}

export async function updateCardTagApi(
	cardID: number,
	tagID: number,
	option_id: number,
	value: string
): Promise<void> {
	const response = await api.put(`/v1/cards/${cardID}/tags/${tagID}`, {
		option_id: option_id,
		value: value
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update card tag');
		return Promise.reject();
	}
}
