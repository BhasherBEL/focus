<script lang="ts">
	import axios from 'axios';
	import { backend } from '../stores/config';
	import type { Tag } from '../stores/interfaces';

	export let tag: Tag;
	let newValue: string = tag.value;

	export let removeTag: (id: number) => void;

	function saveTag() {
		if (tag.value === newValue) return;
		// DELETE
		if (tag.value !== '' && newValue === '') {
			axios.delete(`${backend}/api/v1/cards/${tag.card_id}/tags/${tag.tag_id}`);
			return;
		}
		// CREATE
		if (tag.value === '' && newValue !== '') {
			axios.post(`${backend}/api/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
				value: newValue
			});
			return;
		}
		// UPDATE
		axios.put(`${backend}/api/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
			value: newValue
		});

		tag.value = newValue;
	}
</script>

<tr class="tag">
	<td class="tag-title">{tag.tag_title}</td>
	<td>
		<input bind:value={newValue} on:blur={saveTag} />
	</td>
	<td>
		<button on:click={() => removeTag(tag.tag_id)} class="remove-tag-button">
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="16"
				height="16"
				viewBox="0 0 24 24"
				fill="none"
				stroke="white"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<line x1="18" y1="6" x2="6" y2="18"></line>
				<line x1="6" y1="6" x2="18" y2="18"></line>
			</svg>
		</button>
	</td>
</tr>
