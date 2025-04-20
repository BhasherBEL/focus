<script lang="ts">
	import Menu from '$lib/components/menu/Menu.svelte';
	import { getTagTypeFromId, tagTypes } from '$lib/utils/tagTypes';

	export let type: number;
	export let isTagMenuOpen: boolean = false;
	export let onChoice: (id: number) => void;

	async function onLocalChoice(id: number) {
		onChoice(id);

		isTagMenuOpen = false;
	}
</script>

<button on:click={() => (isTagMenuOpen = !isTagMenuOpen)}>
	Type: {getTagTypeFromId(type)?.name}
</button>
<div class="menu">
	<Menu
		isOpen={isTagMenuOpen}
		onLeave={() => {
			isTagMenuOpen = false;
		}}
	>
		{#each Object.values(tagTypes) as tagType}
			<div
				class="menu-item"
				role="button"
				on:click={() => onLocalChoice(tagType.id)}
				tabindex="0"
				on:keypress={(e) => {
					if (e.key === 'Enter') {
						onLocalChoice(tagType.id);
					}
				}}
			>
				<span>{tagType.name}</span>
				{#if type === tagType.id}
					<span class="mark"> ✓ </span>
				{:else}
					<span class="mark"> </span>
				{/if}
			</div>
		{/each}
	</Menu>
</div>

<style lang="less">
	button {
		padding: 3px 5px;
		text-align: left;
		width: 80%;
	}

	.menu {
		position: relative;
		bottom: 30px;
		left: 105px;
	}

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
