import { writable } from 'svelte/store';
import Project from './Project';
import ProjectTag from './ProjectTag';
import viewsApi from '$lib/api/viewsApi';

const views = writable([] as View[]);

export default class View {
	private _id: number;
	private _project: Project;
	private _pimaryTag: ProjectTag | null;
	private _secondaryTag: ProjectTag | null;
	private _title: string;
	private _sortTag: ProjectTag | null;
	private _sortDirection: number | null;

	private constructor(
		id: number,
		project: Project,
		primaryTag: ProjectTag | null,
		secondaryTag: ProjectTag | null,
		title: string,
		sortTag: ProjectTag | null,
		sortDirection: number | null
	) {
		this._id = id;
		this._project = project;
		this._pimaryTag = primaryTag;
		this._secondaryTag = secondaryTag;
		this._title = title;
		this._sortTag = sortTag;
		this._sortDirection = sortDirection;
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

	static async create(project: Project) {
		const id = await viewsApi.create(project);

		if (!id) return null;

		const view = new View(id, project, null, null, 'New view', null, null);

		views.update((views) => [...views, view]);

		return view;
	}

	async delete(): Promise<boolean> {
		if (!(await viewsApi.delete(this.id))) return false;

		views.update((views) => views.filter((view) => view.id !== this.id));

		return true;
	}

	static parse(json: any): View | null;
	static parse(json: any, project: Project | null | undefined): View | null;

	static parse(json: any, project?: Project | null | undefined): View | null {
		if (!json) return null;

		if (!project) project = Project.fromId(json.project_id);
		if (!project) return null;

		const primaryTag = ProjectTag.fromId(json.primary_tag_id);
		if (!primaryTag) return null;

		const secondaryTag = ProjectTag.fromId(json.secondary_tag_id);
		if (!secondaryTag) return null;

		const sortTag = ProjectTag.fromId(json.sort_tag_id);
		if (!sortTag) return null;

		const view = new View(
			json.id,
			project,
			primaryTag,
			secondaryTag,
			json.title,
			sortTag,
			json.sort_direction
		);

		views.update((views) => [...views, view]);

		return view;
	}
}
