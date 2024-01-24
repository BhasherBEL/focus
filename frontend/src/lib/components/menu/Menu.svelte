<script lang="ts">
	import { onDestroy, onMount } from 'svelte';

	export let isOpen: boolean;
	let menuElement: HTMLElement;
	export let onLeave: () => void = () => {};

	function clickOutside(event: MouseEvent) {
		if (
			isOpen &&
			menuElement &&
			!menuElement.contains(event.target as Node) &&
			menuElement !== event.target
		) {
			isOpen = false;
			onLeave();
		}
	}

	onMount(() => {
		window.addEventListener('click', clickOutside, true);
	});

	onDestroy(() => {
		window.removeEventListener('click', clickOutside, true);
	});
</script>

{#if isOpen}
	<div class="menu" bind:this={menuElement}>
		<slot />
	</div>
{/if}

<style lang="less">
	.menu {
		z-index: 100;
		position: absolute;
		display: flex;
		flex-direction: column;
		background-color: #222;
		border: 1px solid #666;
		padding: 10px 0;
		font-size: 1.3rem;
	}
</style>
