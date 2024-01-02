<script lang="ts">
	import type { Card } from '../../../stores/interfaces';
	import api, { processError } from '../../../utils/api';
	import status from '../../../utils/status';
	import AddIcon from '../../icons/addIcon.svelte';
	import ModalTag from './modal_tag.svelte';

	export let card: Card;

	let newTagName = '';

	async function addTag() {
		if (newTagName === '') return;

		const response = await api.post(`/v1/tags`, {
			project_id: card.projectId,
			title: newTagName,
			type: 0
		});

		if (response.status !== status.Created) {
			processError(response, 'Failed to create tag');
			return;
		}
		const id = response.data.id;

		card.tags = [...card.tags, { card_id: card.id, tag_id: id, option_id: -1, value: '' }];
		newTagName = '';
	}

	function addTagHandler(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			addTag();
		}
	}

	function removeTag(id: number) {
		card.tags = card.tags.filter((tag) => tag.tag_id !== id);
	}
</script>

<table>
	{#if card.tags}
		{#each card.tags as tag}
			<ModalTag bind:tag />
		{/each}
	{/if}
	<tr class="tag">
		<td class="tag-title"
			><input placeholder="Add a new tag" on:keydown={addTagHandler} bind:value={newTagName} /></td
		>
		<td
			><button on:click={addTag}>
				<AddIcon />
			</button>
		</td>
	</tr>
</table>
