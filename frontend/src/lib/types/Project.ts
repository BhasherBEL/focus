import { get, writable } from 'svelte/store';

const projects = writable([] as Project[]);

export default class Project {
	private _id: number;
	private _title: string;

	private constructor(id: number, title: string) {
		this._id = id;
		this._title = title;
	}

	get id(): number {
		return this._id;
	}

	get title(): string {
		return this._title;
	}

	static getAll(): Project[] {
		return get(projects);
	}

	static fromId(id: number): Project | null {
		for (const project of get(projects)) {
			if (project.id === id) {
				return project;
			}
		}

		return null;
	}

	static parse(json: any): Project | null {
		if (!json) return null;

		return new Project(json.id, json.title);
	}

	static parseAll(json: any): Project[] {
		if (!json) return [];

		const projects: Project[] = [];

		for (const project of json) {
			const parsed = Project.parse(project);
			if (parsed) projects.push(parsed);
		}

		return projects;
	}
}
