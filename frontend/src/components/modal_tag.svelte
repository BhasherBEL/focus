<script lang="ts">
	import type { Tag } from '../stores/interfaces';
	import api, { processError } from '../utils/api';
	import status from '../utils/status';

	export let tag: Tag;
	let newValue: string = tag.value;

	export let removeTag: (id: number) => void;

	async function saveTag() {
		if (tag.value === newValue) return;
		// DELETE
		if (tag.value !== '' && newValue === '') {
			const response = await api.delete(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`);

			if (response.status !== status.NoContent) {
				processError(response, 'Failed to delete tag');
				return;
			}
		}
		// CREATE
		else if (tag.value === '' && newValue !== '') {
			const response = await api.post(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
				value: newValue
			});
			if (response.status !== status.Created) {
				processError(response, 'Failed to create tag');
				return;
			}
		}
		// UPDATE
		else {
			const response = await api.put(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
				value: newValue
			});
			if (response.status !== status.NoContent) {
				processError(response, 'Failed to update tag');
				return;
			}
		}

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
