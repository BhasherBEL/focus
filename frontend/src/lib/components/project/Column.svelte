<script lang="ts">
	import Card from '$lib/types/Card';
	import CardComponent from '../card/Card.svelte';
	import AddIcon from '../icons/AddIcon.svelte';
	import type TagOption from '$lib/types/TagOption';
	import type ProjectTag from '$lib/types/ProjectTag';
	import type Project from '$lib/types/Project';

	export let project: Project;
	export let option: TagOption | null = null;
	export let primaryTag: ProjectTag | null = null;
	export let title: string;
	export let columnCards: Card[] = [];

	let lastTitle = title;

	// async function onDrop(e: DragEvent) {
	// 	e.preventDefault();
	// 	if (!$currentDraggedCard || !$currentDraggedCard.cardTags) return;

	// 	$currentDraggedCard;

	// 	for (let tag of $currentDraggedCard.cardTags) {
	// 		if (tag.projectTag !== primaryTag) continue;
	// 		if (tag.option == option) return;

	// 		if (!tag.option && !tag.value) await tag.delete();
	// 		else if (tag.option && optionId)
	// 			await createCardTagApi(tag.card_id, tag.tag_id, tag.option_id, tag.value);
	// 		else await updateCardTagApi(tag.card_id, tag.tag_id, optionId, tag.value);

	// 		tag.option_id = optionId;
	// 		cards.reload();
	// 	}
	// 	currentDraggedCard.set(null);
	// }

	async function addCard() {
		const card = await Card.create(project);

		if (!card) return;
		if (!primaryTag) return;
		if (!option) return;

		await card.addTag(primaryTag, option, null);
	}
</script>

<!-- on:drop={onDrop} -->
<div
	class="column"
	role="listbox"
	tabindex="-1"
	on:dragover={(e) => {
		e.preventDefault();
	}}
>
	<header>
		<input
			bind:value={title}
			type="text"
			on:blur={async () => {
				if (lastTitle === title) return;
				if (!option || !primaryTag) return;
				// option;
				// await updateTagOptionAPI({
				// 	id: optionId,
				// 	tag_id: primary_tag_id,
				// 	value: title
				// });
				// lastTitle = title;
				// cards.reload();
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
