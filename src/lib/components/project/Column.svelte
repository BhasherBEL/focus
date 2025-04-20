<script lang="ts">
	import currentDraggedCard from '$lib/stores/currentDraggedCard';
	import currentView from '$lib/stores/currentView';
	import Card, { cards } from '$lib/types/Card';
	import type Project from '$lib/types/Project';
	import ProjectTag, { projectTags } from '$lib/types/ProjectTag';
	import type TagOption from '$lib/types/TagOption';
	import CardComponent from '../card/Card.svelte';
	import AddIcon from '../icons/AddIcon.svelte';

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

	let openModalCard: Card | null = null;

	async function addCard() {
		const card = await Card.create(project);

		if (!card) return;

		if ($currentView?.filters && $currentView.filters.length > 0) {
			for (const projectTag of $projectTags) {
				for (const filter of $currentView.filters) {
					if (projectTag !== filter.projectTag) continue;
					if (!filter.tagOption) continue;
					if (filter.filterType !== 0) continue;

					if (await card.addTag(projectTag, filter.tagOption, null)) break;
				}
			}
		}

		if (primaryTag && option) {
			await card.addTag(primaryTag, option, null);
		}

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
		{#each columnCards as card (card.id)}
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
