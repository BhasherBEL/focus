export default interface View {
	id: number;
	project_id: number;
	primary_tag_id: number | null;
	secondary_tag_id: number | null;
	title: string;
	sort_tag_id: number | null;
	sort_direction: number | null;
}
