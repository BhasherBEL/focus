<script lang="ts">
	import type { Card } from '$lib/stores/interfaces';
	import ModalTag from './modal_tag.svelte';
	import ModalNewTag from './modal_tag/modal_new_tag.svelte';
	import ProjectTags from '$lib/stores/project_tags';

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
