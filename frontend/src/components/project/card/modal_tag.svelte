<script lang="ts">
	import type { Card, MeTag, TagValue } from '../../../stores/interfaces';
	import ModalTagTitle from './modal_tag/modal_tag_title.svelte';
	import ModalTagValue from './modal_tag/modal_tag_value.svelte';

	export let projectTag: MeTag;
	export let tagValue: TagValue | undefined;
	export let card: Card;
</script>

{#if projectTag}
	<tr>
		<ModalTagTitle {projectTag} />
		<ModalTagValue {projectTag} bind:tagValue {card} />
	</tr>
{/if}
