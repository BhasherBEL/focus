<script lang="ts">
	import ProjectTags from '$lib/stores/projectTags';
	import type Card from '$lib/types/Card';
	import ModalNewTag from './ModalNewTag.svelte';
	import ModalTag from './ModalTag.svelte';

	export let card: Card;
</script>

<table>
	{#if card.tags}
		{#each Object.values($ProjectTags) as projectTag}
			<ModalTag tagValue={card.tags.find((t) => t.tag_id === projectTag.id)} {projectTag} {card} />
		{/each}
	{/if}
	<ModalNewTag projectId={card.project_id} />
</table>
