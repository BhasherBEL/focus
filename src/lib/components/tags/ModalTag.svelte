<script lang="ts">
	import type Card from '$lib/types/Card';
	import type CardTag from '$lib/types/CardTag';
	import type ProjectTag from '$lib/types/ProjectTag';
	import ModalTagTitle from './ModalTagTitle.svelte';
	import ModalTagValue from './ModalTagValue.svelte';

	export let projectTag: ProjectTag;
	export let cardTag: CardTag | undefined;
	export let card: Card;
</script>

<tr>
	<ModalTagTitle {projectTag} />
	<ModalTagValue {projectTag} {cardTag} {card} />
</tr>
