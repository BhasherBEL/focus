import cardsTagsApi from '$lib/api/cardsTagsApi';
import { toastAlert } from '$lib/utils/toasts';
import Card from './Card';
import ProjectTag from './ProjectTag';
import TagOption from './TagOption';

export default class CardTag {
	private _card: Card;
	private _projectTag: ProjectTag;
	private _option: TagOption | null;
	private _value: string | null;

	private constructor(
		card: Card,
		projectTag: ProjectTag,
		option: TagOption | null,
		value: string | null
	) {
		this._card = card;
		this._projectTag = projectTag;
		this._option = option;
		this._value = value;
	}

	get card(): Card {
		return this._card;
	}

	get projectTag(): ProjectTag {
		return this._projectTag;
	}

	get option(): TagOption | null {
		return this._option;
	}

	get value(): string | null {
		return this._value;
	}

	static async create(
		card: Card,
		tag: ProjectTag,
		option: TagOption | null,
		value: string | null
	): Promise<CardTag | null> {
		const res = await cardsTagsApi.create(card.id, tag.id, option ? option.id : null, value);

		if (!res) return null;

		return new CardTag(card, tag, option, value);
	}

	async delete(): Promise<boolean> {
		return cardsTagsApi.delete(this._card.id, this._projectTag.id);
	}

	async update(option: TagOption | null, value: string | null): Promise<boolean> {
		const res = await cardsTagsApi.update(
			this._card.id,
			this._projectTag.id,
			option ? option.id : null,
			value
		);

		if (!res) return false;

		this._option = option;
		this._value = value;

		return true;
	}

	static parse(json: any): CardTag | null;
	static parse(json: any, card: Card | null | undefined): CardTag | null;

	static parse(json: any, card?: Card | null | undefined): CardTag | null {
		if (!json) {
			toastAlert('Failed to parse card tag: json is null');
			return null;
		}

		if (!card) card = Card.fromId(json.card_id);
		if (!card) {
			toastAlert('Failed to parse card tag: card is null');
			return null;
		}
		const tag = ProjectTag.fromId(json.tag_id);
		if (!tag) {
			toastAlert('Failed to parse card tag: tag is null');
			return null;
		}

		const option = tag.options.find((option) => option.id === json.option_id);
		if (!option) {
			toastAlert('Failed to parse card tag: option is null');
			return null;
		}

		return new CardTag(card, tag, option, json.value);
	}

	static parseAll(json: any): CardTag[];
	static parseAll(json: any, card: Card | null): CardTag[];

	static parseAll(json: any, card?: Card | null): CardTag[] {
		if (!json) return [];

		const cardTags: CardTag[] = [];

		for (const cardTag of json) {
			const parsed = this.parse(cardTag, card);
			if (parsed) {
				cardTags.push(parsed);
			}
		}

		return cardTags;
	}
}
