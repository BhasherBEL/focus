<script lang="ts">
	import type { Card } from '../stores/interfaces';
	import projectTags from '../stores/projectTags';
	import ModalCard from './modal_card.svelte';

	export let card: Card;
	export let showModal: boolean;
	export let onDelete: () => void;

	function editCard() {
		showModal = true;
	}

	function cancelEdit() {
		showModal = false;
	}

	function editCardHandler(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			editCard();
		}
	}
</script>

<div
	class="card"
	tabindex="0"
	draggable={true}
	on:click={editCard}
	role="button"
	on:keydown={editCardHandler}
>
	<div class="title">{card.title}</div>
	{#if card.tags}
		<div class="tags">
			{#each card.tags as tag}
				{#if tag.option_id && tag.option_id !== -1}
					{#if $projectTags[tag.tag_id]}
						<span class="tag" style="border: 1px solid #333"
							>{$projectTags[tag.tag_id]?.options.find((o) => o.id == tag.option_id)?.value}</span
						>
					{/if}
				{:else if tag.value}
					<span class="tag" style="border: 1px solid #333">{tag.value}</span>
				{/if}
			{/each}
		</div>
	{/if}
</div>

<ModalCard bind:show={showModal} bind:card onCancel={cancelEdit} {onDelete} />
