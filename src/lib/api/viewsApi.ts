import Filter from '$lib/types/Filter';
import type Project from '$lib/types/Project';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function create(project: Project): Promise<number | null> {
	const response = await api.post('/v1/views', {
		project: project.id
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create view');
		return null;
	}

	return response.data.id;
}

async function update(
	id: number,
	projectId: number,
	primaryTagId: number | null,
	secondaryTagId: number | null,
	title: string,
	sortTagId: number | null,
	sortDirection: number | null
): Promise<boolean> {
	const response = await api.put(`/v1/views/${id}`, {
		project_id: projectId,
		primary_tag_id: primaryTagId,
		secondary_tag_id: secondaryTagId,
		title: title,
		sort_tag_id: sortTagId,
		sort_direction: sortDirection
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update view');
		return false;
	}

	return true;
}

async function delete_(viewId: number): Promise<boolean> {
	const response = await api.delete(`/v1/views/${viewId}`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to delete view');
		return false;
	}

	return true;
}

async function getFilters(viewId: number): Promise<Filter[]> {
	const response = await api.get(`/v1/views/${viewId}/filters`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to get view filters');
		return [];
	}

	return Filter.parseAll(response.data);
}

export default {
	create,
	update,
	delete: delete_,
	getFilters
};
