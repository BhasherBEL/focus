<script lang="ts">
	import type { MeTag, TagValue, Card } from '$lib/stores/interfaces';
	import { getTagTypeFromId } from '$lib/utils/tagTypes';
	import SelectTags from './select_tags.svelte';

	export let projectTag: MeTag;
	export let tagValue: TagValue | undefined;
	export let card: Card;

	let tagType = getTagTypeFromId(projectTag.type);
</script>

{#if tagType}
	<td>
		{#if tagType?.hasOptions}
			<SelectTags multiple={false} {projectTag} {card} {tagValue} />
		{:else if !tagType?.hasOptions}
			<input />
		{/if}
	</td>
{/if}
