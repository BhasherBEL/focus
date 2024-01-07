// import type TagOption from './TagOption';

import { get, writable } from 'svelte/store';
import TagOption from './TagOption';
import Project from './Project';

export const projectTags = writable([] as ProjectTag[]);

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
}
