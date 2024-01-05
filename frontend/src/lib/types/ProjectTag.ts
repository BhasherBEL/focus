// import type TagOption from './TagOption';

import { get, writable } from 'svelte/store';
import TagOption from './TagOption';

const projectTags = writable([] as ProjectTag[]);

export default class ProjectTag {
	private _id: number;
	private _projectId: number;
	private _title: string;
	private _type: number;
	private _options: TagOption[];

	private constructor(
		id: number,
		projectId: number,
		title: string,
		type: number,
		options: TagOption[]
	) {
		this._id = id;
		this._projectId = projectId;
		this._title = title;
		this._type = type;
		this._options = options;
	}

	get id(): number {
		return this._id;
	}

	get projectId(): number {
		return this._projectId;
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

	static getAll(): ProjectTag[] {
		return get(projectTags);
	}

	static fromId(id: number): ProjectTag | null {
		for (const projectTag of get(projectTags)) {
			if (projectTag.id === id) {
				return projectTag;
			}
		}

		return null;
	}

	static parse(json: any): ProjectTag | null {
		if (!json) return null;

		const projectTag = new ProjectTag(json.id, json.project_id, json.title, json.type, []);

		const options = TagOption.parseAll(json.options, projectTag);

		projectTag._options = options;

		return projectTag;
	}

	static parseAll(json: any): ProjectTag[] {
		if (!json) return [];

		const projectTags: ProjectTag[] = [];

		for (const projectTag of json) {
			const parsed = ProjectTag.parse(projectTag);
			if (parsed) projectTags.push(parsed);
		}

		return projectTags;
	}
}
