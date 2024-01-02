<script lang="ts">
	import { onDestroy, onMount } from 'svelte';

	export let isOpen: boolean;
	let menuElement: HTMLElement;
	export let onLeave: () => void = () => {};

	function handleFocus(event: FocusEvent) {
		if (isOpen && !menuElement.contains(event.target as Node) && menuElement !== event.target) {
			isOpen = false;
			onLeave();
		}
	}

	onMount(() => {
		window.addEventListener('focus', handleFocus, true);
	});

	onDestroy(() => {
		window.removeEventListener('focus', handleFocus, true);
	});
</script>

{#if isOpen}
	<div class="menu" bind:this={menuElement}>
		<slot />
	</div>
{/if}

<style lang="less">
	.menu {
		position: absolute;
		display: flex;
		flex-direction: column;
		background-color: #222;
		border: 1px solid #666;
		padding: 10px 0;
	}
</style>
