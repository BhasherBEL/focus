import cardsApi from '$lib/api/cardsApi';
import { get, writable } from 'svelte/store';
import CardTag from './CardTag';
import Project from './Project';
import { toastAlert } from '$lib/utils/toasts';

export const cards = writable([] as Card[]);

export default class Card {
	private _id: number;
	private _project: Project;
	private _title: string;
	private _content: string;
	private _tags: CardTag[];

	private constructor(
		id: number,
		project: Project,
		title: string,
		content: string,
		tags: CardTag[]
	) {
		this._id = id;
		this._project = project;
		this._title = title;
		this._content = content;
		this._tags = tags;
	}

	get id(): number {
		return this._id;
	}

	get project(): Project {
		return this._project;
	}

	get title(): string {
		return this._title;
	}

	get content(): string {
		return this._content;
	}

	get tags(): CardTag[] {
		return this._tags;
	}

	static fromId(id: number): Card | null {
		for (const card of get(cards)) {
			if (card.id === id) {
				return card;
			}
		}

		return null;
	}

	static async create(project: Project): Promise<Card | null> {
		const id = await cardsApi.create(project.id);

		if (!id) return null;

		const card = new Card(id, project, 'Untilted', '', []);

		cards.update((cards) => [...cards, card]);

		return card;
	}

	async delete(): Promise<boolean> {
		const res = await cardsApi.delete(this.id);

		if (!res) return false;

		cards.update((cards) => cards.filter((c) => c.id !== this.id));

		return true;
	}

	static parse(json: any): Card | null;
	static parse(json: any, project: Project | null | undefined): Card | null;

	static parse(json: any, project?: Project | null | undefined): Card | null {
		if (json === null) {
			return null;
		}

		if (!project) project = Project.fromId(json.project_id);
		if (!project) {
			toastAlert('Failed to parse card: project not found');
			return null;
		}

		const card = new Card(json.id, project, json.title, json.content, []);

		card._tags = CardTag.parseAll(json.tags, card);

		cards.update((cards) => {
			if (!cards.find((c) => c.id === card.id)) {
				return [...cards, card];
			}

			return cards.map((c) => (c.id === card.id ? card : c));
		});

		return card;
	}

	static parseAll(json: any): Card[];
	static parseAll(json: any, project: Project | null): Card[];

	static parseAll(json: any, project?: Project | null): Card[] {
		if (json === null) {
			return [];
		}

		const cards: Card[] = [];

		for (const jsonCard of json) {
			const card = this.parse(jsonCard, project);
			if (!card) continue;
			cards.push(card);
		}

		return cards;
	}
}
