import { get, writable } from 'svelte/store';
import Project from './Project';
import ProjectTag from './ProjectTag';
import viewsApi from '$lib/api/viewsApi';
import { toastAlert } from '$lib/utils/toasts';
import Filter from './Filter';
import type TagOption from './TagOption';

export const views = writable([] as View[]);

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

		return true;
	}

	async addFilter(projectTag: ProjectTag, filterType: number, option: TagOption): Promise<boolean> {
		const filter = await Filter.create(this, projectTag, filterType, option);
		if (!filter) return false;

		this._filters = [...this._filters, filter];

		return true;
	}

	async removeFilter(filter: Filter): Promise<boolean> {
		if (!(await filter.delete())) return false;

		this._filters = this._filters.filter((f) => f.id !== filter.id);

		return true;
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
}
