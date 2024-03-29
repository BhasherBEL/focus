import Card from '$lib/types/Card';
import Project from '$lib/types/Project';
import ProjectTag from '$lib/types/ProjectTag';
import View from '$lib/types/View';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

async function getAll(): Promise<Project[]> {
	const response = await api.get('/v1/projects');

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch projects');
		return [];
	}

	return Project.parseAll(response.data);
}

async function create(title: string): Promise<number | null> {
	const response = await api.post('/v1/projects', {
		title
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to create project');
		return null;
	}

	return response.data.id;
}

async function get(projectId: number): Promise<Project | null> {
	const response = await api.get(`/v1/projects/${projectId}`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch project');
		return null;
	}

	return Project.parse(response.data);
}

async function update(projectId: number, title: string): Promise<boolean> {
	const response = await api.put(`/v1/projects/${projectId}`, {
		title
	});

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update project');
		return false;
	}

	return true;
}

async function delete_(projectId: number): Promise<boolean> {
	const response = await api.delete(`/v1/projects/${projectId}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete project');
		return false;
	}

	return true;
}

async function getCards(project: Project): Promise<Card[]> {
	const response = await api.get(`/v1/projects/${project.id}/cards`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch cards');
		return [];
	}

	return Card.parseAll(response.data, project);
}

async function getTags(project: Project): Promise<ProjectTag[]> {
	const response = await api.get(`/v1/projects/${project.id}/tags`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch tags');
		return [];
	}

	return ProjectTag.parseAll(response.data, project);
}

async function getViews(project: Project): Promise<View[]> {
	const response = await api.get(`/v1/projects/${project.id}/views`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to fetch views');
		return [];
	}

	return View.parseAll(response.data, project);
}

export default {
	create,
	get,
	update,
	delete: delete_,
	getAll,
	getCards,
	getTags,
	getViews
};
