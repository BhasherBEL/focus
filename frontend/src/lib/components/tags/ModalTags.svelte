<script lang="ts">
	import type Card from '$lib/types/Card';
	import { projectTags } from '$lib/types/ProjectTag';
	import ModalNewTag from './ModalNewTag.svelte';
	import ModalTag from './ModalTag.svelte';

	export let card: Card;
</script>

<table>
	{#each $projectTags as projectTag}
		<ModalTag
			cardTag={card.cardTags.find((t) => t.projectTag === projectTag)}
			{projectTag}
			{card}
		/>
	{/each}
	<ModalNewTag project={card.project} />
</table>

<style lang="less">
	table {
		font-size: 1.3rem;
	}
</style>
