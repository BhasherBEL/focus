<script lang="ts">
	import Menu from '../../../../utils/menu.svelte';

	export let isOpen = false;
	export let choices: { id: number; value: string }[] = [];
	export let onChoice = (id: number) => {};
	export let currentChoice: number | null;
	export let currentDirection: number | null;
</script>

<Menu bind:isOpen>
	{#each choices as choice}
		<div
			class="menu-item"
			on:click={() => onChoice(choice.id)}
			role="button"
			tabindex="0"
			on:keypress={(e) => {
				if (e.key === 'Enter') {
					onChoice(choice.id);
				}
			}}
		>
			<span>{choice.value}</span>
			{#if currentChoice === choice.id}
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
