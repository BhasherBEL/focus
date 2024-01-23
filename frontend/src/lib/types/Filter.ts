import filtersApi from '$lib/api/filtersApi';
import { toastAlert } from '$lib/utils/toasts';
import ProjectTag from './ProjectTag';
import TagOption from './TagOption';
import View from './View';

export default class Filter {
	private _id: number;
	private _view: View;
	private _projectTag: ProjectTag;
	private _filterType: number;
	private _tagOption: TagOption | null;

	private constructor(
		id: number,
		view: View,
		projectTag: ProjectTag,
		filterType: number,
		tagOption: TagOption | null
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

	get tagOption(): TagOption | null {
		return this._tagOption;
	}

	static async create(
		view: View,
		projectTag: ProjectTag,
		filterType: number,
		tagOption: TagOption | null
	): Promise<Filter | null> {
		const id = await filtersApi.create(view.id, projectTag.id, filterType, tagOption?.id || null);
		if (!id) return null;

		return new Filter(id, view, projectTag, filterType, tagOption);
	}

	async delete(): Promise<boolean> {
		return await filtersApi.delete(this.id);
	}

	async setProjectTag(projectTag: ProjectTag): Promise<boolean> {
		const res = await filtersApi.update(
			this.id,
			this.view.id,
			projectTag.id,
			this.filterType,
			this.tagOption?.id || null
		);

		if (!res) return false;

		this._projectTag = projectTag;

		return true;
	}

	async setFilterType(filterType: number): Promise<boolean> {
		const res = await filtersApi.update(
			this.id,
			this.view.id,
			this.projectTag.id,
			filterType,
			this.tagOption?.id || null
		);

		if (!res) return false;

		this._filterType = filterType;

		return true;
	}

	async setTagOption(tagOption: TagOption | null): Promise<boolean> {
		const res = await filtersApi.update(
			this.id,
			this.view.id,
			this.projectTag.id,
			this.filterType,
			tagOption?.id || null
		);

		if (!res) return false;

		this._tagOption = tagOption;

		return true;
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
			toastAlert(`Failed to parse filter: view ${json.view_id} not found`);
			return null;
		}

		const projectTag = ProjectTag.fromId(json.tag_id);
		if (!projectTag) {
			toastAlert(`Failed to parse filter: projectTag ${json.tag_id} not found`);
			return null;
		}

		const tagOption = projectTag.options.find((option) => option.id === json.option_id);

		return new Filter(json.id, view, projectTag, json.filter_type, tagOption || null);
	}

	parseUpdate(changes: any) {
		if (changes.tag_id) {
			const projectTag = ProjectTag.fromId(changes.tag_id);
			if (projectTag) this._projectTag = projectTag;
		}

		if (changes.filter_type) {
			this._filterType = changes.filter_type;
		}

		if (changes.option_id) {
			this._tagOption =
				this.projectTag.options.find((option) => option.id === changes.option_id) || null;
		}
	}
}
