export const tagTypes: { [name: string]: { id: number; name: string; hasOptions: boolean } } = {
	SELECT: {
		id: 0,
		name: 'Select',
		hasOptions: true
	},
	MULTISELECT: {
		id: 1,
		name: 'Multiselect',
		hasOptions: true
	},
	TEXT: {
		id: 2,
		name: 'Text',
		hasOptions: false
	}
};

export const getTagTypeFromId = (id: number) => {
	for (let type in tagTypes) {
		if (tagTypes[type].id === id) {
			return tagTypes[type];
		}
	}
	return null;
};
