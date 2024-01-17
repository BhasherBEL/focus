<script lang="ts">
	import Menu from '$lib/components/menu/Menu.svelte';
	import type Card from '$lib/types/Card';
	import { cards } from '$lib/types/Card';
	import CardTag from '$lib/types/CardTag';
	import type ProjectTag from '$lib/types/ProjectTag';
	import { projectTags } from '$lib/types/ProjectTag';
	import type TagOption from '$lib/types/TagOption';
	import TrashIcon from '../icons/TrashIcon.svelte';

	export const multiple: boolean = false;
	export let card: Card;
	export let projectTag: ProjectTag;
	export let cardTag: CardTag | undefined;

	let newOption: string = '';
	let isOpen = false;

	async function selectOption(option: TagOption | null) {
		if (cardTag?.option === option) {
			isOpen = false;
			return;
		}

		if (cardTag) {
			if (option) await cardTag.update(option, null);
			else await card.deleteTag(cardTag);
		} else {
			if (option) await card.addTag(projectTag, option, null);
		}
		isOpen = false;
		cards.reload();

		card.showModal = true;
	}

	async function createOption() {
		if (!newOption) return;
		if (!(await projectTag.addOption(newOption))) return;
		newOption = '';
		projectTags.reload();
	}
</script>

<div
	class="select"
	on:click={() => (isOpen = !isOpen)}
	tabindex="0"
	role="button"
	on:keydown={(e) => {
		if (e.key === 'Enter') {
			isOpen = !isOpen;
		}
	}}
>
	<div class="tags">
		{#if cardTag}
			<span class="tag">
				{cardTag.option?.value}
				<button class="real" on:click={() => selectOption(null)}>✗</button>
			</span>
		{/if}
	</div>
</div>
<Menu
	{isOpen}
	onLeave={() => {
		isOpen = false;
	}}
>
	{#each projectTag.options as option (option.id)}
		<div
			class="option"
			on:click={() => selectOption(option)}
			tabindex="0"
			role="button"
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					selectOption(option);
				}
			}}
		>
			<span class="value">{option.value}</span>
			<button
				on:click|stopPropagation={async () => {
					await projectTag.deleteOption(option);
					projectTags.reload();
				}}><TrashIcon size={16} /></button
			>
		</div>
	{/each}
	<div class="new">
		<input
			type="text"
			placeholder="New option"
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					createOption();
				}
			}}
			bind:value={newOption}
		/>
		<button on:click={createOption}>✓</button>
	</div>
</Menu>

<style lang="less">
	.select {
		display: flex;
		align-items: center;
		min-width: 150px;
		border-radius: 5px;
		padding: 3px 20px 3px 10px;
		height: 30px;
		cursor: pointer;
	}

	.select:hover {
		background-color: #fff2;
	}

	.tag {
		background-color: #444;
		border-radius: 3px;
		padding: 1px 2px;
		cursor: default;

		button {
			padding: 2px 4px;
			background-color: transparent;
			border: none;
			color: inherit;
			cursor: pointer;
		}
	}

	.option {
		display: flex;
		align-items: center;
		justify-content: space-between;
		min-width: 150px;
		cursor: pointer;

		&:hover {
			background-color: #fff2;
		}

		.value {
			margin: 5px 10px;
			border-radius: 5px;
			padding: 5px 8px;
			background-color: #fff2;
		}
		button {
			display: flex;
			align-items: center;
			margin-left: 15px;
			padding: 10px;
			background-color: transparent;

			&:hover {
				background-color: #fff2;
			}
		}
	}

	.new {
		display: flex;
		align-items: center;
		justify-content: space-between;

		input {
			width: 50%;
			margin: 5px 10px;
			padding: 5px;
			border-radius: 5px;
			border: 1px solid #666;
			background-color: transparent;
			color: inherit;
		}

		button {
			margin-left: 15px;
			padding: 10px;
		}
	}
</style>
