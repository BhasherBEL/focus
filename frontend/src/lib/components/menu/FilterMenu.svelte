<script lang="ts">
	import currentView from '$lib/stores/currentView';
	import type Filter from '$lib/types/Filter';
	import FilterMenuItem from './FilterMenuItem.svelte';
	import Menu from './Menu.svelte';

	export let isOpen = false;
</script>

{#if $currentView}
	<Menu bind:isOpen>
		{#each $currentView.filters as filter}
			<FilterMenuItem {filter} />
		{/each}
		<FilterMenuItem />
	</Menu>
{/if}

<style lang="less">
</style>
