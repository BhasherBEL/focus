<script lang="ts">
	export let isOpen = false;
	export let choices: { id: number; value: string }[] = [];
	export let onChoice = (id: number) => {};
	export let currentChoice: number = -1;
</script>

{#if isOpen}
	<div class="menu">
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
					<span class="mark"> ✓ </span>
				{/if}
			</div>
		{/each}
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
	}
</style>
