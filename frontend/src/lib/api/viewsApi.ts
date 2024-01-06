import type Project from '$lib/types/Project';
import type View from '$lib/types/View';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function create(project: Project): Promise<number | null> {
	const response = await api.post('/views', {
		project: project.id
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create view');
		return null;
	}

	return response.data.id;
}

async function update(view: View): Promise<boolean> {
	const response = await api.put(`/views/${view.id}`, {
		project: view.project.id,
		primary_tag_id: view.primaryTag?.id,
		secondary_tag_id: view.secondaryTag?.id,
		title: view.title,
		sort_tag_id: view.sortTag?.id,
		sort_direction: view.sortDirection
	});

	if (response.status !== status.OK) {
		processError(response, 'Failed to update view');
		return false;
	}

	return true;
}

async function delete_(viewId: number): Promise<boolean> {
	const response = await api.delete(`/views/${viewId}`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to delete view');
		return false;
	}

	return true;
}

export default {
	create,
	update,
	delete: delete_
};
