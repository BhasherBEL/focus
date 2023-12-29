<script lang="ts">
	import axios from 'axios';
	import type { Card } from '../stores/interfaces';
	import ModalCard from './modal_card.svelte';
	import { backend } from '../stores/config';

	export let card: Card = {
		id: 0,
		project_id: 0,
		title: 'No title',
		content: 'Nocontent',
		tags: []
	};

	let showModal = false;

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
				{#if tag.value}
					<span class="tag" style="border: 1px solid #333">{tag.value}</span>
				{/if}
			{/each}
		</div>
	{/if}
</div>

<ModalCard bind:show={showModal} {card} onCancel={cancelEdit} />
