<script lang="ts">
	import type { Card } from '../stores/interfaces';
	import projectTags from '../stores/projectTags';
	import { currentModalCard } from '../stores/smallStore';
	import ModalCard from './modal_card.svelte';

	export let card: Card;
</script>

<div
	class="card"
	tabindex="0"
	draggable={true}
	on:click={() => ($currentModalCard = card.id)}
	role="button"
	on:keydown={(e) => {
		if (e.key === 'Enter') {
			$currentModalCard = card.id;
		}
	}}
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

<ModalCard bind:card />
