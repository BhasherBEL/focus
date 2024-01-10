<script lang="ts">
	import type Filter from '$lib/types/Filter';
	import FilterMenuItem from './FilterMenuItem.svelte';
	import Menu from './Menu.svelte';

	export let isOpen = false;
	export let filters: Filter[] = [];
</script>

<Menu bind:isOpen>
	{#each filters as filter}
		<FilterMenuItem {filter} />
	{/each}
	<FilterMenuItem />
</Menu>

<style lang="less">
</style>
