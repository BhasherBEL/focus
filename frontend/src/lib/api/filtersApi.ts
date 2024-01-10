import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function create(
	viewId: number,
	projectTagId: number,
	filterType: number,
	tagOptionId: number | null
): Promise<number | null> {
	const response = await api.post(`/v1/filters`, {
		view_id: viewId,
		tag_id: projectTagId,
		filter_type: filterType,
		option_id: tagOptionId
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create filter');
		return null;
	}

	return response.data.id;
}

async function update(
	filterId: number,
	viewId: number,
	projectTagId: number,
	filterType: number,
	tagOptionId: number | null
): Promise<boolean> {
	const response = await api.put(`/v1/filters/${filterId}`, {
		view_id: viewId,
		tag_id: projectTagId,
		filter_type: filterType,
		option_id: tagOptionId
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update filter');
		return false;
	}

	return true;
}

async function delete_(filterId: number): Promise<boolean> {
	const response = await api.delete(`/v1/filters/${filterId}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete filter');
		return false;
	}

	return true;
}

export default {
	create,
	update,
	delete: delete_
};
