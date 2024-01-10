import projectsApi from '$lib/api/projectsApi';

interface Shortcut {
	name: string;
	description: string;
	url: string;
	icons: {
		src: string;
		sizes: string;
	}[];
}

export async function GET() {
	const icon = {
		src: '/img/icon.svg',
		type: 'image/svg+xml',
		sizes: 'any'
	};

	const projects = await projectsApi.getAll();

	const shortcuts: Shortcut[] = projects.map((project) => {
		return {
			name: `Project ${project.title}`,
			description: `Shortcut for project ${project.title}`,
			url: `/${project.id}`,
			icons: [icon]
		};
	});

	const manifest = {
		short_name: 'Focus',
		name: 'Focus',
		start_url: '/',
		display: 'standalone',
		icons: [icon],
		shortcuts
	};

	return new Response(JSON.stringify(manifest), {
		headers: {
			'Content-Type': 'application/manifest+json'
		}
	});
}
