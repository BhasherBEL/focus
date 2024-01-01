import { writable } from 'svelte/store';
import { parseCards, type Card, type View, type TagValue } from './interfaces';
import { deleteCardApi, newCardApi } from '../api/cards';
import { getProjectCardsAPI } from '../api/projects';

export const currentView = writable(null as View | null);

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
