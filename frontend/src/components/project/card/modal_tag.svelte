<script lang="ts">
	import type { TagValue } from '../../../stores/interfaces';
	import projectTags from '../../../stores/projectTags';
	import { cards } from '../../../stores/smallStore';
	import api, { processError } from '../../../utils/api';
	import status from '../../../utils/status';
	import ModalTagTitle from './modal_tag_title.svelte';

	export let tag: TagValue;
	// export let removeTag: (id: number) => void;
	let newValue: string = tag.value;
	let newOption: number = tag.option_id;

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

		if (value) projectTags.addOption(tag.tag_id, value);
	}

	const projectTag = $projectTags[tag.tag_id];
</script>

{#if projectTag}
	<tr>
		<ModalTagTitle projectTag={$projectTags[tag.tag_id]} />
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
	</tr>
{/if}
