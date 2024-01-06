import type Project from '$lib/types/Project';
import type ProjectTag from '$lib/types/ProjectTag';
import TagOption from '$lib/types/TagOption';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function create(project: Project, title: string, type: number): Promise<number | null> {
	const response = await api.post(`/v1/tags/`, {
		project_id: project.id,
		title,
		type
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create tag');
		return null;
	}

	return response.data.id;
}

async function update(tag: ProjectTag): Promise<boolean> {
	const response = await api.put(`/v1/tags/${tag.id}`, tag);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update tag');
		return false;
	}
	return true;
}

async function delete_(tagId: number): Promise<boolean> {
	const response = await api.delete(`/v1/tags/${tagId}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete tag');
		return false;
	}
	return true;
}

async function getOptions(projectTag: ProjectTag): Promise<TagOption[]> {
	const response = await api.get(`/v1/tags/${projectTag.id}/options`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to get tag options');
		return [];
	}

	return TagOption.parseAll(response.data, projectTag);
}

export default {
	create,
	update,
	delete: delete_
};
