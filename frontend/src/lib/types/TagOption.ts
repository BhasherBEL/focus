import tagsOptions from '$lib/api/tagsOptions';
import { toastAlert } from '$lib/utils/toasts';
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

	get projectTag(): ProjectTag {
		return this._projectTag;
	}

	get value(): string {
		return this._value;
	}

	static async create(projectTag: ProjectTag, value: string): Promise<TagOption | null> {
		const id = await tagsOptions.create(projectTag.id, value);

		if (!id) return null;

		return new TagOption(id, projectTag, value);
	}

	async delete(): Promise<boolean> {
		return await tagsOptions.delete(this._id, this.projectTag.id);
	}

	async setValue(value: string): Promise<boolean> {
		const res = await tagsOptions.update(this._id, this.projectTag.id, value);

		if (!res) return false;

		this._value = value;

		return res;
	}

	static parse(json: any): TagOption | null;
	static parse(json: any, projectTag: ProjectTag | null | undefined): TagOption | null;

	static parse(json: any, projectTag?: ProjectTag | null | undefined): TagOption | null {
		if (!json) {
			toastAlert('Failed to parse TagOption');
			return null;
		}

		if (!projectTag) projectTag = ProjectTag.fromId(json.tag_id);
		if (!projectTag) {
			toastAlert('Failed to parse TagOption: ProjectTag not found');
			return null;
		}

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

	parseUpdate(dict: any) {
		if (dict.value) this._value = dict.value;
	}
}
