<script lang="ts">
	import Menu from '$lib/components/menu/Menu.svelte';
	import type ProjectTag from '$lib/types/ProjectTag';

	export let isOpen = false;
	export let choices: ProjectTag[] = [];
	export let onChoice = (projectTag: ProjectTag) => {};
	export let currentChoice: ProjectTag | null;
	export let currentDirection: number | null;
</script>

<Menu bind:isOpen>
	{#each choices as choice}
		<div
			class="menu-item"
			on:click={() => onChoice(choice)}
			role="button"
			tabindex="0"
			on:keypress={(e) => {
				if (e.key === 'Enter') {
					onChoice(choice);
				}
			}}
		>
			<span>{choice.title}</span>
			{#if currentChoice === choice}
				<span class="mark">
					{#if currentDirection === 1}
						↑
					{:else}
						↓
					{/if}
				</span>
			{/if}
		</div>
	{/each}
</Menu>

<style lang="less">
	.menu-item {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 5px 20px;

		&:hover {
			cursor: pointer;
			background-color: #333;
		}
	}

	.mark {
		margin-left: 20px;
	}
</style>
