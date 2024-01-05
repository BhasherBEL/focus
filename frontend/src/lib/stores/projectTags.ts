import type MeTag from '$lib/types/MeTag';
import type TagOption from '$lib/types/TagOption';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';
import { get, writable } from 'svelte/store';
import cards from './cards';

const { subscribe, set, update } = writable({} as { [key: number]: MeTag });

export default {
	subscribe,
	init: async (projectID: number): Promise<boolean> => {
		const response = await api.get(`/v1/projects/${projectID}/tags`);

		if (response.status !== 200) {
			processError(response);
			return false;
		}

		const metags: MeTag[] = response.data;

		const tags: { [key: number]: MeTag } = {};

		if (metags) {
			metags.forEach((tag: MeTag) => {
				if (tag.options === null) tag.options = [];
				tags[tag.id] = tag;
			});
		}

		set(tags);

		return true;
	},
	addOption: async (tag_id: number, value: string) => {
		const response = await api.post(`/v1/tags/${tag_id}/options`, {
			value
		});

		if (response.status !== status.Created) {
			processError(response, 'Failed to create tag option');
			return;
		}

		const option: TagOption = {
			id: response.data.id,
			tag_id,
			value
		};

		update((tags) => {
			tags[tag_id].options.push(option);
			return tags;
		});
	},
	deleteOption: async (tag_id: number, option_id: number) => {
		const response = await api.delete(`/v1/tags/${tag_id}/options/${option_id}`);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to delete tag option');
			return;
		}

		update((tags) => {
			tags[tag_id].options = tags[tag_id].options.filter((option) => option.id !== option_id);
			return tags;
		});

		for (const card of get(cards)) {
			//TODO: same in db
			card.tags.filter((tag) => tag.tag_id !== tag_id || tag.option_id !== option_id);
		}
		cards.reload();
	},
	delete: async (tag_id: number) => {
		const response = await api.delete(`/v1/tags/${tag_id}`);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to delete tag');
			return;
		}

		update((tags) => {
			delete tags[tag_id];
			return tags;
		});
	},
	update: async (tag: MeTag) => {
		const response = await api.put(`/v1/tags/${tag.id}`, tag);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to update tag');
			return;
		}

		update((tags) => {
			tags[tag.id] = tag;
			return tags;
		});
	},
	add: async (ProjectId: number, title: string, type: number) => {
		const response = await api.post(`/v1/tags`, {
			project_id: ProjectId,
			title,
			type
		});

		if (response.status !== status.Created) {
			processError(response, 'Failed to create tag');
			return;
		}

		const tag: MeTag = {
			id: response.data.id,
			project_id: ProjectId,
			title,
			type,
			options: []
		};

		update((tags) => {
			tags[tag.id] = tag;
			return tags;
		});
	}
};
