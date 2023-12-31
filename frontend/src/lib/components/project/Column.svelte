<script lang="ts">
	import Card, { cards } from '$lib/types/Card';
	import CardComponent from '../card/Card.svelte';
	import AddIcon from '../icons/AddIcon.svelte';
	import type TagOption from '$lib/types/TagOption';
	import type ProjectTag from '$lib/types/ProjectTag';
	import type Project from '$lib/types/Project';
	import currentDraggedCard from '$lib/stores/currentDraggedCard';

	export let project: Project;
	export let option: TagOption | null = null;
	export let primaryTag: ProjectTag | null = null;
	export let columnCards: Card[] = [];

	let newOptionValue = option?.value || `No ${primaryTag?.title}`;

	async function onDrop(e: DragEvent) {
		e.preventDefault();
		if (!primaryTag || !$currentDraggedCard || !$currentDraggedCard.cardTags) return;

		const currentCardTag =
			$currentDraggedCard.cardTags.find((tag) => tag.projectTag === primaryTag) || null;
		const currentOption = currentCardTag?.option || null;

		if (currentOption === option) return;

		if (!currentOption && option) {
			await $currentDraggedCard.addTag(primaryTag, option, null);
		} else if (currentOption && !option) {
			if (!currentCardTag) return;
			await $currentDraggedCard.deleteTag(currentCardTag);
		} else if (currentOption && option) {
			if (!currentCardTag) return;
			await $currentDraggedCard.updateTag(currentCardTag, option, null);
		}

		currentDraggedCard.set(null);

		cards.reload();
	}

	async function addCard() {
		const card = await Card.create(project);

		if (!card) return;
		if (!primaryTag) return;
		if (!option) return;

		await card.addTag(primaryTag, option, null);

		cards.reload();
	}
</script>

<div
	class="column"
	role="listbox"
	tabindex="-1"
	on:drop={onDrop}
	on:dragover={(e) => {
		e.preventDefault();
	}}
>
	<header>
		<input
			bind:value={newOptionValue}
			type="text"
			on:blur={async () => {
				if (!option || !primaryTag) return;
				if (newOptionValue === option.value) return;

				await option.setValue(newOptionValue);
				newOptionValue = option.value;
			}}
			disabled={option === null}
		/>
		<span>
			<span>{columnCards.length}</span>
			<span
				class="add"
				on:click={addCard}
				role="button"
				tabindex="0"
				on:keypress={(e) => {
					if (e.key === 'Enter') {
						addCard();
					}
				}}
			>
				<AddIcon />
			</span>
		</span>
	</header>
	<ul>
		{#each columnCards as card}
			<CardComponent {card} />
		{/each}
	</ul>
	<div
		class="addEnd"
		on:click={addCard}
		role="button"
		tabindex="0"
		on:keypress={(e) => {
			if (e.key === 'Enter') {
				addCard();
			}
		}}
	>
		+
	</div>
</div>

<style lang="less">
	.column {
		margin: 20px 10px;
		width: 250px;
	}

	header {
		margin-bottom: 20px;
		display: flex;
		flex-direction: row;
		justify-content: space-between;

		input {
			background-color: #444;
			flex: 1;
			margin-right: 10px;
			padding: 3px 10px;
			border-radius: 3px;
			overflow: hidden;
			color: inherit;
			border: none;
		}

		:last-child {
			display: flex;
		}

		.add {
			cursor: pointer;
			border-radius: 4px;
			width: 20px;
			height: 20px;
			display: flex;
			justify-content: center;
			align-items: center;
			margin-left: 10px;

			&:hover {
				background-color: #fff1;
			}
		}
	}

	.addEnd {
		cursor: pointer;
		border-radius: 4px;
		height: 40px;
		font-size: 150%;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 10px;
		margin-bottom: 10px;

		&:hover {
			background-color: #fff1;
		}
	}
</style>
