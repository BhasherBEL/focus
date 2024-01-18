import projectsApi from '$lib/api/projectsApi';
import { get, writable } from 'svelte/store';

const { subscribe, set, update } = writable([] as Project[]);

export const projects = {
	subscribe,
	set,
	update,
	reload: () => update((projects) => projects)
};

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

	static fromId(id: number): Project | null {
		for (const project of get(projects)) {
			if (project.id === id) {
				return project;
			}
		}

		return null;
	}

	static async create(): Promise<Project | null> {
		const id = await projectsApi.create('New project');

		if (!id) return null;

		const project = new Project(id, 'New project');

		projects.update((projects) => [...projects, project]);

		return project;
	}

	async setTitle(title: string): Promise<boolean> {
		if (!(await projectsApi.update(this._id, title))) return false;

		this._title = title;

		return true;
	}

	async delete(): Promise<boolean> {
		if (!(await projectsApi.delete(this._id))) return false;

		projects.update((projects) => projects.filter((project) => project.id !== this._id));

		return true;
	}

	updateFromDict(dict: any) {
		if (dict.title) this._title = dict.title;
	}

	static parse(json: any): Project | null {
		if (!json) return null;

		const project = new Project(json.id, json.title);

		projects.update((projects) => {
			if (!projects.find((p) => p.id === project.id)) return [...projects, project];

			return projects.map((p) => (p.id === project.id ? project : p));
		});

		return project;
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
