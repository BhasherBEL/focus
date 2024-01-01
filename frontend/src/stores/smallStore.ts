import { writable } from 'svelte/store';
import { parseCards, type Card, type View, type TagValue } from './interfaces';
import { deleteCardApi, newCardApi } from '../api/cards';
import { getProjectCardsAPI } from '../api/projects';
import api, { processError } from '../utils/api';
import status from '../utils/status';

export const currentView = (() => {
	const { subscribe, set, update } = writable(null as View | null);

	return {
		subscribe,
		set,
		update: async (view: View): Promise<boolean> => {
			const response = await api.put(`/v1/views/${view.id}`, view);

			if (response.status !== status.NoContent) {
				processError(response, 'Failed to update view');
				return false;
			}

			set(view);

			return true;
		}
	};
})();

export const currentModalCard = writable(-1);

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
				currentModalCard.set(-1);
			});
		},
		reload: () => {
			update((cards) => cards);
		}
	};
})();
