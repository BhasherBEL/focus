<script lang="ts">
	import AddIcon from '$lib/components/icons/addIcon.svelte';
	import Menu from '$lib/components/utils/menu.svelte';
	import project_tags from '$lib/stores/projectTags';
	import { toastAlert } from '$lib/utils/toasts';
	import { tick } from 'svelte';
	import ModalTagTypes from './modal_tag_types.svelte';

	export let projectId: number;

	let isOpen = false;

	let titleInput: HTMLInputElement;

	let title: string = '';
	let typeId: number = 0;

	async function openMenu() {
		isOpen = !isOpen;
		await tick();
		if (isOpen && titleInput) titleInput.focus();
	}

	async function createTag() {
		if (title == '') return;
		if (!projectId) {
			toastAlert('Failed to create tag', `ProjectId is ${projectId}`);
			return;
		}
		await project_tags.add(projectId, title, typeId);
		isOpen = false;
	}
</script>

<tr>
	<td>
		<button
			class="add-button"
			on:click={openMenu}
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					openMenu();
				}
			}}
		>
			<AddIcon />
		</button>
		<Menu
			{isOpen}
			onLeave={() => {
				isOpen = false;
			}}
		>
			<div class="menu-items">
				<input bind:this={titleInput} bind:value={title} />
				<ModalTagTypes type={typeId} onChoice={(id) => (typeId = id)} />
				<button on:click={createTag}>Create</button>
			</div>
		</Menu>
	</td>
</tr>

<style lang="less">
	.add-button {
		margin-top: 10px;
		padding: 3px 5px;
		border-radius: 5px;
		width: 100%;
		height: 100%;
	}

	.menu-items {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;

		input {
			margin-bottom: 5px;
			width: 70%;
			background-color: inherit;
			color: inherit;
			border: 1px solid #666;
			padding: 3px 5px;
			font-size: inherit;
			border-radius: 5px;
		}

		button {
			padding: 3px 5px;
			text-align: left;
			width: 80%;
		}
	}
</style>
