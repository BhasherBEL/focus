import type { TagOption } from './TagOption';

export default interface MeTag {
	id: number;
	project_id: number;
	title: string;
	type: number;
	options: TagOption[];
}
