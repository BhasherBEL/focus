import { writable } from 'svelte/store';
import View from './View';
import ProjectTag from './ProjectTag';
import TagOption from './TagOption';
import { toastAlert } from '$lib/utils/toasts';
import filtersApi from '$lib/api/filtersApi';

export default class Filter {
	private _id: number;
	private _view: View;
	private _projectTag: ProjectTag;
	private _filterType: number;
	private _tagOption: TagOption;

	private constructor(
		id: number,
		view: View,
		projectTag: ProjectTag,
		filterType: number,
		tagOption: TagOption
	) {
		this._id = id;
		this._view = view;
		this._projectTag = projectTag;
		this._filterType = filterType;
		this._tagOption = tagOption;
	}

	get id(): number {
		return this._id;
	}

	get view(): View {
		return this._view;
	}

	get projectTag(): ProjectTag {
		return this._projectTag;
	}

	get filterType(): number {
		return this._filterType;
	}

	get tagOption(): TagOption {
		return this._tagOption;
	}

	static async create(
		view: View,
		projectTag: ProjectTag,
		filterType: number,
		tagOption: TagOption
	): Promise<Filter | null> {
		const id = await filtersApi.create(view.id, projectTag.id, filterType, tagOption.id);
		if (!id) return null;

		return new Filter(id, view, projectTag, filterType, tagOption);
	}

	async delete(): Promise<boolean> {
		return await filtersApi.delete(this.id);
	}

	static parseAll(json: any): Filter[];
	static parseAll(json: any, view: View | null): Filter[];

	static parseAll(json: any[], view?: View | null): Filter[] {
		if (!json) return [];

		const filters: Filter[] = [];

		for (const filter of json) {
			const parsed = Filter.parse(filter, view);
			if (parsed) filters.push(parsed);
		}

		return filters;
	}

	static parse(json: any): Filter | null;
	static parse(json: any, view: View | null | undefined): Filter | null;

	static parse(json: any, view?: View | null | undefined): Filter | null {
		if (!json) {
			toastAlert('Failed to parse filter: json is null');
			return null;
		}

		if (!view) view = View.fromId(json.view_id);
		if (!view) {
			toastAlert('Failed to parse filter: view not found');
			return null;
		}

		const projectTag = ProjectTag.fromId(json.tag_id);
		if (!projectTag) {
			toastAlert('Failed to parse filter: projectTag not found');
			return null;
		}

		const tagOption = projectTag.options.find((option) => option.id === json.option_id);
		if (!tagOption) {
			toastAlert('Failed to parse filter: tagOption not found');
			return null;
		}

		return new Filter(json.id, view, projectTag, json.filter_type, tagOption);
	}
}
