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

	const consistant_tags = [];

	for (const tag of tags) {
		if ((tag.option_id === -1 && tag.value == '') || tag.tag_id === -1) continue;
		await createCardTagApi(id, tag.tag_id, tag.option_id, tag.value);
		consistant_tags.push({ ...tag, card_id: id });
	}

	return {
		id: id,
		project_id: projectId,
		title: 'Untitled',
		content: '',
		tags: consistant_tags
	};
}

export async function updateCardApi(card: Card): Promise<boolean> {
	const response = await api.put(`/v1/cards/${card.id}`, {
		project_id: card.project_id,
		title: card.title,
		content: card.content
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update card');
		return false;
	}

	return true;
}

export async function deleteCardApi(cardID: number): Promise<void> {
	const response = await api.delete(`/v1/cards/${cardID}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete card');
		return Promise.reject();
	}
}

export async function createCardTagApi(
	cardId: number,
	tagId: number,
	optionId: number | null,
	value: string | null
): Promise<boolean> {
	const response = await api.post(`/v1/cards/${cardId}/tags/${tagId}`, {
		option_id: optionId,
		value: value
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create card tag');
		return false;
	}

	return true;
}

export async function deleteCardTagApi(cardID: number, tagID: number): Promise<void> {
	const response = await api.delete(`/v1/cards/${cardID}/tags/${tagID}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete tag');
		return Promise.reject();
	}
}

export async function updateCardTagApi(
	cardID: number,
	tagID: number,
	option_id: number | null,
	value: string | null
): Promise<void> {
	const response = await api.put(`/v1/cards/${cardID}/tags/${tagID}`, {
		option_id: option_id == -1 ? null : option_id,
		value: value == '' ? null : value
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update card tag');
		return Promise.reject();
	}
}
