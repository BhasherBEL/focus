import ProjectTag from './ProjectTag';

export default class TagOption {
	private _id: number;
	private _projectTag: ProjectTag;
	private _value: string;

	private constructor(id: number, projectTag: ProjectTag, value: string) {
		this._id = id;
		this._projectTag = projectTag;
		this._value = value;
	}

	get id(): number {
		return this._id;
	}

	get tagId(): ProjectTag {
		return this._projectTag;
	}

	get value(): string {
		return this._value;
	}

	static parse(json: any): TagOption | null;
	static parse(json: any, projectTag: ProjectTag | null | undefined): TagOption | null;

	static parse(json: any, projectTag?: ProjectTag | null | undefined): TagOption | null {
		if (!json) return null;

		if (!projectTag) projectTag = ProjectTag.fromId(json.tag_id);
		if (!projectTag) return null;

		return new TagOption(json.id, projectTag, json.value);
	}

	static parseAll(json: any): TagOption[];
	static parseAll(json: any, projectTag: ProjectTag | null): TagOption[];

	static parseAll(json: any, projectTag?: ProjectTag | null): TagOption[] {
		if (!json) return [];

		const options: TagOption[] = [];

		for (const option of json) {
			const parsed = TagOption.parse(option, projectTag);
			if (parsed) options.push(parsed);
		}

		return options;
	}
}
