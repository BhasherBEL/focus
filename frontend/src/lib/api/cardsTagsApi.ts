import api, { processError } from '../utils/api';
import status from '../utils/status';

async function create(
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

async function update(
	cardId: number,
	tagId: number,
	optionId: number | null,
	value: string | null
): Promise<boolean> {
	const response = await api.put(`/v1/cards/${cardId}/tags/${tagId}`, {
		option_id: optionId,
		value: value
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update card tag');
		return false;
	}

	return true;
}

async function delete_(cardId: number, tagId: number): Promise<boolean> {
	const response = await api.delete(`/v1/cards/${cardId}/tags/${tagId}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete tag');
		return false;
	}

	return true;
}

export default {
	create,
	update,
	delete: delete_
};
