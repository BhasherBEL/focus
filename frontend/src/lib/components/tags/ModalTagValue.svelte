<script lang="ts">
	import type Card from '$lib/types/Card';
	import type CardTag from '$lib/types/CardTag';
	import type ProjectTag from '$lib/types/ProjectTag';
	import { getTagTypeFromId } from '$lib/utils/tagTypes';
	import SelectTags from './SelectTags.svelte';

	export let projectTag: ProjectTag;
	export let cardTag: CardTag | undefined;
	export let card: Card;

	let tagType = getTagTypeFromId(projectTag.type);
</script>

{#if tagType}
	<td>
		{#if tagType?.hasOptions}
			<SelectTags multiple={false} {projectTag} {card} {cardTag} />
		{:else if !tagType?.hasOptions}
			<input />
		{/if}
	</td>
{/if}
