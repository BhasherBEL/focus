import type TagOption from '$lib/types/TagOption';
import api, { processError } from '$lib/utils/api';
import status from '$lib/utils/status';

export async function updateTagAPI(option: TagOption): Promise<boolean> {
	const response =
		option.value === ''
			? await api.delete(`/v1/tags/${option.tag_id}/options/${option.id}`)
			: await api.put(`/v1/tags/${option.tag_id}/options/${option.id}`, option);

	if (response.status !== status.NoContent) {
		processError(response, 'Failed to update tag option');
		return false;
	}

	return true;
}
