import type { MeTag, TagOption } from '../stores/interfaces';
import api, { processError } from '../utils/api';
import status from '../utils/status';

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
