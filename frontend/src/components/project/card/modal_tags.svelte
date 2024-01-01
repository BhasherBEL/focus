<script lang="ts">
	import type { Card } from '../../../stores/interfaces';
	import api, { processError } from '../../../utils/api';
	import status from '../../../utils/status';
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
				<svg
					xmlns="http://www.w3.org/2000/svg"
					width="16"
					height="16"
					viewBox="0 0 24 24"
					stroke="white"
					stroke-width="2"
					stroke-linecap="round"
					stroke-linejoin="round"
				>
					<line x1="12" y1="5" x2="12" y2="19"></line>
					<line x1="5" y1="12" x2="19" y2="12"></line>
				</svg>
			</button>
		</td>
	</tr>
</table>
