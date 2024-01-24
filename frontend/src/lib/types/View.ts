import viewsApi from '$lib/api/viewsApi';
import currentView from '$lib/stores/currentView';
import { toastAlert } from '$lib/utils/toasts';
import { get, writable } from 'svelte/store';
import Filter from './Filter';
import Project from './Project';
import ProjectTag from './ProjectTag';
import type TagOption from './TagOption';

const { subscribe, set, update } = writable([] as View[]);

export const views = {
	subscribe,
	set,
	update,
	reload: (view?: View | null) => {
		update((views) => views);
		if (view && view === get(currentView)) {
			currentView.reload();
		}
	}
};

export default class View {
	private _id: number;
	private _project: Project;
	private _pimaryTag: ProjectTag | null;
	private _secondaryTag: ProjectTag | null;
	private _title: string;
	private _sortTag: ProjectTag | null;
	private _sortDirection: number | null;
	private _filters: Filter[];

	private constructor(
		id: number,
		project: Project,
		primaryTag: ProjectTag | null,
		secondaryTag: ProjectTag | null,
		title: string,
		sortTag: ProjectTag | null,
		sortDirection: number | null,
		filters: Filter[]
	) {
		this._id = id;
		this._project = project;
		this._pimaryTag = primaryTag;
		this._secondaryTag = secondaryTag;
		this._title = title;
		this._sortTag = sortTag;
		this._sortDirection = sortDirection;
		this._filters = filters;
	}

	static compare(a: View, b: View): number {
		return a.title.localeCompare(b.title);
	}

	get id(): number {
		return this._id;
	}

	get project(): Project {
		return this._project;
	}

	get primaryTag(): ProjectTag | null {
		return this._pimaryTag;
	}

	get secondaryTag(): ProjectTag | null {
		return this._secondaryTag;
	}

	get title(): string {
		return this._title;
	}

	get sortTag(): ProjectTag | null {
		return this._sortTag;
	}

	get sortDirection(): number | null {
		return this._sortDirection;
	}

	get filters(): Filter[] {
		return this._filters;
	}

	static fromId(id: number): View | null {
		for (const view of get(views)) {
			if (view.id === id) {
				return view;
			}
		}

		return null;
	}

	async setPrimaryTag(projectTag: ProjectTag): Promise<boolean> {
		const response = await viewsApi.update(
			this.id,
			this.project.id,
			projectTag.id,
			this.secondaryTag?.id || null,
			this.title,
			this.sortTag?.id || null,
			this.sortDirection || null
		);

		if (!response) return false;

		this._pimaryTag = projectTag;

		return true;
	}

	async setSortTag(projectTag: ProjectTag, direction: number): Promise<boolean> {
		const response = await viewsApi.update(
			this.id,
			this.project.id,
			this.primaryTag?.id || null,
			this.secondaryTag?.id || null,
			this.title,
			projectTag.id,
			direction
		);

		if (!response) return false;

		this._sortTag = projectTag;
		this._sortDirection = direction;

		return true;
	}

	static async create(project: Project) {
		const id = await viewsApi.create(project);

		if (!id) return null;

		const view = new View(id, project, null, null, 'New view', null, null, []);

		views.update((views) => [...views, view]);

		return view;
	}

	async delete(): Promise<boolean> {
		if (!(await viewsApi.delete(this.id))) return false;

		views.update((views) => views.filter((view) => view.id !== this.id));

		return true;
	}

	async setTitle(title: string): Promise<boolean> {
		if (
			!(await viewsApi.update(
				this.id,
				this.project.id,
				this.primaryTag?.id || null,
				this.secondaryTag?.id || null,
				title,
				this.sortTag?.id || null,
				this.sortDirection || null
			))
		)
			return false;

		this._title = title;

		views.reload(this);

		return true;
	}

	async addFilter(
		projectTag: ProjectTag,
		filterType: number,
		option: TagOption | null
	): Promise<Filter | null> {
		const filter = await Filter.create(this, projectTag, filterType, option);
		if (!filter) return null;

		this._filters = [...this._filters, filter];

		return filter;
	}

	async removeFilter(filter: Filter): Promise<boolean> {
		if (!(await filter.delete())) return false;

		this._filters = this._filters.filter((f) => f.id !== filter.id);

		return true;
	}

	parseUpdate(changes: any) {
		if (changes.primary_tag_id) {
			this._pimaryTag = ProjectTag.fromId(changes.primary_tag_id);
		}

		if (changes.secondary_tag_id) {
			this._secondaryTag = ProjectTag.fromId(changes.secondary_tag_id);
		}

		if (changes.title) {
			this._title = changes.title;
		}

		if (changes.sort_tag_id) {
			this._sortTag = ProjectTag.fromId(changes.sort_tag_id);
		}

		if (changes.sort_direction) {
			this._sortDirection = changes.sort_direction;
		}
	}

	static parseDelete(id: any) {
		views.update((views) => views.filter((view) => view.id !== id));
	}

	static parse(json: any): View | null;
	static parse(json: any, project: Project | null | undefined): View | null;

	static parse(json: any, project?: Project | null | undefined): View | null {
		if (!json) {
			toastAlert('Failed to parse view');
			return null;
		}

		if (!project) project = Project.fromId(json.project_id);
		if (!project) {
			toastAlert('Failed to find project');
			return null;
		}

		const primaryTag = ProjectTag.fromId(json.primary_tag_id);
		const secondaryTag = ProjectTag.fromId(json.secondary_tag_id);
		const sortTag = ProjectTag.fromId(json.sort_tag_id);

		const view = new View(
			json.id,
			project,
			primaryTag,
			secondaryTag,
			json.title,
			sortTag,
			json.sort_direction,
			[]
		);

		view._filters = Filter.parseAll(json.filters, view);

		views.update((views) => {
			if (!views.find((view) => view.id === json.id)) {
				return [...views, view];
			}

			return views.map((v) => (v.id === json.id ? view : v));
		});

		return view;
	}

	static parseAll(json: any): View[];
	static parseAll(json: any, project: Project | null | undefined): View[];

	static parseAll(json: any, project?: Project | null | undefined): View[] {
		if (!json) {
			toastAlert('Failed to parse views');
			return [];
		}

		if (!project) project = Project.fromId(json.project_id);
		if (!project) {
			toastAlert('Failed to find project');
			return [];
		}

		const views: View[] = [];

		for (const viewJson of json) {
			const view = View.parse(viewJson, project);

			if (view) views.push(view);
		}

		return views;
	}

	static parseFilter(json: any) {
		const filter = Filter.parse(json);
		if (!filter) return;

		filter.view._filters = [...filter.view._filters, filter];
	}

	static parseFilterUpdate(json: any) {
		const view = get(views).find((v) => v._filters.map((f) => f.id).includes(json.id));
		if (!view) return;

		const filter = view._filters.find((f) => f.id === json.id);
		if (!filter) return;

		filter.parseUpdate(json);
	}

	static parseFilterDelete(id: number) {
		const view = get(views).find((v) => v._filters.map((f) => f.id).includes(id));
		if (!view) return;

		view._filters = view._filters.filter((f) => f.id !== id);
	}
}
