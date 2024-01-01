<script lang="ts">
	import type { TagValue } from '../../../stores/interfaces';
	import projectTags from '../../../stores/projectTags';
	import { cards } from '../../../stores/smallStore';
	import api, { processError } from '../../../utils/api';
	import status from '../../../utils/status';

	export let tag: TagValue;
	// export let removeTag: (id: number) => void;
	let newValue: string = tag.value;
	let newOption: number = tag.option_id;

	let projectTag = $projectTags[tag.tag_id];

	async function saveTag() {
		if (tag.value === newValue && tag.option_id == newOption) return;
		// DELETE
		if ((tag.value !== '' && newValue === '') || (tag.option_id !== -1 && newOption === -1)) {
			const response = await api.delete(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`);

			if (response.status !== status.NoContent) {
				processError(response, 'Failed to delete tag');
				return;
			}
		}
		// CREATE
		else if ((tag.value === '' && newValue !== '') || (tag.option_id == -1 && newOption !== -1)) {
			const response = await api.post(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
				value: newValue,
				option_id: newOption
			});
			if (response.status !== status.Created) {
				processError(response, 'Failed to create tag');
				return;
			}
		}
		// UPDATE
		else {
			const response = await api.put(`/v1/cards/${tag.card_id}/tags/${tag.tag_id}`, {
				value: newValue,
				option_id: newOption
			});
			if (response.status !== status.NoContent) {
				processError(response, 'Failed to update tag');
				return;
			}
		}

		tag.value = newValue;
		tag.option_id = newOption;
		cards.reload();
	}

	async function newTagOption() {
		const value = prompt('New option value');

		if (value) projectTags.add(tag.tag_id, value);
	}
</script>

{#if projectTag}
	<tr class="tag">
		<td class="tag-title">{projectTag.title}</td>
		<td>
			{#if projectTag.type == 0}
				<select bind:value={newOption} on:change={saveTag}>
					<option value={-1}></option>
					{#each projectTag.options as option}
						<option value={option.id}>{option.value}</option>
					{/each}
				</select>
				<button style="color: white; font-size: 20px;" on:click={newTagOption}>+</button>
			{:else}
				<input bind:value={newValue} on:blur={saveTag} />
			{/if}
		</td>
		<!-- <td>
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
		</td> -->
	</tr>
{/if}
