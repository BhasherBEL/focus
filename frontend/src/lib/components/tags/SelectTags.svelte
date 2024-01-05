<script lang="ts">
	import { updateCardTagApi } from '$lib/api/cards';
	import TrashIcon from '$lib/components/icons/TrashIcon.svelte';
	import Menu from '$lib/components/menu/Menu.svelte';
	import cards from '$lib/stores/cards';
	import project_tags from '$lib/stores/projectTags';
	import type Card from '$lib/types/Card';
	import type MeTag from '$lib/types/MeTag';
	import type TagValue from '$lib/types/TagValue';
	import api, { processError } from '$lib/utils/api';
	import status from '$lib/utils/status';

	export const multiple: boolean = false;
	export let card: Card;
	export let projectTag: MeTag;
	export let tagValue: TagValue | undefined;

	let lastTagValue = { ...tagValue };
	let newOption: string = '';

	$: tagOption = projectTag.options.find((o) => o.id === tagValue?.option_id);

	let isOpen = false;

	async function selectOption(option_id: number | null) {
		if (lastTagValue.option_id === option_id) {
			isOpen = false;
			return;
		}

		if (tagValue) {
			await updateCardTagApi(card.id, projectTag.id, option_id, tagValue.value);

			card.tags = card.tags.map((t) => {
				if (t.tag_id === projectTag.id) {
					t.option_id = option_id;
				}
				return t;
			});

			tagValue = { ...tagValue, option_id };
		} else {
			const response = await api.post(`/v1/cards/${card.id}/tags/${projectTag.id}`, {
				option_id,
				value: ''
			});

			if (response.status !== status.Created) {
				processError(response, 'Failed to create tag');
				return;
			}

			tagValue = {
				card_id: card.id,
				tag_id: projectTag.id,
				option_id,
				value: ''
			};

			card.tags.push(tagValue);
		}
		lastTagValue = { ...tagValue };

		isOpen = false;

		cards.reload();
	}

	async function deleteOption() {
		const response = await api.delete(`/v1/cards/${card.id}/tags/${projectTag.id}`);

		if (response.status !== status.NoContent) {
			processError(response, 'Failed to delete tag');
			return;
		}

		tagValue = undefined;

		card.tags = card.tags.filter((t) => t.tag_id !== projectTag.id);

		cards.reload();
	}

	function createOption() {
		if (!newOption) return;
		project_tags.addOption(projectTag.id, newOption);
		newOption = '';
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
		{#if tagValue}
			<span class="tag">
				{tagOption?.value}
				<button class="real" on:click={() => deleteOption()}>✗</button>
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
	{#each projectTag.options as option}
		<div
			class="option"
			on:click={() => selectOption(option.id)}
			tabindex="0"
			role="button"
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					selectOption(option.id);
				}
			}}
		>
			<span class="value">{option.value}</span>
			<button
				on:click|stopPropagation={() => {
					project_tags.deleteOption(projectTag.id, option.id);
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
