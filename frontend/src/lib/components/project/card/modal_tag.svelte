<script lang="ts">
	import type Card from '$lib/types/Card';
	import type MeTag from '$lib/types/MeTag';
	import type TagValue from '$lib/types/TagValue';
	import ModalTagTitle from './modal_tag/modal_tag_title.svelte';
	import ModalTagValue from './modal_tag/modal_tag_value.svelte';

	export let projectTag: MeTag;
	export let tagValue: TagValue | undefined;
	export let card: Card;
</script>

{#if projectTag}
	<tr>
		<ModalTagTitle {projectTag} />
		<ModalTagValue {projectTag} {tagValue} {card} />
	</tr>
{/if}
