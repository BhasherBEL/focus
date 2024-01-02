<script lang="ts">
	import type { Card, MeTag, TagValue } from '../../../../stores/interfaces';
	import { getTagTypeFromId } from '../../../../utils/tagTypes';
	import SelectTags from './select_tags.svelte';

	export let projectTag: MeTag;
	export let tagValue: TagValue | undefined;
	export let card: Card;

	let tagType = getTagTypeFromId(projectTag.type);
</script>

<td>
	{#if tagType?.hasOptions}
		<SelectTags multiple={false} {projectTag} {card} bind:tagValue />
	{:else if !tagType?.hasOptions}
		<input />
	{/if}
</td>
