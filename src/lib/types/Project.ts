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

	static parse(json: any): Project | null {
		if (!json) return null;

		const project = new Project(json.id, json.title);

		return project;
	}

	static parseAll(json: any): Project[] {
		if (!json) return [];

		// const projects: Project[] = [];
		//
		// for (const project of json) {
		// 	const parsed = Project.parse(project);
		// 	if (parsed) projects.push(parsed);
		// }

		return json
			.map((project: any) => Project.parse(project))
			.filter((project: Project | null) => project !== null) as Project[];
	}
}
