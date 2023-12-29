<script lang="ts">
	import axios from 'axios';
	import ModalTag from './modal_tag.svelte';
	import { backend } from '../stores/config';
	import type { Card } from '../stores/interfaces';

	export let tempCard: Card = {
		id: 0,
		project_id: 0,
		title: 'No title',
		content: 'Nocontent',
		tags: []
	};

	let newTagName = '';

	async function addTag() {
		if (newTagName === '') return;

		const response = await axios.post(`${backend}/api/tag`, {
			project_id: tempCard.project_id,
			title: newTagName,
			type: 0
		});

		const { id } = response.data.id;

		tempCard.tags = [
			...tempCard.tags,
			{ card_id: tempCard.id, tag_id: id, tag_title: newTagName, value: '' }
		];
		newTagName = '';
	}

	function addTagHandler(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			addTag();
		}
	}

	function removeTag(id: number) {
		tempCard.tags = tempCard.tags.filter((tag) => tag.tag_id !== id);
	}
</script>

<table>
	{#if tempCard.tags}
		{#each tempCard.tags as tag}
			<ModalTag {tag} {removeTag} />
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
