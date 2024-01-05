<script lang="ts">
	import type Card from '$lib/types/Card';
	import type MeTag from '$lib/types/MeTag';
	import type TagValue from '$lib/types/TagValue';
	import { getTagTypeFromId } from '$lib/utils/tagTypes';
	import SelectTags from './SelectTags.svelte';

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
