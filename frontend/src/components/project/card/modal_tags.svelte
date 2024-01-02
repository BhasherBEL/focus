<script lang="ts">
	import type { Card } from '../../../stores/interfaces';
	import projectTags from '../../../stores/projectTags';
	import ModalTag from './modal_tag.svelte';
	import ModalNewTag from './modal_tag/modal_new_tag.svelte';

	export let card: Card;
</script>

<table>
	{#if card.tags}
		{#each Object.values($projectTags) as projectTag}
			<ModalTag
				tagValue={card.tags.find((t) => t.tag_id === projectTag.id)}
				bind:projectTag
				{card}
			/>
		{/each}
	{/if}
	<ModalNewTag projectId={card.project_id} />
</table>
