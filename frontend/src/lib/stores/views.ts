import type View from '$lib/types/View';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';
import { writable } from 'svelte/store';

const { subscribe, set, update } = writable([] as View[]);

async function init(projectId: number): Promise<boolean> {
	const response = await api.get(`/v1/projects/${projectId}/views`);

	if (response.status !== status.OK) {
		processError(response, 'Failed to get views');
		return false;
	}

	if (response.data) {
		set(response.data);
	}

	return true;
}
async function add(projectId: number, title: string, primaryTagId: number | null): Promise<View> {
	const response = await api.post(`/v1/views`, {
		title,
		project_id: projectId,
		primary_tag_id: primaryTagId
	});

	if (response.status !== status.Created) {
		processError(response, 'Failed to add view');
		return Promise.reject();
	}

	const view: View = {
		id: response.data.id,
		title: title,
		project_id: projectId,
		primary_tag_id: primaryTagId,
		secondary_tag_id: null,
		sort_tag_id: null,
		sort_direction: null
	};

	update((views) => [...views, view]);

	return view;
}
async function remove(view: View): Promise<boolean> {
	const response = await api.delete(`/v1/views/${view.id}`);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to delete view');
		return false;
	}

	update((views) => views.filter((v) => v.id !== view.id));

	return true;
}

async function edit(view: View): Promise<boolean> {
	if (view.title === '') {
		if (confirm('Are you sure you want to delete this view?')) {
			return remove(view);
		} else {
			return false;
		}
	}

	const response = await api.put(`/v1/views/${view.id}`, view);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update view');
		return false;
	}

	update((views) => views.map((v) => (v.id === view.id ? view : v)));

	return true;
}

export default {
	subscribe,
	init,
	add,
	remove,
	edit
};
