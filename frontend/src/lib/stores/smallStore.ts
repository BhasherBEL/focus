import { writable } from 'svelte/store';
import { parseCards, type Card, type View, type TagValue } from './interfaces';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';
import { getProjectCardsAPI } from '$lib/api/projects';
import { deleteCardApi, newCardApi, updateCardApi } from '$lib/api/cards';

export const currentView = (() => {
	const { subscribe, set, update } = writable(null as View | null);

	return {
		subscribe,
		set
		// update: async (view: View): Promise<boolean> => {
		// 	const response = await api.put(`/v1/views/${view.id}`, view);

		// 	if (response.status !== status.NoContent) {
		// 		processError(response, 'Failed to update view');
		// 		return false;
		// 	}

		// 	set(view);

		// 	return true;
		// }
	};
})();

export const currentModalCard = writable(null as number | null);

export const currentDraggedCard = writable(null as Card | null);

export const cards = (() => {
	const { subscribe, set, update } = writable([] as Card[]);

	return {
		subscribe,
		init: async (projectId: number) => {
			getProjectCardsAPI(projectId).then((c) => {
				set(parseCards(c));
			});
		},
		add: async (projectId: number, tags: TagValue[]) => {
			await newCardApi(projectId, tags).then((card) => {
				update((cards) => [...cards, card]);
				currentModalCard.set(card.id);
			});
		},
		remove: async (card: Card) => {
			await deleteCardApi(card.id).then(() => {
				update((cards) => cards.filter((c) => c.id !== card.id));
				currentModalCard.set(null);
			});
		},
		edit: async (card: Card): Promise<boolean> => {
			if (await updateCardApi(card)) {
				update((cards) => cards.map((c) => (c.id === card.id ? card : c)));
				return true;
			}
			return false;
		},
		reload: () => {
			update((cards) => cards);
		}
	};
})();

export const views = (() => {
	const { subscribe, set, update } = writable([] as View[]);

	const init = async (projectId: number): Promise<boolean> => {
		const response = await api.get(`/v1/projects/${projectId}/views`);

		if (response.status !== status.OK) {
			processError(response, 'Failed to get views');
			return false;
		}

		if (response.data) {
			set(response.data);
		}

		return true;
	};
	const add = async (
		projectId: number,
		title: string,
		primaryTagId: number | null
	): Promise<View> => {
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
	};
	const remove = async (view: View): Promise<boolean> => {
		const response = await api.delete(`/v1/views/${view.id}`);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to delete view');
			return false;
		}

		update((views) => views.filter((v) => v.id !== view.id));

		return true;
	};

	const edit = async (view: View): Promise<boolean> => {
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
	};

	return {
		subscribe,
		init,
		add,
		remove,
		edit
	};
})();
