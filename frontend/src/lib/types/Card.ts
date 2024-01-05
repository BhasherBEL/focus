import type { TagValue } from './TagValue';

export default interface Card {
	id: number;
	project_id: number;
	title: string;
	content: string;
	tags: TagValue[];
}
