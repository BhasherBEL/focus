import type ProjectTag from '$lib/types/ProjectTag';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function create(projectTagId: number, value: string): Promise<number | null> {
	const response = await api.post(`/v1/tags/${projectTagId}/options`, {
		value
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create tag option');
		return null;
	}

	return response.data.id;
}

async function update(tagOptionId: number, projectTagId: number, value: string): Promise<boolean> {
	const response = await api.put(`/v1/tags/${projectTagId}/options/${tagOptionId}`, {
		value
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update tag option');
		return false;
	}
	return true;
}

async function delete_(tagOptionId: number, projectTagId: number): Promise<boolean> {
	const response = await api.delete(`/v1/tags/${projectTagId}/options/${tagOptionId}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete tag option');
		return false;
	}
	return true;
}

export default {
	create,
	update,
	delete: delete_
};
