// import type TagOption from './TagOption';

import projectTagsApi from '$lib/api/projectTagsApi';
import { get, writable } from 'svelte/store';
import Project from './Project';
import TagOption from './TagOption';

const { subscribe, update, set } = writable([] as ProjectTag[]);

export const projectTags = {
	subscribe,
	update,
	set,
	reload: () => update((ps) => ps)
};

export const ProjectTagTypes = {};

export default class ProjectTag {
	private _id: number;
	private _project: Project;
	private _title: string;
	private _type: number;
	private _options: TagOption[];

	private constructor(
		id: number,
		project: Project,
		title: string,
		type: number,
		options: TagOption[]
	) {
		this._id = id;
		this._project = project;
		this._title = title;
		this._type = type;
		this._options = options;
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

	get type(): number {
		return this._type;
	}

	get options(): TagOption[] {
		return this._options;
	}

	static fromId(id: number | null | undefined): ProjectTag | null {
		if (!id) return null;
		for (const projectTag of get(projectTags)) {
			if (projectTag.id === id) {
				return projectTag;
			}
		}

		return null;
	}

	static async create(project: Project, title: string, type: number): Promise<ProjectTag | null> {
		const response = await projectTagsApi.create(project.id, title, type);

		if (!response) return null;

		const projectTag = new ProjectTag(0, project, title, type, []);

		projectTags.update((projectTags) => [...projectTags, projectTag]);

		return projectTag;
	}

	async delete(): Promise<boolean> {
		return await projectTagsApi.delete(this.id);
	}

	async update(title: string, type: number): Promise<boolean> {
		return await projectTagsApi.update(this.id, title, type);
	}

	async addOption(title: string): Promise<TagOption | null> {
		const option = await TagOption.create(this, title);

		if (!option) return null;

		this._options = [...this._options, option];

		return option;
	}

	async deleteOption(option: TagOption): Promise<boolean> {
		const success = await option.delete();

		if (!success) return false;

		this._options = this._options.filter((o) => o.id !== option.id);

		return true;
	}

	static parse(json: any): ProjectTag | null;
	static parse(json: any, project: Project | null | undefined): ProjectTag | null;

	static parse(json: any, project?: Project | null): ProjectTag | null {
		if (!json) return null;

		if (!project) project = Project.fromId(json.project);
		if (!project) return null;

		const projectTag = new ProjectTag(json.id, json.project, json.title, json.type, []);

		const options = TagOption.parseAll(json.options, projectTag);

		projectTag._options = options;

		projectTags.update((projectTags) => {
			if (!projectTags.find((projectTag) => projectTag.id === json.id)) {
				return [...projectTags, projectTag];
			}

			return projectTags.map((pt) => (pt.id === json.id ? projectTag : pt));
		});

		return projectTag;
	}

	static parseAll(json: any): ProjectTag[];
	static parseAll(json: any, project: Project | null): ProjectTag[];

	static parseAll(json: any, project?: Project | null): ProjectTag[] {
		if (!json) return [];

		const projectTags: ProjectTag[] = [];

		for (const projectTag of json) {
			const parsed = ProjectTag.parse(projectTag, project);
			if (parsed) projectTags.push(parsed);
		}

		return projectTags;
	}

	static parseDelete(id: number) {
		projectTags.update((projectTags) => projectTags.filter((projectTag) => projectTag.id !== id));
	}

	parseUpdate(json: any) {
		if (!json) return;

		if (json.title) this._title = json.title;
		if (json.type) this._type = json.type;
	}

	parseOption(json: any) {
		if (!json) return;

		const option = TagOption.parse(json, this);
		if (!option) return;

		this._options = [...this._options, option];
	}

	parseOptionDelete(id: number) {
		this._options = this._options.filter((option) => option.id !== id);
	}

	parseOptionUpdate(json: any) {
		if (!json) return;

		const option = this._options.find((option) => option.id === json.option_id);
		if (!option) return;

		option.parseUpdate(json);
	}
}
